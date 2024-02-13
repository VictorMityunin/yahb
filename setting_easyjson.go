// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package yahb

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

func easyjson1f3eed6DecodeGithubComVictorMityuninYahb(in *jlexer.Lexer, out *Setting) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "currency":
			out.Currency = string(in.String())
		case "windowSize":
			if in.IsNull() {
				in.Skip()
				out.WinSize = nil
			} else {
				if out.WinSize == nil {
					out.WinSize = new(Size)
				}
				easyjson1f3eed6DecodeGithubComVictorMityuninYahb1(in, out.WinSize)
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
func easyjson1f3eed6EncodeGithubComVictorMityuninYahb(out *jwriter.Writer, in Setting) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Currency != "" {
		const prefix string = ",\"currency\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Currency))
	}
	if in.WinSize != nil {
		const prefix string = ",\"windowSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson1f3eed6EncodeGithubComVictorMityuninYahb1(out, *in.WinSize)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Setting) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1f3eed6EncodeGithubComVictorMityuninYahb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Setting) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1f3eed6EncodeGithubComVictorMityuninYahb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Setting) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1f3eed6DecodeGithubComVictorMityuninYahb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Setting) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1f3eed6DecodeGithubComVictorMityuninYahb(l, v)
}
func easyjson1f3eed6DecodeGithubComVictorMityuninYahb1(in *jlexer.Lexer, out *Size) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "width":
			out.Width = int(in.Int())
		case "height":
			out.Height = int(in.Int())
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
func easyjson1f3eed6EncodeGithubComVictorMityuninYahb1(out *jwriter.Writer, in Size) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Width != 0 {
		const prefix string = ",\"width\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Width))
	}
	if in.Height != 0 {
		const prefix string = ",\"height\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Height))
	}
	out.RawByte('}')
}
