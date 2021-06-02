package hhooking

import (
	"github.com/gofiber/fiber/v2"
	jsonitor "github.com/json-iterator/go"
)

type InteractionServer struct {
	rawSrv   *fiber.App
	handlers []StandaloneInteractionHandler
}

type StandaloneInteractionHandler func(*fiber.Ctx)

func NewInteractionServer() *InteractionServer {
	srv := InteractionServer{
		rawSrv: fiber.New(),
	}

	srv.rawSrv.Post("/", func(c *fiber.Ctx) error {
		var body Interaction
		jsonitor.ConfigCompatibleWithStandardLibrary.Unmarshal(c.Body(), &body)

		if body.Type == 1 {
			rep, err := jsonitor.ConfigCompatibleWithStandardLibrary.MarshalToString(
				struct {
					Type InteractionCallbackType `json:"type"`
				}{
					Type: IctPong,
				},
			)
			if err != nil {
				// TODO: err handling
			}

			return c.SendString(rep)
		}

		for _, h := range srv.handlers {
			h(c)
		}

		return nil
	})

	return &srv
}

func (s *InteractionServer) RegisterHandlers(hs ...Handler) {
	s.handlers = append(s.handlers, hs...)
}
