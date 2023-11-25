package bolt

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/opoccomaxao-go/xhttp"
)

type Client struct {
	client  *xhttp.Client
	headers http.Header
	query   url.Values
}

func New(login, password string) *Client {
	res := Client{
		client: xhttp.New(http.Client{
			Timeout: time.Minute,
		}),
		headers: http.Header{
			"Authorization": []string{fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", login, password))))},
		},
		query: url.Values{
			"version":           []string{"CB.3.77"},
			"deviceId":          []string{"1"},
			"deviceType":        []string{"web"},
			"device_name":       []string{"Linux x86_64"},
			"device_os_version": []string{"5.0"},
			"country":           []string{"ua"},
			"language":          []string{"en-US"},
		},
	}

	return &res
}
