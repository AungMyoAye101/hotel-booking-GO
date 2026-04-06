package pagination

import (
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Params struct {
	Page  int
	Limit int
}

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type Response[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}

func Parse(c echo.Context) (Params, error) {
	page := 1
	limit := 10

	if v := c.QueryParam("page"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil || p < 1 {
			return Params{}, echo.NewHTTPError(http.StatusBadRequest, "invalid page")
		}
		page = p
	}

	if v := c.QueryParam("limit"); v != "" {
		l, err := strconv.Atoi(v)
		if err != nil || l < 1 {
			return Params{}, echo.NewHTTPError(http.StatusBadRequest, "invalid limit")
		}
		if l > 100 {
			l = 100
		}
		limit = l
	}

	return Params{Page: page, Limit: limit}, nil
}

func NewMeta(params Params, total int64) Meta {
	totalPages := 0
	if params.Limit > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(params.Limit)))
	}
	return Meta{
		Page:       params.Page,
		Limit:      params.Limit,
		Total:      total,
		TotalPages: totalPages,
	}
}

