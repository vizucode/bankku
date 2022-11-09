package usercore

type IRepoUser interface {
	Insert(activityCore Core) (Core, error)
	Update(activityCore Core) (Core, error)
	GetByEmail(activityCore Core) (bool, error)
}

type IServiceUser interface {
	Create(activityCore Core) Core
	Verify(activityCore Core) Core
}
