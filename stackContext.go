package parser

import (
	"context"
)

type stackContext struct {}

func StackFromContext(ctx context.Context)  (*ErrorStack) {
	s, _ := ctx.Value(stackContext{}).(*ErrorStack)
	return s
}

func StackToContext(ctx context.Context) (context.Context, *ErrorStack) {
	var path string
	parentStack := StackFromContext(ctx)
	if parentStack == nil {
		path = "/"
	} else {
		path = parentStack.path()
	}
	stack := NewErrorStack(path)
	return context.WithValue(ctx, stackContext{}, stack), stack
}