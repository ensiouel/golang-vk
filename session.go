package vk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Session struct {
	client *http.Client
	token  string
	debug  bool
	//Limiter *Limiter
}

type AuthConfig struct {
	Login, Password, Token string
}

func NewEmptySession() *Session {
	return &Session{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		//Limiter: NewLimiter(3),
	}
}

func NewSession(cfg AuthConfig, debug bool) (*Session, error) {
	session := NewEmptySession()

	switch {
	case cfg.Token != "":
		session.token = cfg.Token
	case cfg.Login != "" && cfg.Password != "":
	default:
		return nil, errors.New("[GOVK] Authorization error")
	}

	session.debug = debug

	if session.debug {
		fmt.Printf("[GOVK] %s | Successful authorization!\n", time.Now().Format("2006/01/02 - 15:04:05"))
	}

	return session, nil
}

func (session *Session) Call(method string, params url.Values) *Request {
	if params == nil {
		params = url.Values{}
	}
	return &Request{
		session: session,
		params:  params,
		method:  method,
	}
}

func (session *Session) MakeRequest(r *http.Request) ([]byte, error) {
	response, err := session.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
