// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bombs

import (
	json "encoding/json"
	objects "github.com/bombergame/multiplayer-service/game/objects"
	state "github.com/bombergame/multiplayer-service/game/objects/bombs/state"
	physics "github.com/bombergame/multiplayer-service/game/physics"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonDe73ae40DecodeGithubComBombergameMultiplayerServiceGameObjectsBombs(in *jlexer.Lexer, out *MessageData) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "state":
			out.State = state.State(in.String())
		case "transform":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Transform).UnmarshalJSON(data))
			}
		case "explosion_radius":
			out.ExplosionRadius = physics.Integer(in.Int32())
		case "explosion_timeout":
			out.ExplosionTimeout = float64(in.Float64())
		case "object_id":
			out.ObjectID = objects.ObjectID(in.Int64())
		case "object_type":
			out.ObjectType = objects.ObjectType(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonDe73ae40EncodeGithubComBombergameMultiplayerServiceGameObjectsBombs(out *jwriter.Writer, in MessageData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"transform\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Transform).MarshalJSON())
	}
	{
		const prefix string = ",\"explosion_radius\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ExplosionRadius))
	}
	{
		const prefix string = ",\"explosion_timeout\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.ExplosionTimeout))
	}
	{
		const prefix string = ",\"object_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ObjectID))
	}
	{
		const prefix string = ",\"object_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ObjectType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe73ae40EncodeGithubComBombergameMultiplayerServiceGameObjectsBombs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessageData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe73ae40EncodeGithubComBombergameMultiplayerServiceGameObjectsBombs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe73ae40DecodeGithubComBombergameMultiplayerServiceGameObjectsBombs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessageData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe73ae40DecodeGithubComBombergameMultiplayerServiceGameObjectsBombs(l, v)
}