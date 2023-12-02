package api

import (
	"github.com/gin-gonic/gin"

	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/storage"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {

	handler := handler.NewHandler(cfg, strg)

	// Country ...
	r.POST("/country", handler.CreateCountry)
	r.GET("/country/:id", handler.GetByIDCountry)
	r.GET("/country", handler.GetListCountry)
	r.PUT("/country/:id", handler.UpdateCountry)
	r.DELETE("/country/:id", handler.DeleteCountry)

	//City
	r.POST("/city", handler.CreateCity)
	r.GET("/city/:id", handler.GetByIDCity)
	r.GET("/city", handler.GetListCity)
	r.PUT("/city/:id", handler.UpdateCity)
	r.DELETE("/city/:id", handler.DeleteCity)

	//Aeroport
	r.POST("/aeroport", handler.CreateAeroport)
	r.GET("/aeroport/:id", handler.GetByIDAeroport)
	r.GET("/aeroport", handler.GetListAeroport)
	r.PUT("/aeroport/:id", handler.UpdateAeroport)
	r.DELETE("/aeroport/:id", handler.DeleteAeroport)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/upload", handler.Upload)
}
