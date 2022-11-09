package customerservice

import (
	"bankku/config"
	customercore "bankku/domains/customer/core"
	mocks "bankku/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCustomer(t *testing.T) {
	t.Run("failed get customer cause error repo", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{}, errors.New(config.INTERNAL_SERVER_ERROR))

		service := New(repo)
		result, err := service.Login("havis")

		assert.Equal(t, config.INTERNAL_SERVER_ERROR, err.Error())
		assert.Equal(t, "", result.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed get customer cause not found customer", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{}, errors.New("not found error"))

		service := New(repo)
		result, err := service.Login("havis")

		assert.Equal(t, "not found error", err.Error())
		assert.Equal(t, "", result.Name)
		repo.AssertExpectations(t)
	})

	t.Run("success get customer", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 1000000,
		}, nil)

		service := New(repo)
		result, err := service.Login("havis")

		assert.Equal(t, "havis", result.Name)
		assert.Equal(t, float64(1000000), result.Ballance)
		assert.Equal(t, nil, err)
		repo.AssertExpectations(t)
	})
}

func TestCreateCustomer(t *testing.T) {
	t.Run("failed create customer duplicate name", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("InsertCustomer", mock.Anything).Return(errors.New(config.DUPLICATE_NAME)).Once()

		service := New(repo)
		err := service.CreateCustomer("havis")

		assert.Equal(t, config.DUPLICATE_NAME, err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("success create customer", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("InsertCustomer", mock.Anything).Return(nil).Once()

		service := New(repo)
		err := service.CreateCustomer("havis")

		assert.Equal(t, nil, err)
		repo.AssertExpectations(t)
	})
}

func TestTopUp(t *testing.T) {
	t.Run("failed topup cause zero price", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 0,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.TopUp("havis", 5000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.MINIMAL_TOP_UP, err.Error())
	})

	t.Run("failed topup cause error find customer", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{}, errors.New(config.INTERNAL_SERVER_ERROR)).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.TopUp("havis", 100000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.INTERNAL_SERVER_ERROR, err.Error())
	})

	t.Run("failed topup cause error repo", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 0,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(errors.New(config.INTERNAL_SERVER_ERROR)).Once()

		service := New(repo)
		ballance, err := service.TopUp("havis", 100000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.INTERNAL_SERVER_ERROR, err.Error())
	})

	t.Run("success topup", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 0,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.TopUp("havis", 50000)

		assert.Equal(t, float64(50000), ballance)
		assert.Equal(t, nil, err)
	})
}

func TestWithDraw(t *testing.T) {
	t.Run("failed minimal wd", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 0,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.Withdraw("havis", 5000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.MINIMAL_WD, err.Error())
	})

	t.Run("failed withdraw cause ballance not insuffience", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 0,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.Withdraw("havis", 50000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.BALLANCE_NOT_ENOUGH, err.Error())
	})

	t.Run("failed withdraw case error find customer", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{}, errors.New(config.INTERNAL_SERVER_ERROR)).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.Withdraw("havis", 50000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.INTERNAL_SERVER_ERROR, err.Error())
	})

	t.Run("failed withdraw case error repo", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 100000,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(errors.New(config.INTERNAL_SERVER_ERROR)).Once()

		service := New(repo)
		ballance, err := service.Withdraw("havis", 50000)

		assert.Equal(t, float64(0), ballance)
		assert.Equal(t, config.INTERNAL_SERVER_ERROR, err.Error())
	})

	t.Run("success withdraw", func(t *testing.T) {
		repo := new(mocks.CustomerRepo)
		repo.On("FindCustomer", mock.Anything).Return(customercore.Core{
			Name:     "havis",
			Ballance: 100000,
		}, nil).Once()
		repo.On("UpdateSaldo", mock.Anything).Return(nil).Once()

		service := New(repo)
		ballance, err := service.Withdraw("havis", 50000)

		assert.Equal(t, float64(50000), ballance)
		assert.Equal(t, nil, err)
	})
}
