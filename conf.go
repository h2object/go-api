package api

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"

	"github.com/h2object/httpRPC"
)
var H2O_HOST = "h2object.io"

var ACCESS_KEY string
var SECRET_KEY string

var version = "0.0.1"

var userPattern = regexp.MustCompile("^[a-zA-Z0-9_.-]*$")

// user should be [A-Za-z0-9]*
func SetUser(user string) error {
	if !userPattern.MatchString(user) {
		return errors.New("invalid user format")
	}
	rpc.UserAgent = formatUserAgent(user)
	return nil
}

func formatUserAgent(user string) string {
	return fmt.Sprintf("H2OGO/%s (%s; %s; %s) %s", version, runtime.GOOS, runtime.GOARCH, user, runtime.Version())
}

func init() {
	SetUser("")
}
