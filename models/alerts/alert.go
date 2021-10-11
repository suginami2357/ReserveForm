package alerts

type Alert struct {
	Type    string
	Message string
}

const (
	_ = iota
	Success
	Info
	Warning
	Danger
)

func New(t int, m string) Alert {
	var vs string
	switch t {
	case Success:
		vs = "success"
	case Info:
		vs = "info"
	case Warning:
		vs = "warning"
	case Danger:
		vs = "danger"
	default:
		panic("argument is undefined")
	}
	return Alert{Type: vs, Message: m}
}
