package database

import (
	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"log"
	"os"
	"strings"
)

func MakeMigrations() {
	sb := &strings.Builder{}
	loadModels(sb)
}

func loadModels(sb *strings.Builder) {
	var models []interface{} // add models here
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		log.Printf("failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	sb.WriteString(stmts)
	sb.WriteString(";\n")
}
