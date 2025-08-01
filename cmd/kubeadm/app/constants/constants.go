/*
Copyright 2019 The Kubernetes Authors.

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

package constants

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/version"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	componentversion "k8s.io/component-base/version"
	netutils "k8s.io/utils/net"

	"k8s.io/kubernetes/cmd/kubeadm/app/util/errors"
)

const (
	// KubernetesDir is the directory Kubernetes owns for storing various configuration files
	KubernetesDir = "/etc/kubernetes"
	// ManifestsSubDirName defines directory name to store manifests
	ManifestsSubDirName = "manifests"
	// TempDir defines temporary directory for kubeadm
	// should be joined with KubernetesDir.
	TempDir = "tmp"

	// CertificateBackdate defines the offset applied to notBefore for CA certificates generated by kubeadm
	CertificateBackdate = time.Minute * 5
	// CertificateValidityPeriod defines the validity period for all the signed certificates generated by kubeadm
	CertificateValidityPeriod = time.Hour * 24 * 365
	// CACertificateValidityPeriod defines the validity period for all the signed CA certificates generated by kubeadm
	CACertificateValidityPeriod = time.Hour * 24 * 365 * 10

	// DefaultCertificateDir defines default certificate directory
	DefaultCertificateDir = "pki"

	// CACertAndKeyBaseName defines certificate authority base name
	CACertAndKeyBaseName = "ca"
	// CACertName defines certificate name
	CACertName = "ca.crt"
	// CAKeyName defines certificate name
	CAKeyName = "ca.key"

	// APIServerCertAndKeyBaseName defines API's server certificate and key base name
	APIServerCertAndKeyBaseName = "apiserver"
	// APIServerCertName defines API's server certificate name
	APIServerCertName = "apiserver.crt"
	// APIServerKeyName defines API's server key name
	APIServerKeyName = "apiserver.key"
	// APIServerCertCommonName defines API's server certificate common name (CN)
	APIServerCertCommonName = "kube-apiserver"

	// APIServerKubeletClientCertAndKeyBaseName defines kubelet client certificate and key base name
	APIServerKubeletClientCertAndKeyBaseName = "apiserver-kubelet-client"
	// APIServerKubeletClientCertName defines kubelet client certificate name
	APIServerKubeletClientCertName = "apiserver-kubelet-client.crt"
	// APIServerKubeletClientKeyName defines kubelet client key name
	APIServerKubeletClientKeyName = "apiserver-kubelet-client.key"
	// APIServerKubeletClientCertCommonName defines kubelet client certificate common name (CN)
	APIServerKubeletClientCertCommonName = "kube-apiserver-kubelet-client"

	// EtcdCACertAndKeyBaseName defines etcd's CA certificate and key base name
	EtcdCACertAndKeyBaseName = "etcd/ca"
	// EtcdCACertName defines etcd's CA certificate name
	EtcdCACertName = "etcd/ca.crt"
	// EtcdCAKeyName defines etcd's CA key name
	EtcdCAKeyName = "etcd/ca.key"

	// EtcdServerCertAndKeyBaseName defines etcd's server certificate and key base name
	EtcdServerCertAndKeyBaseName = "etcd/server"
	// EtcdServerCertName defines etcd's server certificate name
	EtcdServerCertName = "etcd/server.crt"
	// EtcdServerKeyName defines etcd's server key name
	EtcdServerKeyName = "etcd/server.key"

	// EtcdListenClientPort defines the port etcd listen on for client traffic
	EtcdListenClientPort = 2379
	// EtcdMetricsPort is the port at which to obtain etcd metrics and health status
	EtcdMetricsPort = 2381

	// EtcdPeerCertAndKeyBaseName defines etcd's peer certificate and key base name
	EtcdPeerCertAndKeyBaseName = "etcd/peer"
	// EtcdPeerCertName defines etcd's peer certificate name
	EtcdPeerCertName = "etcd/peer.crt"
	// EtcdPeerKeyName defines etcd's peer key name
	EtcdPeerKeyName = "etcd/peer.key"

	// EtcdListenPeerPort defines the port etcd listen on for peer traffic
	EtcdListenPeerPort = 2380

	// EtcdHealthcheckClientCertAndKeyBaseName defines etcd's healthcheck client certificate and key base name
	EtcdHealthcheckClientCertAndKeyBaseName = "etcd/healthcheck-client"
	// EtcdHealthcheckClientCertName defines etcd's healthcheck client certificate name
	EtcdHealthcheckClientCertName = "etcd/healthcheck-client.crt"
	// EtcdHealthcheckClientKeyName defines etcd's healthcheck client key name
	EtcdHealthcheckClientKeyName = "etcd/healthcheck-client.key"
	// EtcdHealthcheckClientCertCommonName defines etcd's healthcheck client certificate common name (CN)
	EtcdHealthcheckClientCertCommonName = "kube-etcd-healthcheck-client"

	// APIServerEtcdClientCertAndKeyBaseName defines apiserver's etcd client certificate and key base name
	APIServerEtcdClientCertAndKeyBaseName = "apiserver-etcd-client"
	// APIServerEtcdClientCertName defines apiserver's etcd client certificate name
	APIServerEtcdClientCertName = "apiserver-etcd-client.crt"
	// APIServerEtcdClientKeyName defines apiserver's etcd client key name
	APIServerEtcdClientKeyName = "apiserver-etcd-client.key"
	// APIServerEtcdClientCertCommonName defines apiserver's etcd client certificate common name (CN)
	APIServerEtcdClientCertCommonName = "kube-apiserver-etcd-client"

	// ServiceAccountKeyBaseName defines SA key base name
	ServiceAccountKeyBaseName = "sa"
	// ServiceAccountPublicKeyName defines SA public key base name
	ServiceAccountPublicKeyName = "sa.pub"
	// ServiceAccountPrivateKeyName defines SA private key base name
	ServiceAccountPrivateKeyName = "sa.key"

	// FrontProxyCACertAndKeyBaseName defines front proxy CA certificate and key base name
	FrontProxyCACertAndKeyBaseName = "front-proxy-ca"
	// FrontProxyCACertName defines front proxy CA certificate name
	FrontProxyCACertName = "front-proxy-ca.crt"
	// FrontProxyCAKeyName defines front proxy CA key name
	FrontProxyCAKeyName = "front-proxy-ca.key"

	// FrontProxyClientCertAndKeyBaseName defines front proxy certificate and key base name
	FrontProxyClientCertAndKeyBaseName = "front-proxy-client"
	// FrontProxyClientCertName defines front proxy certificate name
	FrontProxyClientCertName = "front-proxy-client.crt"
	// FrontProxyClientKeyName defines front proxy key name
	FrontProxyClientKeyName = "front-proxy-client.key"
	// FrontProxyClientCertCommonName defines front proxy certificate common name
	FrontProxyClientCertCommonName = "front-proxy-client" //used as subject.commonname attribute (CN)

	// AdminKubeConfigFileName defines name for the kubeconfig aimed to be used by the admin of the cluster
	AdminKubeConfigFileName = "admin.conf"
	// SuperAdminKubeConfigFileName defines name for the kubeconfig aimed to be used by the super-admin of the cluster
	SuperAdminKubeConfigFileName = "super-admin.conf"

	// KubeletBootstrapKubeConfigFileName defines the file name for the kubeconfig that the kubelet will use to do
	// the TLS bootstrap to get itself an unique credential
	KubeletBootstrapKubeConfigFileName = "bootstrap-kubelet.conf"

	// KubeletKubeConfigFileName defines the file name for the kubeconfig that the control-plane kubelet will use for talking
	// to the API server
	KubeletKubeConfigFileName = "kubelet.conf"
	// ControllerManagerKubeConfigFileName defines the file name for the controller manager's kubeconfig file
	ControllerManagerKubeConfigFileName = "controller-manager.conf"
	// SchedulerKubeConfigFileName defines the file name for the scheduler's kubeconfig file
	SchedulerKubeConfigFileName = "scheduler.conf"

	// Some well-known users, groups, roles and clusterrolebindings in the core Kubernetes authorization system

	// ControllerManagerUser defines the well-known user the controller-manager should be authenticated as
	ControllerManagerUser = "system:kube-controller-manager"
	// SchedulerUser defines the well-known user the scheduler should be authenticated as
	SchedulerUser = "system:kube-scheduler"
	// NodesUserPrefix defines the user name prefix as requested by the Node authorizer.
	NodesUserPrefix = "system:node:"
	// SystemPrivilegedGroup defines the well-known group for the apiservers. This group is also superuser by default
	// (i.e. bound to the cluster-admin ClusterRole)
	SystemPrivilegedGroup = "system:masters"
	// NodesGroup defines the well-known group for all nodes.
	NodesGroup = "system:nodes"
	// NodeBootstrapTokenAuthGroup specifies which group a Node Bootstrap Token should be authenticated in
	NodeBootstrapTokenAuthGroup = "system:bootstrappers:kubeadm:default-node-token"
	// KubeProxyClusterRoleName sets the name for the kube-proxy ClusterRole
	KubeProxyClusterRoleName = "system:node-proxier"
	// NodeBootstrapperClusterRoleName defines the name of the auto-bootstrapped ClusterRole for letting someone post a CSR
	NodeBootstrapperClusterRoleName = "system:node-bootstrapper"
	// CSRAutoApprovalClusterRoleName defines the name of the auto-bootstrapped ClusterRole for making the csrapprover controller auto-approve the CSR
	// Starting from v1.8, CSRAutoApprovalClusterRoleName is automatically created by the API server on startup
	CSRAutoApprovalClusterRoleName = "system:certificates.k8s.io:certificatesigningrequests:nodeclient"
	// NodeSelfCSRAutoApprovalClusterRoleName is a role defined in default 1.8 RBAC policies for automatic CSR approvals for automatically rotated node certificates
	NodeSelfCSRAutoApprovalClusterRoleName = "system:certificates.k8s.io:certificatesigningrequests:selfnodeclient"
	// NodesClusterRoleBinding defines the well-known ClusterRoleBinding which binds the too permissive system:node
	// ClusterRole to the system:nodes group. Since kubeadm is using the Node Authorizer, this ClusterRoleBinding's
	// system:nodes group subject is removed if present.
	NodesClusterRoleBinding = "system:node"

	// KubeletBaseConfigMapRole defines the base kubelet configuration ConfigMap.
	KubeletBaseConfigMapRole = "kubeadm:kubelet-config"
	// KubeProxyClusterRoleBindingName sets the name for the kube-proxy ClusterRoleBinding
	KubeProxyClusterRoleBindingName = "kubeadm:node-proxier"
	// NodeKubeletBootstrap defines the name of the ClusterRoleBinding that lets kubelets post CSRs
	NodeKubeletBootstrap = "kubeadm:kubelet-bootstrap"
	// GetNodesClusterRoleName defines the name of the ClusterRole and ClusterRoleBinding to get nodes
	GetNodesClusterRoleName = "kubeadm:get-nodes"
	// NodeAutoApproveBootstrapClusterRoleBinding defines the name of the ClusterRoleBinding that makes the csrapprover approve node CSRs
	NodeAutoApproveBootstrapClusterRoleBinding = "kubeadm:node-autoapprove-bootstrap"
	// NodeAutoApproveCertificateRotationClusterRoleBinding defines name of the ClusterRoleBinding that makes the csrapprover approve node auto rotated CSRs
	NodeAutoApproveCertificateRotationClusterRoleBinding = "kubeadm:node-autoapprove-certificate-rotation"
	// ClusterAdminsGroupAndClusterRoleBinding is the name of the Group used for kubeadm generated cluster
	// admin credentials and the name of the ClusterRoleBinding that binds the same Group to the "cluster-admin"
	// built-in ClusterRole.
	ClusterAdminsGroupAndClusterRoleBinding = "kubeadm:cluster-admins"

	// KubernetesAPICallTimeout specifies how long kubeadm should wait for API calls
	KubernetesAPICallTimeout = 1 * time.Minute
	// KubernetesAPICallRetryInterval defines how long kubeadm should wait before retrying a failed API operation
	KubernetesAPICallRetryInterval = 500 * time.Millisecond

	// DiscoveryTimeout specifies the default discovery timeout for kubeadm (used unless one is specified in the JoinConfiguration)
	DiscoveryTimeout = 5 * time.Minute
	// DiscoveryRetryInterval specifies how long kubeadm should wait before retrying to connect to the control-plane when doing discovery
	DiscoveryRetryInterval = 5 * time.Second

	// TLSBootstrapTimeout specifies how long kubeadm should wait for the kubelet to perform the TLS Bootstrap
	TLSBootstrapTimeout = 5 * time.Minute
	// TLSBootstrapRetryInterval specifies how long kubeadm should wait before retrying the TLS Bootstrap check
	TLSBootstrapRetryInterval = 1 * time.Second

	// EtcdAPICallTimeout specifies how much time to wait for completion of requests against the etcd API.
	EtcdAPICallTimeout = 2 * time.Minute
	// EtcdAPICallRetryInterval specifies how frequently to retry requests against the etcd API.
	EtcdAPICallRetryInterval = 500 * time.Millisecond

	// ControlPlaneComponentHealthCheckTimeout specifies the default control plane component health check timeout
	ControlPlaneComponentHealthCheckTimeout = 4 * time.Minute

	// KubeletHealthCheckTimeout specifies the default kubelet timeout
	KubeletHealthCheckTimeout = 4 * time.Minute

	// UpgradeManifestsTimeout specifies the default timeout for upgrading static Pod manifests
	UpgradeManifestsTimeout = 5 * time.Minute

	// PullImageRetry specifies how many times ContainerRuntime retries when pulling image failed
	PullImageRetry = 5
	// RemoveContainerRetry specifies how many times ContainerRuntime retries when removing container failed
	RemoveContainerRetry = 5

	// MinimumAddressesInServiceSubnet defines minimum amount of nodes the Service subnet should allow.
	// We need at least ten, because the DNS service is always at the tenth cluster clusterIP
	MinimumAddressesInServiceSubnet = 10

	// MaximumBitsForServiceSubnet defines maximum possible size of the service subnet in terms of bits.
	// For example, if the value is 20, then the largest supported service subnet is /12 for IPv4 and /108 for IPv6.
	// Note however that anything in between /108 and /112 will be clamped to /112 due to the limitations of the underlying allocation logic.
	// TODO: https://github.com/kubernetes/enhancements/pull/1881
	MaximumBitsForServiceSubnet = 20

	// MinimumAddressesInPodSubnet defines minimum amount of pods in the cluster.
	// We need at least more than services, an IPv4 /28 or IPv6 /128 subnet means 14 util addresses
	MinimumAddressesInPodSubnet = 14

	// PodSubnetNodeMaskMaxDiff is limited to 16 due to an issue with uncompressed IP bitmap in core:
	// xref: #44918
	// The node subnet mask size must be no more than the pod subnet mask size + 16
	PodSubnetNodeMaskMaxDiff = 16

	// DefaultCertTokenDuration specifies the default amount of time that the token used by upload certs will be valid
	// Default behaviour is 2 hours
	DefaultCertTokenDuration = 2 * time.Hour

	// CertificateKeySize specifies the size of the key used to encrypt certificates on uploadcerts phase
	CertificateKeySize = 32

	// LabelNodeRoleControlPlane specifies that a node hosts control-plane components
	LabelNodeRoleControlPlane = "node-role.kubernetes.io/control-plane"

	// LabelExcludeFromExternalLB can be set on a node to exclude it from external load balancers.
	// This is added to control plane nodes to preserve backwards compatibility with a legacy behavior.
	LabelExcludeFromExternalLB = "node.kubernetes.io/exclude-from-external-load-balancers"

	// AnnotationKubeadmCRISocket specifies the annotation kubeadm uses to preserve the crisocket information given to kubeadm at
	// init/join time for use later. kubeadm annotates the node object with this information
	AnnotationKubeadmCRISocket = "kubeadm.alpha.kubernetes.io/cri-socket"

	// KubeadmConfigConfigMap specifies in what ConfigMap in the kube-system namespace the `kubeadm init` configuration should be stored
	KubeadmConfigConfigMap = "kubeadm-config"

	// ClusterConfigurationConfigMapKey specifies in what ConfigMap key the cluster configuration should be stored
	ClusterConfigurationConfigMapKey = "ClusterConfiguration"

	// KubeProxyConfigMap specifies in what ConfigMap in the kube-system namespace the kube-proxy configuration should be stored
	KubeProxyConfigMap = "kube-proxy"

	// KubeProxyConfigMapKey specifies in what ConfigMap key the component config of kube-proxy should be stored
	KubeProxyConfigMapKey = "config.conf"

	// KubeletBaseConfigurationConfigMap specifies in what ConfigMap in the kube-system namespace the initial remote configuration of kubelet should be stored
	KubeletBaseConfigurationConfigMap = "kubelet-config"

	// KubeletBaseConfigurationConfigMapKey specifies in what ConfigMap key the initial remote configuration of kubelet should be stored
	KubeletBaseConfigurationConfigMapKey = "kubelet"

	// KubeletRunDirectory specifies the directory where the kubelet runtime information is stored.
	KubeletRunDirectory = "/var/lib/kubelet"

	// KubeletConfigurationFileName specifies the file name on the node which stores initial remote configuration of kubelet
	// This file should exist under KubeletRunDirectory
	KubeletConfigurationFileName = "config.yaml"

	// KubeletInstanceConfigurationFileName is the name of the kubelet instance configuration file written
	// to all nodes. This file should exist under KubeletRunDirectory.
	KubeletInstanceConfigurationFileName = "instance-config.yaml"

	// KubeletEnvFileName is a file "kubeadm init" writes at runtime. Using that interface, kubeadm can customize certain
	// kubelet flags conditionally based on the environment at runtime. Also, parameters given to the configuration file
	// might be passed through this file. "kubeadm init" writes one variable, with the name ${KubeletEnvFileVariableName}.
	// This file should exist under KubeletRunDirectory
	KubeletEnvFileName = "kubeadm-flags.env"

	// KubeletEnvFileVariableName specifies the shell script variable name "kubeadm init" should write a value to in KubeletEnvFile
	KubeletEnvFileVariableName = "KUBELET_KUBEADM_ARGS"

	// KubeletHealthzPort is the port of the kubelet healthz endpoint
	KubeletHealthzPort = 10248

	// MinExternalEtcdVersion indicates minimum external etcd version which kubeadm supports
	MinExternalEtcdVersion = "3.5.21-0"

	// DefaultEtcdVersion indicates the default etcd version that kubeadm uses
	DefaultEtcdVersion = "3.6.4-0"

	// Etcd defines variable used internally when referring to etcd component
	Etcd = "etcd"
	// KubeAPIServer defines variable used internally when referring to kube-apiserver component
	KubeAPIServer = "kube-apiserver"
	// KubeControllerManager defines variable used internally when referring to kube-controller-manager component
	KubeControllerManager = "kube-controller-manager"
	// KubeScheduler defines variable used internally when referring to kube-scheduler component
	KubeScheduler = "kube-scheduler"
	// KubeProxy defines variable used internally when referring to kube-proxy component
	KubeProxy = "kube-proxy"
	// CoreDNS defines variable used internally when referring to the CoreDNS component
	CoreDNS = "CoreDNS"
	// Kubelet defines variable used internally when referring to the Kubelet
	Kubelet = "kubelet"
	// Kubeadm defines variable used internally when referring to the kubeadm component
	Kubeadm = "kubeadm"

	// KubeCertificatesVolumeName specifies the name for the Volume that is used for injecting certificates to control plane components (can be both a hostPath volume or a projected, all-in-one volume)
	KubeCertificatesVolumeName = "k8s-certs"

	// KubeConfigVolumeName specifies the name for the Volume that is used for injecting the kubeconfig to talk securely to the api server for a control plane component if applicable
	KubeConfigVolumeName = "kubeconfig"

	// DefaultCIImageRepository points to image registry where CI uploads images from ci build job
	DefaultCIImageRepository = "gcr.io/k8s-staging-ci-images"

	// CoreDNSConfigMap specifies in what ConfigMap in the kube-system namespace the CoreDNS config should be stored
	CoreDNSConfigMap = "coredns"

	// CoreDNSDeploymentName specifies the name of the Deployment for CoreDNS add-on
	CoreDNSDeploymentName = "coredns"

	// CoreDNSImageName specifies the name of the image for CoreDNS add-on
	CoreDNSImageName = "coredns"

	// CoreDNSVersion is the version of CoreDNS to be deployed if it is used
	CoreDNSVersion = "v1.12.1"

	// ClusterConfigurationKind is the string kind value for the ClusterConfiguration struct
	ClusterConfigurationKind = "ClusterConfiguration"

	// InitConfigurationKind is the string kind value for the InitConfiguration struct
	InitConfigurationKind = "InitConfiguration"

	// JoinConfigurationKind is the string kind value for the JoinConfiguration struct
	JoinConfigurationKind = "JoinConfiguration"

	// ResetConfigurationKind is the string kind value for the ResetConfiguration struct
	ResetConfigurationKind = "ResetConfiguration"

	// YAMLDocumentSeparator is the separator for YAML documents
	// TODO: Find a better place for this constant
	YAMLDocumentSeparator = "---\n"

	// CIKubernetesVersionPrefix is the prefix for CI Kubernetes version
	CIKubernetesVersionPrefix = "ci/"

	// DefaultAPIServerBindAddress is the default bind address for the API Server
	DefaultAPIServerBindAddress = "0.0.0.0"

	// ControlPlaneNumCPU is the number of CPUs required on control-plane
	ControlPlaneNumCPU = 2

	// ControlPlaneMem is the number of megabytes of memory required on the control-plane
	// Below that amount of RAM running a stable control plane would be difficult.
	ControlPlaneMem = 1700

	// KubeadmCertsSecret specifies in what Secret in the kube-system namespace the certificates should be stored
	KubeadmCertsSecret = "kubeadm-certs"

	// KubeletPort is the default port for the kubelet server on each host machine.
	// May be overridden by a flag at startup.
	KubeletPort = 10250
	// KubeSchedulerPort is the default port for the scheduler status server.
	// May be overridden by a flag at startup.
	KubeSchedulerPort = 10259
	// KubeControllerManagerPort is the default port for the controller manager status server.
	// May be overridden by a flag at startup.
	KubeControllerManagerPort = 10257
	// KubeAPIServerPort is the default port for the apiserver.
	// May be overridden by a flag at startup.
	KubeAPIServerPort = 6443

	// EtcdAdvertiseClientUrlsAnnotationKey is the annotation key on every etcd pod, describing the
	// advertise client URLs
	EtcdAdvertiseClientUrlsAnnotationKey = "kubeadm.kubernetes.io/etcd.advertise-client-urls"
	// KubeAPIServerAdvertiseAddressEndpointAnnotationKey is the annotation key on every apiserver pod,
	// describing the API endpoint (advertise address and bind port of the api server)
	KubeAPIServerAdvertiseAddressEndpointAnnotationKey = "kubeadm.kubernetes.io/kube-apiserver.advertise-address.endpoint"
	// ComponentConfigHashAnnotationKey holds the config map annotation key that kubeadm uses to store
	// a SHA256 sum to check for user changes
	ComponentConfigHashAnnotationKey = "kubeadm.kubernetes.io/component-config.hash"

	// ControlPlaneTier is the value used in the tier label to identify control plane components
	ControlPlaneTier = "control-plane"

	// Mode* constants were copied from pkg/kubeapiserver/authorizer/modes
	// to avoid kubeadm dependency on the internal module
	// TODO: share Mode* constants in component config

	// ModeAlwaysAllow is the mode to set all requests as authorized
	ModeAlwaysAllow string = "AlwaysAllow"
	// ModeAlwaysDeny is the mode to set no requests as authorized
	ModeAlwaysDeny string = "AlwaysDeny"
	// ModeABAC is the mode to use Attribute Based Access Control to authorize
	ModeABAC string = "ABAC"
	// ModeWebhook is the mode to make an external webhook call to authorize
	ModeWebhook string = "Webhook"
	// ModeRBAC is the mode to use Role Based Access Control to authorize
	ModeRBAC string = "RBAC"
	// ModeNode is an authorization mode that authorizes API requests made by kubelets.
	ModeNode string = "Node"

	// PauseVersion indicates the default pause image version for kubeadm
	PauseVersion = "3.10.1"

	// CgroupDriverSystemd holds the systemd driver type
	CgroupDriverSystemd = "systemd"

	// KubeControllerManagerUserName is the username of the user that kube-controller-manager runs as.
	KubeControllerManagerUserName string = "kubeadm-kcm"
	// KubeAPIServerUserName is the username of the user that kube-apiserver runs as.
	KubeAPIServerUserName string = "kubeadm-kas"
	// KubeSchedulerUserName is the username of the user that kube-scheduler runs as.
	KubeSchedulerUserName string = "kubeadm-ks"
	// EtcdUserName is the username of the user that etcd runs as.
	EtcdUserName string = "kubeadm-etcd"
	// ServiceAccountKeyReadersGroupName is the group of users that are allowed to read the service account private key.
	ServiceAccountKeyReadersGroupName string = "kubeadm-sa-key-readers"
	// UpgradeConfigurationKind is the string kind value for the UpgradeConfiguration struct
	UpgradeConfigurationKind = "UpgradeConfiguration"

	// EnvVarInitDryRunDir has the environment variable for init dry run directory override.
	EnvVarInitDryRunDir = "KUBEADM_INIT_DRYRUN_DIR"
	// EnvVarJoinDryRunDir has the environment variable for join dry run directory override.
	EnvVarJoinDryRunDir = "KUBEADM_JOIN_DRYRUN_DIR"
	// EnvVarUpgradeDryRunDir has the environment variable for upgrade dry run directory override.
	EnvVarUpgradeDryRunDir = "KUBEADM_UPGRADE_DRYRUN_DIR"

	// ProbePort is a general named port to be used in pod manifests.
	ProbePort = "probe-port"
)

var (
	// ControlPlaneTaint is the taint to apply on the PodSpec for being able to run that Pod on the control-plane
	ControlPlaneTaint = v1.Taint{
		Key:    LabelNodeRoleControlPlane,
		Effect: v1.TaintEffectNoSchedule,
	}

	// ControlPlaneToleration is the toleration to apply on the PodSpec for being able to run that Pod on the control-plane
	ControlPlaneToleration = v1.Toleration{
		Key:    LabelNodeRoleControlPlane,
		Effect: v1.TaintEffectNoSchedule,
	}

	// ControlPlaneComponents defines the control-plane component names
	ControlPlaneComponents = []string{KubeAPIServer, KubeControllerManager, KubeScheduler}

	// MinimumControlPlaneVersion specifies the minimum control plane version kubeadm can deploy
	MinimumControlPlaneVersion = getSkewedKubernetesVersion(-1)

	// MinimumKubeletVersion specifies the minimum version of kubelet which kubeadm supports
	MinimumKubeletVersion = getSkewedKubernetesVersion(-3)

	// CurrentKubernetesVersion specifies current Kubernetes version supported by kubeadm
	CurrentKubernetesVersion = getSkewedKubernetesVersion(0)

	// SupportedEtcdVersion lists officially supported etcd versions with corresponding Kubernetes releases
	SupportedEtcdVersion = map[uint8]string{
		31: "3.5.21-0",
		32: "3.5.21-0",
		33: "3.5.21-0",
		34: "3.6.4-0",
	}

	// KubeadmCertsClusterRoleName sets the name for the ClusterRole that allows
	// the bootstrap tokens to access the kubeadm-certs Secret during the join of a new control-plane
	KubeadmCertsClusterRoleName = fmt.Sprintf("kubeadm:%s", KubeadmCertsSecret)

	// DefaultKubernetesPlaceholderVersion is a placeholder version in case the component-base
	// version was not populated during build.
	DefaultKubernetesPlaceholderVersion = version.MustParseSemantic("v1.0.0-placeholder-version")
)

// getSkewedKubernetesVersion returns the current MAJOR.(MINOR+n).0 Kubernetes version with a skew of 'n'
// It uses the kubeadm version provided by the 'component-base/version' package. This version must be populated
// by passing linker flags during the kubeadm build process. If the version is empty, assume that kubeadm
// was either build incorrectly or this code is running in unit tests.
func getSkewedKubernetesVersion(n int) *version.Version {
	versionInfo := componentversion.Get()
	return getSkewedKubernetesVersionImpl(&versionInfo, n)
}

func getSkewedKubernetesVersionImpl(versionInfo *apimachineryversion.Info, n int) *version.Version {
	// TODO: update if the kubeadm version gets decoupled from the Kubernetes version.
	// This would require keeping track of the supported skew in a table.
	// More changes would be required if the kubelet version one day decouples from that of Kubernetes.
	var ver *version.Version
	if len(versionInfo.Major) == 0 {
		return DefaultKubernetesPlaceholderVersion
	}
	ver = version.MustParseSemantic(versionInfo.GitVersion)
	// Append the MINOR version skew.
	// TODO: handle the case of Kubernetes moving to v2.0 or having MAJOR version updates in the future.
	// This would require keeping track (in a table) of the last MINOR for a particular MAJOR.
	minor := uint(int(ver.Minor()) + n)
	return version.MustParseSemantic(fmt.Sprintf("v%d.%d.0", ver.Major(), minor))
}

// EtcdSupportedVersion returns officially supported version of etcd for a specific Kubernetes release
// If passed version is not in the given list, the function returns the nearest version with a warning
func EtcdSupportedVersion(supportedEtcdVersion map[uint8]string, versionString string) (etcdVersion *version.Version, warning, err error) {
	kubernetesVersion, err := version.ParseSemantic(versionString)
	if err != nil {
		return nil, nil, err
	}
	desiredVersion, etcdStringVersion := uint8(kubernetesVersion.Minor()), ""

	min, max := ^uint8(0), uint8(0)
	for k, v := range supportedEtcdVersion {
		if desiredVersion == k {
			etcdStringVersion = v
			break
		}
		if k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}

	if len(etcdStringVersion) == 0 {
		if desiredVersion < min {
			etcdStringVersion = supportedEtcdVersion[min]
		}
		if desiredVersion > max {
			etcdStringVersion = supportedEtcdVersion[max]
		}
		warning = errors.Errorf("could not find officially supported version of etcd for Kubernetes %s, falling back to the nearest etcd version (%s)",
			versionString, etcdStringVersion)
	}

	etcdVersion, err = version.ParseSemantic(etcdStringVersion)
	if err != nil {
		return nil, nil, err
	}

	return etcdVersion, warning, nil
}

// GetStaticPodDirectory returns the location on the disk where the Static Pod should be present
func GetStaticPodDirectory() string {
	return filepath.Join(KubernetesDir, ManifestsSubDirName)
}

// GetStaticPodFilepath returns the location on the disk where the Static Pod should be present
func GetStaticPodFilepath(componentName, manifestsDir string) string {
	return filepath.Join(manifestsDir, componentName+".yaml")
}

// GetAdminKubeConfigPath returns the location on the disk where admin kubeconfig is located by default
func GetAdminKubeConfigPath() string {
	return filepath.Join(KubernetesDir, AdminKubeConfigFileName)
}

// GetKubeletKubeConfigPath returns the location on the disk where kubelet kubeconfig is located by default
func GetKubeletKubeConfigPath() string {
	return filepath.Join(KubernetesDir, KubeletKubeConfigFileName)
}

// GetDryRunDir creates a temporary directory under /etc/kubernetes/tmp.
// If the environment variable with name stored in envVar is set, it is used instead.
// msgFunc will be used to print a message to the user that they can use envVar for override.
func GetDryRunDir(envVar, dirName string, msgFunc func(format string, args ...interface{})) (string, error) {
	envVarDir := os.Getenv(envVar)
	if len(envVarDir) > 0 {
		return envVarDir, nil
	}
	tempDir := filepath.Join(KubernetesDir, TempDir)
	generatedDir, err := createTmpDir(tempDir, dirName)
	if err != nil {
		return "", err
	}

	msgFunc("Using dry-run directory %s. To override it, set the environment variable %s",
		generatedDir, envVar)

	return generatedDir, nil
}

// CreateTempDir creates a temporary directory under /etc/kubernetes/tmp
// or under the provided parent directory if it's set.
func CreateTempDir(parent, dirName string) (string, error) {
	tempDir := filepath.Join(KubernetesDir, TempDir)
	if len(parent) > 0 {
		tempDir = filepath.Join(parent, TempDir)
	}
	return createTmpDir(tempDir, dirName)
}

func createTmpDir(tempDir, dirName string) (string, error) {
	if err := os.MkdirAll(tempDir, 0700); err != nil {
		return "", errors.Wrapf(err, "failed to create directory %q", tempDir)
	}
	tempDir, err := os.MkdirTemp(tempDir, dirName)
	if err != nil {
		return "", errors.Wrapf(err, "could not create a temporary directory in %q", tempDir)
	}
	return tempDir, nil
}

// CreateTimestampDir is a function that creates a temporary directory under /etc/kubernetes/tmp formatted with the current date
func CreateTimestampDir(kubernetesDir, dirName string) (string, error) {
	tempDir := filepath.Join(KubernetesDir, TempDir)
	if len(kubernetesDir) != 0 {
		tempDir = filepath.Join(kubernetesDir, TempDir)
	}

	// creates target folder if not already exists
	if err := os.MkdirAll(tempDir, 0700); err != nil {
		return "", errors.Wrapf(err, "failed to create directory %q", tempDir)
	}

	timestampDirName := fmt.Sprintf("%s-%s", dirName, time.Now().Format("2006-01-02-15-04-05"))
	timestampDir := filepath.Join(tempDir, timestampDirName)
	if err := os.Mkdir(timestampDir, 0700); err != nil {
		return "", errors.Wrap(err, "could not create timestamp directory")
	}

	return timestampDir, nil
}

// GetDNSIP returns a dnsIP, which is 10th IP in svcSubnet CIDR range
func GetDNSIP(svcSubnetList string) (net.IP, error) {
	// Get the service subnet CIDR
	svcSubnetCIDR, err := GetKubernetesServiceCIDR(svcSubnetList)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get internal Kubernetes Service IP from the given service CIDR (%s)", svcSubnetList)
	}

	// Selects the 10th IP in service subnet CIDR range as dnsIP
	dnsIP, err := netutils.GetIndexedIP(svcSubnetCIDR, 10)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get internal Kubernetes Service IP from the given service CIDR")
	}

	return dnsIP, nil
}

// GetKubernetesServiceCIDR returns the default Service CIDR for the Kubernetes internal service
func GetKubernetesServiceCIDR(svcSubnetList string) (*net.IPNet, error) {
	// The default service address family for the cluster is the address family of the first
	// service cluster IP range configured via the `--service-cluster-ip-range` flag
	// of the kube-controller-manager and kube-apiserver.
	svcSubnets, err := netutils.ParseCIDRs(strings.Split(svcSubnetList, ","))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse ServiceSubnet %v", svcSubnetList)
	}
	if len(svcSubnets) == 0 {
		return nil, errors.New("received empty ServiceSubnet")
	}
	return svcSubnets[0], nil
}

// GetAPIServerVirtualIP returns the IP of the internal Kubernetes API service
func GetAPIServerVirtualIP(svcSubnetList string) (net.IP, error) {
	svcSubnet, err := GetKubernetesServiceCIDR(svcSubnetList)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get internal Kubernetes Service IP from the given service CIDR")
	}
	internalAPIServerVirtualIP, err := netutils.GetIndexedIP(svcSubnet, 1)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get the first IP address from the given CIDR: %s", svcSubnet.String())
	}
	return internalAPIServerVirtualIP, nil
}
