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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = &metav1.LabelSelector{}

// +kubebuilder:object:root=true

// PrometheusRole is a namespaced, logical grouping of PrometheusPolicyRule that can be referenced as a unit by a PrometheusRoleBinding.
type PrometheusRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Rules holds all the PrometheusPolicyRules for this Role
	// +optional
	Rules []PrometheusPolicyRule `json:"rules"`
}

// +kubebuilder:object:root=true

// PrometheusRoleList contains a list of PrometheusRole
type PrometheusRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrometheusRole `json:"items"`
}

// PrometheusPolicyRule holds information that describes a policy rule, but does not contain information
// about who the rule applies to or which namespace the rule applies to.
type PrometheusPolicyRule struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
	Verbs []string `json:"verbs"`

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path
	// Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
	// +optional
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`

	// A label query over metrics that should match this role
	// +optional
	MetricSelector *metav1.LabelSelector `json:"metricSelector,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PrometheusRole{}, &PrometheusRoleList{})
}
