package decoder_test

import (
	"os"
	"path/filepath"
	"testing"

	decoder "github.com/chrainwang/data-decoder"
	"github.com/spf13/viper"
)

func TestDecodeStringToMap(t *testing.T) {
	v := viper.New()
	cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "fixture", "test-data.yaml")
	v.SetConfigFile(cfgFilePath)
	if err := v.ReadInConfig(); err != nil {
		t.Fatal(err)
	}
	var m map[string]interface{}
	d := decoder.NewDecoder(v)
	if err := d.Decode("testmap", &m); err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}

func TestDecodeToStruct(t *testing.T) {
	v := viper.New()
	cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "fixture", "test-data.yaml")
	v.SetConfigFile(cfgFilePath)
	if err := v.ReadInConfig(); err != nil {
		t.Fatal(err)
	}
	d := decoder.NewDecoder(v)
	var output TestMapDataModel
	if err := d.Decode("testmap", &output); err != nil {
		t.Fatal(err)
	}
	t.Log(output)
}
