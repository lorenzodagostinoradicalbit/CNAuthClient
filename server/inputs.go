package server

type UserInput struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Password  string `json:"password"`
}

type ListUserInput struct {
	Namespace string `json:"namespace"`
}
