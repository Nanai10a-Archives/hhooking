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
	url, err := url.Parse(BaseAPIUrl)
	if err != nil {
		// TODO: err handling
	}

	url.Path = path.Join(url.Path, "applications", fmt.Sprintf("%d", id), "commands")

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		// TODO: err handling
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", token))

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		// TODO: err handling
	}

	var v []ApplicationCommand
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		// TODO: err handling
	}

	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, &v)
	if err != nil {
		// TODO: err handling
	}

	return v
}

func sendRequest(targetPath string, method string, content []byte, token string, rep interface {}) {
    url, err := url.Parse(BaseAPIUrl)
    if err != nil {
        // TODO: err handling
    }

    url.Path = path.Join(url.Path, targetPath)

    buf := bytes.NewBuffer(content)
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

