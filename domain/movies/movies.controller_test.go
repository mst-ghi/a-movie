package movies

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/core"

	"github.com/gin-gonic/gin"
)

type fakeService struct {
	movies []Movie
	err    error
}

func (f *fakeService) Search(ctx context.Context, search string) ([]Movie, error) {
	return f.movies, f.err
}

func (f *fakeService) List(ctx context.Context, filter string, page int) ([]Movie, error) {
	return f.movies, f.err
}

func newTestRouter(service MoviesServiceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)

	ctrl := &MoviesController{
		root:    core.GetController(),
		service: service,
	}

	router := gin.New()
	router.GET("/api/v1/movies", ctrl.FindAll)
	router.GET("/api/v1/movies/search", ctrl.Search)

	return router
}

func TestSearch_RequiresQuery(t *testing.T) {
	router := newTestRouter(&fakeService{})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies/search", nil)
	router.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for missing query, got %d: %s", res.Code, res.Body.String())
	}
}

func TestSearch_UpstreamErrorBecomes502(t *testing.T) {
	router := newTestRouter(&fakeService{err: context.DeadlineExceeded})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies/search?query=test", nil)
	router.ServeHTTP(res, req)

	if res.Code != http.StatusBadGateway {
		t.Fatalf("expected 502 for upstream failure, got %d: %s", res.Code, res.Body.String())
	}
}

func TestFindAll_ReturnsEnvelope(t *testing.T) {
	router := newTestRouter(&fakeService{movies: []Movie{{ID: 7, Title: "Blade Runner"}}})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies?filter=imdb&page=1", nil)
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", res.Code, res.Body.String())
	}

	var body struct {
		Message string `json:"message"`
		Errors  map[string]string
		Data    struct {
			Movies []Movie `json:"movies"`
		} `json:"data"`
	}
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(body.Data.Movies) != 1 || body.Data.Movies[0].Title != "Blade Runner" {
		t.Fatalf("unexpected movies in response: %s", res.Body.String())
	}
}

func TestFindAll_UpstreamErrorBecomes502(t *testing.T) {
	router := newTestRouter(&fakeService{err: context.DeadlineExceeded})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies", nil)
	router.ServeHTTP(res, req)

	if res.Code != http.StatusBadGateway {
		t.Fatalf("expected 502 for upstream failure, got %d: %s", res.Code, res.Body.String())
	}
}
