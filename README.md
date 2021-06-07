# data-decoder
A data decoding util based on mapstructure package
The main target of this package is to solve the problem when requiring nested values from viper. viper doesn't process environment values as structured value, this causes problem when unmarshalling value to structs.
The decoder requires value for nested struct fields particularly if it's not found from the data map required for it's parent struct. Environment values are used if the value is not devined in map-typed data in config files.
