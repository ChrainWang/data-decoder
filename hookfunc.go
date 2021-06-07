package decoder

import (
	"encoding/json"
	"reflect"
	"strings"
)

func DecodeStructuredString(ctx *DecodeContext, targetType reflect.Type) error {
	if targetKind := targetType.Kind(); ctx.Data == nil ||
		(targetKind != reflect.Map && targetKind != reflect.Struct) ||
		reflect.TypeOf(ctx.Data).Kind() != reflect.String {
		return nil
	}

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(ctx.Data.(string)), &m); err != nil {
		return err
	}

	ctx.Data = m
	return nil
}

func UnmarshalMapToStruct(ctx *DecodeContext, targetType reflect.Type) error {
	if ctx.Data != nil {
		return unmarshalMapToStruct(ctx, targetType)
	}

	if ctx.QueryOption == QO_NO {
		return nil
	}

	m := make(map[string]interface{})
	for i := 0; i != targetType.NumField(); i++ {
		field := targetType.Field(i)
		if field.Anonymous {
			continue
		}
		fieldName := field.Name
		if firstRune := fieldName[0]; firstRune < byte('A') || firstRune > byte('Z') {
			continue
		}
		fieldName = strings.ToLower(field.Name)
		queryKey := strings.Join([]string{ctx.QueryKey, fieldName}, ".")
		fieldDecodeCtx := &DecodeContext{
			QueryKey:    queryKey,
			QueryOption: QueryOptionType(field.Tag.Get("query-data")),
		}
		m[fieldName] = fieldDecodeCtx
	}
	ctx.Data = m
	return nil
}

func unmarshalMapToStruct(ctx *DecodeContext, targetType reflect.Type) error {
	if dataType := reflect.TypeOf(ctx.Data); dataType.Kind() != reflect.Map || targetType.Kind() != reflect.Struct {
		return nil
	}

	m := make(map[string]interface{})
	dataVal := reflect.ValueOf(ctx.Data)
	for i := 0; i != targetType.NumField(); i++ {
		// get each field from target type
		field := targetType.Field(i)
		fieldName := strings.ToLower(field.Name)
		// key value used to get field value from decode context
		mapKeyVal := reflect.ValueOf(fieldName)
		// envalue decode context for field
		fieldDecodeCtx := &DecodeContext{
			QueryKey: strings.Join([]string{ctx.QueryKey, fieldName}, "."),
		}
		dataVal := dataVal.MapIndex(mapKeyVal)
		if dataVal.IsValid() {
			fieldDecodeCtx.Data = dataVal.Interface()
		}
		// set up query option for field
		if ctx.QueryOption == QO_NO {
			fieldDecodeCtx.QueryOption = QO_NO
		} else {
			fieldDecodeCtx.QueryOption = QueryOptionType(field.Tag.Get("query-data"))
		}
		m[fieldName] = fieldDecodeCtx
	}
	ctx.Data = m
	return nil
}
