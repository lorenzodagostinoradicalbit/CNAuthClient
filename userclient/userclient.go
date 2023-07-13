package userclient

import (
	"context"
	"os"

	"github.com/lorenzodagostinoradicalbit/CNAuth/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type UserClient struct {
	client    *dynamic.DynamicClient
	namespace string
}

func buildConfiguration() (*rest.Config, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	var clusterConfig *rest.Config
	var err error
	if kubeconfig == "" {
		clusterConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		clusterConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return clusterConfig, nil
}

func NewUserClientFromNamespace(ns string) (*UserClient, error) {
	config, err := buildConfiguration()
	if err != nil {
		return nil, err
	}

	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	ret := &UserClient{
		client:    clientset,
		namespace: ns,
	}
	return ret, nil
}

func NewUserClient() (*UserClient, error) {
	config, err := buildConfiguration()
	if err != nil {
		return nil, err
	}

	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	ret := &UserClient{
		client:    clientset,
		namespace: "",
	}
	return ret, nil
}

func (uc *UserClient) Get(name string) (*v1alpha1.User, error) {
	if uc.namespace == "" {
		return nil, &NamespaceNotSetError{}
	}
	userSchema := v1alpha1.GroupVersion.WithResource("users")

	obj, err := uc.client.
		Resource(userSchema).
		Namespace(uc.namespace).
		Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	user := &v1alpha1.User{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserClient) List() (*v1alpha1.UserList, error) {
	if uc.namespace == "" {
		return nil, &NamespaceNotSetError{}
	}
	userSchema := v1alpha1.GroupVersion.WithResource("users")

	list, err := uc.client.
		Resource(userSchema).
		Namespace(uc.namespace).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	users := &v1alpha1.UserList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(list.UnstructuredContent(), users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserClient) SetNamespace(ns string) {
	uc.namespace = ns
}
