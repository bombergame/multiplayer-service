package rest

import (
	"github.com/bombergame/common/errs"
	"net/http"
)

func (srv *Service) joinRoom(w http.ResponseWriter, r *http.Request) {
	_, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		err = errs.NewServiceError(err)
		srv.WriteErrorWithBody(w, err)
		return
	}

	//TODO
}
