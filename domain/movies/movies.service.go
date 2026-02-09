package movies

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type MoviesServiceInterface interface {
	Search(search string) ([]Movie, error)
	List(search, filter string, page int) ([]Movie, error)
}

type MoviesService struct {
}

const API_URL = "https://server-hi-speed-iran.info"
const API_KEY = "4F5A9C3D9A86FA54EACEDDD635185"

func NewMoviesService() *MoviesService {
	return &MoviesService{}
}

func (service *MoviesService) Search(search string) ([]Movie, error) {
	url := fmt.Sprintf("%s/api/search/%s/%s", API_URL, url.QueryEscape(search), API_KEY)

	fmt.Println(url, search)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error occurred")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string][]Movie
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data["posters"], nil
}

func (service *MoviesService) List(search, filter string, page int) ([]Movie, error) {
	url := fmt.Sprintf("%s/api/movie/by/filtres/0/%s/%d/%s", API_URL, filter, page, API_KEY)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error occurred")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []Movie
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
