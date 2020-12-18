package accessors

type DbNotification interface {
	addNotification() (interface{}, error)
}
