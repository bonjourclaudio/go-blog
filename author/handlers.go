package author

import (
	"encoding/json"
	"github.com/claudioontheweb/go-blog/db"
	customHTTP "github.com/claudioontheweb/go-blog/http"
	"github.com/gorilla/mux"
	"net/http"
)

// Get All Authors
func GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authors []Author
	db.DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

// Get Author By ID
func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var author Author
	if db.DB.Where("id = ?", params["authorId"]).First(&author).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Author with ID: " + params["authorId"])
	} else {
		json.NewEncoder(w).Encode(author)
	}
}

// Create Author
func CreateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author Author
	json.NewDecoder(r.Body).Decode(&author)
	err := db.DB.Create(&author).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(&author)
	customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully created new Author")
}

// Delete Author
func DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var author Author

	if db.DB.Where("id = ?", params["authorId"]).First(&author).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Author with ID: " + params["authorId"])
	} else {
		db.DB.Delete(&author, params["authorId"])
		customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully deleted Author with ID: " + params["authorId"])
	}

}