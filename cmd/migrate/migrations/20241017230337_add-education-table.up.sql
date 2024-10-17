CREATE TABLE IF NOT EXISTS education (
	`id` INT AUTO_INCREMENT,
	`user_id` VARCHAR(255) NOT NULL, -- Reference to users table
	`institution` VARCHAR(255) NOT NULL,
	`degree` VARCHAR(255) NOT NULL,

	`field` VARCHAR(255), -- Field of study, eg. major or field of study
	`minor` VARCHAR(255), -- Minor, if applicable
	`gpa` DECIMAL(3, 2), -- GPA, if applicable

	`country` VARCHAR(255),
	`state` VARCHAR(255),
	`city` VARCHAR(255),

	`current` BOOLEAN NOT NULL DEFAULT FALSE,
	`start_date` DATE,
	`end_date` DATE,

	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);
