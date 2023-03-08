package fxuntils

import "github.com/gofiber/fiber/v2"

// A struct to return normal responses.
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ReturnResp struct {
	StatusCode   int    `json:"statusCode"`
	StatusReason string `json:"statusReason"`
}

const (
	FxOk   = "successd"
	FxFail = "faild"
)

// A fuction to return beautiful responses.
func Send(c *fiber.Ctx, resp Response) error {
	// Set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.Code)
	// Return JSON
	return c.JSON(resp)
}

const (
	InvalidJWT   = 4100
	MissingJWT   = 4101
	EmptyParam   = 4111
	ParamMissing = 4112

	LoginFailed         = 4201
	QueryAccountFailed  = 4202
	AccountNotExist     = 4203
	PasswordNotMatch    = 4204
	GenerateTokenFailed = 4205

	QueryAccountServiceFailed  = 4301
	CreateAccountServiceFailed = 4302
	UpdateAccountServiceFailed = 4303
)

func SuccessStatusMessage(status int) string {
	return successStatusMessage[status]
}

var successStatusMessage = []string{
	2000: "Successd.", // StatusContinue
}

func FailStatusMessage(status int) string {
	return failStatusMessage[status]
}

// 业务状态码
var failStatusMessage = []string{
	4100: "invalid or expired JWT",
	4101: "missing or malformed JWT.", // StatusContinue
	4111: "empty parameter.",
	4112: "missing parameter.",

	4201: "login failed.",
	4202: "query account failed.",
	4203: "account does not exist.",
	4204: "passwords do not match.",
	4205: "failed to generate token.",

	4301: "query account service failed.",
	4302: "failed to create account service.",
	4303: "failed to modify account service.",
}

const (
	On  = 1
	Off = 0
)
