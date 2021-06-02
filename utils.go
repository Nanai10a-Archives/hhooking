package hhooking

import (
	"crypto/ed25519"
	"net/http"
    "encoding/json"
    "io"

	external "github.com/bsdlp/discord-interactions-go/interactions"
)

func SignatureVerify(r *http.Request, key ed25519.PublicKey) bool {
    verified := external.Verify(r, key)
    return verified
}

func decodeJSON(r io.Reader, v *interface {}) {
    bytes, err := io.ReadAll(r)

    if err != nil {
        // TODO: err handling
    }

    json.Unmarshal(bytes, v)
}

