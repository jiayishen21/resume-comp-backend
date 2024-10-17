CREATE TABLE IF NOT EXISTS users (
	`id` VARCHAR(255) NOT NULL, -- auth0 userid
	`email` VARCHAR(255) NOT NULL,
	`display_name` VARCHAR(255) NOT NULL,	

	`private` BOOLEAN NOT NULL DEFAULT FALSE,

	`company` VARCHAR(255),
	`position` VARCHAR(255),

	`country` VARCHAR(255),
	`state` VARCHAR(255),
	`city` VARCHAR(255),

	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (id),
	UNIQUE KEY (email)
);
