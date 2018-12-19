package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//WithRecover recovers after panic in request handler
func (srv *Service) WithRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					srv.components.Logger.Error(r)
					srv.WriteError(w, http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}

//WithLogs adds logs for each http request
func (srv *Service) WithLogs(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wr := &LoggingResponseWriter{
				writer: w,
			}
			h.ServeHTTP(wr, r)

			srv.components.Logger.Info(
				fmt.Sprintf("%s %s %d", r.Method, r.RequestURI, wr.status),
			)
		},
	)
}

//WithAuth parses user profile id from auth token
func (srv *Service) WithAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userAgent, err := srv.ReadUserAgent(r)
			if err != nil {
				srv.WriteErrorWithBody(w, err)
				return
			}

			authToken, err := srv.ReadAuthToken(r)
			if err != nil {
				srv.WriteErrorWithBody(w, err)
				return
			}

			info, err := srv.components.AuthManager.GetProfileInfo(authToken, userAgent)
			if err != nil {
				srv.WriteErrorWithBody(w, err)
				return
			}

			srv.setAuthProfileID(r, info.ID)
			h.ServeHTTP(w, r)
		},
	)
}

//CORS contains CORS response headers
type CORS struct {
	Origins     []string
	Methods     []string
	Headers     []string
	Credentials bool
}

//WithCORS adds CORS headers to each response
func (srv *Service) WithCORS(h http.Handler, cors CORS) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(
				"Access-Control-Allow-Origin",
				strings.Join(cors.Origins, ","),
			)
			w.Header().Set(
				"Access-Control-Allow-Credentials",
				strconv.FormatBool(cors.Credentials),
			)
			w.Header().Set(
				"Access-Control-Allow-Methods",
				strings.Join(cors.Methods, ","),
			)
			w.Header().Set(
				"Access-Control-Allow-Headers",
				strings.Join(cors.Headers, ","),
			)
			h.ServeHTTP(w, r)
		},
	)
}
