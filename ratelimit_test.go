package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	h2o := NewH2Object("127.0.0.1", 9000)
	log := ConsoleLogger{}
	auth := NewAdminAuth("liujianping","h2object.io")
	
	rate1 := &RateLimit{
		URI: "/test",
		Seconds: 10,
		Limit: 100,
	}
	assert.Nil(t, h2o.DelRateLimit(log, auth, rate1))
	assert.Nil(t, h2o.SetRateLimit(log, auth, rate1))

	var rate2 RateLimit
	rate2.URI = "/test"
	assert.Nil(t, h2o.GetRateLimit(log, auth, &rate2))	

	log.Info("ratelimit get: %v", rate2)	
}