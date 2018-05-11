package prom

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var (
	_ Test   = (*Suite)(nil)
	_ Test   = (*Expectation)(nil)
	_ Test   = (List)(nil)
	_ Test   = (*BeforeFuncs)(nil)
	_ Test   = (*AfterFuncs)(nil)
	_ Result = (*baseResult)(nil)
	_ Result = (*RSWithNode)(nil)
)

type Test interface {
	Exec()
	run()
}

// Describe describe what you want to test.
func Describe(desc string, tc ...Test) Test {
	t := &Suite{Desc: desc}
	for _, v := range tc {
		switch e := v.(type) {
		case *BeforeFuncs:
			if t.BeforeFuncs != nil {
				t.BeforeFuncs.Funcs =
					append(t.BeforeFuncs.Funcs, e.Funcs...)
			} else {
				t.BeforeFuncs = e
			}
		case *AfterFuncs:
			if t.AfterFuncs != nil {
				t.AfterFuncs.Funcs =
					append(t.AfterFuncs.Funcs, e.Funcs...)
			} else {
				t.AfterFuncs = e
			}
		case *Suite:
			e.Parent = t
			t.Cases = append(t.Cases, e)
		case *Expectation:
			e.Parent = t
			t.Cases = append(t.Cases, e)
		}
	}
	return t
}

type List []Test

func (ls List) run() {}
func (ls List) Exec() {
	for _, v := range ls {
		v.Exec()
	}
}

type Suite struct {
	Parent      *Suite
	Desc        string
	BeforeFuncs *BeforeFuncs
	AfterFuncs  *AfterFuncs
	Cases       List

	FailedExpectations []*Expectation
	MarkedSKip         bool
	MarkedSkipMessage  string
	Duration           time.Duration
	Expectations       []*Expectation
}

func (s *Suite) FullName() string {
	var names []string
	p := s
	for p != nil {
		names = append(names, p.Desc)
		p = p.Parent
	}
	size := len(names)
	rvs := make([]string, size)
	for i := 0; i < size; i++ {
		rvs[i] = names[size-i-1]
	}
	return strings.Join(rvs, " ")
}

func (s *Suite) Exec() {
	start := time.Now()
	if s.BeforeFuncs != nil {
		s.BeforeFuncs.Exec()
	}
	defer func() {
		s.Duration = time.Now().Sub(start)
	}()
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(*Error); ok {
				if err.Pending {
					s.MarkedSKip = true
					s.MarkedSkipMessage = err.Message.Error()
				}
			}
		} else {
			if s.AfterFuncs != nil {
				s.AfterFuncs.Exec()
			}
		}

	}()
	for i := 0; i < len(s.Cases); i++ {
		v := s.Cases[i]
		switch e := v.(type) {
		case *Suite:
			e.Exec()
		case *Expectation:
			e.Exec()
		}
	}
}

func Exec(ts ...Test) Test {
	ls := List(ts)
	ls.Exec()
	return ls
}

// Skip marks test suite as skipped
func Skip(message string) {
	panic(&Error{Pending: true, Message: errors.New(message)})
}

func defaultResultFn() Result {
	return &baseResult{}
}

func (*Suite) run() {}

func It(desc string, fn func(Result)) Test {
	return &Expectation{Desc: desc, Func: fn}
}

type Result interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	FatalF(string, ...interface{})
	Errors() []error
}

type Expectation struct {
	Parent       *Suite
	Desc         string
	Func         func(Result)
	Passed       bool
	FailMessages []string
}

func (*Expectation) run() {}

func (e *Expectation) Exec() {
	defer func() {
		if ev := recover(); ev != nil {
			if err, ok := ev.(*Error); ok {
				if !err.Pending {
					e.Passed = false
					e.FailMessages = append(e.FailMessages, err.Message.Error())
				}
			}
		}
	}()
	rs := &baseResult{}
	if e.Func != nil {
		e.Func(rs)
	}
	errs := rs.Errors()
	if errs != nil {
		for _, v := range errs {
			e.FailMessages = append(e.FailMessages, v.Error())
		}
	} else {
		e.Passed = true
	}
}

type ResultInfo struct {
	Case         string   `json:"case"`
	Failed       bool     `json:"failed"`
	FailMessages []string `json:"fail_messages"`
}

type Error struct {
	Message error
	Pending bool
}

func (e *Error) Error() string {
	if e.Message != nil {
		return e.Error()
	}
	return ""
}

type baseResult struct {
	err []error
}

func (b *baseResult) Error(v ...interface{}) {
	b.err = append(b.err, errors.New(fmt.Sprint(v...)))
}
func (b *baseResult) Fatal(v ...interface{}) {
	panic(&Error{Message: errors.New(fmt.Sprint(v...))})
}

func (b *baseResult) Errorf(s string, v ...interface{}) {
	b.err = append(b.err, fmt.Errorf(s, v...))
}

func (b *baseResult) FatalF(s string, v ...interface{}) {
	panic(&Error{Message: fmt.Errorf(s, v...)})
}

func (b *baseResult) Errors() []error {
	return b.err
}

type Component struct {
	ID        string
	Component func() interface{}
	IsBody    bool
	Cases     List
}

func (c *Component) runIntegration() {}

type Integration interface {
	runIntegration()
}

// Node is an interface for retrieving a rendered Component node.
type Node interface {
	Node() *js.Object
}

type RSWithNode struct {
	baseResult
	Object *js.Object
}

func (rs *RSWithNode) Node() *js.Object {
	return rs.Object
}

// Render returns an integration test for non body Components. Use this to test
// Components that renders spans,div etc.
//
// NOTE: func()interface{} was supposed to be func()vecty.ComponentOrHTML this
// is a workaround, because importing vecty in this package will make it
// impossible to run the commandline tools since vecty only works with the
// browser.
func Render(desc string, c func() interface{}, cases ...Test) Integration {
	return &Component{
		ID: desc, Component: c, Cases: cases,
	}
}

// RenderBody is like Render but the Component is expected to be elem.Body
func RenderBody(desc string, c func() interface{}, cases ...Test) Integration {
	return &Component{
		ID: desc, Component: c, Cases: cases, IsBody: true,
	}
}

type BeforeFuncs struct {
	Funcs []func()
}

func (*BeforeFuncs) run() {}
func (b *BeforeFuncs) Exec() {
	for _, v := range b.Funcs {
		v()
	}
}

// Before is a list of functions that will be executed before the actual test
// suite is run.
func Before(fn ...func()) Test {
	return &BeforeFuncs{Funcs: fn}
}

type AfterFuncs struct {
	Funcs []func()
}

func (*AfterFuncs) run() {}
func (b *AfterFuncs) Exec() {
	for _, v := range b.Funcs {
		v()
	}
}

// After is a list of functions that will be executed after the actual test
// suite is run.
// You can use this to release resources/cleanup after the tests are done.
func After(fn ...func()) Test {
	return &AfterFuncs{Funcs: fn}
}
