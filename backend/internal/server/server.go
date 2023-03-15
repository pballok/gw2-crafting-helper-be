package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pballok/gw2-crafting-helper/backend/internal/item"
)

func StartServer() {
	router := createRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Failed to start up server : %v", err)
	}

	fmt.Print("Listening...\n")
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId, err := strconv.Atoi(vars["itemId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := item.NewItem(itemId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if item.Id != itemId {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reply, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(reply))
}

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/items/{itemId}", itemsHandler).Methods("GET")

	return router
}
