package api

import (
	"github.com/h2object/rpc"
	"path"
	"net/url"
)

//! --- Security ---
type Security struct{
	URI				string  	`json:"-"`	
	Method 			string	    `json:"-"`
	Require 		string  	`json:"require"`
	AuthID_EQ_Key	bool		`json:"authid_eq_key"`
	// AuthID_EQ_Field string 		`json:"authid_eq_field" object:"authid_eq_field" structs:"authid_eq_field"`	
}

func NewNoneSecurity(uri string, method string) *Security {
	return &Security{
		URI: uri,
		Method: method,
		Require: "none",
	}
}

func NewSignedSecurity(uri string, method string, authid_eq_key bool) *Security {
	return &Security{
		URI: uri,
		Method: method,
		Require: "signed",
		AuthID_EQ_Key: authid_eq_key,
	}	
}

func NewAdminsSecurity(uri string, method string) *Security {
	return &Security{
		URI: uri,
		Method: method,
		Require: "admin",
	}	
}

// `{"require":"none"}`
// `{"require":"signed"}`
// `{"require":"signed", "authid_eq_key":false}`
// `{"require":"signed", "authid_eq_key":true}`
// `{"require":"admin"}`

func (h2o *H2Object) SetSecurity(l Logger, auth Auth, securtiy *Security) error {
	params := url.Values{
		"method": {securtiy.Method},
	}

	URL := rpc.BuildHttpURL(h2o.addr, path.Join(securtiy.URI, ".security"), params)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	ret := map[string]interface{}{}
	if err := h2o.conn.PostJson(l, URL, securtiy, &ret); err != nil {
		return err
	}
	if l != nil {
		l.Debug("ret: %v", ret)
	}
	return nil
}

func (h2o *H2Object) DelSecurity(l Logger, auth Auth, securtiy *Security) error {
	params := url.Values{
		"method": {securtiy.Method},
	}

	URL := rpc.BuildHttpURL(h2o.addr, path.Join(securtiy.URI, ".security"), params)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Delete(l, URL, nil); err != nil {
		return err
	}
	return nil	
}

func (h2o *H2Object) GetSecurity(l Logger, auth Auth, securtiy *Security) error {
	URL := rpc.BuildHttpURL(h2o.addr,path.Join("/securities", path.Join(securtiy.URI, securtiy.Method + ".mtd.system")), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, securtiy); err != nil {
		return err
	}
	return nil
}

