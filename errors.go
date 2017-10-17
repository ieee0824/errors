package errors

import "fmt"

const (
	Unknown = iota
	Info
	Warn
	Error
	Critical
)

type List struct {
	Errors map[int]*Errors
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
	Msg   string
	Code  int
	Level int
	Name  string
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
			Level: options[1],
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
