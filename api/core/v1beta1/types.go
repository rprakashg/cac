package v1beta1

// A username on the system
type Username string

func (u Username) String() string {
	return string(u)
}
