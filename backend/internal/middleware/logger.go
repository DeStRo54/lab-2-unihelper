package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (mw *MwManager) RequestLogger(logger *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Infow("Incoming request",
			"method", c.Method(),
			"url", c.OriginalURL(),
			"remote_ip", c.IP(),
		)

		return c.Next()
	}
}

func (mw *MwManager) ErrorLogger(logger *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			logger.Errorw("Unhandled error occurred",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.Error(err),
			)
			return err
		}

		return nil
	}
}
