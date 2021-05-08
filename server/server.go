package server

import "go-run-python/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
