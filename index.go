package api

import (
	"github.com/h2object/rpc"
	"path"
	"net/url"
)

//! --- Index ---
type Index struct{
	Bucket string
	Field  string
}

func (h2o *H2Object) CreateIndex(l Logger, auth Auth, index Index) error {
	params := url.Values{
		"field": {index.Field},
	}

	URL := rpc.BuildHttpURL(h2o.addr, path.Join(index.Bucket, ".index"), params)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	ret := map[string]interface{}{}
	if err := h2o.conn.PostJson(l, URL, nil, &ret); err != nil {
		return err
	}
	if l != nil {
		l.Debug("ret: %v", ret)
	}
	return nil
}

func (h2o *H2Object) DeleteIndex(l Logger, auth Auth, index Index) error {
	params := url.Values{
		"field": {index.Field},
	}

	URL := rpc.BuildHttpURL(h2o.addr, path.Join(index.Bucket, ".index"), params)
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
