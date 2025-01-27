package e2e

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/gavv/httpexpect/v2"
	"github.com/xlzd/gotp"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	expectLogin *httpexpect.Expect
)

func TestLogin(t *testing.T) {
	client := &http.Client{}
	expectLogin = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8081",
		Client:   client,
		Reporter: httpexpect.NewRequireReporter(nil),
	})

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeLoginScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/login.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run login feature tests")
	}
}

type OTPVerifier struct {
	totp *gotp.TOTP
}

func NewOTPVerifier(secret string) *OTPVerifier {
	return &OTPVerifier{
		totp: gotp.NewDefaultTOTP(secret),
	}
}

func (v OTPVerifier) Verify(code string) error {
	if v.totp.Verify(code, time.Now().Unix()) {
		return nil
	} else {
		return fmt.Errorf("OTP verification failed")
	}
}

func loginWith2FA(ctx *godog.ScenarioContext) {
	var response *httpexpect.Response
	secret := os.Getenv("TOTP_SECRET")
	totp := NewOTPVerifier(secret)

	ctx.Step(`^User send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectLogin.Request(method, endpoint).
			WithJSON(map[string]string{
				"login":    "user",
				"password": "user",
			}).
			Expect()
		return nil
	})

	ctx.Step(`^the response on /login code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /login should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"success": "now should enter 2fa",
		})
		return nil
	})

	ctx.Step(`^user send "([^"]*)" request to "([^"]*)"$`, func(method, endpoint string) error {
		response = expectLogin.Request(method, endpoint).
			WithJSON(map[string]string{
				"login": "user",
				"code":  totp.totp.Now(),
			}).Expect()
		return nil
	})

	ctx.Step(`^the response on /verify code should be (\d+)$`, func(statusCode int) error {
		response.Status(statusCode)
		return nil
	})

	ctx.Step(`^the response on /verify should match json:$`, func(expectedJSON *godog.DocString) error {
		response.JSON().Object().IsEqual(map[string]interface{}{
			"success": "successfully logged in",
		})
		return nil
	})
}

func InitializeLoginScenario(ctx *godog.ScenarioContext) {
	loginWith2FA(ctx)
}
