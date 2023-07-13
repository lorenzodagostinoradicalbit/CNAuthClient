package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lorenzodagostinoradicalbit/CNAuth/api/v1alpha1"
)

func getUser(name string, namespace string) (*v1alpha1.User, error) {
	uc, err := UserClientInstance(namespace)
	if err != nil {
		return nil, err
	}
	us, err := uc.Get(name)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func listUsers(namespace string) (*v1alpha1.UserList, error) {
	uc, err := UserClientInstance(namespace)
	if err != nil {
		return nil, err
	}
	users, err := uc.List()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserToken(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := getUser(input.Name, input.Namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Spec.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "bad password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": user.Status.Token})
}

func ListUser(c *gin.Context) {
	var input ListUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := listUsers(input.Namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var usersNames []string
	for _, user := range users.Items {
		usersNames = append(usersNames, user.Name)
	}
	c.JSON(http.StatusOK, gin.H{"users": usersNames, "namespace": input.Namespace})
}
