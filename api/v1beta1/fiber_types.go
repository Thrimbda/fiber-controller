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

package v1beta1

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FiberSpec defines the desired state of Fiber
type FiberSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Fiber. Edit Fiber_types.go to remove/update
	FiberSpeak string               `json:"foo,omitempty"`
	Deployment appsv1.Deployment    `json:"deployment,omitempty"`
	CronJob    batchv1beta1.CronJob `json:"cronJob,omitempty"`
}

// FiberStatus defines the observed state of Fiber
type FiberStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LastFib    int64 `json:"lastFib"`
	CurrentFib int64 `json:"currentFib"`
}

// +kubebuilder:object:root=true

// Fiber is the Schema for the fibers API
type Fiber struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FiberSpec   `json:"spec,omitempty"`
	Status FiberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FiberList contains a list of Fiber
type FiberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fiber `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Fiber{}, &FiberList{})
}
