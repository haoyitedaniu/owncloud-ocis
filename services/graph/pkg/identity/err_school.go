package identity

import (
	"context"
	"net/url"

	libregraph "github.com/owncloud/libre-graph-api-go"
)

// ErrEducationBackend is a dummy EducationBackend, doing nothing
type ErrEducationBackend struct{}

// CreateEducationSchool creates the supplied school in the identity backend.
func (i *ErrEducationBackend) CreateEducationSchool(ctx context.Context, school libregraph.EducationSchool) (*libregraph.EducationSchool, error) {
	return nil, errNotImplemented
}

// DeleteEducationSchool deletes a given school, identified by id
func (i *ErrEducationBackend) DeleteEducationSchool(ctx context.Context, id string) error {
	return errNotImplemented
}

// GetEducationSchool implements the EducationBackend interface for the ErrEducationBackend backend.
func (i *ErrEducationBackend) GetEducationSchool(ctx context.Context, nameOrID string, queryParam url.Values) (*libregraph.EducationSchool, error) {
	return nil, errNotImplemented
}

// GetEducationSchools implements the EducationBackend interface for the ErrEducationBackend backend.
func (i *ErrEducationBackend) GetEducationSchools(ctx context.Context, queryParam url.Values) ([]*libregraph.EducationSchool, error) {
	return nil, errNotImplemented
}

// GetEducationSchoolMembers implements the EducationBackend interface for the ErrEducationBackend backend.
func (i *ErrEducationBackend) GetEducationSchoolMembers(ctx context.Context, id string) ([]*libregraph.EducationUser, error) {
	return nil, errNotImplemented
}

// AddMembersToEducationSchool adds new members (reference by a slice of IDs) to supplied school in the identity backend.
func (i *ErrEducationBackend) AddMembersToEducationSchool(ctx context.Context, schoolID string, memberID []string) error {
	return errNotImplemented
}

// RemoveMemberFromEducationSchool removes a single member (by ID) from a school
func (i *ErrEducationBackend) RemoveMemberFromEducationSchool(ctx context.Context, schoolID string, memberID string) error {
	return errNotImplemented
}

// CreateEducationUser creates a given education user in the identity backend.
func (i *ErrEducationBackend) CreateEducationUser(ctx context.Context, user libregraph.EducationUser) (*libregraph.EducationUser, error) {
	return nil, errNotImplemented
}

// DeleteEducationUser deletes a given educationuser, identified by username or id, from the backend
func (i *ErrEducationBackend) DeleteEducationUser(ctx context.Context, nameOrID string) error {
	return errNotImplemented
}

// UpdateEducationUser applies changes to given education user, identified by username or id
func (i *ErrEducationBackend) UpdateEducationUser(ctx context.Context, nameOrID string, user libregraph.EducationUser) (*libregraph.EducationUser, error) {
	return nil, errNotImplemented
}

// GetEducationUser implements the EducationBackend interface for the ErrEducationBackend backend.
func (i *ErrEducationBackend) GetEducationUser(ctx context.Context, nameOrID string, queryParam url.Values) (*libregraph.EducationUser, error) {
	return nil, errNotImplemented
}

// GetEducationUsers implements the EducationBackend interface for the ErrEducationBackend backend.
func (i *ErrEducationBackend) GetEducationUsers(ctx context.Context, queryParam url.Values) ([]*libregraph.EducationUser, error) {
	return nil, errNotImplemented
}