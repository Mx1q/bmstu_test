package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/xlzd/gotp"
)

var login = "user"
var pass = "user"

var changePassReqs map[string]bool = make(map[string]bool)
var loginReqs map[string]bool = make(map[string]bool)

func main() {
	//conf, err := config.ReadConfig("./config.yaml")
	//if err != nil {
	//	log.Fatal(err)
	//}

	secret := os.Getenv("TOTP_SECRET")
	totp := NewOTPVerifier(secret)

	r := gin.Default()
	//log.Println(totp.totp.Now()) // TODO: check code in e2e tests + ci cd

	r.POST("/login", LoginHandler)
	r.POST("/reset", ResetPasswordHandler)
	r.POST("/verifyLogin", func(c *gin.Context) {
		type Req struct {
			Login string `json:"login"`
			Code  string `json:"code"`
		}
		var req Req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if !loginReqs[req.Login] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Should enter pass first"})
			return
		}
		delete(loginReqs, req.Login)

		err := totp.Verify(req.Code)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid 2fa code"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "successfully logged in"})
	})
	r.POST("/verifyResetPassword", func(c *gin.Context) {
		type Req struct {
			Code        string `json:"code"`
			Login       string `json:"login"`
			NewPassword string `json:"newPassword"`
		}
		var req Req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if !changePassReqs[req.Login] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Should enter pass first"})
			return
		}
		delete(changePassReqs, req.Login)

		err := totp.Verify(req.Code)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid 2fa code"})
			return
		}

		pass = req.NewPassword
		c.JSON(http.StatusOK, gin.H{"success": "successful changing pass"})
	})

	fmt.Println("server was started")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func LoginHandler(c *gin.Context) {
	type Req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Login != login || req.Password != pass {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login"})
		return
	}

	loginReqs[req.Login] = true
	c.JSON(http.StatusOK, gin.H{"success": "now should enter 2fa"})
}

func ResetPasswordHandler(c *gin.Context) {
	var request struct {
		Login       string `json:"login"`
		OldPassword string `json:"oldPassword"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.Login != login || request.OldPassword != pass {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	changePassReqs[request.Login] = true
	c.JSON(http.StatusOK, gin.H{"success": "Now should enter 2fa code to reset password"})
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
