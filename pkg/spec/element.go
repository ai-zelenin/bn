package spec

type Element interface {
	ID() string
}

type Activity interface {
	Definition() *Task
}
