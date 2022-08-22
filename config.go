package main

type FileConfig struct {
	sourceFileName string
	schema         []map[string]string
}

// Table name as a key and source filename and schema as a value. Schema with mapping field name and it's type by appearing
// in source file order
var ConverterConfig = map[string]FileConfig{
	"example": {
		sourceFileName: "example.csv",
		schema: []map[string]string{
			{"created_at": "string"},
			{"updated_at": "string"},
			{"id": "number"},
			{"name": "string"},
			{"enum_value": "string"},
			{"email": "string"},
			{"confirmed": "boolean"},
		},
	},
}

var OutputDirName = "output"
var InputDirName = "input"
