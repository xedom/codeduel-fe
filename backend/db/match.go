package db

func (m *MariaDB) InitMatchTables() error {
	if err := m.createStatusTable(); err != nil { return err }
	if err := m.createModeTable(); err != nil { return err }
	if err := m.createChallengeTable(); err != nil { return err }

	if err := m.createMatchTable(); err != nil { return err }
	if err := m.createMatchUserLinkTable(); err != nil { return err }

	return nil
}

func (m *MariaDB) createMatchTable() error {
	query := `CREATE TABLE IF NOT EXISTS `+"`match`"+` (
		id INT unique AUTO_INCREMENT,
		owner_id INT NOT NULL,
		challenge_id INT NOT NULL,
		mode_id INT NOT NULL,
		max_users INT NOT NULL,
		max_duration INT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		
		PRIMARY KEY (id),
		FOREIGN KEY (owner_id) REFERENCES user(id),
		FOREIGN KEY (challenge_id) REFERENCES `+"`challenge`"+`(id),
		FOREIGN KEY (mode_id) REFERENCES mode(id),
		UNIQUE INDEX (id)
	);`
	_, err := m.db.Exec(query)
	return err
}

func (m *MariaDB) createMatchUserLinkTable() error {
	query := `CREATE TABLE IF NOT EXISTS match_user_link (
		id INT unique AUTO_INCREMENT,
		match_id INT NOT NULL,
		user_id INT NOT NULL,
		status_id INT NOT NULL,
		`+"`rank`"+` INT NOT NULL,
		duration INT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		
		PRIMARY KEY (id),
		FOREIGN KEY (match_id) REFERENCES `+"`match`"+`(id),
		FOREIGN KEY (user_id) REFERENCES user(id),
		FOREIGN KEY (status_id) REFERENCES status(id),
		UNIQUE INDEX (id)
	);`
	_, err := m.db.Exec(query)
	return err
}

func (m *MariaDB) createChallengeTable() error {
	query := `CREATE TABLE IF NOT EXISTS `+"`challenge`"+` (
		id INT unique AUTO_INCREMENT,
		owner_id INT NOT NULL,
		title VARCHAR(50) NOT NULL,
		description VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		
		PRIMARY KEY (id),
		FOREIGN KEY (owner_id) REFERENCES user(id),
		UNIQUE INDEX (id)
	);`
	_, err := m.db.Exec(query)
	return err
}

func (m *MariaDB) createModeTable() error {
	query := `CREATE TABLE IF NOT EXISTS mode (
		id INT unique AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
		description VARCHAR(255) NOT NULL,
		
		PRIMARY KEY (id),
		UNIQUE INDEX (id)
	);`
	_, err := m.db.Exec(query)
	if err != nil { return err }

	queryDefaultValues := `INSERT IGNORE INTO mode
	(id, name, description) VALUES	
	(1, 'speed', 'The shortest time wins.'),
	(2, 'size', 'The shortest code wins.'),
	(3, 'efficiency', 'The most efficient code wins.'),
	(4, 'memory', 'The most memory efficient code wins.'),
	(5, 'readability', 'The most readable code wins.'),
	(6, 'style', 'The most stylish code wins.');
	`
	_, err = m.db.Exec(queryDefaultValues)
	return err
}

func (m *MariaDB) createStatusTable() error {
	query := `CREATE TABLE IF NOT EXISTS status (
		id INT unique AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,

		PRIMARY KEY (id),
		UNIQUE INDEX (id)
	);`
	_, err := m.db.Exec(query)
	if err != nil { return err }

	queryDefaultValues := `INSERT IGNORE INTO status
	(id, name) VALUES	
	(0, 'not ready'),
	(1, 'ready'),
	(2, 'in match'),
	(3, 'abandoned'),
	(4, 'finished');
	`
	_, err = m.db.Exec(queryDefaultValues)
	return err
}
