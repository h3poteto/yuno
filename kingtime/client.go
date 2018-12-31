package kingtime

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Kingtime struct {
	URL   string
	token string
}

var URL = "https://api.kingtime.jp/v1.0"

func New(token string) *Kingtime {
	return &Kingtime{
		URL,
		token,
	}
}

func (k *Kingtime) get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", k.URL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+k.token)
	log.Info(k.token)
	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("Parse error")
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("Status is invalid: %s", body)
	}

	return body, nil
}

func (k *Kingtime) post(path string, params []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", k.URL+path, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+k.token)

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("Parse error")
		return nil, err
	}

	if res.StatusCode != http.StatusOK || res.StatusCode != http.StatusCreated {
		return nil, errors.Errorf("Status is invalid: %s", body)
	}

	return body, nil
}
