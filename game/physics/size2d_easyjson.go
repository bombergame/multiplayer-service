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

func easyjson97fc65f9DecodeGithubComBombergameMultiplayerServiceGamePhysics(in *jlexer.Lexer, out *Size2D) {
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
			out.Width = Integer(in.Int32())
		case "height":
			out.Height = Integer(in.Int32())
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
func easyjson97fc65f9EncodeGithubComBombergameMultiplayerServiceGamePhysics(out *jwriter.Writer, in Size2D) {
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
func (v Size2D) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97fc65f9EncodeGithubComBombergameMultiplayerServiceGamePhysics(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Size2D) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97fc65f9EncodeGithubComBombergameMultiplayerServiceGamePhysics(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Size2D) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97fc65f9DecodeGithubComBombergameMultiplayerServiceGamePhysics(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Size2D) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97fc65f9DecodeGithubComBombergameMultiplayerServiceGamePhysics(l, v)
}