package hhooking

import (
	"io"
	"net/http"

	jsonitor "github.com/json-iterator/go"
)

type GCFInteractionFunction func(http.ResponseWriter, *http.Request)
type GCFInteractionHandler func(*Interaction)

func CreateInteractionHandler(fs ...GCFInteractionHandler) GCFInteractionFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		var body Interaction
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			// TODO: err handling
		}

		jsonitor.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, &body)

		if body.Type == ItPing {
			rep, err := jsonitor.ConfigCompatibleWithStandardLibrary.Marshal(
				struct {
					Type InteractionCallbackType `json:"type"`
				}{
					Type: IctPong,
				},
			)
			if err != nil {
				// TODO: err handling
			}

			w.Write(rep)
		}

		for _, h := range fs {
			h(&body)
		}
	}
}
