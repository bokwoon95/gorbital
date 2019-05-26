package erro

import (
	"fmt"
	"io"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

// Wrap lorem ipsum
func Wrap(err error) error {
	pc, filename, linenr, _ := runtime.Caller(1)
	return errors.Wrapf(err, "\n\nerror in function[%s] file[%s] line[%d]", runtime.FuncForPC(pc).Name(), filename, linenr)
}

// WrapX lorem ipsum
func WrapX(err *error) {
	if *err != nil {
		pc, filename, linenr, _ := runtime.Caller(1)
		*err = errors.Wrapf(*err, "\n\nerror in function[%s] file[%s] line[%d]", runtime.FuncForPC(pc).Name(), filename, linenr)
	}
}

// WrapDebug lorem ipsum
func WrapDebug(err error, vars ...interface{}) error {
	pc, filename, linenr, _ := runtime.Caller(1)
	return errors.Wrapf(err, `

error in function[%s] file[%s] line[%d]
↓↓↓ Debug variables below ↓↓↓
--------------------------------------------------------------------------------
%s--------------------------------------------------------------------------------
cause of error`, runtime.FuncForPC(pc).Name(), filename, linenr, spew.Sdump(vars...))
}

// WrapDebugX lorem ipsum
func WrapDebugX(err *error, vars ...interface{}) {
	if *err != nil {
		pc, filename, linenr, _ := runtime.Caller(1)
		*err = errors.Wrapf(*err, `

error in function[%s] file[%s] line[%d]
↓↓↓ Debug variables below ↓↓↓
--------------------------------------------------------------------------------
%s--------------------------------------------------------------------------------
cause of error`, runtime.FuncForPC(pc).Name(), filename, linenr, spew.Sdump(vars...))
	}
}

// Dump lorem ipsum
func Dump(w io.Writer, err error) {
	fmt.Fprintf(w, err.Error())
}

// DumpVars lorem ipsum
func DumpVars(w io.Writer, vars ...interface{}) {
	pc, filename, linenr, _ := runtime.Caller(1)
	fmt.Fprintf(w, `

function[%s] file[%s] line[%d]
↓↓↓ Debug variables below ↓↓↓
--------------------------------------------------------------------------------
%s--------------------------------------------------------------------------------


`, runtime.FuncForPC(pc).Name(), filename, linenr, spew.Sdump(vars...))
}

// LogFile will append a pretty-printed variable into a file (relative to
// project root). If the file does not exist, one will be created. It takes in
// an error pointer which is modified inplace so that any errors upon writing
// can be added to the existing error stack. If you do not wish to pass in any
// error variable, you can pass in nil.
func LogFile(filename string, variable interface{}, err *error) {
}

// WrapReturn lorem ipsum
func WrapReturn(err error) func() error {
	return func() error {
		if err != nil {
			pc, filename, linenr, _ := runtime.Caller(0)
			return errors.Wrapf(err, "\n\nerror in function[%s] file[%s] line[%d]", runtime.FuncForPC(pc).Name(), filename, linenr)
		}
		return nil
	}
}

// WrapReturnMulti lorem ipsum
func WrapReturnMulti(err *error) func() {
	return func() {
		if *err != nil {
			pc, filename, linenr, _ := runtime.Caller(0)
			*err = errors.Wrapf(*err, "\n\nerror in function[%s] file[%s] line[%d]", runtime.FuncForPC(pc).Name(), filename, linenr)
			return
		}
		return
	}
}
