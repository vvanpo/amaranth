CREATE TABLE person (
	id serial PRIMARY KEY,
	email text NOT NULL UNIQUE,
	name text NOT NULL,
	telephone text,
	registered timestamp(0) with time zone NOT NULL DEFAULT now(),
	password_key character(64),
	password_salt character(64) UNIQUE
);
CREATE TABLE session (
	token character(64) PRIMARY KEY,
	person integer NOT NULL REFERENCES person,
	sign_in_time timestamp(0) with time zone NOT NULL DEFAULT now(),
	last_seen timestamp(0) with time zone NOT NULL DEFAULT now()
);
