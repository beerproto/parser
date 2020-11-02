package parser

import (
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/go-multierror"
	"reflect"
	"strings"
)

type ErrorStack struct {
	prefix string
	error error
}

func NewErrorStack(prefix string) *ErrorStack {
	return &ErrorStack{
		prefix: prefix,
	}
}


func (c *ErrorStack) Push(i interface{}) {
	var name string
	if reflect.TypeOf(i).Comparable() {
		name = reflect.TypeOf(i).Elem().Name()
	} else {
		name = reflect.TypeOf(i).Name()
	}

	c.prefix = fmt.Sprintf("%s/%s", c.prefix, name)
}

func (c *ErrorStack) getName(field reflect.StructField) string {
	if protobuf, ok := field.Tag.Lookup("protobuf"); ok {
		arr := strings.Split( protobuf, ",")
		var name string
		for _, s := range arr {
			if strings.HasPrefix(s, "json=") {
				return s[5:]
			}
			if strings.HasPrefix(s, "name=") {
				name = s[5:]
			}
		}
		return name
	}

	if json, ok := field.Tag.Lookup("json"); ok {
		i := strings.Index(json, ",")
		return json[0:i]
	}

	return field.Name
}

func (c *ErrorStack) PushMethod(index int,  field reflect.StructField) {
	name := c.getName(field)
	if index >= 0 {
		c.prefix = fmt.Sprintf("%s/%s[%v]", c.prefix, name, index)
	} else {
		c.prefix = fmt.Sprintf("%s/%s", c.prefix, name)
	}
}

func (c *ErrorStack) Pop() {
	index := strings.LastIndex(c.prefix, "/")
	if index >= 0 {
		c.prefix = c.prefix[0:index]
	}
}

func (c *ErrorStack) path() string {
	return c.prefix
}

func (c *ErrorStack) AppendError(err error) {
	c.error = multierror.Append(Prefix(err, c.prefix), c.error)
}

func Prefix(err error, prefix string) error {
	if err == nil {
		return nil
	}

	format := fmt.Sprintf("%s.{{err}}", prefix)
	switch err := err.(type) {
	case *multierror.Error:
		if err == nil {
			err = new(multierror.Error)
		}

		for i, e := range err.Errors {
			err.Errors[i] = e
		}

		return err
	default:
		return errwrap.Wrapf(format, err)
	}
}

func (c *ErrorStack) Errors() error {
	if c.error == nil {
		return nil
	}
	if result, ok := c.error.(*multierror.Error); ok {
		if result != nil {
			result.ErrorFormat = func(es []error) string {
				if len(es) == 1 {
					return fmt.Sprintf("1 error occurred:\n\t* %s\n\n", es[0])
				}

				points := make([]string, len(es))
				for i, err := range es {
					points[i] = fmt.Sprintf("* %s", err)
				}

				return fmt.Sprintf(
					"%d errors occurred:\n\t%s\n\n",
					len(es), strings.Join(points, "\n\t"))
			}
		}
	}

	return c.error
}
