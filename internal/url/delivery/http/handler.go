package http

import (
	"log/slog"
	"net/http"
	"url-shortener/internal/url/delivery/dto"
	"url-shortener/internal/url/interface"
	"url-shortener/internal/util/response"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlUseCase _interface.UseCase
}

func NewUrlHandler(urlUseCase _interface.UseCase) _interface.Handler {
	return &urlHandler{
		urlUseCase: urlUseCase,
	}
}

func (u *urlHandler) Shorten(ctx *gin.Context) {
	c := ctx.Request.Context()

	var req dto.ShortenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("error binding json", err)

		response.Error(
			ctx,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)

		return
	}

	result, err := u.urlUseCase.Shorten(c, req.LongURL)
	if err != nil {
		slog.Error("error shortening url", err)

		response.Error(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)

		return
	}

	response.Success(
		ctx,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	)
}

func (u *urlHandler) Redirect(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
