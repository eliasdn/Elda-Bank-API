package utils

import (
	"fmt"

	"github.com/eliasdn/Elda-Bank-API/pkg/repository"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		// Admin credentials (all access).
		credentials = []string{
			repository.UserCreateUserCredential,
			repository.UserCreateBankerCredential,

			repository.UserUpdateUserCredential,
			repository.UserUpdatebankerCredential,

			repository.UserDisableUserCredential,
			repository.UserDisablebnakerCredential,
		}
	case repository.BankerRoleName:
		// Moderator credentials (only book creation and update).
		credentials = []string{
			repository.UserCreateUserCredential,
			repository.UserUpdateUserCredential,
		}
	case repository.UserRoleName:
		// Simple user credentials (only book creation).
		credentials = []string{
			//repository.UserCreateUserCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
