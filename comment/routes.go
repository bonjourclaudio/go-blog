package comment

import (
	"github.com/claudioontheweb/go-blog/auth"
	"github.com/claudioontheweb/go-blog/router"
)

var Routes = router.RoutePrefix{
	"/comments",
	auth.JwtVerify,
	[]router.Route{
		{
			"GetComments",
			"GET",
			"",
			GetCommentsHandler,
		},
		{
			"GetComment",
			"GET",
			"/{commentId}",
			GetCommentHandler,
		},
		{
			"CreateComment",
			"POST",
			"",
			CreateCommentHandler,
		},
		{
			"DeleteComment",
			"DELETE",
			"/{commentId}",
			DeleteCommentHandler,
		},
		{
			"GetCommentsOfPost",
			"GET",
			"/post/{postId}",
			GetCommentsOfPost,
		},
		{
			"GetCommentsOfUser",
			"GET",
			"/user/{userId}",
			GetCommentsOfUser,
		},
	},
}
