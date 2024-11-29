package session

import (
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/golang-jwt/jwt/v5"
)

func SessionJwt(jwtkey string, expire int) core.Handle {
	if expire == 0 {
		expire = 1800
	}
	jwtSignKey := []byte(jwtkey)
	return func(c *core.GContent) {
		// log.Debug(c, "jwt")
		token := c.GetRequest().Header.Get("Token")
		if token == "" {
			token = c.GetRequest().URL.Query().Get("Token")
		}
		data := &jwt.RegisteredClaims{}
		if token != "" {
			t, e := jwt.ParseWithClaims(token, data, func(t *jwt.Token) (interface{}, error) {
				return jwtSignKey, nil
			})
			if e == nil && t.Valid { // && data.ExpiresAt != nil && data.ExpiresAt.Sub(time.Now()) > 0 {
				c.SetUserID(data.ID)
			}

		}
		c.Next()
		if c.IsLogin() {
			// log.Debug(c, "jwt login")
			data.ID = fmt.Sprintf("%d", c.GetUserID())
			data.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expire)))
			t := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
			sendToken, e := t.SignedString(jwtSignKey)
			if e == nil {
				// log.Debug(c, "jwt login")
				c.GetResponseWriter().Header().Add("Token", sendToken)
			} else {
				// log.Debug(c, "jwt e -> %s", e.Error())
			}
		}
		// log.Debug(c, "jwt end")
	}
}
