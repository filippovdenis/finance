package main

import "os"

type Configuration struct{
	ListenTo string
	DBType string
	ConnectionString string
}

func LoadConfig() Configuration{
	var cfg Configuration

	cfg.ListenTo = os.Getenv("FIN_LISTEN_TO")
	cfg.DBType = os.Getenv("FIN_DBTYPE")
	cfg.ConnectionString = os.Getenv("FIN_CONNECTION_STRING")


	return cfg
}