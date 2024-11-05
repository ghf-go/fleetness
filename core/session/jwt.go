package session

import (
	"fmt"

	"github.com/ghf-go/fleetness/core"
	"github.com/golang-jwt/jwt/v5"
)

func SessionJwt(jwtkey string) core.Handle {
	jwtSignKey := []byte(jwtkey)
	return func(c *core.GContent) {
		// log.Debug(c, "jwt")
		token := c.GetRequest().Header.Get("Token")
		data := &jwt.RegisteredClaims{}
		if token != "" {
			t, e := jwt.ParseWithClaims(token, data, func(t *jwt.Token) (interface{}, error) {
				return jwtSignKey, nil
			})
			if e == nil && t.Valid {
				c.SetUserID(data.ID)
			}

		}
		c.Next()
		if c.IsLogin() {
			// log.Debug(c, "jwt login")
			data.ID = fmt.Sprintf("%d", c.GetUserID())
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
