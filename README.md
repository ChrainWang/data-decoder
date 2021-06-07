# data-decoder
A data decoding util based on mapstructure package
The main target of this package is to solve the problem when requiring nested values from *viper*. *Viper* doesn't process **environment value**s as structured value, this causes problem when unmarshalling value to structs.
*Decoder* requires value for nested struct fields particularly if it's not found from the data map required for it's parent struct. **Environment value**s are used if the value is not defined in map-typed data in config files.

# Example
## Environments
```
LOGGING_FILE_PATH=decoder.log
```
## Contents in yaml file
```
logging:
  level: debug
```
## Logging Config Model
```
type LoggingConfig struct {
    Level string
    File struct{
        Path string
    }
}
```
## Using viper
### Code
```
v := viper.New()
cfgFilePath := "Config File Path"
v.SetConfigFile(cfgFilePath)
v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
v.AutomaticEnv()
v.MergeInConfig()
cfgModel := LoggingConfigModel{}
fmt.Println(v.Get("logging.file.path"))
v.UnmarshalKey("logging", &cfgModel)
fmt.Println(cfgModel)
```
### Result
```
decoder.log
{debug {}}
```
## Using data-decoder
### Code
```
v := viper.New()
cfgFilePath := filepath.Join(os.Getenv("PROJECT_PATH"), "test", "fixture", "logging_config.yaml")
v.SetConfigFile(cfgFilePath)
v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
v.AutomaticEnv()
v.MergeInConfig()

d := decoder.NewDecoder(v)

cfgModel := LoggingConfigModel{}
fmt.Println(v.Get("logging.file.path"))
d.Decode("logging", &cfgModel)
fmt.Println(cfgModel)
```
### Result
```
decoder.log
{debug {decoder.log}}
```
