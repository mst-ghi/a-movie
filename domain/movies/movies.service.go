package movies

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"app/core/config"
)

type MoviesServiceInterface interface {
	Search(ctx context.Context, search string) ([]Movie, error)
	List(ctx context.Context, filter string, page int) ([]Movie, error)
}

type MoviesService struct {
}

const upstreamTimeout = 10 * time.Second

var httpClient = &http.Client{Timeout: upstreamTimeout}

var validFilters = map[string]bool{
	"created": true,
	"imdb":    true,
	"year":    true,
}

func NewMoviesService() *MoviesService {
	return &MoviesService{}
}

func getJSON(ctx context.Context, endpoint string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upstream returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}

func (service *MoviesService) Search(ctx context.Context, search string) ([]Movie, error) {
	endpoint := fmt.Sprintf(
		"%s/api/search/%s/%s",
		config.Get("MOVIE_API_URL"),
		url.QueryEscape(search),
		config.Get("MOVIE_API_KEY"),
	)

	var data map[string][]Movie
	if err := getJSON(ctx, endpoint, &data); err != nil {
		return nil, err
	}

	return data["posters"], nil
}

func (service *MoviesService) List(ctx context.Context, filter string, page int) ([]Movie, error) {
	if !validFilters[filter] {
		filter = "created"
	}

	endpoint := fmt.Sprintf(
		"%s/api/movie/by/filtres/0/%s/%d/%s",
		config.Get("MOVIE_API_URL"),
		filter,
		page,
		config.Get("MOVIE_API_KEY"),
	)

	var data []Movie
	if err := getJSON(ctx, endpoint, &data); err != nil {
		return nil, err
	}

	return data, nil
}
