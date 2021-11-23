package api

import (
	"encoding/base64"
	"fmt"

	"github.com/opoccomaxao-go/request"
)

type API struct {
	client  *request.Client
	headers map[string]string
}

func NewAPI(login, password string) *API {
	res := API{
		client: request.New(1),
		headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", login, password)))),
		},
	}

	return &res
}
