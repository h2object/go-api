package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}
	auth := NewAdminAuth("liujianping","h2object.io")
	
	var app Application
	assert.Nil(t, h2o.GetApplication(log, auth, &app))
	log.Info("Get Application: (%v)", app)
	
	app.Admin = true
	assert.Nil(t, h2o.SetApplication(log, auth, &app))
	
	var service Service
	assert.Nil(t, h2o.GetService(log, auth, &service))
	log.Info("Get Service: (%v)", service)
	
	service.Token = false
	assert.Nil(t, h2o.SetService(log, auth, &service))

	var session Session
	assert.Nil(t, h2o.GetSession(log, auth, &session))
	log.Info("Get Session: (%v)", session)
	
	session.Expire = "20s"
	assert.Nil(t, h2o.SetSession(log, auth, &session))

	var wechat AuthWechat
	assert.Nil(t, h2o.GetAuthWechat(log, auth, &wechat))
	log.Info("Get Wechat: (%v)", wechat)
	
	wechat.Secret = "wechat secret"
	assert.Nil(t, h2o.SetAuthWechat(log, auth, &wechat))

}