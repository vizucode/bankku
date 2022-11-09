package customerrepo

import (
	customercore "bankku/domains/customer/core"
	customermodel "bankku/domains/customer/model"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *customerRepo {
	return &customerRepo{
		db: db,
	}
}

func (r *customerRepo) FindUser(customerCore customercore.Core) (customercore.Core, error) {
	model := customermodel.Customer{}

	tx := r.db.Model(customermodel.Customer{}).Where("name", customerCore.Name).First(&model)
	if tx.Error != nil {
		return customercore.Core{}, tx.Error
	}

	return customercore.Core{
		Name:     model.Name,
		Ballance: model.Ballance,
	}, nil
}

func (r *customerRepo) UpdateSaldo(customerCore customercore.Core) error {
	model := customermodel.Customer{
		Ballance: customerCore.Ballance,
	}

	tx := r.db.Model(customermodel.Customer{}).Where("name", customerCore.Name).Updates(&model)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
