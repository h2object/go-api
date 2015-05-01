package api

import (
	"net/http"
)

type HeaderAuth struct{
	headers http.Header
}

func NewHeaderAuth() *HeaderAuth {
	return &HeaderAuth{
		headers: http.Header{},
	}
}

func (hds *HeaderAuth) Add(key, value string) {
	hds.headers.Add(key, value)
}

func (hds *HeaderAuth) Set(key, value string) {
	hds.headers.Set(key, value)
}

func (hds *HeaderAuth) Del(key string) {
	hds.headers.Del(key)
}

func (hds *HeaderAuth) Do(req *http.Request) *http.Request {
	for k, vs := range hds.headers {
		req.Header.Del(k)
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	return req
}

func NewAdminAuth(appid, secret string) Auth {
	auth := NewHeaderAuth()
	auth.Add("X-H2O-AppId", appid)
	auth.Add("X-H2O-Secret", secret)
	return auth
}

func NewTokenAuth(token string) Auth {
	auth := NewHeaderAuth()
	auth.Del("Authorization")
	auth.Add("Authorization", fmt.Sprintf("H2O %s", token))
	return auth
}

func (h2o *H2Object) SignInPassword(email string, password string, remember bool) (Auth, error) {
	return nil, nil
} 

