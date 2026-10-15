// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
)

func TestKubernetesClusterKubeconfigConnectionDetail(t *testing.T) {
	cluster, err := GetProvider(context.Background())
	if err != nil {
		t.Fatalf("GetProvider(): %v", err)
	}
	namespaced, err := GetProviderNamespaced(context.Background())
	if err != nil {
		t.Fatalf("GetProviderNamespaced(): %v", err)
	}
	providers := map[string]*ujconfig.Provider{"Cluster": cluster, "Namespaced": namespaced}

	cases := map[string]struct {
		args map[string]any
		want map[string][]byte
	}{
		"KubeconfigFetched": {
			args: map[string]any{"kubeconfig": map[string]any{
				"kubeconfig": "apiVersion: v1\nkind: Config\n",
				"host":       "https://lb-0a19c0ca.upcloudlb.com:6443",
			}},
			want: map[string][]byte{"kubeconfig": []byte("apiVersion: v1\nkind: Config\n")},
		},
		"FetchKubeconfigDisabled": {
			args: map[string]any{"kubeconfig": nil},
			want: nil,
		},
		"EmptyKubeconfig": {
			args: map[string]any{"kubeconfig": map[string]any{"kubeconfig": ""}},
			want: nil,
		},
	}
	for scope, p := range providers {
		fn := p.Resources["upcloud_kubernetes_cluster"].Sensitive.AdditionalConnectionDetailsFn
		for name, tc := range cases {
			t.Run(scope+"/"+name, func(t *testing.T) {
				got, err := fn(tc.args)
				if err != nil {
					t.Fatalf("AdditionalConnectionDetailsFn(): %v", err)
				}
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("AdditionalConnectionDetailsFn(): -want, +got:\n%s", diff)
				}
			})
		}
	}
}
