package post

import (
	"encoding/json"
	"github.com/claudioontheweb/go-blog/db"
	customHTTP "github.com/claudioontheweb/go-blog/http"
	"github.com/gorilla/mux"
	"net/http"
)

/*

##################################
Posts

*/

// Get All Posts
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	db.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

// Get Post By Id
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post Post
	if db.DB.Where("id = ?", params["postId"]).First(&post).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Post with ID: " + params["postId"])
	} else {
		json.NewEncoder(w).Encode(post)
	}
}

// Create Post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	err := db.DB.Create(&post).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(&post)
	customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully created new Post")
}


// Delete Post
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post Post

	if db.DB.Where("id = ?", params["postId"]).First(&post).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Post with ID: " + params["postId"])
	} else {
		db.DB.Delete(&post, params["postId"])
		customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully deleted Post with ID: " + params["postId"])
	}
}

// Get All Posts from Author
func GetPostsByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var posts []Post

	if db.DB.Where("author_id = ?", params["authorId"]).Find(&posts).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Posts from Author with ID: " + params["authorId"])
	} else {
		json.NewEncoder(w).Encode(posts)
	}
}


/*

##################################
Tags

*/

// Get All Tags
func GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tags []Tag
	db.DB.Find(&tags)
	json.NewEncoder(w).Encode(tags)
}

// Get Tag By Id
func GetTagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tag Tag
	if db.DB.Where("id = ?", params["tagId"]).First(&tag).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Tag with ID: " + params["tagId"])
	} else {
		json.NewEncoder(w).Encode(tag)
	}
}

// Get All Tags of specific Post
func GetTagsOfPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var post Post
	var tags []Tag

	if db.DB.Where("id = ?", params["postId"]).Find(&post).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Tags found of Post with ID: " + params["postId"])
	} else {

		db.DB.Model(&post).Related(&tags, "Tags")

		db.DB.Preload("Tags").First(&post)

		json.NewEncoder(w).Encode(tags)

	}
}

// Create Tag
func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var post Post
	var tag Tag

	json.NewDecoder(r.Body).Decode(&tag)


	if db.DB.Where("id = ?", params["postId"]).Find(&post).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Post found with ID: " + params["postId"])
	} else {
		err := db.DB.Model(&post).Association("Tags").Append(&tag).Error
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(tag)

		// customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully created new Tag")
	}
}

// Delete Tag
func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tag Tag

	if db.DB.Where("id = ?", params["tagId"]).First(&tag).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Tag with ID: " + params["tagId"])
	} else {
		db.DB.Delete(&tag, params["tagId"])
		customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully deleted Tag with ID: " + params["tagId"])
	}
}

/*

##############################
Likes

 */
// Increment Like of Post
func IncrementLikeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var post Post
	var like Like
	json.NewDecoder(r.Body).Decode(&like)

	if db.DB.Where("id = ?", params["postId"]).First(&post).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Post with ID: " + params["postId"])
	} else {

		err := db.DB.Model(&post).Association("Likes").Append(&like).Error
		if err != nil {
			panic(err)
		}
		customHTTP.NewSuccessResponse(w, http.StatusOK, "Successfully increased likes of Post with ID: " + params["postId"])
	}
}

func GetLikesOfPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var post Post
	var likes []Like

	if db.DB.Where("id = ?", params["postId"]).Find(&post).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Likes found of Post with ID: " + params["postId"])
	} else {

		db.DB.Model(&post).Related(&likes, "Likes")

		db.DB.Preload("Likes").First(&post)

		json.NewEncoder(w).Encode(likes)

	}
}


func GetLikesOfAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var likes []Like

	if db.DB.Where("author_id = ?", params["authorId"]).Find(&likes).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No Likes found of Author with ID: " + params["authorId"])
	} else {

		json.NewEncoder(w).Encode(likes)

	}
}