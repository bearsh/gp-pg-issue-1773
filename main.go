package main

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
)

var createStmt string = `DROP TABLE IF EXISTS public."table" CASCADE;
CREATE TABLE public."table" (
	id serial NOT NULL,
	data bytea,
	CONSTRAINT table_pk PRIMARY KEY (id)
);`

func main() {
	options, err := pg.ParseURL(os.Args[1])
	if err != nil {
		panic(">>> parsing connection url error")
	}
	db := pg.Connect(options)
	db.AddQueryHook(pgdebug.DebugHook{Verbose: true})

	fmt.Printf(">>> create table 'table' in db '%v' on '%v'\n", db.Options().Database, db.String())
	if _, err := db.Exec(createStmt); err != nil {
		fmt.Printf(">>> could not create table: %v\nexit\n", err)
		os.Exit(1)
	}

	t1 := &Table{
		Data: []byte{1, 2, 3, 4},
	}
	fmt.Printf(">>> insert t1: '%v' into 'table'\n", t1)
	if _, err := db.Model(t1).Insert(); err != nil {
		fmt.Printf(">>> t1: Insert: %v\nexit\n", err)
		os.Exit(1)
	}
	fmt.Printf(">>> t1 after inserting: '%v'\n", t1)

	t2 := &Table{
		ID: t1.ID,
	}
	fmt.Printf(">>> select the just inserted row\n")
	if err := db.Model(t2).WherePK().Select(); err != nil {
		fmt.Printf(">>> error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf(">>> returned value: %v\n", t2)
}
