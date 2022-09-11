package domain

const (
	ErrNotInDB              = "sql: no rows in result set"
	ErrBadRequest           = "bad request"
	ErrStillInGym           = "still in gym"
	ErrIsNotInGym           = "not in gym"
	ErrEmptyAuthHeader      = "empty auth header"
	ErrInvalidAuthHeader    = "invalid auth header"
	ErrEmptyToken           = "token is empty"
	ErrManagerIdNotFound    = "manager id not found"
	ErrDoesntHaveMembership = "doesn't have membership"
	ErrExpiredMembership    = "membership expired"

	ErrNotInDBMessage              = "No results found in database"
	ErrBadRequestMessage           = "Bad request: input data is incorrect"
	ErrStillInGymMessage           = "Cannot start the visit: person is still in the gym"
	ErrIsNotInGymMessage           = "Cannot end the visit: person is not in the gym"
	ErrNotAuthMessage              = "Not authorized"
	ErrInternalServerMessage       = "Server is not responding at the moment"
	ErrManagerIdNotFoundMessage    = "Cannot do the operation: manager ID not found. Try signing in again"
	ErrDoesntHaveMembershipMessage = "This member doesn't have any active memberships"
	ErrExpiredMembershipMessage    = "Cannot do the operation: member's membership is expired"

	MessageMemberCreated            = "Member created successfully"
	MessageMemberUpdated            = "Member updated successfully"
	MessageMemberFound              = "Member found successfully"
	MessageMemberDeleted            = "Member deleted successfully"
	MessageMembersMembershipSet     = "Membership set successfully"
	MessageMembersMembershipDeleted = "Membership deleted successfully"

	MessageMembershipCreated = "Membership created successfully"
	MessageMembershipUpdated = "Membership updated successfully"
	MessageMembershipFound   = "Membership found successfully"
	MessageMembershipDeleted = "Membership deleted successfully"

	MessageTrainerCreated = "Trainer created successfully"
	MessageTrainerUpdated = "Trainer updated successfully"
	MessageTrainerFound   = "Trainer found successfully"
	MessageTrainerDeleted = "Trainer deleted successfully"

	MessageVisitSet   = "Visit started successfully"
	MessageVisitEnded = "Visit ended successfully"
)
