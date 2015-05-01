package api

import (
	"github.com/h2object/rpc"
	"path"
	"net/url"
	"errors"
)

type Object struct{
	Bucket string
	Key string
	Value interface{}
}

//! --- Object ---
func (h2o *H2Object) PostObject(l Logger, auth Auth, obj *Object) error {
	if obj == nil {
		return errors.New("object is nil")
	}

	var params url.Values = nil
	if obj.Key != "" {
		params = url.Values{
			"key": {obj.Key},
		}
	}

	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, ".json"), params)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	ret := map[string]interface{}{}
	if err := h2o.conn.PostJson(l, URL, obj.Value, &ret); err != nil {
		return err
	}
	
	obj.Key = ret["key"].(string)
	if l != nil {
		l.Debug("obj: %v", obj)
	}
	return nil	
}

func (h2o *H2Object) PutObject(l Logger, auth Auth, obj *Object) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".json"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, obj.Value, nil); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) PatchObject(l Logger, auth Auth, obj *Object) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".json"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PatchJson(l, URL, obj.Value, nil); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) DeleteObject(l Logger, auth Auth, obj *Object) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".json"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Delete(l, URL, nil); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) GetObject(l Logger, auth Auth, obj *Object) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".json"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	var ret interface{}
	if err := h2o.conn.Get(l, URL, &ret); err != nil {
		return err
	}

	obj.Value = ret
	return nil
}

func (h2o *H2Object) Size(l Logger, auth Auth, obj *Object) (int64, error) {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".size"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	ret := map[string]interface{}{}
	if err := h2o.conn.Get(l, URL, &ret); err != nil {
		return 0, err
	}

	size := ret["size"].(float64)
	return int64(size), nil
}

func (h2o *H2Object) Keys(l Logger, auth Auth, obj *Object) (interface{}, error) {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".keys"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	var ret interface{}
	if err := h2o.conn.Get(l, URL, &ret); err != nil {
		return 0, err
	}
	return ret, nil
}

func (h2o *H2Object) Total(l Logger, auth Auth, obj *Object) (interface{}, error) {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".total"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	var ret interface{}
	if err := h2o.conn.Get(l, URL, &ret); err != nil {
		return 0, err
	}
	return ret, nil
}