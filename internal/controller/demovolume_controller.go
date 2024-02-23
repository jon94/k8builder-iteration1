/*
Copyright 2024.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "operator-demo/api/v1"
)

// DemoVolumeReconciler reconciles a DemoVolume object
type DemoVolumeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.demo.jonlimpw.io,resources=demovolumes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.demo.jonlimpw.io,resources=demovolumes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.demo.jonlimpw.io,resources=demovolumes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DemoVolume object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *DemoVolumeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("Enter Reconcile", "req", req)

	// TODO(user): your logic here

	volume := &demov1.DemoVolume{}
	r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, volume)

	l.Info("Enter Reconcile", "spec", volume.Spec, "status", volume.Status)

	if volume.Spec.Name != volume.Status.Name {
		volume.Status.Name = volume.Spec.Name
		r.Status().Update(ctx, volume)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DemoVolumeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.DemoVolume{}).
		Complete(r)
}
