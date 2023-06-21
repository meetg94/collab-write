package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
)

const (
	documentFile = "document.json"
)

type Document struct {
	Content string `json:"content"`
}

func main() {
	// Create a router
	r := mux.NewRouter()

	r.HandleFunc("/document", getDocumentHandler).Methods("GET")
	r.HandleFunc("/edit", editDocumentHandler).Methods("POST")

	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getDocumentHandler(w http.ResponseWriter, r *http.Request) {
	//Read the document
	content, err := readDocument()
	if err != nil {
		log.Println("Error reading document:", err)
		http.Error(w, "Failed to read document", http.StatusInternalServerError)
		return
	}
	//Send doc as response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
}

func editDocumentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming edit request from the user editing the document
	// Apply edit to the document
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update the document content with the new content
	newContent := string(body)
	err = writeDocument(newContent)
	if err != nil {
		log.Println("Error writing document: ", err)
		http.Error(w, "Failed to write document", http.StatusInternalServerError)
		return
	}
	//Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Document updated successfully"))
}

func readDocument() (string, error) {
	//Read content
	fileContent, err := ioutil.ReadFile(documentFile)
	if err != nil {
		return "", err
	}

	var doc Document
	err = json.Unmarshal(fileContent, &doc)
	if err != nil {
		return "", err
	}

	return doc.Content, nil
}

func writeDocument(content string) error {
	//Create a document
	doc := Document {
		Content: content,
	}

	//Marshal the document into JSON format
	fileContent, err := json.MarshalIndent(doc, "", " ")
	if err != nil {
		return err
	}

	//Write the JSON content to the document
	err = ioutil.WriteFile(documentFile, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}