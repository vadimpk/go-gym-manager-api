package tests

const (
	createManagerQuery    = "INSERT INTO managers (first_name, last_name, phone_number, email, password) VALUES ('string', 'string', 'string', 'string@test.com', 'string');"
	resetSequenceQuery    = "ALTER SEQUENCE %s_id_seq RESTART; UPDATE %s SET id = DEFAULT;"
	deleteDataQuery       = "DELETE FROM %s;"
	createMemberQuery     = "INSERT INTO members (first_name, last_name, phone_number, joined_at) VALUES (%s, %s, %s, now());"
	createMembershipQuery = "INSERT INTO memberships (short_name, description, price, duration) VALUES (%s, %s, %d, %s);"
	setMembershipQuery    = "INSERT INTO members_memberships (member_id, membership_id, membership_expiration) VALUES (%d, %d, %s);"
)
