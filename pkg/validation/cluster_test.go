/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package validation

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"testing"

	semverlib "github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/features"
	"k8c.io/kubermatic/v2/pkg/version"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"
)

var (
	dc = &kubermaticv1.Datacenter{
		Spec: kubermaticv1.DatacenterSpec{
			Openstack: &kubermaticv1.DatacenterSpecOpenstack{
				// Used for a test case
				EnforceFloatingIP: true,
			},
		},
	}
)

func TestValidateClusterSpec(t *testing.T) {
	tests := []struct {
		name  string
		spec  *kubermaticv1.ClusterSpec
		valid bool
	}{
		{
			name:  "empty spec should not cause a panic",
			valid: false,
			spec:  &kubermaticv1.ClusterSpec{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateClusterSpec(test.spec, dc, features.FeatureGate{}, version.New([]*version.Version{{
				Version: semverlib.MustParse("1.2.3"),
			}}, nil, nil), nil).ToAggregate()

			if (err == nil) != test.valid {
				t.Errorf("Extected err to be %v, got %v", test.valid, err)
			}
		})
	}
}

func TestValidateCloudSpec(t *testing.T) {
	tests := []struct {
		name  string
		spec  kubermaticv1.CloudSpec
		valid bool
	}{
		{
			name:  "valid openstack spec",
			valid: true,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "some-datacenter",
				ProviderName:   string(kubermaticv1.OpenstackCloudProvider),
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					Project:  "some-project",
					Username: "some-user",
					Password: "some-password",
					Domain:   "some-domain",
					// Required due to the above defined DC
					FloatingIPPool: "some-network",
				},
			},
		},
		{
			name:  "valid openstack spec - only projectID specified",
			valid: true,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "some-datacenter",
				ProviderName:   string(kubermaticv1.OpenstackCloudProvider),
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					ProjectID: "some-project",
					Username:  "some-user",
					Password:  "some-password",
					Domain:    "some-domain",
					// Required due to the above defined DC
					FloatingIPPool: "some-network",
				},
			},
		},
		{
			name:  "invalid openstack spec - no datacenter specified",
			valid: false,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "",
				ProviderName:   string(kubermaticv1.OpenstackCloudProvider),
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					Project:  "some-project",
					Username: "some-user",
					Password: "some-password",
					Domain:   "some-domain",
					// Required due to the above defined DC
					FloatingIPPool: "some-network",
				},
			},
		},
		{
			name:  "invalid openstack spec - no floating ip pool defined but required by dc",
			valid: false,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "some-datacenter",
				ProviderName:   string(kubermaticv1.OpenstackCloudProvider),
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					Project:        "some-project",
					Username:       "some-user",
					Password:       "some-password",
					Domain:         "some-domain",
					FloatingIPPool: "",
				},
			},
		},
		{
			name:  "specifies multiple cloud providers",
			valid: false,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "some-datacenter",
				ProviderName:   string(kubermaticv1.OpenstackCloudProvider),
				Digitalocean: &kubermaticv1.DigitaloceanCloudSpec{
					Token: "a-token",
				},
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					Project:        "some-project",
					Username:       "some-user",
					Password:       "some-password",
					Domain:         "some-domain",
					FloatingIPPool: "",
				},
			},
		},
		{
			name:  "valid provider name",
			valid: true,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "some-datacenter",
				ProviderName:   string(kubermaticv1.OpenstackCloudProvider),
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					Project:        "some-project",
					Username:       "some-user",
					Password:       "some-password",
					Domain:         "some-domain",
					FloatingIPPool: "some-network",
				},
			},
		},
		{
			name:  "invalid provider name",
			valid: false,
			spec: kubermaticv1.CloudSpec{
				DatacenterName: "some-datacenter",
				ProviderName:   "closedstack", // *giggle*
				Openstack: &kubermaticv1.OpenstackCloudSpec{
					Project:        "some-project",
					Username:       "some-user",
					Password:       "some-password",
					Domain:         "some-domain",
					FloatingIPPool: "some-network",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateCloudSpec(test.spec, dc, kubermaticv1.IPFamilyIPv4, nil, true).ToAggregate()

			if (err == nil) != test.valid {
				t.Errorf("Extected err to be %v, got %v", test.valid, err)
			}
		})
	}
}

