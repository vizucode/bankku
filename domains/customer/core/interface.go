package customercore

type IRepoCustomer interface {
	InsertCustomer(customerCore Core) error
	FindCustomer(customerCore Core) (Core, error)
	UpdateSaldo(customerCore Core) error
}

type IServiceCustomer interface {
	CreateCustomer(username string) error
	Login(username string) (Core, error)
	TopUp(username string, price float64) (float64, error)
	Withdraw(username string, price float64) (float64, error)
}
