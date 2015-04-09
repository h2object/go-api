package api

import (
	"github.com/h2object/httpRPC"
)

type Client struct{
	token 	string
	isLogin bool
	isAdmin bool
	Conn rpc.Client
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetAdmin(access_key, secret_key string) {
	if access_key != "" && secret_key != "" {
		ACCESS_KEY = access_key
		SECRET_KEY = secret_key
	}
	isAdmin = true
}

func (c *Client) SignIn(provider string, auth string, credental string) error {

}

func (c *Client) SignOff() error {

}
