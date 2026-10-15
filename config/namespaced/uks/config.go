// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package uks

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the uks (UpCloud Kubernetes Service) group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("upcloud_kubernetes_cluster", func(r *config.Resource) {
		r.UseAsync = true

		r.References["network"] = config.Reference{
			TerraformName: "upcloud_network",
		}
		// Publish the kubeconfig under the key provider-kubernetes and
		// provider-helm expect, in addition to upjet's
		// attribute.kubeconfig.kubeconfig. Upstream only fills the attribute
		// when fetch_kubeconfig is true.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			details, ok := attr["kubeconfig"].(map[string]any)
			if !ok {
				return nil, nil
			}
			kubeconfig, ok := details["kubeconfig"].(string)
			if !ok || kubeconfig == "" {
				return nil, nil
			}
			return map[string][]byte{"kubeconfig": []byte(kubeconfig)}, nil
		}
	})

	p.AddResourceConfigurator("upcloud_kubernetes_node_group", func(r *config.Resource) {
		r.AddSingletonListConversion("custom_plan", "customPlan")
		r.AddSingletonListConversion("gpu_plan", "gpuPlan")
		r.AddSingletonListConversion("cloud_native_plan", "cloudNativePlan")

		r.References["cluster"] = config.Reference{
			TerraformName: "upcloud_kubernetes_cluster",
		}
	})
}
