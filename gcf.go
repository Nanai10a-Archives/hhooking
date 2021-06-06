package hhooking

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type GCFInteractionFunction func(http.ResponseWriter, *http.Request)
type GCFInteractionHandler func(Interaction) InteractionReponse

func CreateInteractionHandler(hexEncodedKey string, h GCFInteractionHandler) GCFInteractionFunction {
	decodeKey, err := hex.DecodeString(hexEncodedKey)
	if err != nil {
		// TODO: err handling
		panic("failed decode key.")
	}
	key := ed25519.PublicKey(decodeKey)

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received: %v", r)

		if !SignatureVerify(r, key) {
			http.Error(w, "Signature checking failed.", http.StatusUnauthorized)
			log.Printf("failed signature verify\nkey: %v\nHeaders: %v\n", key, r.Header)
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
				panic("failed body reading.")
			}

			buf = append(buf, v...)

			if resSize < size {
				break
			}
		}

		err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(buf, &body)
		if err != nil {
			// TODO: err handling
			panic("failed unmarshal body.")
		}
		log.Printf("raw Body :%v\nparsed Body: %v\n", buf, body)

		if body.Type == ItPing {
			rep, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(
				InteractionReponse{
					Type: IctPong,
				},
			)
			if err != nil {
				// TODO: err handling
				panic("failed marshal reply.")
			}

			log.Printf("Ping/Pong! response: %v\n", rep)

			w.Header().Add("Content-Type", "application/json")
			w.Write(rep)
			return
		}

		repStruct := h(body)

		rep, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(repStruct)
		if err != nil {
			// TODO: err handling
			panic("failed marshal reply.")
		}

		log.Printf("Response: %v\n", rep)

		w.Header().Add("Content-Type", "application/json")
		w.Write(rep)
	}
}
