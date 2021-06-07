# data-decoder
A data decoding util based on mapstructure package
The main target of this package is to solve the problem when requiring nested values from *viper*. *Viper* doesn't process **environment values** as structured value, this causes problem when unmarshalling value to structs.
*Decoder* requires value for nested struct fields particularly if it's not found from the data map required for it's parent struct. **Environment values** are used if the value is not defined in map-typed data in config files.

# Examples
## Environments
```
LOGGING_FILE_PATH=decoder.log
```

## *viper* vs *data-decoder* 
### Contents in yaml file
```
logging:
  level: debug
```
### Logging Config Model
```
type LoggingConfigModel1 struct {
	Level string
	File  struct {
		Path string
	}
}
```
### Using viper
#### Code
```
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
```
#### Result
```
logging.file.path: decoder.log
cfgModel: {debug {}}
```
### Using data-decoder
#### Code
```
v := viper.New()
cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "decode_logging_config", "config.yaml")
v.SetConfigFile(cfgFilePath)
v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
v.AutomaticEnv()
v.MergeInConfig()

d := decoder.NewDecoder(v)

cfgModel := LoggingConfigModel1{}
fmt.Println("logging.file.path:", v.Get("logging.file.path"))
d.Decode("logging", &cfgModel)
fmt.Println("cfgModel:", cfgModel)
```
#### Result
```
logging.file.path: decoder.log
cfgModel: {debug {decoder.log}}
```

## Using query options
### Disable nested query
#### Models
```
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
```
#### Code
```
v := viper.New()
cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "decode_logging_config", "config.yaml")
v.SetConfigFile(cfgFilePath)
v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
v.AutomaticEnv()
v.MergeInConfig()

d := decoder.NewDecoder(v)

cfgModel2 := LoggingConfigModel2{}
fmt.Println(v.Get("logging.file.path"))
d.Decode("logging", &cfgModel2)
fmt.Println("cfgModel2:", cfgModel2)

cfgModel3 := LoggingConfigModel3{}
fmt.Println(v.Get("logging.file.path"))
d.Decode("logging", &cfgModel3)
fmt.Println("cfgModel3:", cfgModel3)
```
#### Result
```
decoder.log
cfgModel2: {debug {}}
decoder.log
cfgModel3: {debug {}}
```

### Use *bypass* together with *viper alias*
#### Environment
```
LOGGING_CONFIG_FILE=decoder.log
```
#### Model
```
type LoggingConfigModel4 struct {
	Level string
	File  struct {
		Path string
	} `query-option:"bypass"`
}
```
#### Code
```
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
d.Decode("logging", &cfgModel4)
fmt.Println("cfgModel4:", cfgModel4)
```
#### Result
```
decoder.log
cfgModel4: {debug {decoder.log}}
```

