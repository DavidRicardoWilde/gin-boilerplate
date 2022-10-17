package tasks

type ITask interface {
	Name() string
	Exec()
}
