package GoWaves

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type api struct {
	addr string
}

var nodeAPI = api{}

// Get http request
func (a *api) get(url string, target interface{}) error {
	r, err := http.Get(a.addr + url)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// Post http request
func (a *api) post(url string, target interface{}, form url.Values) error {
	body := bytes.NewBufferString(form.Encode())
	r, err := http.Post(a.addr+url, "json/application", body)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
