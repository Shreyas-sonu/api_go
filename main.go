package main

import (
	"github.com/gin-contrib/cors"
)

func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000",}
	corsConfig.AllowCredentials = true

}
