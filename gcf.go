package hhooking

import (
	"crypto/ed25519"
	"encoding/hex"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type GCFInteractionFunction func(http.ResponseWriter, *http.Request)
type GCFInteractionHandler func(Interaction) InteractionReponse

func CreateInteractionHandler(hexEncodedKey string, h GCFInteractionHandler) GCFInteractionFunction {
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

		buf := make([]byte, 0)
		defer r.Body.Close()
		for {
			size := 64
			v := make([]byte, size)
			resSize, err := r.Body.Read(v)
			if err != nil {
				// TODO: err handling
			}

			buf = append(buf, v...)

			if resSize < size {
				break
			}
		}

		jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(buf, &body)

		if body.Type == ItPing {
			rep, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(
				InteractionReponse{
					Type: IctPong,
				},
			)
			if err != nil {
				// TODO: err handling
			}

			w.Write(rep)
			return
		}

		repStruct := h(body)

		rep, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(repStruct)
		if err != nil {
			// TODO: err handling
		}

		w.Write(rep)
	}
}
