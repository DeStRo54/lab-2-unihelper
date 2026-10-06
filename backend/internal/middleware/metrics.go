package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/errs"
	"homewormanager/pkg/metrics"
	"time"
)

func (mw *MwManager) Metrics() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		duration := time.Since(start)
		path := c.Path()
		method := c.Method()
		status := c.Response().StatusCode()
		if err != nil {
			status = errs.GetStatusFromError(err)
		}

		statusStr := fmt.Sprintf("%d", status)

		metrics.HTTPRequestsTotal.WithLabelValues(method, path, statusStr).Inc()
		metrics.HTTPRequestDuration.WithLabelValues(method, path, statusStr).Observe(duration.Seconds())
		return err
	}
}
