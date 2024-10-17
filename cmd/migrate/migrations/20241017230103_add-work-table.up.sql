CREATE TABLE IF NOT EXISTS work (
	`id` INT AUTO_INCREMENT,
	`user_id` VARCHAR(255) NOT NULL, -- Reference to users table
	`company` VARCHAR(255) NOT NULL,
	`position` VARCHAR(255) NOT NULL,

	`country` VARCHAR(255),
	`state` VARCHAR(255),
	`city` VARCHAR(255),

	`current` BOOLEAN NOT NULL DEFAULT FALSE, -- whether the user is currently working at this company
	`start_date` DATE,
	`end_date` DATE,

	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);
