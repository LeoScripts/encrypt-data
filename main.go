package main

import (
	"encrypt-data/util"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/encrypt", Encrypt).Methods("POST")

	log.Println("start http encrypt-data server.....")
	log.Fatal(http.ListenAndServe(":7777", router))
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	//------
	file, _, err := r.FormFile("file") //no formulario deve ter o campo file
	if err != nil {
		log.Printf("erro read file: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileName := uuid.New().String()
	tempFileName := uuid.New().String()

	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("erro read file: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("read file success")
	tempFile, err := os.CreateTemp("./", fileName)
	if err != nil {
		log.Printf("error create temp file: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	log.Println("temp file created")

	tempFile.Write(fileBytes)
	//------
	secretKey := []byte("1234567890asdfgh")                         // deve vim de uma variavel de ambiente
	util.DecryptLargeFile(tempFile.Name(), tempFileName, secretKey) //ecriptografa
	returnFileBytes, err := os.ReadFile(tempFileName)               //ler o arquivo do gravado no disco
	if err != nil {
		log.Printf("error create temp file: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnFileBytes)
	defer os.Remove(tempFileName)
	defer os.Remove(tempFile.Name())

}
