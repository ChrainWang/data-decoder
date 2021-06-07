package decoder_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	decoder "github.com/chrainwang/data-decoder"
	"github.com/spf13/viper"
)

type LoggingConfigModel1 struct {
	Level string
	File  struct {
		Path string
	}
}

type LoggingConfigModel2 struct {
	Level string
	File  struct {
		Path string `query-option:"no"`
	}
}

type LoggingConfigModel3 struct {
	Level string
	File  struct {
		Path string
	} `query-option:"no"`
}

type LoggingConfigModel4 struct {
	Level string
	File  struct {
		Path string
	} `query-option:"bypass"`
}

func TestDecodeLoggingConfigWithViper(t *testing.T) {
	v := viper.New()
	cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "decode_logging_config", "config.yaml")
	v.SetConfigFile(cfgFilePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.MergeInConfig()
	cfgModel := LoggingConfigModel1{}
	fmt.Println("logging.file.path:", v.Get("logging.file.path"))
	v.UnmarshalKey("logging", &cfgModel)
	fmt.Println("cfgModel:", cfgModel)
}

func TestDecodeLoggingConfigWithDataDecoder(t *testing.T) {
	v := viper.New()
	cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "decode_logging_config", "config.yaml")
	v.SetConfigFile(cfgFilePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.MergeInConfig()

	d := decoder.NewDecoder(v)

	cfgModel := LoggingConfigModel1{}
	fmt.Println("logging.file.path:", v.Get("logging.file.path"))
	if err := d.Decode("logging", &cfgModel); err != nil {
		t.Fatal(err)
	}
	fmt.Println("cfgModel:", cfgModel)
}

func TestDecodeLoggingConfigNoQuery(t *testing.T) {
	v := viper.New()
	cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "decode_logging_config", "config.yaml")
	v.SetConfigFile(cfgFilePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.MergeInConfig()

	d := decoder.NewDecoder(v)

	cfgModel2 := LoggingConfigModel2{}
	fmt.Println(v.Get("logging.file.path"))
	if err := d.Decode("logging", &cfgModel2); err != nil {
		t.Fatal(err)
	}
	fmt.Println("cfgModel2:", cfgModel2)

	cfgModel3 := LoggingConfigModel3{}
	fmt.Println(v.Get("logging.file.path"))
	if err := d.Decode("logging", &cfgModel3); err != nil {
		t.Fatal(err)
	}
	fmt.Println("cfgModel3:", cfgModel3)
}

func TestDecodeLoggingConfigBypass(t *testing.T) {
	v := viper.New()
	cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "decode_logging_config", "config.yaml")
	v.SetConfigFile(cfgFilePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.MergeInConfig()

	v.RegisterAlias("logging.file.path", "logging.file")
	d := decoder.NewDecoder(v)
	cfgModel4 := LoggingConfigModel4{}
	fmt.Println(v.Get("logging.file.path"))
	if err := d.Decode("logging", &cfgModel4); err != nil {
		t.Fatal(err)
	}
	fmt.Println("cfgModel4:", cfgModel4)

}
