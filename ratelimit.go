package api

import (
	"github.com/h2object/rpc"
	"path"
)

type RateLimit struct{
	URI 	 string 	`json:"-"`
	Seconds  int64		`json:"seconds"`
	Limit    int64 		`json:"limit"`
}

func (h2o *H2Object) SetRateLimit(l Logger, auth Auth,ratelimit *RateLimit) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(ratelimit.URI, ".ratelimit"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	ret := map[string]interface{}{}
	if err := h2o.conn.PostJson(l, URL, ratelimit, &ret); err != nil {
		return err
	}
	if l != nil {
		l.Debug("ret: %v", ret)
	}
	return nil
}

func (h2o *H2Object) DelRateLimit(l Logger, auth Auth,ratelimit *RateLimit) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(ratelimit.URI, ".ratelimit"), nil)
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

func (h2o *H2Object) GetRateLimit(l Logger, auth Auth,ratelimit *RateLimit) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(ratelimit.URI, ".ratelimit"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, ratelimit); err != nil {
		return err
	}
	return nil	
}
