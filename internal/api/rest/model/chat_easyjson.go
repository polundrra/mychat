// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	model "mychat/internal/repo/model"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9b8f5552DecodeMychatInternalApiRestModel(in *jlexer.Lexer, out *Chats) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Chats, 0, 0)
			} else {
				*out = Chats{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 model.Chat
			easyjson9b8f5552DecodeMychatInternalRepoModel(in, &v1)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9b8f5552EncodeMychatInternalApiRestModel(out *jwriter.Writer, in Chats) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			easyjson9b8f5552EncodeMychatInternalRepoModel(out, v3)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Chats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b8f5552EncodeMychatInternalApiRestModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Chats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b8f5552EncodeMychatInternalApiRestModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Chats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b8f5552DecodeMychatInternalApiRestModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Chats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b8f5552DecodeMychatInternalApiRestModel(l, v)
}
func easyjson9b8f5552DecodeMychatInternalRepoModel(in *jlexer.Lexer, out *model.Chat) {
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
		case "ID":
			out.ID = uint64(in.Uint64())
		case "Name":
			out.Name = string(in.String())
		case "Users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]uint32, 0, 16)
					} else {
						out.Users = []uint32{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v4 uint32
					v4 = uint32(in.Uint32())
					out.Users = append(out.Users, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
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
func easyjson9b8f5552EncodeMychatInternalRepoModel(out *jwriter.Writer, in model.Chat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Users {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Uint32(uint32(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson9b8f5552DecodeMychatInternalApiRestModel1(in *jlexer.Lexer, out *ChatReq) {
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
		case "Name":
			out.Name = string(in.String())
		case "Users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]uint32, 0, 16)
					} else {
						out.Users = []uint32{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v7 uint32
					v7 = uint32(in.Uint32())
					out.Users = append(out.Users, v7)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson9b8f5552EncodeMychatInternalApiRestModel1(out *jwriter.Writer, in ChatReq) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Users {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Uint32(uint32(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatReq) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b8f5552EncodeMychatInternalApiRestModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b8f5552EncodeMychatInternalApiRestModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatReq) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b8f5552DecodeMychatInternalApiRestModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b8f5552DecodeMychatInternalApiRestModel1(l, v)
}
func easyjson9b8f5552DecodeMychatInternalApiRestModel2(in *jlexer.Lexer, out *ChatMessagesReq) {
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
		case "ID":
			out.ID = uint64(in.Uint64())
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
func easyjson9b8f5552EncodeMychatInternalApiRestModel2(out *jwriter.Writer, in ChatMessagesReq) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatMessagesReq) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b8f5552EncodeMychatInternalApiRestModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatMessagesReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b8f5552EncodeMychatInternalApiRestModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatMessagesReq) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b8f5552DecodeMychatInternalApiRestModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatMessagesReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b8f5552DecodeMychatInternalApiRestModel2(l, v)
}
