// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package users

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

func easyjson3e56fa40DecodeGithubComMtokuDiUsers(in *jlexer.Lexer, out *UserCreateResult) {
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
		case "Status":
			out.Status = string(in.String())
		case "Message":
			out.Message = string(in.String())
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
func easyjson3e56fa40EncodeGithubComMtokuDiUsers(out *jwriter.Writer, in UserCreateResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix[1:])
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"Message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserCreateResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e56fa40EncodeGithubComMtokuDiUsers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCreateResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e56fa40EncodeGithubComMtokuDiUsers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCreateResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e56fa40DecodeGithubComMtokuDiUsers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCreateResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e56fa40DecodeGithubComMtokuDiUsers(l, v)
}
func easyjson3e56fa40DecodeGithubComMtokuDiUsers1(in *jlexer.Lexer, out *APIUserCreateRequest) {
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
		case "UserID":
			out.UserID = string(in.String())
		case "Password":
			out.Password = string(in.String())
		case "Nickname":
			out.Nickname = string(in.String())
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
func easyjson3e56fa40EncodeGithubComMtokuDiUsers1(out *jwriter.Writer, in APIUserCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"Password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v APIUserCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e56fa40EncodeGithubComMtokuDiUsers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v APIUserCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e56fa40EncodeGithubComMtokuDiUsers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *APIUserCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e56fa40DecodeGithubComMtokuDiUsers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *APIUserCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e56fa40DecodeGithubComMtokuDiUsers1(l, v)
}
