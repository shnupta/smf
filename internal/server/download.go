package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// DownloadHandler checks the file is present, and then generates a unique
// id for the file, returning the download link in the response.
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Download handler")

	vars := mux.Vars(r)
	if _, ok := vars["id"]; !ok {
		w.Write([]byte("Something went wrong"))
		return
	}

	if vars["id"] == "" {
		return
	}

	id := vars["id"]

	// TODO: Do all sorts of checks and database stuff
	http.ServeFile(w, r, "files/"+id)
}
