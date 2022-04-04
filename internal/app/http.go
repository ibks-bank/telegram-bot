package app

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type body interface {
	Marshall() []byte
}

func (a *app) get(url, token string) ([]byte, error) {
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-auth-token", token)

	client := http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyBytes, nil
}

func (a *app) post(url, token string, body body) ([]byte, error) {
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body.Marshall()))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-auth-token", token)

	client := http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyBytes, nil
}
