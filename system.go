package api

import (
	"github.com/h2object/rpc"
	"path"
)

//! --- System ---
func (h2o *H2Object) GetSystem(l Logger, auth Auth, obj *Object) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join(obj.Bucket, obj.Key+ ".system"), nil)

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