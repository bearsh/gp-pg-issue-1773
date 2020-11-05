package main

var Columns = struct {
	Table struct {
		ID, Data string
	}
}{
	Table: struct {
		ID, Data string
	}{
		ID:   "id",
		Data: "data",
	},
}

var Tables = struct {
	Table struct {
		Name, Alias string
	}
}{
	Table: struct {
		Name, Alias string
	}{
		Name:  "table",
		Alias: "t",
	},
}

type Table struct {
	tableName struct{} `pg:"table,alias:t,discard_unknown_columns"`

	ID   int    `pg:"id,pk"`
	Data []byte `pg:"data,array"`
}
