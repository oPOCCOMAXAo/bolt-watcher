package bolt

import (
	"context"
	"net/http"

	"github.com/opoccomaxao-go/xhttp"
	"github.com/opoccomaxao-go/xhttp/request"
	"github.com/opoccomaxao-go/xhttp/response"
	"github.com/pkg/errors"
)

type PollRequest struct {
	Stage            Stage         `json:"stage"`
	ViewPort         Viewport      `json:"viewport"`
	PickupStop       Point         `json:"pickup_stop"`
	DestinationStops []Point       `json:"destination_stops"`
	PaymentMethod    PaymentMethod `json:"payment_method"`
}

type PollResponse struct {
	CommonResponse
	Data struct {
		Vehicles struct {
			Taxi map[string]Vehicle `json:"taxi"`
		} `json:"vehicles"`
	} `json:"data"`
}

func (c *Client) Poll(
	ctx context.Context,
	req *PollRequest,
) (*PollResponse, error) {
	var res PollResponse

	err := c.client.Do(ctx, xhttp.Params{
		Method:       http.MethodPost,
		URL:          "https://user.live.boltsvc.net/mobility/search/poll",
		Headers:      c.headers,
		Query:        c.query,
		RequestBody:  request.JSON(req),
		ResponseBody: response.JSON(&res),
	})
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.WithMessage(ErrFailed, res.Message)
	}

	return &res, nil
}
