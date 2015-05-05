package api

import (
	"github.com/h2object/rpc"
	"net/url"
	"path"
)

//! --- Security ---
type Valid struct{
	Bucket string
	Field  string
	Valid  map[string]interface{}
}

//! valid definitions
/*
	valid["required"] = true
	valid["email"] = true
	valid["min"] = 8
	valid["max"] = 88
	valid["minsize"] = 2
	valid["maxsize"] = 4
	valid["length"] = 6
	valid["match"] = ""
*/

func (h2o *H2Object) AddValid(l Logger, auth Auth, valid Valid) error {
	params := url.Values{
		"field": {valid.Field},
	}

	URL := rpc.BuildHttpURL(h2o.addr, valid.Bucket + ".valid", params)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PostJson(l, URL, valid.Valid, nil); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) DelValid(l Logger, auth Auth, valid Valid) error {
	params := url.Values{
		"field": {valid.Field},
	}

	URL := rpc.BuildHttpURL(h2o.addr, valid.Bucket + ".valid", params)
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

func (h2o *H2Object) GetValid(l Logger, auth Auth, valid *Valid) error {
	URL := rpc.BuildHttpURL(h2o.addr,path.Join("/valids", path.Join(valid.Bucket, valid.Field + ".vld.system")), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	var ret map[string]interface{}
	if err := h2o.conn.Get(l, URL, &ret); err != nil {
		return err
	}
	valid.Valid = ret
	return nil
}
