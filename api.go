package decoder

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type QueryOptionType string

const (
	QO_ALWAYS QueryOptionType = "always" // always query data from data provider
	QO_AUTO                   = ""       // query data if data is nil
	QO_NO                     = "no"     // never query data from data provider for current field node and its children
	QO_BYPASS                 = "bypass" // never query data for the current field node
)

type DataProvider interface {
	Get(key string) interface{}
}

type DecodeHookFunc func(*DecodeContext, reflect.Type) error

func ComposeDecodeHookFunc(fs ...DecodeHookFunc) mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data interface{}) (interface{}, error) {
		ctx, ok := data.(*DecodeContext)
		if !ok {
			return data, nil
		}

		for _, hook := range fs {
			if err := hook(ctx, t); err != nil {
				return nil, err
			}
			if ctx.done {
				return ctx.Data, nil
			}
		}
		return ctx.Data, nil
	}
}
