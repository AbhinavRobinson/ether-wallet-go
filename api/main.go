package api

import (
	app "ether-wallet-go/app"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Init() {
	app := fiber.New()

	app.Get("/", statusCheck)
	app.Post("/simple/new", createNewSimpleWallet)

	app.Listen("127.0.0.1:3000")
}

func statusCheck(c *fiber.Ctx) error {
	return c.SendString("available")
}

func createNewSimpleWallet(c *fiber.Ctx) error {
	bodyParams := new(CreateSimpleWalletBody)
	if err := c.BodyParser(bodyParams); err != nil {
		return err
	}
	wallet, err := app.GenerateSimpleWallet(bodyParams.Nickname, bodyParams.ChainConfig)
	if err != nil {
		c.SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("%v", wallet.Address))
}
