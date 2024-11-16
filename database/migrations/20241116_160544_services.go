package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Services_20241116_160544 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Services_20241116_160544{}
	m.Created = "20241116_160544"

	migration.Register("Services_20241116_160544", m)
}

// Run the migrations
func (m *Services_20241116_160544) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE services(`service_id` int(11) NOT NULL AUTO_INCREMENT,`service_name` varchar(100) NOT NULL,`service_description` varchar(500) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,PRIMARY KEY (`service_id`))")
}

// Reverse the migrations
func (m *Services_20241116_160544) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `services`")
}
