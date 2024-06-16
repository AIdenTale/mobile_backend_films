package models

type DataBaseError struct {
	Error string `json:"error"`
}

type FilmReviewDbResponse struct {
	Data []FilmReview `json:"data"`
}

type FilmReview struct {
	ID          int    `json:"id"`
	Dt          string `json:"dt"`
	Rating      int    `json:"rating"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ToFilm      struct {
		Title       string  `json:"title"`
		Year        string  `json:"year"`
		PosterLink  *string `json:"poster_link"`
		TrailerLink *string `json:"trailer_link"`
		Description string  `json:"description"`
		Country     string  `json:"country"`
		Rating      int     `json:"rating"`
	} `json:"to_film"`
	User struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	} `json:"user"`
	Likes    int `json:"likes"`
	Comments int `json:"comments"`
}

type FilmReviewCommentDbResp struct {
	Data []FilmReviewComment `json:"data"`
}

type FilmReviewComment struct {
	ID       int `json:"id"`
	Comments struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Dt       string `json:"dt"`
		Text     string `json:"text"`
		Likes    int    `json:"likes"`
	} `json:"comments"`
}
