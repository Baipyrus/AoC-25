package registry

type Challenge struct {
	Name string
	Exec func(string)
}

var challenges []Challenge

func Register(name string, exec func(string)) {
	challenges = append(
		challenges,
		Challenge{name, exec})
}

func Get() []Challenge {
	return challenges
}
