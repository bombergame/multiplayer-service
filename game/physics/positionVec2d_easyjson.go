// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package physics

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

func easyjsonD57d4c9DecodeGithubComBombergameMultiplayerServiceGamePhysics(in *jlexer.Lexer, out *PositionVec2D) {
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
		case "x":
			out.X = Coordinate(in.Int32())
		case "y":
			out.Y = Coordinate(in.Int32())
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
func easyjsonD57d4c9EncodeGithubComBombergameMultiplayerServiceGamePhysics(out *jwriter.Writer, in PositionVec2D) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Y))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionVec2D) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD57d4c9EncodeGithubComBombergameMultiplayerServiceGamePhysics(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionVec2D) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD57d4c9EncodeGithubComBombergameMultiplayerServiceGamePhysics(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionVec2D) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD57d4c9DecodeGithubComBombergameMultiplayerServiceGamePhysics(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionVec2D) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD57d4c9DecodeGithubComBombergameMultiplayerServiceGamePhysics(l, v)
}
