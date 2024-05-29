-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION audit_time_stamp() RETURNS TRIGGER AS
$$
BEGIN
	IF TG_OP = 'UPDATE' THEN
		NEW.updated_at := current_timestamp;
        RETURN NEW;
	ELSIF TG_OP = 'INSERT' THEN
			NEW.updated_at := current_timestamp;
			NEW.created_at := current_timestamp;
	        RETURN NEW;
	END IF;
END;
$$
LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TABLE users (
		id BIGSERIAL PRIMARY KEY NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		user_name TEXT NOT NULL
);

CREATE OR REPLACE TRIGGER users_timestamp_audit_trigger BEFORE INSERT OR UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE audit_time_stamp();

-- +goose Down
DROP TABLE users;
DROP FUNCTION audit_time_stamp;
