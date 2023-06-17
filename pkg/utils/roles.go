package utils

import (
	"fmt"
	"time"

	"github.com/eliasdn/Elda-Bank-API/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string, c *fiber.Ctx) (string, error) {
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		tokenRole := claims.Role
		if tokenRole != "admin" {
			// Return status 403 and permission denied error message.
			return "", c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": true,
				"msg":   "permission denied, check credentials of your token",
			})
		}
	case repository.BankerRoleName:
		tokenRole := claims.Role
		if tokenRole != "banker" {
			// Return status 403 and permission denied error message.
			return "", c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": true,
				"msg":   "permission denied, check credentials of your token",
			})
		}
	case repository.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
