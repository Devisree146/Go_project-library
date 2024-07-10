package main

import (
	"github.com/Devisree146/Go_project-library.git/api_handler"
)

func main() {
	inMemoryRouter := api_handler.SetupInMemoryRouter()
	redisCacheRouter := api_handler.SetupRedisCacheRouter()
	multiCacheRouter := api_handler.SetupMultiCacheRouter()

	// Assuming you want to start each router on different ports
	go inMemoryRouter.Run(":8081")
	go redisCacheRouter.Run(":8082")
	multiCacheRouter.Run(":8080")
}
