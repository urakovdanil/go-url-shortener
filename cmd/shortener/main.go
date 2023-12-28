package main

import (
	"fmt"
	"github.com/urakovdanil/go-url-shortener/internal/app/handlers"
	"github.com/urakovdanil/go-url-shortener/internal/storage"
	"github.com/urakovdanil/go-url-shortener/internal/storage/common"
	"github.com/urakovdanil/go-url-shortener/internal/storage/inmem"
	"net/http"
	"os"
)

func main() {
	im := common.Storage(inmem.New())
	storage.SetUsed(&im)

	http.HandleFunc(handlers.ShortenPath, handlers.HandlerFunc)
	if err := http.ListenAndServe(":8080", http.HandlerFunc(handlers.HandlerFunc)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
