package movies

type ResponseType map[string]any

type Movie struct {
	ID             int         `json:"id"`
	Type           string      `json:"type"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Year           int         `json:"year"`
	Imdb           float64     `json:"imdb"`
	Comment        bool        `json:"comment"`
	Rating         float64     `json:"rating"`
	Duration       interface{} `json:"duration"`
	Downloadas     string      `json:"downloadas"`
	Playas         string      `json:"playas"`
	Classification interface{} `json:"classification"`
	Image          string      `json:"image"`
	Cover          string      `json:"cover"`
	Genres         []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"genres"`
	Sources []struct {
		ID      int    `json:"id"`
		Quality string `json:"quality"`
		Type    string `json:"type"`
		URL     string `json:"url"`
	} `json:"sources"`
	Country []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
}

type MovieResponseType struct {
	Movie Movie `json:"movie"`
}

func MovieResponse(movie Movie) ResponseType {
	return ResponseType{
		"movie": movie,
	}
}

type MoviesResponseType struct {
	Movies []Movie `json:"movies"`
}

func MoviesResponse(movies []Movie) ResponseType {
	return ResponseType{
		"movies": movies,
	}
}
