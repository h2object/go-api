package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValide(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}
	auth := NewAdminAuth("liujianping","h2object.io")
	

	valid1 := Valid{
		Bucket: "/test",
		Field: "name",
		Valid: map[string]interface{}{
			"required": true,
		},
	}

	valid2 := Valid{
		Bucket: "/test",
		Field: "mailbox",
		Valid: map[string]interface{}{
			"email": true,
		},
	}

	assert.Nil(t, h2o.AddValid(log, auth, valid1))
	assert.Nil(t, h2o.AddValid(log, auth, valid2))

	var v3 Valid
	v3.Bucket = "/test"
	v3.Field = "mailbox"

	assert.Nil(t, h2o.GetValid(log, auth, &v3))
	log.Info("Get Valid:(%v)", v3.Valid)

	assert.Nil(t, h2o.DelValid(log, auth, v3))
}