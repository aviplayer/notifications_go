package models

type NType int

const (
	Email = iota
	SMS
	WebHook
)

var Types = map[NType]int{
	Email:   1,
	SMS:     2,
	WebHook: 3,
}

type Notification struct {
	Id          int8
	Title       string
	Type        NType
	Template    string
	Description string
}

func (n Notification) getType() int {
	return Types[n.Type]
}

func (Notification) NewNotification(title string, nType NType, template string, description string) *Notification {
	return &Notification{
		0, title, nType, template, description,
	}

}

type Notify interface {
	sendMessage() (bool, error)
	createMessage() string
}
