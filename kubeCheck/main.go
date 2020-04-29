package kubeCheck

import (
	"context"
	"fmt"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	"time"
)

var (
	Duration   time.Duration
	Labels     []string
	Namespace  string
	KubeConfig string
)

func Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), Duration)
	defer cancel()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", KubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deployAppNum, readyAppNum := len(Labels), 0

	watch, err := clientSet.ExtensionsV1beta1().Deployments(Namespace).Watch(metav1.ListOptions{Watch: true, LabelSelector: fmt.Sprintf("app in (%s)", strings.Join(Labels, ","))})

	if err != nil {
		panic(err.Error())
	}

	ticker := time.NewTicker(5 * time.Second)

	var status v1beta1.DeploymentStatus

	for {
		select {
		case <-ticker.C:
			fmt.Printf("%d of %d replicas are ready\n", status.ReadyReplicas, status.Replicas)
		case event := <-watch.ResultChan():
			if event.Object == nil {
				continue
			}
			d := event.Object.(*v1beta1.Deployment)
			status = d.Status
			fmt.Printf("Checking %s deploy status\n", d.Name)
			if DeploymentComplete(d, &status) {
				fmt.Printf("%s deploy success\n", d.Name)
				readyAppNum++
				if deployAppNum == readyAppNum {
					fmt.Printf("All deployment success\n")
					ticker.Stop()
					watch.Stop()
					return nil
				}
			}
			fmt.Printf("%d of %d replicas are ready\n", status.ReadyReplicas, status.Replicas)
		case <-ctx.Done():
			ticker.Stop()
			watch.Stop()
			return nil
		}
	}

}

func DeploymentComplete(deployment *v1beta1.Deployment, newStatus *v1beta1.DeploymentStatus) bool {
	return newStatus.UpdatedReplicas == *(deployment.Spec.Replicas) &&
		newStatus.Replicas == *(deployment.Spec.Replicas) &&
		newStatus.AvailableReplicas == *(deployment.Spec.Replicas) &&
		newStatus.ObservedGeneration >= deployment.Generation
}
