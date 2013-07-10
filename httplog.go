package httplog

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var (
  /* This default icon is empty with a long lived cache */
	DefaultFavIcon FavIcon = defaultFavIcon{}
)

type defaultFavIcon struct {
}

func (dfi defaultFavIcon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	LogRequest(r, 200)
	w.Header().Set("Cache-Control", "max-age=315360000")
}

/* simple interface for a favicon */
type FavIcon interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// for debugging request headers
func LogHeaders(r *http.Request) {
	fmt.Printf("HEADERS:\n")
	for k, v := range r.Header {
		fmt.Printf("\t%s\n", k)
		for i, _ := range v {
			fmt.Printf("\t\t%s\n", v[i])
		}
	}
}

/* kindof a common log type output */
func LogRequest(r *http.Request, statusCode int) {
	var addr string
	var user_agent string

	user_agent = ""
	addr = RealIP(r)

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
		time.Now().Format(time.RFC1123Z),
		r.Method,
		r.URL.String(),
		user_agent,
		statusCode,
		r.ContentLength)
}

func RealIP(r *http.Request) (ip string) {
	ip = r.RemoteAddr

	port_pos := strings.LastIndex(ip, ":")
	if port_pos != -1 {
		ip = ip[0:port_pos]
	}

	for k, v := range r.Header {
		if k == "X-Forwarded-For" {
			ip = strings.Join(v, " ")
		}
	}

	return ip
}
