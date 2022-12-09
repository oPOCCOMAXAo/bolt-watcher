package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type API struct {
	headers map[string]string
	client  *http.Client

	mu sync.Mutex
}

func NewAPI(login, password string) *API {
	res := API{
		client: &http.Client{
			Timeout: time.Minute,
		},
		headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", login, password)))),
			"Content-Type":  "application/json",
		},
	}

	return &res
}

func (api *API) request(
	ctx context.Context,
	url string,
	params url.Values,
	body interface{},
	resPtr interface{},
	debug bool,
) error {
	api.mu.Lock()
	defer api.mu.Unlock()

	var bodyReader io.Reader
	buf := bytes.Buffer{}

	if body == nil {
		bodyReader = http.NoBody
	} else {
		_ = json.NewEncoder(&buf).Encode(body)

		bodyReader = &buf
	}

	url += "?" + params.Encode()

	if debug {
		log.Printf("%s %s\n%s\n\n", http.MethodPost, url, buf.String())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
	if err != nil {
		return errors.WithStack(err)
	}

	for k, v := range api.headers {
		req.Header.Add(k, v)
	}

	res, err := api.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	defer res.Body.Close()

	if resPtr != nil {
		buf.Reset()

		_, err = io.Copy(&buf, res.Body)
		if err != nil {
			return errors.WithStack(err)
		}

		if debug {
			log.Printf("%s\n\n", buf.String())
		}

		err = json.NewDecoder(&buf).Decode(resPtr)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if res.StatusCode != http.StatusOK {
		return errors.Errorf(res.Status)
	}

	return nil
}
