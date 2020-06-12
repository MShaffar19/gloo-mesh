// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"context"

	v1alpha1_2 "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1_2 "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/aws_creds"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/clients"
	aws_utils "github.com/solo-io/service-mesh-hub/pkg/common/aws/parser"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/selection"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/settings"
	mc_wire "github.com/solo-io/service-mesh-hub/pkg/common/compute-target/wire"
	"github.com/solo-io/service-mesh-hub/pkg/common/container-runtime/docker"
	"github.com/solo-io/service-mesh-hub/pkg/common/csr/installation"
	"github.com/solo-io/service-mesh-hub/pkg/common/filesystem/files"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/helm"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/kubeconfig"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/compute-target/aws"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/compute-target/aws/clients/eks"
	event_watcher_factories "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/compute-target/event-watcher-factories"
	appmesh_tenancy "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/cluster-tenancy/k8s/appmesh"
	eks2 "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/k8s-cluster/rest/eks"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh-workload/k8s"
	appmesh2 "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh-workload/k8s/appmesh"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh/k8s/consul"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh/k8s/istio"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh/k8s/linkerd"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh/rest/appmesh"
)

// Injectors from wire.go:

func InitializeDiscovery(ctx context.Context) (DiscoveryContext, error) {
	config, err := mc_wire.LocalKubeConfigProvider()
	if err != nil {
		return DiscoveryContext{}, err
	}
	asyncManager, err := mc_wire.LocalManagerProvider(ctx, config)
	if err != nil {
		return DiscoveryContext{}, err
	}
	fileReader := files.NewDefaultFileReader()
	converter := kubeconfig.NewConverter(fileReader)
	asyncManagerController := mc_wire.KubeClusterCredentialsHandlerProvider(converter)
	secretAwsCredsConverter := aws_creds.DefaultSecretAwsCredsConverter()
	client := mc_wire.DynamicClientProvider(asyncManager)
	meshClientFactory := v1alpha1.MeshClientFactoryProvider()
	arnParser := aws_utils.NewArnParser()
	appmeshRawClientFactory := clients.AppmeshRawClientFactoryProvider()
	settingsClient := v1alpha1_2.SettingsClientProvider(client)
	settingsHelperClient := settings.NewAwsSettingsHelperClient(settingsClient)
	awsSelector := selection.NewAwsSelector(arnParser)
	appMeshDiscoveryReconciler := appmesh.NewAppMeshDiscoveryReconciler(client, meshClientFactory, arnParser, appmeshRawClientFactory, settingsHelperClient, awsSelector)
	kubernetesClusterClient := v1alpha1.KubernetesClusterClientProvider(client)
	eksClientFactory := eks.EksClientFactoryProvider()
	eksConfigBuilderFactory := eks.EksConfigBuilderFactoryProvider()
	secretClientFromConfigFactory := v1.SecretClientFromConfigFactoryProvider()
	kubernetesClusterClientFromConfigFactory := v1alpha1.KubernetesClusterClientFromConfigFactoryProvider()
	namespaceClientFromConfigFactory := v1.NamespaceClientFromConfigFactoryProvider()
	helmClientForFileConfigFactory := helm.HelmClientForFileConfigFactoryProvider()
	helmClientForMemoryConfigFactory := helm.HelmClientForMemoryConfigFactoryProvider()
	deploymentClientFromConfigFactory := v1_2.DeploymentClientFromConfigFactoryProvider()
	imageNameParser := docker.NewImageNameParser()
	deployedVersionFinder, err := DeployedVersionFinderProvider(config, deploymentClientFromConfigFactory, imageNameParser)
	if err != nil {
		return DiscoveryContext{}, err
	}
	csrAgentInstallerFactory := installation.NewCsrAgentInstallerFactory(helmClientForFileConfigFactory, helmClientForMemoryConfigFactory, deployedVersionFinder)
	clusterRegistrationClient, err := ClusterRegistrationClientProvider(config, secretClientFromConfigFactory, kubernetesClusterClientFromConfigFactory, namespaceClientFromConfigFactory, converter, csrAgentInstallerFactory)
	if err != nil {
		return DiscoveryContext{}, err
	}
	eksDiscoveryReconciler := eks2.NewEksDiscoveryReconciler(kubernetesClusterClient, eksClientFactory, eksConfigBuilderFactory, clusterRegistrationClient, settingsHelperClient, awsSelector)
	v := AwsDiscoveryReconcilersProvider(appMeshDiscoveryReconciler, eksDiscoveryReconciler)
	stsClientFactory := clients.STSClientFactoryProvider()
	awsCredsHandler := aws.NewAwsAPIHandler(secretAwsCredsConverter, v, stsClientFactory)
	v2 := ComputeTargetCredentialsHandlersProvider(asyncManagerController, awsCredsHandler)
	asyncManagerStartOptionsFunc := mc_wire.LocalManagerStarterProvider(v2)
	multiClusterDependencies := mc_wire.MulticlusterDependenciesProvider(ctx, asyncManager, asyncManagerController, asyncManagerStartOptionsFunc)
	configMapClientFactory := v1.ConfigMapClientFactoryProvider()
	istioMeshScanner := istio.NewIstioMeshScanner(imageNameParser, configMapClientFactory)
	consulConnectInstallationScanner := consul.NewConsulConnectInstallationScanner(imageNameParser)
	consulConnectMeshScanner := consul.NewConsulMeshScanner(imageNameParser, consulConnectInstallationScanner)
	linkerdMeshScanner := linkerd.NewLinkerdMeshScanner(imageNameParser)
	replicaSetClientFactory := v1_2.ReplicaSetClientFactoryProvider()
	deploymentClientFactory := v1_2.DeploymentClientFactoryProvider()
	ownerFetcherFactory := k8s.OwnerFetcherFactoryProvider()
	serviceClientFactory := v1.ServiceClientFactoryProvider()
	meshServiceClientFactory := v1alpha1.MeshServiceClientFactoryProvider()
	meshWorkloadClientFactory := v1alpha1.MeshWorkloadClientFactoryProvider()
	podEventWatcherFactory := event_watcher_factories.NewPodEventWatcherFactory()
	serviceEventWatcherFactory := event_watcher_factories.NewServiceEventWatcherFactory()
	meshWorkloadEventWatcherFactory := event_watcher_factories.NewMeshWorkloadEventWatcherFactory()
	deploymentEventWatcherFactory := event_watcher_factories.NewDeploymentEventWatcherFactory()
	podClientFactory := v1.PodClientFactoryProvider()
	meshEventWatcherFactory := event_watcher_factories.NewMeshEventWatcherFactory()
	appMeshScanner := aws_utils.NewAppMeshScanner(arnParser)
	awsAccountIdFetcher := aws_utils.NewAwsAccountIdFetcher(arnParser, configMapClientFactory)
	meshWorkloadScannerFactory := appmesh2.AppMeshWorkloadScannerFactoryProvider(appMeshScanner, awsAccountIdFetcher)
	clusterTenancyScannerFactory := appmesh_tenancy.AppMeshTenancyScannerFactoryProvider(appMeshScanner, awsAccountIdFetcher)
	discoveryContext := DiscoveryContextProvider(multiClusterDependencies, istioMeshScanner, consulConnectMeshScanner, linkerdMeshScanner, replicaSetClientFactory, deploymentClientFactory, ownerFetcherFactory, serviceClientFactory, meshServiceClientFactory, meshWorkloadClientFactory, podEventWatcherFactory, serviceEventWatcherFactory, meshWorkloadEventWatcherFactory, deploymentEventWatcherFactory, meshClientFactory, podClientFactory, meshEventWatcherFactory, meshWorkloadScannerFactory, clusterTenancyScannerFactory)
	return discoveryContext, nil
}
