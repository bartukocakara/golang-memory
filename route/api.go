package route

import (
	"golang-memory/handle"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	modelHandler := handle.NewModelHandler()

	r := gin.Default()

	apiRoutes := r.Group("/api")

	{
		apiRoutes.GET("/", modelHandler.GetAll)
	}

	return r.Run(address)

}