package comment

import (
	"encoding/json"
	customHTTP "github.com/claudioontheweb/go-blog/http"
	"github.com/claudioontheweb/go-blog/db"
	"github.com/gorilla/mux"
	"net/http"
)

// Get All Comments
func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comments []Comment
	db.DB.Find(&comments)
	json.NewEncoder(w).Encode(comments)
}

// Get Comment By ID
func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var comment Comment
	if db.DB.Where("id = ?", params["commentId"]).First(&comment).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Comment with ID: " + params["commentId"])
	} else {
		json.NewEncoder(w).Encode(comment)
	}
}
// Create Comment
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comment Comment
	json.NewDecoder(r.Body).Decode(&comment)
	err := db.DB.Create(&comment).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(&comment)
	customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully created new Comment")
}

// Delete Comment
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var comment Comment

	if db.DB.Where("id = ?", params["commentId"]).First(&comment).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Comment with ID: " + params["commentId"])
	} else {
		db.DB.Delete(&comment, params["commentId"])
		customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully deleted Comment with ID: " + params["commentId"])
	}

}

// Get All Comments of specific Post
func GetCommentsOfPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var comments []Comment

	if db.DB.Where("post_id = ?", params["postId"]).Find(&comments).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Comments of Post with ID: " + params["postId"])
	} else {
		json.NewEncoder(w).Encode(comments)
	}
}

// Get All Comments of User
func GetCommentsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var comments []Comment

	if db.DB.Where("user_id = ?", params["userId"]).Find(&comments).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Comments of User with ID: " + params["userId"])
	} else {
		json.NewEncoder(w).Encode(comments)
	}
}


