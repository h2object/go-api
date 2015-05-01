package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSystem(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}

	// auth := NewHeaderAuth()
	// auth.Add("X-H2O-AppId", "liujianping")
	// auth.Add("X-H2O-Secret", "h2object.io")
	auth := NewAdminAuth("liujianping","h2object.io")
	root := &Object{
		Bucket: "/",
	}
	assert.Nil(t, h2o.GetSystem(log, auth, root))
	log.Debug("Get: (%v)", root)
}