func TestValidateUpdateWindow(t *testing.T) {
	tests := []struct {
		name         string
		updateWindow kubermaticv1.UpdateWindow
		err          error
	}{
		{
			name: "valid update window",
			updateWindow: kubermaticv1.UpdateWindow{
				Start:  "04:00",
				Length: "1h",
			},
			err: nil,
		},
		{
			name: "invalid start date",
			updateWindow: kubermaticv1.UpdateWindow{
				Start:  "invalid",
				Length: "1h",
			},
			err: errors.New("invalid time of day"),
		},
		{
			name: "invalid length",
			updateWindow: kubermaticv1.UpdateWindow{
				Start:  "Thu 04:00",
				Length: "1",
			},
			err: errors.New("missing unit in duration"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateUpdateWindow(&test.updateWindow)
			if (err != nil) != (test.err != nil) {
				t.Errorf("Extected err to be %v, got %v", test.err, err)
			}

			// loosely validate the returned error message
			if test.err != nil && !strings.Contains(err.Error(), test.err.Error()) {
				t.Errorf("Extected err to contain \"%v\", but got \"%v\"", test.err, err)
			}
		})
	}
}

func TestValidateLeaderElectionSettings(t *testing.T) {
	tests := []struct {
		name                   string
		leaderElectionSettings kubermaticv1.LeaderElectionSettings
		wantErr                bool
	}{
		{
			name:                   "empty leader election settings",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{},
			wantErr:                false,
		},
		{
			name: "valid leader election settings",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{
				LeaseDurationSeconds: pointer.Int32Ptr(int32(10)),
				RenewDeadlineSeconds: pointer.Int32Ptr(int32(5)),
				RetryPeriodSeconds:   pointer.Int32Ptr(int32(10)),
			},
			wantErr: false,
		},
		{
			name: "invalid leader election settings",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{
				LeaseDurationSeconds: pointer.Int32Ptr(int32(5)),
				RenewDeadlineSeconds: pointer.Int32Ptr(int32(10)),
				RetryPeriodSeconds:   pointer.Int32Ptr(int32(10)),
			},
			wantErr: true,
		},
		{
			name: "lease duration only",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{
				LeaseDurationSeconds: pointer.Int32Ptr(int32(10)),
			},
			wantErr: true,
		},
		{
			name: "renew duration only",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{
				RenewDeadlineSeconds: pointer.Int32Ptr(int32(10)),
			},
			wantErr: true,
		},
		{
			name: "retry period only",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{
				RetryPeriodSeconds: pointer.Int32Ptr(int32(10)),
			},
			wantErr: false,
		},
		{
			name: "negative value",
			leaderElectionSettings: kubermaticv1.LeaderElectionSettings{
				LeaseDurationSeconds: pointer.Int32Ptr(int32(10)),
				RenewDeadlineSeconds: pointer.Int32Ptr(int32(-5)),
				RetryPeriodSeconds:   pointer.Int32Ptr(int32(10)),
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errs := ValidateLeaderElectionSettings(&test.leaderElectionSettings, field.NewPath("spec"))

			if test.wantErr == (len(errs) == 0) {
				t.Errorf("Want error: %t, but got: \"%v\"", test.wantErr, errs)
			}
		})
	}
}

