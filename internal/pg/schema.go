package pg

const Schema = `
CREATE TABLE credential
(
	id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	password VARCHAR(128) NOT NULL,
	token VARCHAR(128) NOT NULL,
	email VARCHAR(255) NOT NULL DEFAULT '',
	email_tmp VARCHAR(255) NOT NULL DEFAULT '',
	email_verified BOOLEAN NOT NULL,
	verification_code VARCHAR(8) NOT NULL DEFAULT '',
	verification_code_attempts smallint NOT NULL DEFAULT 0,
	created_at timestamp with time zone DEFAULT now() NOT NULL,
	updated_at timestamp with time zone DEFAULT now() NOT NULL,
	UNIQUE (email),
	UNIQUE (email_tmp),
	UNIQUE (token)
);
`
