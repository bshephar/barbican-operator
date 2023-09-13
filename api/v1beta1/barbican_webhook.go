/*
Copyright 2023.

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

//
// Generated by:
//
// operator-sdk create webhook --group barbican --version v1beta1 --kind Barbican --programmatic-validation --defaulting
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// BarbicanDefaults -
type BarbicanDefaults struct {
	APIContainerImageURL    string
	WorkerContainerImageURL string
}

var barbicanDefaults BarbicanDefaults

// log is for logging in this package.
var barbicanlog = logf.Log.WithName("barbican-resource")

// SetupBarbicanDefaults - initialize Barbican spec defaults for use with either internal or external webhooks
func SetupBarbicanDefaults(defaults BarbicanDefaults) {
	barbicanDefaults = defaults
	barbicanlog.Info("Barbican defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *Barbican) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-barbican-openstack-org-v1beta1-barbican,mutating=true,failurePolicy=fail,sideEffects=None,groups=barbican.openstack.org,resources=barbicans,verbs=create;update,versions=v1beta1,name=mbarbican.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Barbican{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Barbican) Default() {
	barbicanlog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this Barbican spec
func (spec *BarbicanSpec) Default() {
	if spec.BarbicanAPI.ContainerImage == "" {
		spec.BarbicanAPI.ContainerImage = barbicanDefaults.APIContainerImageURL
	}

	if spec.BarbicanWorker.ContainerImage == "" {
		spec.BarbicanWorker.ContainerImage = barbicanDefaults.WorkerContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-barbican-openstack-org-v1beta1-barbican,mutating=false,failurePolicy=fail,sideEffects=None,groups=barbican.openstack.org,resources=barbicans,verbs=create;update,versions=v1beta1,name=vbarbican.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Barbican{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Barbican) ValidateCreate() error {
	barbicanlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Barbican) ValidateUpdate(old runtime.Object) error {
	barbicanlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Barbican) ValidateDelete() error {
	barbicanlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}