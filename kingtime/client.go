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

func New(URL, token string) *Kingtime {
	return &Kingtime{
		URL,
		token,
	}
}

func (k *Kingtime) post(path string, params []byte) ([]byte, error) {
	req, err := http.NewRequest("Post", k.URL+path, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("ContentType", "application/json")
	req.Header.Set("Authorization", "Bearer "+k.token)

	cli := http.Client{}
	res, err := cli.Do(req)
	log.Info("Requesting")
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
