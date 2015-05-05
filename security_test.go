package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSecurity(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}
	auth := NewAdminAuth("liujianping","h2object.io")
	
	sec1 := NewSignedSecurity("/test", "post", true)
	assert.Nil(t, h2o.DelSecurity(log, auth, sec1))
	assert.Nil(t, h2o.SetSecurity(log, auth, sec1))

	var sec2 Security
	sec2.URI = "/test"
	sec2.Method = "post"
	assert.Nil(t, h2o.GetSecurity(log, auth, &sec2))
	log.Debug("Get Security: (%v)", sec2)

	root := &Object{
		Bucket: "/securities/test",
	}
	assert.Nil(t, h2o.GetSystem(log, auth, root))
	log.Debug("Get: (%v)", root)
}