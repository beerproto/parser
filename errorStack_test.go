package parser

import (
	"fmt"
	"github.com/beerproto/beerjson.go"
	"github.com/hashicorp/go-multierror"
	"reflect"
	"testing"
)

func TestErrorStack_push(t *testing.T) {

	stack := NewErrorStack("/")

	beer := &beerjson.Beerjson{
		Mashes: []beerjson.MashProcedureType{},
	}
	stack.Push(beer)

	stack.AppendError(fmt.Errorf("version is required"))

	list := []int{0, 1}
	field, _ := reflect.TypeOf(beer).Elem().FieldByName("Mashes")
	for _, i := range list {
		stack.PushMethod(i, field)

		stack.AppendError(fmt.Errorf("name is required"))

		stack.Pop()
	}

	err := stack.Errors()

	if err == nil {
		t.Error(fmt.Errorf("expected errors"))
	}

	if es, ok := err.(*multierror.Error); ok {
		if es == nil {
			t.Error(fmt.Errorf("expected multierror"))
		}
		if len(es.Errors) != 3 {
			t.Error(fmt.Errorf("expected 3 errors"))
		}
	}

}
