// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domains

import (
	json "encoding/json"
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

func easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceDomains(in *jlexer.Lexer, out *Room) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "title":
			out.Title = string(in.String())
		case "time_limit":
			out.TimeLimit = int64(in.Int64())
		case "max_num_players":
			out.MaxNumPlayers = int64(in.Int64())
		case "allow_anonymous":
			out.AllowAnonymous = bool(in.Bool())
		case "field_size":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.FieldSize).UnmarshalJSON(data))
			}
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
func easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceDomains(out *jwriter.Writer, in Room) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"time_limit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TimeLimit))
	}
	{
		const prefix string = ",\"max_num_players\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxNumPlayers))
	}
	{
		const prefix string = ",\"allow_anonymous\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.AllowAnonymous))
	}
	{
		const prefix string = ",\"field_size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.FieldSize).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Room) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceDomains(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Room) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceDomains(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Room) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceDomains(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Room) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceDomains(l, v)
}
func easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceDomains1(in *jlexer.Lexer, out *FieldSize) {
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
		case "width":
			out.Width = int32(in.Int32())
		case "height":
			out.Height = int32(in.Int32())
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
func easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceDomains1(out *jwriter.Writer, in FieldSize) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"width\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Width))
	}
	{
		const prefix string = ",\"height\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Height))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FieldSize) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceDomains1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FieldSize) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceDomains1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FieldSize) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceDomains1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FieldSize) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceDomains1(l, v)
}
