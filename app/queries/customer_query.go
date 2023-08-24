package queries

import (
	"github.com/eliasdn/Elda-Bank-API/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CustomerQueries struct for queries from customer model.
type CustomerQueries struct {
	*gorm.DB
}

// GetCustomers method for getting all customers.
func (q *CustomerQueries) GetCustomers() ([]models.Customer, error) {
	customers := []models.Customer{}
	err := q.Find(&customers)
	if err != nil {
		return customers, err.Error
	}
	return customers, nil
}

// GetCustomer method for getting one customer by given ID.
func (q *CustomerQueries) GetCustomer(id uuid.UUID) (models.Customer, error) {
	customer := models.Customer{}
	err := q.First(&customer, id)
	if err != nil {
		return customer, err.Error
	}
	return customer, nil
}

// Createcustomer method for creating customer by given customer object.
func (q *CustomerQueries) CreateCustomer(b *models.Customer) error {
	query := q.Create(b)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// UpdateCustomer method for updating customer by given Customer object.
func (q *CustomerQueries) UpdateCustomer(id uuid.UUID, b *models.Customer) error {
	query := q.Where("id = ?", id).Save(b)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// DeleteCustomer method for delete customer by given ID.
func (q *CustomerQueries) DeleteCustomer(id uuid.UUID) error {
	customer := models.Customer{}
	query := q.Where("id = ?", id).Delete(&customer)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
