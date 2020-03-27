package errors

import (
	"errors"
	"fmt"
)

/*
官方博客中出现的一个笔误：
原文链接：https://blog.golang.org/go1.13-errors
原文如下：
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

if e, ok := err.(*NotFoundError); ok {
    // e.Name wasn't found
}

其中“e.Name wasn't found” 错误，实际应该是“e.Name was found”

修复该笔误的PR：https://github.com/golang/blog/pull/33
*/

func ExampleNotFoundError_Error() {
	err1 := errors.New("this is not NotFoundError")
	if e, ok := err1.(*NotFoundError); ok {
		fmt.Printf("err1 is a NotFoundError, e.Name=%s", e.Name)
	}

	err2 := NewNotFoundError("err2")

	if e, ok := err2.(*NotFoundError); ok {
		fmt.Printf("err2 is a NotFoundError, e.Name=%s", e.Name)
	}

	// OutPut:
	// err2 is a NotFoundError, e.Name=err2
}
