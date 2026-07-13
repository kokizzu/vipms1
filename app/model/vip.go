package model

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Vip struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Country    string `db:"country"`
	ETA        string `db:"eta"`
	Arrived    bool   `db:"arrived"`
	PhotoUrl   string `db:"photo_url"`
	Attributes string `db:"attributes"`
}

func (vip *Vip) ValidationError() string {
	if len(vip.Name) < 3 {
		return `name must be minimum 3 characters`
	}
	if len(vip.Country) < 3 {
		return `country must be minimum 3 characters`
	}
	if _, err := time.Parse(`2006-01-02 15:04:05`, vip.ETA); err != nil {
		return `eta must use format YYYY-MM-DD HH:MM:SS`
	}
	if len(vip.PhotoUrl) > 120 {
		return `photo_url must be maximum 120 characters`
	}
	return ``
}

func (v *Vip) GetById(db *sqlx.DB, id int64) error {
	err := db.Get(v, `SELECT id, name, country, eta, arrived, photo_url, attributes
FROM vips
WHERE id = ?`, id)
	return err
}

func (vip *Vip) IsNotValid(c *fiber.Ctx) bool {
	if errMsg := vip.ValidationError(); errMsg != `` {
		NotValid(c, true, errMsg)
		return true
	}
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
   ID INTEGER AUTO_INCREMENT PRIMARY KEY,
   name VARCHAR(40) NOT NULL,
   country VARCHAR(40) NOT NULL,
   eta DATETIME NOT NULL DEFAULT NOW(),
   arrived BOOL DEFAULT false NOT NULL,
   photo_url VARCHAR(120) NOT NULL,
   attributes TEXT NOT NULL
)`)
	return err
}

func VipsAll(db *sqlx.DB) ([]Vip, error) {
	vips := []Vip{}
	err := db.Select(&vips, `SELECT id, name, country, eta, arrived, photo_url, attributes
FROM vips`)
	return vips, err
}
