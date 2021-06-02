package hhooking

import (
	"github.com/gofiber/fiber/v2"
	jsonitor "github.com/json-iterator/go"
)

type InteractionServer struct {
	rawSrv  *fiber.App
	handler func()
}

func NewInteractionServer() InteractionServer {
	srv := InteractionServer{
		rawSrv:  fiber.New(),
		handler: func() {},
	}

	srv.rawSrv.Post("/", func(c *fiber.Ctx) error {
		var body Interaction
		jsonitor.ConfigCompatibleWithStandardLibrary.Unmarshal(c.Body(), &body)

		if body.Type == 1 {
			rep, err := jsonitor.ConfigCompatibleWithStandardLibrary.MarshalToString(
				struct {
					Type InteractionType `json:"type"`
				}{
					Type: ItPing,
				},
			)
			if err != nil {
				// TODO: err handling
			}

			return c.SendString(rep)
		}

		return nil
	})

	return srv
}
