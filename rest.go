package hhooking

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	jsoniter "github.com/json-iterator/go"
)

const (
	BaseAPIUrl = "https://discord.com/api/v8/"
)

func GetGlobalApplicationCommands(id SnowFlake, token string) []ApplicationCommand {
    reqPath := path.Join("applications", fmt.Sprintf("%d", id), "commands")
	var v []ApplicationCommand
    sendRequest(reqPath, "GET", nil, fmt.Sprintf("Bot %s", token), &v)

	return v
}

func sendRequest(targetPath string, method string, content []byte, token string, rep interface {}) {
    url, err := url.Parse(BaseAPIUrl)
    if err != nil {
        // TODO: err handling
    }

    url.Path = path.Join(url.Path, targetPath)

    var buf *bytes.Buffer
    switch content {
    case nil:
        buf = nil
    default:
        buf = bytes.NewBuffer(content)
    }

    req, err := http.NewRequest(method, url.String(), buf)
    if err != nil {
        // TODO: err handling
    }

    req.Header.Add("Authorization", token)

    client := new(http.Client)
    res, err := client.Do(req)
    if err != nil {
        // TODO: err handling
    }

    resBytes, err := io.ReadAll(res.Body)
    if err != nil {
        // TODO: err handling
    }

    err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resBytes, rep)
    if err != nil {
        // TODO: err handling
    }
}

