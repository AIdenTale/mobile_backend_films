package internal

import (
	"courseProject/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ComplexHandler struct {
	service *FilmsService
}

func NewComplexHandler(service *FilmsService) *ComplexHandler {
	return &ComplexHandler{
		service: service,
	}
}

type HandlerMapURLConfig struct {
	URL     string
	Handler func(ctx *gin.Context)
	Method  string
}

func (h *ComplexHandler) RegisterUrls() *gin.Engine {
	g := gin.Default()

	optsURLS := []HandlerMapURLConfig{
		{URL: "/api/reviews", Handler: h.GetAllFilmsReviewsHandler, Method: "GET"},
		{URL: "/api/reviews/comments", Handler: h.GetFilmReviewCommentsHandler, Method: "GET"},
	}

	g.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	for i := range optsURLS {
		switch optsURLS[i].Method {
		case "GET":
			g.GET(optsURLS[i].URL, optsURLS[i].Handler)
			break
		case "POST":
			g.POST(optsURLS[i].URL, optsURLS[i].Handler)
			break
		}
		g.OPTIONS(optsURLS[i].URL, h.OptionsAcceptHandler)
	}

	return g
}

func (h *ComplexHandler) OptionsAcceptHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.IndentedJSON(http.StatusOK, "ok")
}

func (h *ComplexHandler) GetAllFilmsReviewsHandler(ctx *gin.Context) {
	utils.AcceptAllHosts(ctx)
	flats, err := h.service.GetAllFilmsReviews()
	if err != nil {
		utils.InternalServiceError(ctx, err)
		return
	}

	utils.BindObjToContext(ctx, flats)
}

func (h *ComplexHandler) GetFilmReviewCommentsHandler(ctx *gin.Context) {
	utils.AcceptAllHosts(ctx)

	queryParams := ctx.Request.URL.Query()
	param, ok := queryParams["review_id"]
	if !ok {
		utils.ValidationError(ctx, errors.New("cannot get review_id from query params"))
		return
	}
	reviewID, err := strconv.Atoi(param[0])
	if err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	data, err := h.service.GetAllFilmReviewsCommentsByID(reviewID)
	if err != nil {
		utils.InternalServiceError(ctx, err)
		return
	}

	utils.BindObjToContext(ctx, data)
	return
}
