CREATE TABLE IF NOT EXISTS tokens (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	access_token VARCHAR(300) NOT NULL,
	refresh_token VARCHAR(300) NOT NULL,
	sso_access_token VARCHAR(300) NOT NULL,
	token_type VARCHAR(300) NOT NULL,
	expiry TIMESTAMP WITH TIME ZONE NOT NULL,
	email VARCHAR(300) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX tokens_access_token ON tokens (access_token);
CREATE UNIQUE INDEX tokens_sso_access_token ON tokens (sso_access_token);
CREATE UNIQUE INDEX tokens_refresh_token ON tokens (refresh_token);
