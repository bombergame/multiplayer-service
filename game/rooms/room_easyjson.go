// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package rooms

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

func easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceGameRooms(in *jlexer.Lexer, out *TickerMessageData) {
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
		case "value":
			out.Value = float64(in.Float64())
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
func easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceGameRooms(out *jwriter.Writer, in TickerMessageData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TickerMessageData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceGameRooms(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TickerMessageData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceGameRooms(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TickerMessageData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceGameRooms(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TickerMessageData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceGameRooms(l, v)
}
func easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceGameRooms1(in *jlexer.Lexer, out *RoomMessageData) {
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
		case "title":
			out.Title = string(in.String())
		case "state":
			out.State = string(in.String())
		case "players":
			if in.IsNull() {
				in.Skip()
				out.Players = nil
			} else {
				in.Delim('[')
				if out.Players == nil {
					if !in.IsDelim(']') {
						out.Players = make([]int64, 0, 8)
					} else {
						out.Players = []int64{}
					}
				} else {
					out.Players = (out.Players)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int64
					v1 = int64(in.Int64())
					out.Players = append(out.Players, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceGameRooms1(out *jwriter.Writer, in RoomMessageData) {
	out.RawByte('{')
	first := true
	_ = first
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
		const prefix string = ",\"players\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Players == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Players {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v3))
			}
			out.RawByte(']')
		}
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
func (v RoomMessageData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceGameRooms1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoomMessageData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson426b19f9EncodeGithubComBombergameMultiplayerServiceGameRooms1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoomMessageData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceGameRooms1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoomMessageData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson426b19f9DecodeGithubComBombergameMultiplayerServiceGameRooms1(l, v)
}