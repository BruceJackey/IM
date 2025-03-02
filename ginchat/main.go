package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()

	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8080") // Run the server on port 8080
}
