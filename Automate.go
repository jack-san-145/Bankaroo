package main

import "log"

func CreateTables() {
	query := `CREATE TABLE IF NOT EXISTS CustomersList (
    acc_no     INT PRIMARY KEY AUTO_INCREMENT,
    cus_name   VARCHAR(50) NOT NULL,
    password   VARCHAR(100) NOT NULL,
    address    VARCHAR(150),
    mobile_no  VARCHAR(15)
) AUTO_INCREMENT = 1000;`
	_, err := Db.Exec(query)
	if err != nil {
		log.Println("error while creating table CustomersList - ", err)
		return
	}

	query = `CREATE TABLE IF NOT EXISTS CustomerBalance (
    acc_no  INT PRIMARY KEY,
    balance INT
);`
	_, err = Db.Exec(query)
	if err != nil {
		log.Println("error while creating table CustomerBalance - ", err)
		return
	}

	query = `CREATE TABLE IF NOT EXISTS Transaction (
    sender_id        INT NOT NULL,
    receiver_id      INT NOT NULL,
    sender_name      VARCHAR(50) NOT NULL,
    receiver_name    VARCHAR(50) NOT NULL,
    amount           INT NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	_, err = Db.Exec(query)
	if err != nil {
		log.Println("error while creating table Transaction - ", err)
		return
	}

}
