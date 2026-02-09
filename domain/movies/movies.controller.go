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
// @security Bearer
// @router   /api/v1/movies/search [get]
// @summary  get list of movies based on search query
// @accept   json
// @produce  json
// @Param    query query string true "query on movie name"
// @success  200 {object} core.Response[MoviesResponseType]
func (ctrl *MoviesController) Search(c *gin.Context) {
	search := c.DefaultQuery("query", "2026")
	movies, _ := ctrl.service.Search(search)
	ctrl.root.Success(c, MoviesResponse(movies))
}

// @tags     Movies
// @security Bearer
// @router   /api/v1/movies [get]
// @summary  get list of movies
// @accept   json
// @produce  json
// @Param   filter query string false "filter list movie" enums(created, imdb, year)
// @Param    page query string false "pagination page_value, default 0"
// @success  200 {object} core.Response[MoviesResponseType]
func (ctrl *MoviesController) FindAll(c *gin.Context) {
	search, filter, page := core.PaginateQueries(c)
	movies, _ := ctrl.service.List(search, filter, page)
	ctrl.root.Success(c, MoviesResponse(movies))
}
