package model

type TaskStatus string

const (
	ToDo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

func (s TaskStatus) IsValid() bool {
	switch s {
	case ToDo, InProgress, Done:
		return true
	}
	return false
}

func (s TaskStatus) String() string {
	return string(s)
}
