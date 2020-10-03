package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// UploadHandler checks for valid links and serves the relevant file.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload handler")

	// Max 100MB file upload
	r.ParseMultipartForm(100 << 20)

	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte("Error retrieving file\n" + err.Error()))
		return
	}

	defer file.Close()

	tempFile, err := ioutil.TempFile("files", "")
	if err != nil {
		w.Write([]byte("Could not create file"))
		return
	}
	defer tempFile.Close()

	// read contents of file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte("Error reading contents of file"))
		return
	}

	// write to the tempfile
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		w.Write([]byte("Error saving file to server"))
	}

	// Return the name of the file
	w.Write([]byte(strings.Split(tempFile.Name(), "/")[1]))
}
