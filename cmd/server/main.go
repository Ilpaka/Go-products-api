// @title Products API KR_17_03
// @version 1.0
// @description Simple CRUD API for Products with JWT auth.
// @host localhost:8090
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {token}" to authorize.
package main

import (
	"makarov_project/internal/config"
	"makarov_project/internal/handler"
	"makarov_project/internal/repository"
	"makarov_project/internal/service"

	_ "makarov_project/docs"
)

func main() {
	cfg := config.Load()
	repo := repository.NewProductRepo()
	svc := service.NewProductService(repo)

	r := handler.NewRouter(cfg, svc)
	r.Run(":8090")
}
