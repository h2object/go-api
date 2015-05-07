package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFileUpload(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}
	auth := NewAdminAuth("liujianping","h2object.io")
	
	assert.Nil(t, h2o.upload_file(log, auth, "", "/hello.html", "testdata/hello.html"))

}