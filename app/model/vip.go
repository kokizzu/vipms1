package model

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Vip struct {
	ID         int64
	Name       string
	Country    string
	ETA        time.Time
	Arrived    bool
	PhotoUrl   string
	Attributes string
}

func (v *Vip) GetById(db *sqlx.DB, id int64) error {
	err := db.Select(&v, `SELECT id, name, country, eta, arrived, photo_url, attributes
FROM vips
WHERE id = ?`, id)
	return err
}

func (vip *Vip) IsNotValid(c *fiber.Ctx) bool {
	if NotValid(c, len(vip.Name) < 3, `name must be minimum 3 characters`) {
		return true
	}
	if NotValid(c, len(vip.Country) < 3, `country must be minimum 3 characters`)) {
		return true
	}
	// TODO: check other things
	return false
}

func (v *Vip) Insert(db *sqlx.DB) (int64, error) {
	rs, err := db.Exec(`INSERT INTO vips(name, country, eta, arrived, photo_url, attributes)
VALUES(?,?,?,?,?,?)
`, v.Name, v.Country, v.ETA, v.Arrived, v.PhotoUrl, v.Attributes)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (v *Vip) SetArrived(db *sqlx.DB, id int64) error {
	v.Arrived = true
	_, err := db.Exec(`UPDATE vips SET arrived = true WHERE id = ?`, v.ID)
	return err
}

func CreateVipTable(db *sqlx.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS vips (
   ID INTEGER AUTO_INCREMENT PRIMARY KEY 
   name VARCHAR(40) NOT NULL,
   country VARCHAR(40) NOT NULL,
   eta DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
   arrived BOOL DEFAULT false NOT NULL,
   photo_url VARCHAR(120) NOT NULL,
   attributes TEXT DEFAULT '[]' NOT NULL
)`)
	return err
}

func VipsAll(db *sqlx.DB) ([]Vip, error) {
	vips := []Vip{}
	err := db.Select(&vips, `SELECT id, name, country, eta, arrived, photo_url, attributes
FROM vips`)
	return vips, err
}
