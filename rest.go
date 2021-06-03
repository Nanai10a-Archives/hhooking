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

func GetGlobalApplicationCommands(appId SnowFlake, token string) []ApplicationCommand {
    reqPath := path.Join("applications", fmt.Sprintf("%d", appId), "commands")
	var v []ApplicationCommand
    sendRequest(reqPath, "GET", nil, token, &v)

	return v
}

type ApplicationCommandPostData struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Options []ApplicationCommandOption `json:"options,omitempty"`
    DefaultPermisson bool `json:"default_permisson,omitempty"`
}

func CreateGlobalApplicationCommands(appId SnowFlake, token string, data ApplicationCommandPostData) ApplicationCommand {
    reqPath := path.Join("applications", fmt.Sprintf("%d", appId), "commands")
    var v ApplicationCommand
    content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
    if err != nil {
        // TODO: err handling
    }
    sendRequest(reqPath, "POST", content, token, &v)

    return v
}

func GetGlobalApplicationCommand(appId SnowFlake, token string, cmdId SnowFlake) ApplicationCommand {
    reqPath := path.Join("applications", fmt.Sprintf("%d", appId), "commands", fmt.Sprintf("%d", cmdId))
    var v ApplicationCommand
    sendRequest(reqPath, "GET", nil, token, &v)

    return v
}

// FIXME: 内容は同じだけどtagだけ違うので微妙なところではある "...PostData"との統合を検討.
type ApplicationCommandPatchData struct {
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Options []ApplicationCommandOption `json:"options,omitempty"`
    DefaultPermisson bool `json:"default_permisson,omitempty"`
}

func EditGlobalApplicationCommand(appId SnowFlake, token string, cmdId SnowFlake, data ApplicationCommandPatchData) ApplicationCommand {
    reqPath := path.Join("applications", fmt.Sprintf("%d", appId), "commands", fmt.Sprintf("%d", cmdId))
    var v ApplicationCommand
    content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
    if err != nil {
        // TODO: err handling
    }
    sendRequest(reqPath, "PATCH", content, token, &v)

    return v
}

// FIXME: error返しません?
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
