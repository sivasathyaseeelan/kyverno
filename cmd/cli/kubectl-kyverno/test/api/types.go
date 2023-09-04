package api

import (
	policyreportv1alpha2 "github.com/kyverno/kyverno/api/policyreport/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Test struct {
	Name      string        `json:"name"`
	Policies  []string      `json:"policies"`
	Resources []string      `json:"resources"`
	Variables string        `json:"variables,omitempty"`
	UserInfo  string        `json:"userinfo,omitempty"`
	Results   []TestResults `json:"results"`
	Values    *Values       `json:"values,omitempty"`
}

type TestResults struct {
	// Policy mentions the name of the policy.
	Policy string `json:"policy"`
	// Rule mentions the name of the rule in the policy.
	// It's required in case policy is a kyverno policy.
	// +optional
	Rule string `json:"rule,omitempty"`
	// IsValidatingAdmissionPolicy indicates if the policy is a validating admission policy.
	// It's required in case policy is a validating admission policy.
	// +optional
	IsValidatingAdmissionPolicy bool `json:"isValidatingAdmissionPolicy,omitempty"`
	// Result mentions the result that the user is expecting.
	// Possible values are pass, fail and skip.
	Result policyreportv1alpha2.PolicyResult `json:"result"`
	// Status mentions the status that the user is expecting.
	// Possible values are pass, fail and skip.
	Status policyreportv1alpha2.PolicyResult `json:"status,omitempty"`
	// Resource mentions the name of the resource on which the policy is to be applied.
	Resource string `json:"resource,omitempty"`
	// Resources gives us the list of resources on which the policy is going to be applied.
	Resources []string `json:"resources"`
	// Kind mentions the kind of the resource on which the policy is to be applied.
	Kind string `json:"kind"`
	// Namespace mentions the namespace of the policy which has namespace scope.
	Namespace string `json:"namespace,omitempty"`
	// PatchedResource takes a resource configuration file in yaml format from
	// the user to compare it against the Kyverno mutated resource configuration.
	PatchedResource string `json:"patchedResource,omitempty"`
	// GeneratedResource takes a resource configuration file in yaml format from
	// the user to compare it against the Kyverno generated resource configuration.
	GeneratedResource string `json:"generatedResource,omitempty"`
	// CloneSourceResource takes the resource configuration file in yaml format
	// from the user which is meant to be cloned by the generate rule.
	CloneSourceResource string `json:"cloneSourceResource,omitempty"`
}

type Policy struct {
	Name      string     `json:"name"`
	Resources []Resource `json:"resources"`
	Rules     []Rule     `json:"rules"`
}

type Rule struct {
	Name          string                   `json:"name"`
	Values        map[string]interface{}   `json:"values"`
	ForeachValues map[string][]interface{} `json:"foreachValues"`
}

type Values struct {
	Policies           []Policy            `json:"policies"`
	GlobalValues       map[string]string   `json:"globalValues"`
	NamespaceSelectors []NamespaceSelector `json:"namespaceSelector"`
	Subresources       []Subresource       `json:"subresources"`
}

type Resource struct {
	Name   string                 `json:"name"`
	Values map[string]interface{} `json:"values"`
}

type Subresource struct {
	APIResource    metav1.APIResource `json:"subresource"`
	ParentResource metav1.APIResource `json:"parentResource"`
}

type NamespaceSelector struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}