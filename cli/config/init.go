package config

// initialise environment variables for parsing configuration
func init() {
	EnvRegister(&Config.Host, "FAIRSHARE_HOST")
	EnvRegister(&Config.Port, "FAIRSHARE_PORT")
	EnvRegister(&Config.Log, "FAIRSHARE_LOG")
}

// initialise command line flags for parsing configuration
func init() {
	cmd.Var(&Config.Host, "host", "host to listen on [env: FAIRSHARE_HOST]")
	cmd.Var(&Config.Port, "port", "port to listen on [env: FAIRSHARE_PORT]")
	cmd.Var(&Config.Log, "log", "log level above which records get logged [env: FAIRSHARE_LOG]")
}
