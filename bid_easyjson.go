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

func easyjson31527abDecodeGithubComVictorMityuninYahb(in *jlexer.Lexer, out *Bid) {
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
		case "displayUrl":
			out.DisplayUrl = string(in.String())
		case "displayCode":
			out.DisplayCode = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "cpm":
			out.Cpm = float64(in.Float64())
		case "currency":
			out.Currency = string(in.String())
		case "placementId":
			out.PlacementId = string(in.String())
		case "codeType":
			out.CodeType = BidCodeType(in.String())
		case "size":
			if in.IsNull() {
				in.Skip()
				out.Size = nil
			} else {
				if out.Size == nil {
					out.Size = new(Size)
				}
				(*out.Size).UnmarshalEasyJSON(in)
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
func easyjson31527abEncodeGithubComVictorMityuninYahb(out *jwriter.Writer, in Bid) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DisplayUrl != "" {
		const prefix string = ",\"displayUrl\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.DisplayUrl))
	}
	if in.DisplayCode != "" {
		const prefix string = ",\"displayCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DisplayCode))
	}
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if in.Cpm != 0 {
		const prefix string = ",\"cpm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Cpm))
	}
	if in.Currency != "" {
		const prefix string = ",\"currency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Currency))
	}
	if in.PlacementId != "" {
		const prefix string = ",\"placementId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PlacementId))
	}
	if in.CodeType != "" {
		const prefix string = ",\"codeType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CodeType))
	}
	if in.Size != nil {
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Size).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bid) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson31527abEncodeGithubComVictorMityuninYahb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bid) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson31527abEncodeGithubComVictorMityuninYahb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bid) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson31527abDecodeGithubComVictorMityuninYahb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bid) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson31527abDecodeGithubComVictorMityuninYahb(l, v)
}
