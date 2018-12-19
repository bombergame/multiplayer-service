package rest

import (
	"bufio"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"

	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/mailru/easyjson"
)

const (
	//UserAgentHeader is a corresponding header
	UserAgentHeader = "User-Agent"

	//AuthorizationHeader is a corresponding header
	AuthorizationHeader = "Authorization"

	//ProfileIDHeader is a corresponding header
	ProfileIDHeader = "X-Profile-ID"
)

//LoggingResponseWriter wraps ResponseWriter and saves responses info
type LoggingResponseWriter struct {
	status int
	writer http.ResponseWriter
}

//Header wraps corresponding method
func (w *LoggingResponseWriter) Header() http.Header {
	return w.writer.Header()
}

//Write wraps corresponding method
func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

//WriteHeader wraps corresponding method and saves response status
func (w *LoggingResponseWriter) WriteHeader(status int) {
	w.status = status
	w.writer.WriteHeader(status)
}

//Hijack wraps corresponding method
func (w *LoggingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.writer.(http.Hijacker).Hijack()
}

//ReadHeader returns header value or error
func (srv *Service) ReadHeader(r *http.Request, name string) (string, error) {
	v := r.Header.Get(name)
	if v == consts.EmptyString {
		err := errs.NewBadRequestError("header \"" + name + "\" not set")
		return consts.EmptyString, err
	}
	return v, nil
}

//ReadRequestBody parses json body
func (srv *Service) ReadRequestBody(v easyjson.Unmarshaler, r *http.Request) error {
	const wrongFormatMessage = "wrong request body format"

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errs.NewInvalidFormatError(wrongFormatMessage)
	}
	defer func() {
		if err = r.Body.Close(); err != nil {
			panic(err)
		}
	}()

	err = easyjson.Unmarshal(body, v)
	if err != nil {
		return errs.NewInvalidFormatError(wrongFormatMessage)
	}

	return nil
}

//WriteOk writes http OK response
func (srv *Service) WriteOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

//WriteOkWithBody writes http OK response with json body
func (srv *Service) WriteOkWithBody(w http.ResponseWriter, v easyjson.Marshaler) {
	srv.writeJSON(w, v)
}

//WriteError writes error response status only
func (srv *Service) WriteError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

//WriteErrorWithBody writes http error response with json body
func (srv *Service) WriteErrorWithBody(w http.ResponseWriter, err error) {
	var status int

	sErr, ok := err.(*errs.ServiceError)
	if ok {
		switch sErr.ErrorType() {
		case errs.NotAuthorized:
			status = http.StatusUnauthorized

		case errs.AccessDenied:
			status = http.StatusForbidden

		case errs.InvalidFormat:
			status = http.StatusUnprocessableEntity

		case errs.Duplicate:
			status = http.StatusConflict

		case errs.NotFound:
			status = http.StatusNotFound

		case errs.Internal:
			srv.components.Logger.Error(sErr.InnerError())
			status = http.StatusInternalServerError

		default:
			srv.components.Logger.Error(err.Error())
			status = http.StatusInternalServerError
		}

		srv.writeText(w, status, err.Error())
	} else {
		srv.components.Logger.Error(err.Error())
		status = http.StatusInternalServerError
	}

	srv.writeText(w, status, err.Error())
}

//ReadAuthProfileID returns authenticated user profile ID
func (srv *Service) ReadAuthProfileID(r *http.Request) (int64, error) {
	v, err := srv.ReadHeader(r, ProfileIDHeader)
	if err != nil {
		return consts.AnyInt, err
	}

	iv64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return consts.AnyInt, errs.NewNotAuthorizedError()
	}

	return iv64, nil
}

//ReadUserAgent returns User-Agent header value
func (srv *Service) ReadUserAgent(r *http.Request) (string, error) {
	return srv.ReadHeader(r, UserAgentHeader)
}

//ReadAuthToken returns Authorization header value without Bearer prefix
func (srv *Service) ReadAuthToken(r *http.Request) (string, error) {
	const prefix = "Bearer "

	bearer, err := srv.ReadHeader(r, AuthorizationHeader)
	if err != nil {
		return consts.EmptyString, err
	}

	n := len(prefix)
	if len(bearer) <= n || bearer[:n] != prefix {
		return consts.EmptyString, errs.NewNotAuthorizedError()
	}

	return bearer[n:], nil
}

func (srv *Service) writeText(w http.ResponseWriter, status int, txt string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(txt)); err != nil {
		panic(err)
	}
}

func (srv *Service) writeJSON(w http.ResponseWriter, v easyjson.Marshaler) {
	b, err := easyjson.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}

func (srv *Service) setAuthProfileID(r *http.Request, id int64) {
	r.Header.Set(ProfileIDHeader, strconv.FormatInt(id, 10))
}
