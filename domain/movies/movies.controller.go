package movies

import (
	"app/core"

	"github.com/gin-gonic/gin"
)

type MoviesController struct {
	root    *core.Controller
	service MoviesServiceInterface
}

func NewMoviesController() *MoviesController {
	return &MoviesController{
		root:    core.GetController(),
		service: NewMoviesService(),
	}
}

// @tags     Movies
// @router   /api/v1/movies/search [get]
// @summary  get list of movies based on search query
// @accept   json
// @produce  json
// @Param    query query string true "query on movie name"
// @success  200 {object} MoviesEnvelope
func (ctrl *MoviesController) Search(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		ctrl.root.BadRequestError(c, map[string]string{
			"query": "The query is required",
		})
		return
	}

	movies, err := ctrl.service.Search(c.Request.Context(), query)
	if err != nil {
		ctrl.root.UpstreamError(c, err)
		return
	}

	ctrl.root.Success(c, MoviesResponse(movies))
}

// @tags     Movies
// @router   /api/v1/movies [get]
// @summary  get list of movies
// @accept   json
// @produce  json
// @Param   filter query string false "filter list movie" enums(created, imdb, year)
// @Param    page query string false "pagination page_value, default 0"
// @success  200 {object} MoviesEnvelope
func (ctrl *MoviesController) FindAll(c *gin.Context) {
	filter, page := core.PaginateQueries(c)

	movies, err := ctrl.service.List(c.Request.Context(), filter, page)
	if err != nil {
		ctrl.root.UpstreamError(c, err)
		return
	}

	ctrl.root.Success(c, MoviesResponse(movies))
}
