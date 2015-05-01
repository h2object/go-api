package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestObject(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}

	root := &Object{
		Bucket: "/users",
	}

	assert.Nil(t, h2o.DeleteObject(log, nil, root))

	james := &Object{
		Bucket: "/users",
		Key:"james",
		Value: map[string]interface{}{
			"name": "liujianping",
		},
	}
	assert.Nil(t, h2o.PostObject(log, nil, james))

	james.Value = map[string]interface{}{
		"age": 34,
	}
	assert.Nil(t, h2o.PatchObject(log, nil, james))

	james.Value = map[string]interface{}{
		"name": "james",
		"age": 35,
	}
	assert.Nil(t, h2o.PutObject(log, nil, james))

	get := &Object{
		Bucket: "/users",
		Key:"james",
	}
	assert.Nil(t, h2o.GetObject(log, nil, get))
	log.Debug("Get: (%v)", get)

	unknown := &Object{
		Bucket: "/users",
		Value: map[string]interface{}{
			"name": "unknown",
		},
	}
	assert.Nil(t, h2o.PostObject(log, nil, unknown))
	log.Debug("unknown: (%s, %s) => %v", unknown.Bucket, unknown.Key, unknown.Value)
	unknown.Value = map[string]interface{}{
		"age": 34,
	}
	assert.Nil(t, h2o.PatchObject(log, nil, unknown))

	unknown.Value = map[string]interface{}{
		"name": "unknown",
		"age": 35,
	}
	assert.Nil(t, h2o.PutObject(log, nil, unknown))

	size, e := h2o.Size(log, nil, root)
	assert.Nil(t, e)
	log.Debug("/users size: %d", size)	

	keys, e := h2o.Keys(log, nil, root)
	assert.Nil(t, e)
	log.Debug("/users keys: %v", keys)	

	total, e := h2o.Total(log, nil, root)
	assert.Nil(t, e)
	log.Debug("/users total: %v", total)	
}