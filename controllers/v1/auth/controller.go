package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewToken(c *gin.Context) {
	var s Session
	s.GenSessionToken()

	b64Token := base64.RawURLEncoding.EncodeToString(s.Sid)

	c.String(http.StatusOK, b64Token)
}

func VerifyToken(c *gin.Context, redirect func(c *gin.Context)) {

	b64Token := c.Query("token")

	token, berr := base64.RawURLEncoding.DecodeString(b64Token)
	if berr != nil {
		c.String(http.StatusForbidden, "Invalid Token")
	}

	if chk, err := MatchToken(token); chk {
		redirect(c)
	} else {
		c.String(http.StatusForbidden, err.Error())
	}
}
