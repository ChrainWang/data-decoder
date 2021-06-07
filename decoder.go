package decoder

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Decoder struct {
	dp DataProvider
}

func NewDecoder(dp DataProvider) *Decoder {
	d := &Decoder{
		dp: dp,
	}
	return d
}

func (d *Decoder) Decode(key string, output interface{}) error {
	decoder, err := d.getDecoder(output)
	if err != nil {
		return err
	}
	ctx := &DecodeContext{
		QueryKey:    key,
		QueryOption: QO_ALWAYS,
	}
	return decoder.Decode(ctx)
}

func (d *Decoder) getDecoder(output interface{}) (*mapstructure.Decoder, error) {
	c := &mapstructure.DecoderConfig{
		Result: output,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ComposeDecodeHookFunc(
				d.queryData,
				DecodeStructuredString,
				UnmarshalMapToStruct,
			),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToIPHookFunc(),
			mapstructure.StringToIPNetHookFunc(),
		),
	}
	return mapstructure.NewDecoder(c)
}

func (d *Decoder) queryDataHookFunc() mapstructure.DecodeHookFunc {
	return ComposeDecodeHookFunc(d.queryData)
}

func (d *Decoder) queryData(ctx *DecodeContext, targetType reflect.Type) error {
	switch ctx.QueryOption {
	case QO_ALWAYS:
		ctx.Data = d.dp.Get(ctx.QueryKey)
	case QO_AUTO:
		if ctx.Data == nil && ctx.QueryKey != "" {
			ctx.Data = d.dp.Get(ctx.QueryKey)
		}
	case QO_NO, QO_BYPASS:
	default:
		return fmt.Errorf("invalid value for field tag 'query', [%s]", ctx.QueryOption)
	}
	return nil
}
