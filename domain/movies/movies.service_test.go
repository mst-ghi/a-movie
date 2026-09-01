package movies

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpstream(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	t.Setenv("MOVIE_API_URL", server.URL)
	t.Setenv("MOVIE_API_KEY", "test-key")

	return server
}

func TestMoviesService_Search(t *testing.T) {
	setUpstream(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/search/interstellar/test-key" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"posters": [{"id": 1, "title": "Interstellar", "year": 2014}]}`)
	})

	service := NewMoviesService()
	movies, err := service.Search(context.Background(), "interstellar")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(movies) != 1 || movies[0].Title != "Interstellar" {
		t.Fatalf("unexpected movies: %+v", movies)
	}
}

func TestMoviesService_Search_UpstreamError(t *testing.T) {
	setUpstream(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	service := NewMoviesService()
	movies, err := service.Search(context.Background(), "anything")
	if err == nil {
		t.Fatal("expected error for non-200 upstream response")
	}
	if movies != nil {
		t.Fatalf("expected nil movies, got %+v", movies)
	}
}

func TestMoviesService_Search_InvalidJSON(t *testing.T) {
	setUpstream(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "not json")
	})

	service := NewMoviesService()
	if _, err := service.Search(context.Background(), "anything"); err == nil {
		t.Fatal("expected error for invalid upstream JSON")
	}
}

func TestMoviesService_List_FilterWhitelist(t *testing.T) {
	var gotFilter string

	setUpstream(t, func(w http.ResponseWriter, r *http.Request) {
		gotFilter = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id": 2, "title": "Heat", "year": 1995}]`)
	})

	service := NewMoviesService()

	// An injected filter falls back to "created".
	if _, err := service.List(context.Background(), "../admin?x=1", 0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotFilter != "/api/movie/by/filtres/0/created/0/test-key" {
		t.Fatalf("injected filter not sanitized, upstream path: %s", gotFilter)
	}

	// A valid filter is passed through.
	if _, err := service.List(context.Background(), "imdb", 3); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotFilter != "/api/movie/by/filtres/0/imdb/3/test-key" {
		t.Fatalf("valid filter mangled, upstream path: %s", gotFilter)
	}
}

func TestMoviesService_List_Success(t *testing.T) {
	setUpstream(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id": 2, "title": "Heat"}, {"id": 3, "title": "Alien"}]`)
	})

	service := NewMoviesService()
	movies, err := service.List(context.Background(), "year", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(movies) != 2 || movies[1].Title != "Alien" {
		t.Fatalf("unexpected movies: %+v", movies)
	}
}

func TestMoviesResponse_NilBecomesEmptyArray(t *testing.T) {
	payload, err := json.Marshal(MoviesResponse(nil))
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	if string(payload) != `{"movies":[]}` {
		t.Fatalf("expected empty JSON array, got %s", payload)
	}
}
