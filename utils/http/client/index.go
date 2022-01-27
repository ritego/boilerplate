package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

func Request(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, err
}

func Get(url string) ([]byte, error) {
	return Request(http.MethodGet, url, nil)
}

func Post(url string, body io.Reader) ([]byte, error) {
	return Request(http.MethodPost, url, body)
}

func Patch(url string, body io.Reader) ([]byte, error) {
	return Request(http.MethodPatch, url, body)
}

func Put(url string, body io.Reader) ([]byte, error) {
	return Request(http.MethodPut, url, body)
}

func Delete(url string) ([]byte, error) {
	return Request(http.MethodDelete, url, nil)
}
