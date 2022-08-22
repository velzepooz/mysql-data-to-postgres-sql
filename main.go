package main

import (
	"github.com/velzepooz/convert-mysql-data-to-postgres-sql-script/csvparser"
	"github.com/velzepooz/convert-mysql-data-to-postgres-sql-script/scriptsaver"
	"github.com/velzepooz/convert-mysql-data-to-postgres-sql-script/sqlgenerator"
)

func main() {
	for tableName, config := range ConverterConfig {
		csvData, err := csvparser.ParseCSV(config.sourceFileName, InputDirName)

		if err != nil {
			continue
		}

		sqlScript := sqlgenerator.ConvertDataToSQL(*csvData, tableName, config.schema)

		scriptsaver.SaveScriptToFile(OutputDirName, tableName, sqlScript)
	}
}
