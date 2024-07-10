package api_handler

import (
	"github.com/Devisree146/Go_project-library.git/multicache"
	"github.com/gin-gonic/gin"
)

func SetupMultiCacheRouter() *gin.Engine {
	router := multicache.SetupRouter()
	return router
}
