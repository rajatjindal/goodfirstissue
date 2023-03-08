// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package twitter

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

func easyjsonE23b537bDecodeGithubComRajatjindalGoodfirstissueTwitter(in *jlexer.Lexer, out *Tokens) {
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
		case "ConsumerKey":
			out.ConsumerKey = string(in.String())
		case "ConsumerToken":
			out.ConsumerToken = string(in.String())
		case "Token":
			out.Token = string(in.String())
		case "TokenSecret":
			out.TokenSecret = string(in.String())
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
func easyjsonE23b537bEncodeGithubComRajatjindalGoodfirstissueTwitter(out *jwriter.Writer, in Tokens) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ConsumerKey\":"
		out.RawString(prefix[1:])
		out.String(string(in.ConsumerKey))
	}
	{
		const prefix string = ",\"ConsumerToken\":"
		out.RawString(prefix)
		out.String(string(in.ConsumerToken))
	}
	{
		const prefix string = ",\"Token\":"
		out.RawString(prefix)
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"TokenSecret\":"
		out.RawString(prefix)
		out.String(string(in.TokenSecret))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Tokens) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE23b537bEncodeGithubComRajatjindalGoodfirstissueTwitter(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tokens) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE23b537bEncodeGithubComRajatjindalGoodfirstissueTwitter(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tokens) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE23b537bDecodeGithubComRajatjindalGoodfirstissueTwitter(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tokens) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE23b537bDecodeGithubComRajatjindalGoodfirstissueTwitter(l, v)
}
func easyjsonE23b537bDecodeGithubComRajatjindalGoodfirstissueTwitter1(in *jlexer.Lexer, out *Client) {
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
func easyjsonE23b537bEncodeGithubComRajatjindalGoodfirstissueTwitter1(out *jwriter.Writer, in Client) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Client) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE23b537bEncodeGithubComRajatjindalGoodfirstissueTwitter1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Client) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE23b537bEncodeGithubComRajatjindalGoodfirstissueTwitter1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Client) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE23b537bDecodeGithubComRajatjindalGoodfirstissueTwitter1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Client) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE23b537bDecodeGithubComRajatjindalGoodfirstissueTwitter1(l, v)
}