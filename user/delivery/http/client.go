package http

import (
	"encoding/json"
	"github.com/AnyKey/service/user"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type delivery struct{}

func New() *delivery {
	return &delivery{}
}

const (
	baseURL = "https://www.googleapis.com/youtube/v3/subscriptions"
)

func (*delivery) GetSubscriptions(token string) (*user.List, error) {
	http.DefaultClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error create new request")
	}

	query := req.URL.Query()
	query.Add("part", "id")
	query.Add("part", "contentDetails")
	query.Add("part", "snippet")
	query.Add("mine", "true")
	req.URL.RawQuery = query.Encode()
	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error conv bytes")
	}
	var res user.List
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, errors.Wrap(err, "Error Unmarshal")
	}

	return &res, nil
}
