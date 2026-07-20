package starter

type Notifier interface {
	Notify() string
}

type emailNotifier struct {
	address string
}

func (notifier emailNotifier) Notify() string {
	return "email:" + notifier.address
}

func Notify(notifier Notifier) string {
	return notifier.Notify()
}

func IsNilInterface(notifier Notifier) bool {
	return notifier == nil
}
