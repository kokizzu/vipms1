package app

import (
	"errors"
	"log"
	"vipms1/app/config"
	"vipms1/app/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// https://gofiber.io/

func MustOk(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func IsError(c *fiber.Ctx, err error) bool {
	if err != nil {
		c.JSON(model.Response{
			Error: err,
		})
		return true
	}
	return false
}

func Run() {
	db, err := config.ConnectMysql()
	MustOk(err)

	MustOk(model.CreateVipTable(db))

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{ // https://docs.gofiber.io/api/middleware/basicauth
			config.API_AUTH_USER: config.API_AUTH_PASS,
		},
	}))

	app.Get("/vips/all", func(c *fiber.Ctx) error {
		vips, err := model.VipsAll(db)
		return model.CreateResponse(c, vips, err)
	})

	app.Get(`/vips/by-id/:id`, func(c *fiber.Ctx) error {
		id, err := c.ParamsInt(`id`)
		if IsError(c, err) {
			return nil
		}
		vip := model.Vip{}
		err = vip.GetById(db, int64(id))
		return model.CreateResponse(c, vip, err)
	})

	app.Post(`/vips/create`, func(c *fiber.Ctx) error {
		vip := model.Vip{}
		err := c.BodyParser(&vip)
		if IsError(c, err) {
			return nil
		}
		if vip.IsNotValid(c) {
			return nil
		}
		id, err := vip.Insert(db)
		return model.CreateResponse(c, id, err)
	})

	app.Patch(`/vips/arrive/:id`, func(c *fiber.Ctx) error {
		id, err := c.ParamsInt(`id`)
		if IsError(c, err) {
			return nil
		}
		vip := model.Vip{}
		err = vip.GetById(db, int64(id))
		if IsError(c, err) {
			return nil
		}
		if vip.ID == 0 {
			IsError(c, errors.New(`vip not found on database`))
			return nil
		}
		if vip.Arrived {
			IsError(c, errors.New(`vip already arrived`))
			return nil
		}
		err = vip.SetArrived(db, 0)
		return model.CreateResponse(c, vip, err)
	})

	log.Fatal(app.Listen(config.API_PORTT))
}
