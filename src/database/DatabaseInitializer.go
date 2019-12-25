package database

import (
	"os"
)

var path = "resources/database/"
var filename = "appstud.sqlite"
var createUserTableRequest = `create table if not exists users (id varchar(36) not null primary key, login varchar(64), password varchar(64));`
var createSessionsTableRequest = `create table if not exists sessions (id varchar(36) not null primary key, user_id varchar(36));`

func Init() {
	var filepath = path + filename
	checkOrCreateFile(filepath)
	Connect(filepath)
	createTablesIfNotExist()
}

func checkOrCreateFile(filepath string) {
	_, fileErr := os.Stat(filepath)
	if os.IsNotExist(fileErr) {
		os.MkdirAll(path, 7)
		os.Create(filepath)
	}
}

func createTablesIfNotExist() {
	ExecuteSimpleUpdatableQuery(createUserTableRequest)
	ExecuteSimpleUpdatableQuery(createSessionsTableRequest)
}
