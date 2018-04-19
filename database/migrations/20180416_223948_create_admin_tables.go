package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAdminTables_20180416_223948 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAdminTables_20180416_223948{}
	m.Created = "20180416_223948"

	migration.Register("CreateAdminTables_20180416_223948", m)
}

// Run the migrations
func (m *CreateAdminTables_20180416_223948) Up() {
	m.SQL("CREATE TABLE `admins` ("+
	"`id` int(11) unsigned NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(128) NOT NULL DEFAULT '',"+
		"`password` varchar(128) NOT NULL DEFAULT '',"+
		"`status` varchar(128) NOT NULL DEFAULT 'NORMAL',"+
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB DEFAULT CHARSET=utf8")
}

// Reverse the migrations
func (m *CreateAdminTables_20180416_223948) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `admins`")
}
