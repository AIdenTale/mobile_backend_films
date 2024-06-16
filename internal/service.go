package internal

import (
	"courseProject/db"
	"courseProject/models"
)

type FilmsService struct {
	db *db.PostgresDriver
}

func NewComplexService(dbDriver *db.PostgresDriver) *FilmsService {
	return &FilmsService{
		db: dbDriver,
	}
}

func (s *FilmsService) GetAllFilmsReviews() ([]models.FilmReview, error) {
	filmDbData := models.FilmReviewDbResponse{}

	err := s.db.ExecuteSP("get_all_reviews", &filmDbData, nil)
	if err != nil {
		return nil, err
	}

	return filmDbData.Data, nil
}

func (s *FilmsService) GetAllFilmReviewsCommentsByID(reviewID int) ([]models.FilmReviewComment, error) {
	commentsDbData := models.FilmReviewCommentDbResp{}

	type FilmReviewIDParams struct {
		ReviewID int `json:"review_id"`
	}

	reviewIDparams := FilmReviewIDParams{
		ReviewID: reviewID,
	}

	err := s.db.ExecuteSP("get_all_review_comments_by_id", &commentsDbData, reviewIDparams)
	if err != nil {
		return nil, err
	}

	return commentsDbData.Data, nil
}
