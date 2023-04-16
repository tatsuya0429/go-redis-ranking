package timeline

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TimelineController struct {
	repo *TimelineRepository
}

func NewTimelineController() *TimelineController {
	return &TimelineController{
		repo: NewTimelineRepository(),
	}
}

func (con *TimelineController) GetTimeline(c echo.Context) error {
	limit := 50
	offset := 0
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	offset, err = strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	posts, err := con.repo.GetTimeline(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, posts)
}
