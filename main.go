package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const MAX_FILE_SIZE = 10 * 1024 * 1024 // 10MB
const UPLOAD_FOLDER = "uploads"

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_FILE_SIZE)

	if err := r.ParseMultipartForm(MAX_FILE_SIZE); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = os.MkdirAll(fmt.Sprintf("./%s", UPLOAD_FOLDER), os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(fmt.Sprintf("./%s/%s%s", UPLOAD_FOLDER, uuid.New(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	file := fmt.Sprintf("./uploads/%s", id)

	_, err := os.Stat(file)

	if errors.Is(err, os.ErrNotExist) {
		http.Error(w, "Image doesn't exist", http.StatusNotFound)
		return
	}

	img, err := ioutil.ReadFile(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(img)
}

func main() {
	r := chi.NewRouter()

	r.Post("/api/image", UploadFile)
	r.Get("/api/image/{id}", GetImage)

	http.ListenAndServe(":8080", r)
}
