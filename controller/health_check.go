package controller

import (
	"go-enum-example/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	appCentral usecase.AppCentral
}

func NewController() *echo.Echo {
	ctl := &controller{
		appCentral: usecase.AppCentral{},
	}

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		log.Printf("HTTPErrorHandler: %s", err)
		c.String(http.StatusInternalServerError, "internal server error")
	}
	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	})

	e.GET("/health_check", ctl.HealthCheck)

	return e
}

func (ctl *controller) HealthCheck(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := ctl.appCentral.CheckHealthStatus(ctx)
	if err != nil {
		return err // error is handled by (*echo.Echo).HTTPErrorHandler
	}

	return result.Handle(
		usecase.NewCheckHealthStatusHandler(
			func(e usecase.CheckHealthStatusHealthy) error {
				log.Printf("health check succeeded: duration=%s", e.FinishedAt.Sub(e.StartedAt))
				return c.JSON(http.StatusOK, map[string]string{
					"status": "healthy",
				})
			},
			func(e usecase.CheckHealthStatusUnhealthy) error {
				log.Printf("health check failed: cause=%q duration=%s", e.Cause, e.FinishedAt.Sub(e.StartedAt))
				return c.JSON(http.StatusServiceUnavailable, map[string]string{
					"status": "unhealthy",
				})
			},
		),
	)
}
