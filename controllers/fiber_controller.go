/*


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
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	thecoresv1beta1 "github.com/thrimbda/fiber-controller/api/v1beta1"
)

// FiberReconciler reconciles a Fiber object
type FiberReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=thecores.thrimbda.com,resources=fibers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=thecores.thrimbda.com,resources=fibers/status,verbs=get;update;patch

func (r *FiberReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("fiber", req.NamespacedName)

	var fiber *thecoresv1beta1.Fiber

	err := r.Get(context.TODO(), req.NamespacedName, fiber)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	return r.reconcile(fiber)
}

func (r *FiberReconciler) reconcile(fiber *thecoresv1beta1.Fiber) (ctrl.Result, error) {

	return ctrl.Result{}, nil
}

func (r *FiberReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&thecoresv1beta1.Fiber{}).
		Complete(r)
}
