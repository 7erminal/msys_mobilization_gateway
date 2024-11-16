package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Clients_20241116_161515 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Clients_20241116_161515{}
	m.Created = "20241116_161515"

	migration.Register("Clients_20241116_161515", m)
}

// Run the migrations
func (m *Clients_20241116_161515) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE clients(`client_id` int(11) NOT NULL AUTO_INCREMENT,`client_name` varchar(100) NOT NULL,`client_code` varchar(100) NOT NULL,`date_created` datetime NOT NULL,`date_modified` datetime NOT NULL,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,PRIMARY KEY (`client_id`))")
}

// Reverse the migrations
func (m *Clients_20241116_161515) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `clients`")
}
