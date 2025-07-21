package adapters

import (
	"github.com/go-hex/RookieJoel/core"
	"gorm.io/gorm"
)

//secondary adapter implementation of the OrderRepository interface
type GormOrderRepository struct {
	db *gorm.DB // GORM database connection
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db} // return an instance of GormOrderRepository
}

func (r *GormOrderRepository) Save(order core.Order) error {
	if err := r.db.Create(&order).Error; err != nil {
		return err // return error if saving fails
	}
	return nil // return nil if order is saved successfully
}