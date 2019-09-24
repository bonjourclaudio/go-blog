package comment

import (
	"github.com/claudioontheweb/go-blog/router"
)

var Routes = router.RoutePrefix{
	"/comments",
	[]router.Route{
		router.Route{
			"GetComments",
			"GET",
			"",
			GetCommentsHandler,
		},
		router.Route{
			"GetComment",
			"GET",
			"/{commentId}",
			GetCommentHandler,
		},
		router.Route{
			"CreateComment",
			"POST",
			"",
			CreateCommentHandler,
		},
		router.Route{
			"DeleteComment",
			"DELETE",
			"/{commentId}",
			DeleteCommentHandler,
		},
		router.Route{
			"GetCommentsOfPost",
			"GET",
			"/post/{postId}",
			GetCommentsOfPost,
		},
		router.Route{
			"GetCommentsOfAuthor",
			"GET",
			"/author/{authorId}",
			GetCommentsOfAuthor,
		},
	},
}
