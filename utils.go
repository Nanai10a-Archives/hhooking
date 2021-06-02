package hhooking

import (
	"crypto/ed25519"
	"net/http"

	external "github.com/bsdlp/discord-interactions-go/interactions"
)

func SignatureVerify(r *http.Request, key ed25519.PublicKey) bool {
    verified := external.Verify(r, key)
    return verified
}
