package customerservice

import (
	"bankku/config"
	customercore "bankku/domains/customer/core"
	"errors"
)

type customerService struct {
	repo customercore.IRepoCustomer
}

func New(repo customercore.IRepoCustomer) *customerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) Login(username string) (customercore.Core, error) {
	core := customercore.Core{
		Name: username,
	}

	return s.repo.FindCustomer(core)
}

func (s *customerService) CreateCustomer(username string) error {
	core := customercore.Core{
		Name: username,
	}

	err := s.repo.InsertCustomer(core)
	if err != nil {
		return errors.New(config.DUPLICATE_NAME)
	}

	return nil
}

func (s *customerService) TopUp(username string, price float64) (float64, error) {
	core := customercore.Core{
		Name: username,
	}

	if price < 50000 {
		return 0, errors.New(config.MINIMAL_TOP_UP)
	}

	customerData, err := s.repo.FindCustomer(core)
	if err != nil {
		return 0, errors.New(config.INTERNAL_SERVER_ERROR)
	}

	core.Ballance = price + customerData.Ballance

	err = s.repo.UpdateSaldo(core)
	if err != nil {
		return 0, err
	}

	return core.Ballance, err
}

func (s *customerService) Withdraw(username string, price float64) (float64, error) {
	core := customercore.Core{
		Name: username,
	}

	if price < 50000 {
		return 0, errors.New(config.MINIMAL_WD)
	}

	customerData, err := s.repo.FindCustomer(core)
	if err != nil {
		return 0, errors.New(config.INTERNAL_SERVER_ERROR)
	}

	if price > customerData.Ballance {
		return 0, errors.New(config.BALLANCE_NOT_ENOUGH)
	}

	core.Ballance = customerData.Ballance - price

	err = s.repo.UpdateSaldo(core)

	if err != nil {
		return 0, errors.New(config.INTERNAL_SERVER_ERROR)
	}

	return core.Ballance, nil
}
