package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"internal/document"
	"project-root/internal/server"
	"project-root/internal/server/websocket"
)

func main() {
	// Initialize document
	doc := document.NewDocument("Initial document content")

	// Create WebSocket hub
	hub := websocket.NewHub()
	go hub.Run()

	// Create server
	srv := server.NewServer(doc, hub)

	// Create router
	r := mux.NewRouter()
	r.HandleFunc("/document", srv.GetDocumentHandler).Methods("GET")
	r.HandleFunc("/edit", srv.EditDocumentHandler).Methods("POST")

	// Serve static files (if applicable)
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
