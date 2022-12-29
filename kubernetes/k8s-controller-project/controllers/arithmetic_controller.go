/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mathsv1 "demo/api/v1"
)

// ArithmeticReconciler reconciles a Arithmetic object
type ArithmeticReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=maths.controller,resources=arithmetics,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=maths.controller,resources=arithmetics/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=maths.controller,resources=arithmetics/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Arithmetic object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *ArithmeticReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//_ = log.FromContext(ctx)

	// // TODO(user): your logic here
	// // Get the arithmetic object, here - namespaced/name is to uniquely identify the resource object
	var problem mathsv1.Arithmetic
	if err := r.Get(ctx, req.NamespacedName, &problem); err != nil {
		fmt.Println("Could not get arithmetic resource, ", err)
		return ctrl.Result{}, err
	}

	fmt.Printf("Reconciling for %s", req.NamespacedName)
	fmt.Printf("\nExpression : %s", problem.Spec.Expression)

	if problem.Status.Answer == "" {
		pod := corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("job-%s", req.Name),
				Namespace: "default",
			},
			Spec: corev1.PodSpec{
				RestartPolicy: "Never",
				Containers: []corev1.Container{
					{
						Name:  "demo-container",
						Image: "python:latest",
						Args:  []string{"python", "-c", fmt.Sprintf("print(%s)", problem.Spec.Expression)},
					},
				},
			},
		}

		err := r.Create(ctx, &pod)
		if err != nil {
			fmt.Println("\nPod not created, error: ", err)
			return ctrl.Result{}, err
		}

		fmt.Println("\nPod created successfully!")
		time.Sleep(10 * time.Second)

		podLogsAns, err := readPodLogs(pod)
		if err != nil {
			fmt.Println("Unable to read the Pod Logs, error: ", err)
			return ctrl.Result{}, err
		}

		fmt.Println("Pod Logs: ", podLogsAns)

		problem.Status.Answer = podLogsAns
		err = r.Status().Update(ctx, &problem, &client.UpdateOptions{})
		if err != nil {
			fmt.Println("Error occured while updating the status field, error: ", err)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func readPodLogs(pod corev1.Pod) (string, error) {
	config := ctrl.GetConfigOrDie()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Unable to load the config, err:", err)
		return "", err
	}
	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &corev1.PodLogOptions{})
	logReader, err := req.Stream(context.Background())
	if err != nil {
		return "", err
	}

	defer logReader.Close()

	// bcz we don't know how much data comes from the log, we use buffer here
	logs, err := ioutil.ReadAll(logReader)
	if err != nil {
		return "", err
	}

	return string(logs), err
}

// SetupWithManager sets up the controller with the Manager.
func (r *ArithmeticReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mathsv1.Arithmetic{}).
		Complete(r)
}
