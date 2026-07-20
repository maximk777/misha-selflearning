package starter

import "testing"

func TestNotifyUsesBehaviorNotConcreteType(t *testing.T) {
	if got := Notify(emailNotifier{address: "misha@example.test"}); got != "email:misha@example.test" {
		t.Fatalf("Notify()=%q", got)
	}
}

func TestNilInterfaceAndTypedNilAreDifferent(t *testing.T) {
	var none Notifier
	if !IsNilInterface(none) {
		t.Fatal("nil interface must be nil")
	}
	var email *emailNotifier
	var typedNil Notifier = email
	if IsNilInterface(typedNil) {
		t.Fatal("typed nil in interface must not compare equal to nil")
	}
}
