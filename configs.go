package api

import (
	"github.com/robfig/config"
	"strings"
)

// This handles the parsing of app.conf
// It has a "preferred" section that is checked first for option queries.
// If the preferred section does not have the option, the DEFAULT section is
// checked fallback.
type MergedConfig struct {
	config  *config.Config
	section string // Check this section first, then fall back to DEFAULT
}

func NewEmptyConfig() *MergedConfig {
	return &MergedConfig{config.NewDefault(), ""}
}

func LoadConfig(file string) (*MergedConfig, error) {
	var err error
	conf, err := config.ReadDefault(file)
	if err != nil {
		return nil, err
	}
	return &MergedConfig{conf, ""}, nil
}

func (c *MergedConfig) Raw() *config.Config {
	return c.config
}

func (c *MergedConfig) SetSection(section string) {
	c.section = section
}

func (c *MergedConfig) SetOption(name, value string) {
	c.config.AddOption(c.section, name, value)
}

func (c *MergedConfig) Int(option string) (result int, found bool) {
	result, err := c.config.Int(c.section, option)
	if err == nil {
		return result, true
	}
	if _, ok := err.(config.OptionError); ok {
		return 0, false
	}

	// If it wasn't an OptionError, it must have failed to parse.
	return 0, false
}

func (c *MergedConfig) IntDefault(option string, dfault int) int {
	if r, found := c.Int(option); found {
		return r
	}
	return dfault
}

func (c *MergedConfig) Bool(option string) (result, found bool) {
	result, err := c.config.Bool(c.section, option)
	if err == nil {
		return result, true
	}
	if _, ok := err.(config.OptionError); ok {
		return false, false
	}

	// If it wasn't an OptionError, it must have failed to parse.
	return false, false
}

func (c *MergedConfig) BoolDefault(option string, dfault bool) bool {
	if r, found := c.Bool(option); found {
		return r
	}
	return dfault
}

func (c *MergedConfig) String(option string) (result string, found bool) {
	if r, err := c.config.String(c.section, option); err == nil {
		return stripQuotes(r), true
	}
	return "", false
}

func (c *MergedConfig) StringDefault(option, dfault string) string {
	if r, found := c.String(option); found {
		return r
	}
	return dfault
}

func (c *MergedConfig) HasSection(section string) bool {
	return c.config.HasSection(section)
}

// Options returns all configuration option keys.
// If a prefix is provided, then that is applied as a filter.
func (c *MergedConfig) Options(prefix string) []string {
	var options []string
	keys, _ := c.config.Options(c.section)
	for _, key := range keys {
		if strings.HasPrefix(key, prefix) {
			options = append(options, key)
		}
	}
	return options
}

// Helpers

func stripQuotes(s string) string {
	if s == "" {
		return s
	}

	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}

	return s
}



const DefaultFileContentType = "application/octet-stream"

var mimeConfig *MergedConfig

// Load mime-types.conf on init.
func LoadMimeConfig() {
	var err error
	mimeConfig, err = LoadConfig("mime-types.conf")
	if err != nil {
		panic(err)
	}
}

func init() {
	LoadMimeConfig()
}

func ContentTypeByFilename(filename string) string {
	dot := strings.LastIndex(filename, ".")
	if dot == -1 || dot+1 >= len(filename) {
		return DefaultFileContentType
	}

	extension := filename[dot+1:]
	contentType := mimeConfig.StringDefault(extension, "")
	if contentType == "" {
		return DefaultFileContentType
	}

	if strings.HasPrefix(contentType, "text/") {
		return contentType + "; charset=utf-8"
	}

	return contentType
}
