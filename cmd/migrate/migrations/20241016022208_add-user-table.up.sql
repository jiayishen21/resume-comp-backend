CREATE TABLE IF NOT EXISTS users (
	`id` VARCHAR(255) NOT NULL, -- auth0 userid
	`email` VARCHAR(255) NOT NULL,
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (id),
	UNIQUE KEY (email)
);
