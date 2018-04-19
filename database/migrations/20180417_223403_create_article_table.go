package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateArticleTable_20180417_223403 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateArticleTable_20180417_223403{}
	m.Created = "20180417_223403"

	migration.Register("CreateArticleTable_20180417_223403", m)
}

// Run the migrations
func (m *CreateArticleTable_20180417_223403) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `articles` ("+
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT,"+
	"	`user_id` int(128) unsigned NOT NULL DEFAULT '0',"+
	"	`titile` varchar(128) NOT NULL DEFAULT '',"+
	"	`content` text,"+
	"	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,"+
	"	PRIMARY KEY (`id`)"+
	") ENGINE=InnoDB DEFAULT CHARSET=utf8")
}

// Reverse the migrations
func (m *CreateArticleTable_20180417_223403) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE articles")

}