func TestValidateClusterNetworkingConfig(t *testing.T) {
	tests := []struct {
		name          string
		networkConfig kubermaticv1.ClusterNetworkingConfig
		dc            *kubermaticv1.Datacenter
		wantErr       bool
	}{
		{
			name:          "empty network config",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{},
			wantErr:       true,
		},
		{
			name: "valid network config",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "missing pods CIDR",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "missing services CIDR",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "valid dual-stack network config",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "invalid dual-stack network config (IPv6 as primary address)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"fd00::/104", "10.241.0.0/16"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"fd03::/120", "10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "invalid dual-stack network config (missing IPv6 services CIDR)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "valid ip family - IPv4",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyIPv4,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "valid ip family - dual stack",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "invalid ip family",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "valid node CIDR mask sizes",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				NodeCIDRMaskSizeIPv4:     pointer.Int32(26),
				NodeCIDRMaskSizeIPv6:     pointer.Int32(112),
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "invalid node CIDR mask size - IPv4",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				NodeCIDRMaskSizeIPv4:     pointer.Int32(12),
				NodeCIDRMaskSizeIPv6:     pointer.Int32(112),
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "invalid node CIDR mask size - IPv6",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				NodeCIDRMaskSizeIPv4:     pointer.Int32(24),
				NodeCIDRMaskSizeIPv6:     pointer.Int32(64),
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "missing DNS domain",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "missing proxy mode",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20"}},
				DNSDomain:                "cluster.local",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "invalid pod cidr",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Pods: kubermaticv1.NetworkRanges{CIDRBlocks: []string{"192.127.0.0:20"}},
			},
			wantErr: true,
		},
		{
			name: "invalid service cidr",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				Services: kubermaticv1.NetworkRanges{CIDRBlocks: []string{"192.127/20"}},
			},
			wantErr: true,
		},
		{
			name: "invalid DNS domain",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				DNSDomain: "cluster.bla",
			},
			wantErr: true,
		},
		{
			name: "invalid proxy mode",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				ProxyMode: "none",
			},
			wantErr: true,
		},
		{
			name: "valid dual-stack datacenter config",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					AWS: &kubermaticv1.DatacenterSpecAWS{},
				},
			},
			wantErr: false,
		},
		{
			name: "valid dual-stack datacenter config (openstack)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					Openstack: &kubermaticv1.DatacenterSpecOpenstack{
						IPv6Enabled: pointer.BoolPtr(true),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid dual-stack datacenter config (IPv6 not enabled for openstack datacenter)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					Openstack: &kubermaticv1.DatacenterSpecOpenstack{
						IPv6Enabled: pointer.BoolPtr(false),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid dual-stack datacenter config (vsphere)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					VSphere: &kubermaticv1.DatacenterSpecVSphere{
						IPv6Enabled: pointer.BoolPtr(true),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid dual-stack datacenter config (IPv6 not enabled for vsphere datacenter)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					VSphere: &kubermaticv1.DatacenterSpecVSphere{
						IPv6Enabled: pointer.BoolPtr(false),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid dual-stack datacenter config (not known ipv6 cloud provider)",
			networkConfig: kubermaticv1.ClusterNetworkingConfig{
				IPFamily:                 kubermaticv1.IPFamilyDualStack,
				Pods:                     kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.241.0.0/16", "fd00::/104"}},
				Services:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"10.240.32.0/20", "fd03::/120"}},
				DNSDomain:                "cluster.local",
				ProxyMode:                "ipvs",
				NodeLocalDNSCacheEnabled: pointer.BoolPtr(true),
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					Digitalocean: &kubermaticv1.DatacenterSpecDigitalocean{},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errs := ValidateClusterNetworkConfig(&test.networkConfig, test.dc, nil, field.NewPath("spec", "networkConfig"))
			if test.wantErr == (len(errs) == 0) {
				t.Errorf("Want error: %t, but got: \"%v\"", test.wantErr, errs)
			}
		})
	}
}

