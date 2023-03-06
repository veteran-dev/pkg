package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/veteran-dev/pkg/fxuntils"
	"github.com/veteran-dev/pkg/system/configfx"
)

// Protected protect routes
func Protected() func(c *fiber.Ctx) error {
	config := configfx.ProvideConfig()
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.PasswordSalt),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	var resp fxuntils.Response
	var returns fxuntils.ReturnResp
	if err.Error() == "Missing or malformed JWT" {
		returns.StatusCode = fxuntils.MissingJWT
		returns.StatusReason = fxuntils.FailStatusMessage(returns.StatusCode)
	} else {
		returns.StatusCode = fxuntils.InvalidJWT
		returns.StatusReason = fxuntils.FailStatusMessage(returns.StatusCode)
	}
	resp.Data = returns
	resp.Message = fxuntils.FxFail
	return fxuntils.Send(c, resp)
}
