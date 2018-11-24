package rest

import (
	"fmt"
	"net/http"
)

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

func (srv *Service) WithLogs(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wr := &LoggingResponseWriter{
				writer: w,
			}
			h.ServeHTTP(wr, r)

			srv.components.Logger.Info(fmt.Sprintf("%s %s %d", r.Method, r.RequestURI, wr.status))
		},
	)
}

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
