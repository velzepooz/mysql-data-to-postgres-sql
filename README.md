# Golang script to convert csv data imported from MysqlDb to sql script file for Postgres db

1. Export your table data to csv file
2. Put file with data to `input` folder
3. At `config.go` add table name, name of source file, and describe table schema: <column_name>: <data_type> at appearing order at csv file. Column name should be the same as at your file.
4. Run script
5. Look for converted files at `output` folder

## Datatype mapping

`date => string`
`int/float => number`
`enum => string`
`boolean => boolean`
