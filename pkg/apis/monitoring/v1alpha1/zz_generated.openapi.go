// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfg":       schema_pkg_apis_monitoring_v1alpha1_AlertMgrCfg(ref),
		"github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfgSpec":   schema_pkg_apis_monitoring_v1alpha1_AlertMgrCfgSpec(ref),
		"github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfgStatus": schema_pkg_apis_monitoring_v1alpha1_AlertMgrCfgStatus(ref),
		"github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.Param":             schema_pkg_apis_monitoring_v1alpha1_Param(ref),
	}
}

func schema_pkg_apis_monitoring_v1alpha1_AlertMgrCfg(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AlertMgrCfg is the Schema for the alertmgrcfgs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfgSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfgStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfgSpec", "github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.AlertMgrCfgStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_monitoring_v1alpha1_AlertMgrCfgSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AlertMgrCfgSpec defines the desired state of AlertMgrCfg",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"params": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.Param"),
									},
								},
							},
						},
					},
				},
				Required: []string{"type"},
			},
		},
		Dependencies: []string{
			"github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1.Param"},
	}
}

func schema_pkg_apis_monitoring_v1alpha1_AlertMgrCfgStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AlertMgrCfgStatus defines the observed state of AlertMgrCfg",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_monitoring_v1alpha1_Param(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Param is a list of alerting receivers.",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "value"},
			},
		},
		Dependencies: []string{},
	}
}
