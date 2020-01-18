/*
Copyright 2020 loch.

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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "github.com/leisheyoufu/k8svirtops/api/v1"
)

// VirtualMachineReconciler reconciles a VirtualMachine object
type VirtualMachineReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=infra.loch.com,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infra.loch.com,resources=virtualmachines/status,verbs=get;update;patch

func (r *VirtualMachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("virtualmachine", req.NamespacedName)

	// your logic here
	vm := &infrav1.VirtualMachine{}
	ctx := context.Background()
	if err := r.Get(ctx, req.NamespacedName, vm); err != nil {
		r.Log.Error(err, "unable to fetch vm")
	} else {
		vm.Status.Status = "Running"
		if err := r.Status().Update(ctx, vm); err != nil {
			r.Log.Error(err, "unable to update vm status")
		}
		fmt.Println(vm.Spec.CPU, vm.Spec.Memory)
	}
	return ctrl.Result{}, nil
}

func (r *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.VirtualMachine{}).
		Complete(r)
}
