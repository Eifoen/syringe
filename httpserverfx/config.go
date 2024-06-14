package httpserverfx

import (
	"net/http"
)

var defaultConfig = &ServerConfig{
	Server: &http.Server{},
}

// ServerConfig contains the entire configuration used by the httpserverfx module
// it contains basically every field exported by http.Server, except for the Handler,
// whiich should be provided via fx itself.
type ServerConfig struct {
	*http.Server
}
