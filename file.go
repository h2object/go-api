package api

import (
	"io"
	"os"
	"path"
	"net/url"
	"github.com/h2object/rpc"
)

func (h2o *H2Object) upload_file(l Logger, auth Auth, provider string, dest_uri string, file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, filename := path.Split(file)
	contentType := ContentTypeByFilename(filename)

	return h2o.upload(l, auth, provider, dest_uri, contentType, fd)
}

func (h2o *H2Object) upload(l Logger, auth Auth, provider string, dest_uri string, contentType string, rd io.Reader) error {
	params := url.Values{
		"provider": {provider},
	}
	URL := rpc.BuildHttpURL(h2o.addr, dest_uri, params)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Put(l, URL, contentType, rd, 0, nil); err != nil {
		return err
	}
	return nil
}