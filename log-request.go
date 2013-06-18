package httplog

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

/* kindof a common log type output */
func LogRequest(r *http.Request, statusCode int) {
	var addr string
	var user_agent string

	user_agent = ""
	addr = r.RemoteAddr

	for k, v := range r.Header {
		if k == "User-Agent" {
			user_agent = strings.Join(v, " ")
		}
		if k == "X-Forwarded-For" {
			addr = strings.Join(v, " ")
		}
	}

	fmt.Printf("%s - - [%s] \"%s %s\" \"%s\" %d %d\n",
		addr,
		time.Now(),
		r.Method,
		r.URL.String(),
		user_agent,
		statusCode,
		r.ContentLength)
}
