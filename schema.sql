CREATE TABLE todo (
	id UUID PRIMARY KEY,
	name text NOT NULL,
	description text,
	done boolean NOT NULL,
	createAt timestamp NOT NULL
)
