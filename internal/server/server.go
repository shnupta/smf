package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	fmt.Printf("Serving at port 8080\n")

	router := mux.NewRouter()

	router.HandleFunc("/upload", UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/{id}", DownloadHandler)

	http.ListenAndServe(":8080", router)
}
