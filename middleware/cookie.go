package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cookie(c *gin.Context) {
	// get 1
	ck, _ := c.Request.Cookie("foo")
	fmt.Println("cookie is ", ck.String())
	// get all
	cks := c.Request.Cookies()
	fmt.Println("all cookie ", cks)

	//

	// set 1
	var c1 = &http.Cookie{
		Name:  "foo",
		Value: "123",
	}
	http.SetCookie(c.Writer, c1)

	// set 2
	var c2 = &http.Cookie{
		Name:  "bar",
		Value: "456",
	}
	c.Writer.Header().Set("Set-Cookie", c2.String())
	// set 3
	var c3 = &http.Cookie{
		Name:  "baz",
		Value: "789",
	}
	c.Writer.Header().Add("Set-Cookie", c3.String())

	c.Next()

}
