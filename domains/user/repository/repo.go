package activityrepo

import (
	usercore "bankku/domains/user/core"
	usermodel "bankku/domains/user/model"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Insert(userCore usercore.Core) (usercore.Core, error) {
	model := usermodel.ToModel(userCore)
	userCore.IsActive = false
	tx := r.db.Create(&model)

	if tx.Error != nil {
		return usercore.Core{}, tx.Error
	}

	return usermodel.ToCore(model), nil
}

func (r *userRepo) Update(userCore usercore.Core) (usercore.Core, error) {
	model := usermodel.ToModel(userCore)
	model.IsActive = userCore.IsActive
	tx := r.db.Model(usermodel.User{}).Where("email", userCore.Email).Select("is_active").Updates(&model)
	if tx.Error != nil {
		return usercore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return usercore.Core{}, gorm.ErrRecordNotFound
	}

	tx = r.db.Model(usermodel.User{}).First(&model)
	if tx.Error != nil {
		return usercore.Core{}, tx.Error
	}

	return usermodel.ToCore(model), nil
}

func (r *userRepo) GetByEmail(userCore usercore.Core) (bool, error) {
	model := usermodel.ToModel(userCore)
	tx := r.db.Model(usermodel.User{}).Where("email", userCore.Email).Where("password", userCore.Password).First(&model)
	if tx.Error != nil {
		return false, tx.Error
	}

	if tx.RowsAffected < 1 {
		return false, gorm.ErrRecordNotFound
	}

	tx = r.db.Model(usermodel.User{}).First(&model)
	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
