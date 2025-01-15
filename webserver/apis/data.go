package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nickemma/cryptomaster_and_web_server/model"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(model.GetAll())
}

func HandleGetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Extract the ID from the URL path
	idStr := r.URL.Path[len("/api/data/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Fetch the data and return it
	result := model.GetById(id)
	if result.Id == 0 {
		http.Error(w, "Exhibition not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var item model.Exhibition
		decoder := json.NewDecoder(r.Body)

		// Disallow unknown fields in the JSON
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&item)
		if err != nil {
			http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
			return
		}

		model.AddData(item)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
