package author

import (
	"github.com/claudioontheweb/go-blog/router"
)

var Routes = router.RoutePrefix{
	"/authors",
	[]router.Route{
		router.Route{
			"GetAuthors",
			"GET",
			"",
			GetAuthorsHandler,
		},
		router.Route{
			"GetAuthor",
			"GET",
			"/{authorId}",
			GetAuthorHandler,
		},
		router.Route{
			"CreateAuthor",
			"POST",
			"",
			CreateAuthorHandler,
		},
		router.Route{
			"DeleteAuthor",
			"DELETE",
			"/{authorId}",
			DeleteAuthorHandler,
		},
	},
}