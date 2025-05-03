package fiber

import (
	"github.com/RockkleyPushPost/common/config"
	"github.com/gofiber/fiber/v2"
)

func NewFiber(config config.ServerConfig) (*fiber.App, error) {
	if err := config.Validate(); err != nil {

		return nil, err
	}

	app := fiber.App{}

	return &app, nil

}
