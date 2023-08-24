package controllers

import (
	"time"

	"github.com/eliasdn/Elda-Bank-API/app/models"
	"github.com/eliasdn/Elda-Bank-API/pkg/repository"
	"github.com/eliasdn/Elda-Bank-API/pkg/utils"
	"github.com/eliasdn/Elda-Bank-API/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAccounts func gets all exists accounts.
// @Description Get all exists accounts.
// @Summary get all exists accounts
// @Tags Accounts
// @Accept json
// @Produce json
// @Success 200 {array} models.Account
// @Router /v1/accounts [get]
func GetAccounts(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all accounts.
	accounts, err := db.GetAccounts()
	if err != nil {
		// Return, if accounts not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      "accounts were not found",
			"count":    0,
			"accounts": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(accounts),
		"accounts": accounts,
	})
}

// GetAccount func gets account by given ID or 404 error.
// @Description Get account by given ID.
// @Summary get account by given ID
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account
// @Router /v1/account/{id} [get]
func GetAccount(c *fiber.Ctx) error {
	// Catch account ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get account by ID.
	account, err := db.GetAccount(id)
	if err != nil {
		// Return, if account not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "account with the given ID is not found",
			"account": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"account": account,
	})
}

// CreateAccount func for creates a new account.
// @Description Create a new account.
// @Summary create a new account
// @Tags Account
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param user_id body string true "User ID"
// @Param account_attrs body models.AccountAttrs true "Account attributes"
// @Success 200 {object} models.Account
// @Security ApiKeyAuth
// @Router /v1/account [post]
func CreateAccount(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current account.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `account:create` from JWT data of current account.
	credential := claims.Credentials[repository.UserCreateUserCredential]

	// Only user with `account:create` credential can create a new account.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// Create new Account struct
	account := &models.Account{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(account); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Account model.
	validate := utils.NewValidator()

	// Set initialized default data for account:
	account.ID = uuid.New()
	account.CreatedAt = time.Now()
	account.UserID = claims.UserID
<<<<<<< Updated upstream
	account.AccountStatus = 1 // 0 == draft, 1 == active
=======
	account.Disabled = true // 0 == draft, 1 == active
>>>>>>> Stashed changes

	// Validate account fields.
	if err := validate.Struct(account); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create account by given model.
	if err := db.CreateAccount(account); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"account": account,
	})
}

// UpdateAccount func for updates account by given ID.
// @Description Update account.
// @Summary update account
// @Tags Account
// @Accept json
// @Produce json
// @Param id body string true "Account ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param user_id body string true "User ID"
// @Param account_status body integer true "Account status"
// @Param account_attrs body models.AccountAttrs true "Account attributes"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/account [put]
func UpdateAccount(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current account.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `account:update` from JWT data of current account.
	credential := claims.Credentials[repository.UserCreateUserCredential]

	// Only account creator with `account:update` credential can update his account.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// Create new Account struct
	account := &models.Account{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(account); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if account with given ID is exists.
	foundedAccount, err := db.GetAccount(account.ID)
	if err != nil {
		// Return status 404 and account not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "account with this ID not found",
		})
	}

	// Set user ID from JWT data of current user.
	userID := claims.UserID

	// Only the creator can delete his account.
	if foundedAccount.UserID == userID {
		// Set initialized default data for account:
		account.UpdatedAt = time.Now()

		// Create a new validator for a Account model.
		validate := utils.NewValidator()

		// Validate account fields.
		if err := validate.Struct(account); err != nil {
			// Return, if some fields are not valid.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   utils.ValidatorErrors(err),
			})
		}

		// Update account by given ID.
		if err := db.UpdateAccount(foundedAccount.ID, account); err != nil {
			// Return status 500 and error message.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 201.
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error": false,
			"msg":   nil,
		})
	} else {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, only the creator can delete his account",
		})
	}
}

// DeleteAccount func for deletes account by given ID.
// @Description Delete account by given ID.
// @Summary delete account by given ID
// @Tags Account
// @Accept json
// @Produce json
// @Param id body string true "Account ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/account [delete]
func DeleteAccount(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current account.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `account:delete` from JWT data of current account.
	credential := claims.Credentials[repository.UserDisableUserCredential]

	// Only account creator with `account:delete` credential can delete his account.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// Create new Account struct
	account := &models.Account{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(account); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Account model.
	validate := utils.NewValidator()

	// Validate account fields.
	if err := validate.StructPartial(account, "id"); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if account with given ID is exists.
	foundedAccount, err := db.GetAccount(account.ID)
	if err != nil {
		// Return status 404 and account not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "account with this ID not found",
		})
	}

	// Set user ID from JWT data of current user.
	userID := claims.UserID

	// Only the creator can delete his account.
	if foundedAccount.UserID == userID {
		// Delete account by given ID.
		if err := db.DeleteAccount(foundedAccount.ID); err != nil {
			// Return status 500 and error message.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 204 no content.
		return c.SendStatus(fiber.StatusNoContent)
	} else {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, only the creator can delete his account",
		})
	}
}
