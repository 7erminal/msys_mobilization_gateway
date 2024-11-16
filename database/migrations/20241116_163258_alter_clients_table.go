package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AlterClientsTable_20241116_163258 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterClientsTable_20241116_163258{}
	m.Created = "20241116_163258"

	migration.Register("AlterClientsTable_20241116_163258", m)
}

// Run the migrations
func (m *AlterClientsTable_20241116_163258) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE clients CHANGE `date_created` `date_created` datetime DEFAULT CURRENT_TIMESTAMP, CHANGE `date_modified` `date_modified` datetime ON UPDATE CURRENT_TIMESTAMP, CHANGE `active` `active` int(11) DEFAULT 1")
}

// Reverse the migrations
func (m *AlterClientsTable_20241116_163258) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
