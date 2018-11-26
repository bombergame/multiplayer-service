package domains

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/satori/go.uuid"
)

//easyjson:json
type Room struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	CreatedBy      int64     `json:"-"`
	TimeLimit      int64     `json:"time_limit"`
	MaxNumPlayers  int64     `json:"max_num_players"`
	AllowAnonymous bool      `json:"allow_anonymous"`
	FieldSize      FieldSize `json:"field_size"`
}

//easyjson:json
type FieldSize struct {
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
}

func (r *Room) Validate() error {
	if r.Title == consts.EmptyString {
		return errs.NewInvalidFormatError("empty title")
	}
	if r.TimeLimit < 5 || r.TimeLimit > 30 {
		return errs.NewInvalidFormatError("time limit out of range")
	}
	if r.MaxNumPlayers < 1 || r.MaxNumPlayers > 10 {
		return errs.NewInvalidFormatError("max number of players out of range")
	}
	if r.FieldSize.Width < 10 || r.FieldSize.Width > 100 {
		return errs.NewInvalidFormatError("field width out of range")
	}
	if r.FieldSize.Height < 10 || r.FieldSize.Height > 100 {
		return errs.NewInvalidFormatError("field height out of range")
	}
	return nil
}
