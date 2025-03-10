package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSMachineDeploymentSpec is the structure put into the provider spec of the
// Cluster API's MachineDeployment type. There it is tracked as serialized raw
// extension.
//
//     kind: AWSMachineDeploymentSpec
//     apiVersion: cluster.giantswarm.io/v1alpha1
//     metadata:
//       labels:
//         aws-operator.giantswarm.io/version: 6.2.0
//         cluster-operator.giantswarm.io/version: 0.17.0
//         giantswarm.io/cluster: 8y5kc
//         giantswarm.io/organization: "giantswarm"
//         giantswarm.io/machine-deployment: al9qy
//         release.giantswarm.io/version: 7.3.1
//       name: al9qy
//     nodePool:
//       description: my fancy node pool
//       machine:
//         dockerVolumeSizeGB: 100
//         kubeletVolumeSizeGB: 100
//       scaling:
//         max: 3
//         min: 3
//     provider:
//       availabilityZones:
//         - eu-central-1a
//       worker:
//         instanceType: m4.xlarge
//
type AWSMachineDeploymentSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodePool          AWSMachineDeploymentSpecNodePool `json:"nodePool" yaml:"nodePool"`
	Provider          AWSMachineDeploymentSpecProvider `json:"provider" yaml:"provider"`
}

type AWSMachineDeploymentSpecNodePool struct {
	Description string                                  `json:"description" yaml:"description"`
	Machine     AWSMachineDeploymentSpecNodePoolMachine `json:"machine" yaml:"machine"`
	Scaling     AWSMachineDeploymentSpecNodePoolScaling `json:"scaling" yaml:"scaling"`
}

type AWSMachineDeploymentSpecNodePoolMachine struct {
	DockerVolumeSizeGB  int `json:"dockerVolumeSizeGB" yaml:"dockerVolumeSizeGB"`
	KubeletVolumeSizeGB int `json:"kubeletVolumeSizeGB" yaml:"kubeletVolumeSizeGB"`
}

type AWSMachineDeploymentSpecNodePoolScaling struct {
	Max int `json:"max" yaml:"max"`
	Min int `json:"min" yaml:"min"`
}

type AWSMachineDeploymentSpecProvider struct {
	AvailabilityZones []string                               `json:"availabilityZones" yaml:"availabilityZones"`
	Worker            AWSMachineDeploymentSpecProviderWorker `json:"worker" yaml:"worker"`
}

type AWSMachineDeploymentSpecProviderWorker struct {
	InstanceType string `json:"instanceType" yaml:"instanceType"`
}
