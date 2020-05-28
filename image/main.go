package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// FileUpload ...
func FileUpload(r *http.Request) (string, error) {

	mf, fh, err := r.FormFile("catimg")
	if err != nil {
		return "", err
	}
	defer mf.Close()

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := filepath.Join(wd, "view", "assets", "img", fh.Filename)
	image, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer image.Close()
	io.Copy(image, mf)
	return fh.Filename, nil

}
