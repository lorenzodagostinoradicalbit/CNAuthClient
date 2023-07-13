package server

import "auth-client/userclient"

var userClient *userclient.UserClient

func UserClientInstance(namespace string) (*userclient.UserClient, error) {
	var err error
	if userClient == nil {
		userClient, err = userclient.NewUserClientFromNamespace(namespace)
		return userClient, err
	}
	userClient.SetNamespace(namespace)
	return userClient, nil
}
