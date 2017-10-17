package errors

import (
	"fmt"
	"encoding/json"
	"bytes"
	"strings"
)

const (
	Unknown = iota
	Info
	Warn
	Err
	Critical
)

type level int

func (l level) String() string {
	switch l {
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Err:
		return "Err"
	case Critical:
		return "Critical"
	}
	return "Unknown"
}

func (l *level) Set(i int) {
	h := level(i)
	l = &h
}

func (l level) MarshalJSON()([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, l.String())), nil
}

func (l level) UmmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.TrimPrefix(s, `""`)
	s = strings.TrimSuffix(s, `""`)

	switch s {
	case "Info":
		l.Set(Info)
	case "Warn":
		l.Set(Warn)
	case "Err":
		l.Set(Err)
	case "Critical":
		l.Set(Critical)
	default:
		l.Set(Unknown)
	}

	return nil
}

type List struct {
	Errors map[int]*Error
	Index  map[string]int
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(e *Error) error {
	if existingError, ok := l.Errors[e.Code]; ok {
		return New(fmt.Sprintf("error code %d is already existed: %v", e.Code, existingError))
	}
	l.Errors[e.Code] = e
	if e.Name != "" {
		l.Index[e.Name] = e.Code
	}

	return nil
}

func (l *List) Get(name string) *Error {
	e, ok := l.Errors[l.Index[name]]
	if !ok {
		return New("unknown error", -1, Unknown)
	}

	return e
}

type Error struct {
	Msg   string `json:"msg"`
	Code  int `json:"code"`
	Level level `json:"level"`
	Name  string `json:"name"`
}

func New(msg string, options ...int) *Error {
	if len(options) == 1 {
		return &Error{
			Msg:  msg,
			Code: options[0],
		}
	} else if len(options) == 2 {
		return &Error{
			Msg:   msg,
			Code:  options[0],
			Level: level(options[1]),
		}
	}
	return &Error{Msg: msg}
}

func (e *Error) SetName(n string) *Error {
	e.Name = n
	return e
}

func (e *Error) Error() string {
	return e.Msg
}

func (e Error) String() string {
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(e); err != nil {
		return ""
	}
	return buffer.String()
}
