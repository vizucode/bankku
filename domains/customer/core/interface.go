package customercore

type IRepoCustomser interface {
	FindUser(customerCore Core) (Core, error)
	UpdateSaldo(customerCore Core) error
}

type IServiceCustomer interface {
	Login(username string) (Core, error)
	TopUp(username string, price float64) error
	Withdraw(username string, price float64) error
}