func TestValidateGCPCloudSpec(t *testing.T) {
	testCases := []struct {
		name              string
		spec              *kubermaticv1.GCPCloudSpec
		dc                *kubermaticv1.Datacenter
		ipFamily          kubermaticv1.IPFamily
		gcpSubnetworkResp apiv1.GCPSubnetwork
		expectedError     error
	}{
		{
			name: "valid ipv4 gcp spec",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily: kubermaticv1.IPFamilyIPv4,
		},
		{
			name: "invalid gcp spec: service account cannot be empty",
			spec: &kubermaticv1.GCPCloudSpec{
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyIPv4,
			expectedError: errors.New("\"serviceAccount\" cannot be empty"),
		},
		{
			name: "invalid gcp spec: NodePortsAllowedIPRange",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "invalid",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyIPv4,
			expectedError: &net.ParseError{Type: "CIDR address", Text: "invalid"},
		},
		{
			name: "invalid gcp spec: NodePortsAllowedIPRanges",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"invalid",
					},
				},
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyIPv4,
			expectedError: fmt.Errorf("unable to parse CIDR \"invalid\": %w", &net.ParseError{Type: "CIDR address", Text: "invalid"}),
		},
		{
			name: "invalid dual-stack gcp spec: empty network",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyDualStack,
			expectedError: errors.New("network and subnetwork should be defined for GCP dual-stack (IPv4 + IPv6) cluster"),
		},
		{
			name: "invalid dual-stack gcp spec: empty subnetwork",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
				Network: "global/networks/dualstack",
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyDualStack,
			expectedError: errors.New("network and subnetwork should be defined for GCP dual-stack (IPv4 + IPv6) cluster"),
		},
		{
			name: "invalid dual-stack gcp spec: invalid subnetwork path",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
				Network:    "global/networks/dualstack",
				Subnetwork: "invalid",
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyDualStack,
			expectedError: errors.New("invalid GCP subnetwork path"),
		},
		{
			name: "invalid dual-stack gcp spec: wrong region",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
				Network:    "global/networks/dualstack",
				Subnetwork: "projects/kubermatic-dev/regions/europe-west2/subnetworks/dualstack-europe-west2",
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west3",
					},
				},
			},
			ipFamily:      kubermaticv1.IPFamilyDualStack,
			expectedError: errors.New("GCP subnetwork should belong to same cluster region"),
		},
		{
			name: "valid gcp dual-stack spec",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
				Network:    "global/networks/dualstack",
				Subnetwork: "projects/kubermatic-dev/regions/europe-west2/subnetworks/dualstack-europe-west2",
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west2",
					},
				},
			},
			ipFamily: kubermaticv1.IPFamilyDualStack,
			gcpSubnetworkResp: apiv1.GCPSubnetwork{
				IPFamily: kubermaticv1.IPFamilyDualStack,
			},
		},
		{
			name: "invalid gcp dual-stack spec: wrong network stack type",
			spec: &kubermaticv1.GCPCloudSpec{
				ServiceAccount:          "service-account",
				NodePortsAllowedIPRange: "0.0.0.0/0",
				NodePortsAllowedIPRanges: &kubermaticv1.NetworkRanges{
					CIDRBlocks: []string{
						"0.0.0.0/0",
						"::/0",
					},
				},
				Network:    "global/networks/default",
				Subnetwork: "projects/kubermatic-dev/regions/europe-west2/subnetworks/default",
			},
			dc: &kubermaticv1.Datacenter{
				Spec: kubermaticv1.DatacenterSpec{
					GCP: &kubermaticv1.DatacenterSpecGCP{
						Region: "europe-west2",
					},
				},
			},
			ipFamily: kubermaticv1.IPFamilyDualStack,
			gcpSubnetworkResp: apiv1.GCPSubnetwork{
				IPFamily: kubermaticv1.IPFamilyIPv4,
			},
			expectedError: errors.New("GCP subnetwork should belong to same cluster network stack type"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateGCPCloudSpec(tc.spec, tc.dc, tc.ipFamily, func(ctx context.Context, sa, region, subnetworkName string) (apiv1.GCPSubnetwork, error) {
				return tc.gcpSubnetworkResp, nil
			})
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestValidateEncryptionConfiguration(t *testing.T) {
	tests := []struct {
		name        string
		clusterSpec *kubermaticv1.ClusterSpec
		expectErr   field.ErrorList
	}{
		{
			name: "small key",
			clusterSpec: &kubermaticv1.ClusterSpec{
				Features: map[string]bool{
					kubermaticv1.ClusterFeatureEncryptionAtRest: true,
				},
				EncryptionConfiguration: &kubermaticv1.EncryptionConfiguration{
					Enabled: true,
					Secretbox: &kubermaticv1.SecretboxEncryptionConfiguration{
						Keys: []kubermaticv1.SecretboxKey{
							{
								Name:  "small-key",
								Value: "cmLcMbw6gdxPHQ==",
							},
						},
					},
				},
			},
			expectErr: field.ErrorList{
				&field.Error{
					Type:     "FieldValueInvalid",
					Field:    "spec.encryptionConfiguration.secretbox.keys[0]",
					BadValue: "small-key",
					Detail:   "key length should be 32 it is 10",
				},
			},
		},
		{
			name: "bad base64",
			clusterSpec: &kubermaticv1.ClusterSpec{
				Features: map[string]bool{
					kubermaticv1.ClusterFeatureEncryptionAtRest: true,
				},
				EncryptionConfiguration: &kubermaticv1.EncryptionConfiguration{
					Enabled: true,
					Secretbox: &kubermaticv1.SecretboxEncryptionConfiguration{
						Keys: []kubermaticv1.SecretboxKey{
							{
								Name:  "small-key",
								Value: "cmLcMbw6gdxPH$==",
							},
						},
					},
				},
			},
			expectErr: field.ErrorList{
				&field.Error{
					Type:     "FieldValueInvalid",
					Field:    "spec.encryptionConfiguration.secretbox.keys[0]",
					BadValue: "small-key",
					Detail:   "illegal base64 data at input byte 13",
				},
			},
		},
		{
			name: "good key",
			clusterSpec: &kubermaticv1.ClusterSpec{
				Features: map[string]bool{
					kubermaticv1.ClusterFeatureEncryptionAtRest: true,
				},
				EncryptionConfiguration: &kubermaticv1.EncryptionConfiguration{
					Enabled: true,
					Secretbox: &kubermaticv1.SecretboxEncryptionConfiguration{
						Keys: []kubermaticv1.SecretboxKey{
							{
								Name:  "good-key",
								Value: "RGolflgAc+eBbm1lys87pTNQZVf0i67rlpPZGtTkVjQ=",
							},
						},
					},
				},
			},
			expectErr: field.ErrorList{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateEncryptionConfiguration(test.clusterSpec, field.NewPath("spec", "encryptionConfiguration"))
			assert.Equal(t, test.expectErr, err)
		})
	}
}
