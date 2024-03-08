package cli

import "flag"

// Config is the configuration options passed to the fairshare server.
// Means of conveying configuration may be:
//
//   - Command line flags
var Config = struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
	Log  string `json:"log,omitempty"`
}{
	Port: "8080",
	Log:  "INFO",
}

// initialize command line flags for parsing configuration
func init() {
	flag.StringVar(&Config.Host, "host", Config.Host, "host to listen on")
	flag.StringVar(&Config.Port, "port", Config.Port, "port to listen on")
	flag.StringVar(&Config.Log, "log", Config.Log, "log level above which records get logged")
}
