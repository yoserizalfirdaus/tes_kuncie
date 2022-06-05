package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoserizalfirdaus/tes_kuncie/adapter/httphandler"
	"github.com/yoserizalfirdaus/tes_kuncie/adapter/repository"
	"github.com/yoserizalfirdaus/tes_kuncie/infrastructure/postgres"
	"github.com/yoserizalfirdaus/tes_kuncie/usecase/checkout"
)

func main() {
	fmt.Println("starting")

	db, err := postgres.Connect("database=mydb user=myuser password=mypassword host=postgres sslmode=disable")
	if err != nil {
		log.Fatal("error connect postgre", err)
	}

	//init repo
	promoRepo, _ := repository.NewPromoRepo(db)
	productRepo, _ := repository.NewProductRepo(db)
	cartRepo, _ := repository.NewCartRepository()
	//init usecase
	checkoutUsecase := checkout.NewCheckoutUsecase(promoRepo, productRepo, cartRepo)
	//init handler
	httpHandler := httphandler.NewHttpHandler(checkoutUsecase)

	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/cart/add-to-cart").HandlerFunc(httpHandler.AddToCart)
	router.Methods(http.MethodPost).Path("/cart/checkout").HandlerFunc(httpHandler.Checkout)

	server := &http.Server{
		ReadTimeout:  0,
		WriteTimeout: 0,
		Addr:         ":9876",
		Handler:      router,
	}
	server.ListenAndServe()
}
