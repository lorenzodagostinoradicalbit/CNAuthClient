package userclient

type ClientError struct{}

type NamespaceNotSetError ClientError

func (nns *NamespaceNotSetError) Error() string {
	return "Namespace was not set"
}
