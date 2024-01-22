package main

import (
	"github.com/go-chi/chi"
	gocache "github.com/patrickmn/go-cache"
	"log"
	"time"
	"trafilea/cmd/api"
	"trafilea/pkg/repository"
	"trafilea/pkg/service"

	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error running application. %v", err)
	}
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

func run() error {
	log.Println("starting...")

	cache := gocache.New(defaultExpiration, purgeTime)
	cacheData := repository.CacheData{Collection: cache}

	readCacheRepository := repository.NewReadCacheRepository(cacheData)
	updateCacheRepository := repository.NewUpdateCacheRepository(cacheData)

	divisibleService := service.NewDivisibleService()
	valueNumberService := service.NewValueNumberService()
	createNumberService := service.NewCreateNumberService(readCacheRepository, updateCacheRepository)
	getCollectionService := service.NewGetCollectionService(readCacheRepository)

	divisibleHandler := api.NewDivisibleHandler(&divisibleService).Handle
	valueNumberHandler := api.NewValueNumberHandler(&valueNumberService).Handle
	createNumberHandler := api.NewCreateNumberHandler(&createNumberService).Handle
	getCollectionHandler := api.NewGetCollectionHandler(&getCollectionService).Handle

	app := api.AppHandler{
		DivisibleHandler:     divisibleHandler,
		NumberValueHandler:   valueNumberHandler,
		CreateValueHandler:   createNumberHandler,
		GetCollectionHandler: getCollectionHandler,
	}

	router := chi.NewRouter()
	app.InitializeRoutes(router)

	log.Println("listening...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return err
	}

	return nil
}
