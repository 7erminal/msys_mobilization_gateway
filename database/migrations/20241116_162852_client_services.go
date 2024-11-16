package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ClientServices_20241116_162852 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ClientServices_20241116_162852{}
	m.Created = "20241116_162852"

	migration.Register("ClientServices_20241116_162852", m)
}

// Run the migrations
func (m *ClientServices_20241116_162852) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE client_services(`client_service_id` int(11) NOT NULL AUTO_INCREMENT,`client_id` int(11) NOT NULL,`service_id` int(11) NOT NULL,`client_service_path` varchar(100) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT 1,PRIMARY KEY (`client_service_id`), FOREIGN KEY (client_id) REFERENCES clients(client_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (service_id) REFERENCES services(service_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ClientServices_20241116_162852) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `client_services`")
}
