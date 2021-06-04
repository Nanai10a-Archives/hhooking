package hhooking

import (
	"crypto/ed25519"
	"encoding/hex"
	"io"
	"net/http"

	jsonitor "github.com/json-iterator/go"
)

type GCFInteractionFunction func(http.ResponseWriter, *http.Request)
type GCFInteractionHandler func(Interaction)

func CreateInteractionHandler(hexEncodedKey string, fs ...GCFInteractionHandler) GCFInteractionFunction {
    decodeKey, err := hex.DecodeString(hexEncodedKey)
    if err != nil {
        // TODO: err handling
    }
    key := ed25519.PublicKey(decodeKey)

	return func(w http.ResponseWriter, r *http.Request) {
        if !SignatureVerify(r, key) {
            http.Error(w, "Signature checking failed.", http.StatusUnauthorized)
            return
        }

		var body Interaction
		bytes, err := io.ReadAll(r.Body)
        defer r.Body.Close()
		if err != nil {
			// TODO: err handling
		}

		jsonitor.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, &body)

		if body.Type == ItPing {
			rep, err := jsonitor.ConfigCompatibleWithStandardLibrary.Marshal(
				InteractionReponse {
					Type: IctPong,
				},
			)
			if err != nil {
				// TODO: err handling
			}

			w.Write(rep)
            return
		}

		for _, h := range fs {
			h(body)
		}
	}
}
