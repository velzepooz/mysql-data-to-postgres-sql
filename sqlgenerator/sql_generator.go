package sqlgenerator

import (
	"fmt"
	"strings"
)

func ConvertDataToSQL(csvData [][]string, tableName string, tableSchema []map[string]string) string {
	mappedData := *mapDataToDBTableFields(csvData, tableSchema)
	tableSchemaFields := getSchemaFields(&tableSchema)

	return generateSQLScript(&tableSchema, tableName, &tableSchemaFields, &mappedData)
}

func generateSQLScript(schema *[]map[string]string, tableName string, schemaFields *[]string, mappedData *[]map[string]string) (script string) {
	script = fmt.Sprintf("INSERT INTO \"%s\"\n(%s) VALUES\n", tableName, strings.Join(*schemaFields, ", "))

	for lineNumber, line := range *mappedData {
		for fieldNumber, fields := range *schema {
			for field, dataType := range fields {
				if fieldNumber == 0 {
					script += "("
				}

				data := line[field]

				if data == "NULL" {
					script += data
				} else {
					switch dataType {
					case "string":
						script += "'" + data + "'"
					case "number":
						script += line[field]
					case "boolean":
						if data == "0" || data == "1" {
							script += BooleanConversion[data]
						} else {
							script += data
						}
					}
				}

				if fieldNumber == len(*schema)-1 {
					if lineNumber == len(*mappedData)-1 {
						script += ");"
					} else {
						script += "),\n"
					}
				} else {
					script += ", "
				}
			}
		}
	}

	return
}

func getSchemaFields(schema *[]map[string]string) (fields []string) {
	for _, schemaFields := range *schema {
		for key := range schemaFields {
			fields = append(fields, key)
		}
	}

	return
}

func mapDataToDBTableFields(csvData [][]string, schema []map[string]string) *[]map[string]string {
	var table []map[string]string

	for index, line := range csvData {
		// exclude csv headers
		if index == 0 {
			continue
		}

		tempMap := map[string]string{}

		for index, value := range schema {
			for key := range value {
				tempMap[key] = line[index]
			}
		}

		table = append(table, tempMap)
	}

	return &table
}
