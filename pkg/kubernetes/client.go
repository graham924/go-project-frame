package kubernetes

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var K8sClient k8sClient

// K8sClient kubernetes客户端
type k8sClient struct {
	// Config 集群配置对象
	Config *rest.Config
	// ClientSet kubernetes的客户端集合
	ClientSet *kubernetes.Clientset
}

func InitK8sClient() error {
	var err error
	var config *rest.Config
	var kubeConfig *string
	// 获取 kubeConfig 配置文件的路径
	if home := homeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		// 如果没有找到HOME路径，则默认配置文件路径设置为空，让用户自己从命令行设置kubeconfig文件的路径
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	flag.Parse()
	// 首先尝试使用InClusterConfig()方法从集群内部获取配置，如果失败了，则会尝试使用 kubeConfig 文件来创建集群配置。
	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig); err != nil {
			return err
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	K8sClient.Config = config
	K8sClient.ClientSet = clientset
	fmt.Println("initialize kubernetes clientSet success")
	return nil
}

// homeDir 从环境变量中，取出HOME或USERPROFILE的值，即为 用户的目录
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	// 没有HOME，说明是Windows系统，取 USERPROFILE 的值
	return os.Getenv("USERPROFILE") // windows
}
