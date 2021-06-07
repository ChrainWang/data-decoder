package decoder_test

import (
	"strings"
	"testing"

	decoder "github.com/chrainwang/data-decoder"
	"github.com/spf13/viper"
)

type SubDataModel struct {
	Field1 string
	Field2 string
	Field3 int
}

type DataModel struct {
	Fields []SubDataModel
}

func TestUnmarshalStructSlice(t *testing.T) {
	v := viper.New()
	cfgFilePath := "config.yaml"
	v.SetConfigFile(cfgFilePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.ReadInConfig()

	var result []SubDataModel
	d := decoder.NewDecoder(v)
	d.Decode("data.fields", &result)
	t.Log(result)
}

func TestNestedStructSlice(t *testing.T) {
	v := viper.New()
	cfgFilePath := "config.yaml"
	v.SetConfigFile(cfgFilePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.ReadInConfig()

	var result DataModel
	d := decoder.NewDecoder(v)
	d.Decode("data", &result)
	t.Log(result)
}
