package api

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUA(t *testing.T) {
	err := SetUser("")
	assert.Nil(t, err)

	err = SetUser("错误的UA")
	assert.NotNil(t, err)

	err = SetUser("Test0-_.")
	assert.Nil(t, err)
}

func TestFormat(t *testing.T) {
	str := "tesT0.-_"
	v := formatUserAgent(str)
	if !strings.Contains(v, str) {
		t.Fatal("should include user")
	}
	if !strings.HasPrefix(v, "H2OGO/"+version) {
		t.Fatal("invalid format")
	}
}
