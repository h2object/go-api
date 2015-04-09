package api

import (
	"net/url"
	"path"
	"github.com/h2object/httpRPC"
)

type Object struct{
	bucket  string
	key 	string
	value 	interface{}
}

func NewObject(bucket string, key string, val interface{}) *Object {
	return &Object{
		bucket: bucket,
		key: key,
		value: val,
	}	
}

func (obj *Object) BucketURI(suffix string) string {
	return path.Join(obj.bucket, obj.key + suffix)
}

func (obj *Object) URI(suffix string) string {
	return path.Join(obj.bucket, obj.key + suffix)
}

func (c *Client) PostObject(obj *Object) error {
	v := url.Values{}
	if obj.key != "" {
		v.Add("key", obj.key)
	}
	if c.isAdmin {
		v.Add("appid", ACCESS_KEY)
		v.Add("secret", SECRET_KEY)		
	}
	if c.isLogin && c.token != "" {
		v.Add("token", c.token)
	}
	u := rpc.BuildHttpURL(H2O_HOST, obj.BucketURI(".json"), v)

	var ret map[string]interface{}
	if err := c.PostJson(u, obj.value, &ret); err != nil {
		return err
	}
	obj.key = ret["created"].(string)
	return nil
}

func (c *Client) PatchObject(obj *Object) error {
	v := url.Values{}
	if c.isAdmin {
		v.Add("appid", ACCESS_KEY)
		v.Add("secret", SECRET_KEY)		
	}
	if c.isLogin && c.token != "" {
		v.Add("token", c.token)
	}
	u := rpc.BuildHttpURL(H2O_HOST, obj.URI(".json"), v)
	if err := c.PatchJson(u, obj.value, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) PutObject(obj *Object) error {
	v := url.Values{}
	if c.isAdmin {
		v.Add("appid", ACCESS_KEY)
		v.Add("secret", SECRET_KEY)		
	}
	if c.isLogin && c.token != "" {
		v.Add("token", c.token)
	}
	u := rpc.BuildHttpURL(H2O_HOST, obj.URI(".json"), v)
	if err := c.PutJson(u, obj.value, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteObject(obj *Object) error {
	v := url.Values{}
	if c.isAdmin {
		v.Add("appid", ACCESS_KEY)
		v.Add("secret", SECRET_KEY)		
	}
	if c.isLogin && c.token != "" {
		v.Add("token", c.token)
	}
	u := rpc.BuildHttpURL(H2O_HOST, obj.URI(".json"), v)
	if err := c.Delete(u, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) SizeOfObject(obj *Object, size *int64) error {
	v := url.Values{}
	if c.isAdmin {
		v.Add("appid", ACCESS_KEY)
		v.Add("secret", SECRET_KEY)		
	}
	if c.isLogin && c.token != "" {
		v.Add("token", c.token)
	}
	u := rpc.BuildHttpURL(H2O_HOST, obj.URI(".size"), v)
	
	var ret map[string]interface{}
	if err := c.Get(u, &ret); err != nil {
		return err
	}
	*size = ret["size"].(int64)
	return nil
}

func (c *Client) KeysOfObject(obj *Object, keys []string) error {
	v := url.Values{}
	if c.isAdmin {
		v.Add("appid", ACCESS_KEY)
		v.Add("secret", SECRET_KEY)		
	}
	if c.isLogin && c.token != "" {
		v.Add("token", c.token)
	}
	u := rpc.BuildHttpURL(H2O_HOST, obj.URI(".keys"), v)
	
	var ret map[string]interface{}
	if err := c.Get(u, &ret); err != nil {
		return err
	}
	
	return nil
}

func (c *Client) TotalObject(obj *Object, total *int64, objects *int64) error {
	return nil	
}

func (c *Client) GetObject(obj *Object) error {
	return nil
}

