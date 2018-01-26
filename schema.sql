CREATE TABLE person (
	id serial PRIMARY KEY,
	email text NOT NULL UNIQUE,
	registered timestamp(0) with time zone NOT NULL DEFAULT now(),
	password_key character(64),
	password_salt character(64) UNIQUE
);
CREATE TABLE session (
	token character(64) PRIMARY KEY,
	person integer NOT NULL REFERENCES person,
	sign_in_time timestamp(0) with time zone NOT NULL DEFAULT now(),
	host inet,
	last_seen timestamp(0) with time zone NOT NULL DEFAULT now()
);
CREATE TABLE page (
	
);
