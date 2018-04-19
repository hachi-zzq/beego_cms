package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20180417_222717 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20180417_222717{}
	m.Created = "20180417_222717"

	migration.Register("CreateUsersTable_20180417_222717", m)
}

// Run the migrations
func (m *CreateUsersTable_20180417_222717) Up() {
	m.SQL("CREATE TABLE `users` ("+
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(128) NOT NULL DEFAULT '',"+
		"`password` varchar(128) NOT NULL DEFAULT '',"+
		"`avatar` varchar(128) DEFAULT NULL,"+
		"`status` varchar(128) NOT NULL DEFAULT 'NORMAL',"+
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB DEFAULT CHARSET=utf8")
}

// Reverse the migrations
func (m *CreateUsersTable_20180417_222717) Down() {
	m.SQL("DROP TABLE users")
}
