package rest

import (
	"encoding/json"
	"errors"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"net/http"
	"strconv"
)

const (
	UserAgentHeader     = "User-Agent"
	AuthorizationHeader = "Authorization"
	ProfileIDHeader     = "X-Profile-ID"
)

type LoggingResponseWriter struct {
	status int
	writer http.ResponseWriter
}

func (w *LoggingResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func (w *LoggingResponseWriter) WriteHeader(status int) {
	w.status = status
	w.writer.WriteHeader(status)
}

func (srv *Service) ReadHeader(r *http.Request, name string) (string, error) {
	v := r.Header.Get(name)
	if v == consts.EmptyString {
		err := errors.New(name + " header not set")
		return consts.EmptyString, errs.NewServiceError(err)
	}
	return v, nil
}

func (srv *Service) ReadRequestBody(v interface{}, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return errs.NewInvalidFormatError("wrong request body format")
	}
	return nil
}

func (srv *Service) WriteOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (srv *Service) WriteOkWithBody(w http.ResponseWriter, v interface{}) {
	srv.writeJSON(w, http.StatusOK, v)
}

func (srv *Service) WriteError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (srv *Service) WriteErrorWithBody(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	switch srvErr := err.(type) {
	case *errs.NotAuthorizedError:
		status = http.StatusUnauthorized

	case *errs.AccessDeniedError:
		status = http.StatusForbidden

	case *errs.InvalidFormatError:
		status = http.StatusUnprocessableEntity

	case *errs.DuplicateError:
		status = http.StatusConflict

	case *errs.NotFoundError:
		status = http.StatusNotFound

	case *errs.ServiceError:
		srv.components.Logger.Error(srvErr.InnerError())
		status = http.StatusInternalServerError

	default:
		srv.components.Logger.Error(err.Error())
		srv.writeText(w, http.StatusInternalServerError, errs.ServiceErrorMessage)
		return
	}

	srv.writeText(w, status, err.Error())
}

func (srv *Service) writeText(w http.ResponseWriter, status int, txt string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(txt))
}

func (srv *Service) writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (srv *Service) ReadAuthProfileID(r *http.Request) (int64, error) {
	v, err := srv.ReadHeader(r, ProfileIDHeader)
	if err != nil {
		return 0, err
	}

	iv64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, errs.NewServiceError(err)
	}

	return iv64, nil
}

func (srv *Service) ReadUserAgent(r *http.Request) (string, error) {
	return srv.ReadHeader(r, UserAgentHeader)
}

func (srv *Service) ReadAuthToken(r *http.Request) (string, error) {
	bearer, err := srv.ReadHeader(r, AuthorizationHeader)
	if err != nil {
		return consts.EmptyString, err
	}

	n := len("Bearer ")
	if len(bearer) <= n {
		return consts.EmptyString, errs.NewInvalidFormatError("wrong authorization token")
	}

	return bearer[n:], nil
}

func (srv *Service) setAuthProfileID(r *http.Request, id int64) {
	r.Header.Set(ProfileIDHeader, strconv.FormatInt(id, 10))
}
