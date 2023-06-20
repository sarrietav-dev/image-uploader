package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

const MAX_FILE_SIZE = 10 * 1024 * 1024 // 10MB
const UPLOAD_FOLDER = "uploads"

type UploadFileResponse struct {
	ImgId string `json:"imgId"`
}

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

	id := fmt.Sprintf("%s%s", uuid.New(), filepath.Ext(fileHeader.Filename))

	dst, err := os.Create(fmt.Sprintf("./%s/%s", UPLOAD_FOLDER, id))
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

	response, err := json.Marshal(UploadFileResponse{ImgId: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
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

	img, err := os.ReadFile(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(img)
}

func main() {
	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)

	r.Post("/api/image", UploadFile)
	r.Get("/api/image/{id}", GetImage)

	log.Printf("Server up and running localhost:%s", "8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
