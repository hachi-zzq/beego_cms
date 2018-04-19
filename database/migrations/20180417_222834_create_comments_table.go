package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCommentsTable_20180417_222834 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCommentsTable_20180417_222834{}
	m.Created = "20180417_222834"

	migration.Register("CreateCommentsTable_20180417_222834", m)
}

// Run the migrations
func (m *CreateCommentsTable_20180417_222834) Up() {
	m.SQL("CREATE TABLE `comments` ("+
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT,"+
		"`user_id` int(10) unsigned NOT NULL DEFAULT '0',"+
		"`article` int(10) unsigned NOT NULL DEFAULT '0',"+
		"`content` text,"+
		"`status` varchar(128) NOT NULL DEFAULT 'NORMAL',"+
		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB DEFAULT CHARSET=utf8")
}

// Reverse the migrations
func (m *CreateCommentsTable_20180417_222834) Down() {
	m.SQL("DROP TABLE comments")
}
