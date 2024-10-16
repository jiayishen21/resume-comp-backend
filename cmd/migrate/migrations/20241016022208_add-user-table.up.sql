CREATE TABLE IF NOT EXISTS users (
	`id` VARCHAR(255) NOT NULL, -- auth0 userid
	`firstName` VARCHAR(255) NOT NULL,
	`lastName` VARCHAR(255) NOT NULL,
	`email` VARCHAR(255) NOT NULL,
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (id),
	UNIQUE KEY (email)
);
