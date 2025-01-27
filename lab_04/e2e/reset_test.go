package e2e

import (
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gavv/httpexpect/v2"
)

var (
	expectReset *httpexpect.Expect
)

func TestResetPassword(t *testing.T) {
	client := &http.Client{}
	expectReset = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8081",
		Client:   client,
		Reporter: httpexpect.NewRequireReporter(nil),
	})

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeResetPasswordScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/reset.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run reset password feature tests")
	}
}

func resetPasswordWith2FA(ctx *godog.ScenarioContext) {
	var response *httpexpect.Response
	secret := os.Getenv("TOTP_SECRET")
	totp := NewOTPVerifier(secret)

	ctx.Step(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectReset.Request(method, endpoint).
			WithJSON(map[string]string{
				"login":       "user",
				"oldPassword": "user",
			}).
			Expect()
		return nil
	})

	ctx.Step(`^the response on /reset code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /reset should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"success": "Now should enter 2fa code to reset password",
		})
		return nil
	})

	ctx.Step(`^user send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectReset.Request(method, endpoint).
			WithJSON(map[string]string{
				"login":       "user",
				"code":        totp.totp.Now(),
				"newPassword": "test",
			}).Expect()
		return nil
	})

	ctx.Step(`^the response on /verifyResetPassword code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /verifyResetPassword should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"success": "successful changing pass",
		})
		return nil
	})
}

func InitializeResetPasswordScenario(ctx *godog.ScenarioContext) {
	resetPasswordWith2FA(ctx)
}
