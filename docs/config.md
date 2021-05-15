---
title: API Documentation
linktitle: API Documentation
description: Reference of the jx-promote configuration
weight: 10
---
<p>Packages:</p>
<ul>
<li>
<a href="#core.jenkins-x.io%2fv4beta1">core.jenkins-x.io/v4beta1</a>
</li>
</ul>
<h2 id="core.jenkins-x.io/v4beta1">core.jenkins-x.io/v4beta1</h2>
<p>
<p>Package v4 is the v4 version of the API.</p>
</p>
Resource Types:
<ul></ul>
<h3 id="core.jenkins-x.io/v4beta1.AutoUpdateConfig">AutoUpdateConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>AutoUpdateConfig contains auto update config</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Enabled autoupdate</p>
</td>
</tr>
<tr>
<td>
<code>schedule</code></br>
<em>
string
</em>
</td>
<td>
<p>Schedule cron of auto updates</p>
</td>
</tr>
<tr>
<td>
<code>autoMerge</code></br>
<em>
bool
</em>
</td>
<td>
<p>AutoMerge if enabled lets auto merge any generated update PullRequests on the dev cluster git repository</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.AzureConfig">AzureConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.ClusterConfig">ClusterConfig</a>)
</p>
<p>
<p>AzureConfig contains Azure specific requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>registrySubscription</code></br>
<em>
string
</em>
</td>
<td>
<p>RegistrySubscription the registry subscription for defaulting the container registry.
Not used if you specify a Registry explicitly</p>
</td>
</tr>
<tr>
<td>
<code>dns</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.AzureDNSConfig">
AzureDNSConfig
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>secretStorage</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.AzureSecretConfig">
AzureSecretConfig
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storage</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.AzureStorageConfig">
AzureStorageConfig
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.AzureDNSConfig">AzureDNSConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.AzureConfig">AzureConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>tenantId</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>subscriptionId</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resourceGroup</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.AzureSecretConfig">AzureSecretConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.AzureConfig">AzureConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyVaultName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.AzureStorageConfig">AzureStorageConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.AzureConfig">AzureConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>storageAccountName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.ChartRepositoryType">ChartRepositoryType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.DestinationConfig">DestinationConfig</a>)
</p>
<p>
<p>ChartRepositoryType is the type of chart repository used for helm</p>
</p>
<h3 id="core.jenkins-x.io/v4beta1.ClusterConfig">ClusterConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>ClusterConfig contains cluster specific requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>DestinationConfig</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.DestinationConfig">
DestinationConfig
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azure</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.AzureConfig">
AzureConfig
</a>
</em>
</td>
<td>
<p>AzureConfig the azure specific configuration</p>
</td>
</tr>
<tr>
<td>
<code>gke</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.GKEConfig">
GKEConfig
</a>
</em>
</td>
<td>
<p>GKEConfig the gke specific configuration</p>
</td>
</tr>
<tr>
<td>
<code>environmentGitPublic</code></br>
<em>
bool
</em>
</td>
<td>
<p>EnvironmentGitPublic determines whether jx boot create public or private git repos for the environments</p>
</td>
</tr>
<tr>
<td>
<code>gitPublic</code></br>
<em>
bool
</em>
</td>
<td>
<p>GitPublic determines whether jx boot create public or private git repos for the applications</p>
</td>
</tr>
<tr>
<td>
<code>provider</code></br>
<em>
string
</em>
</td>
<td>
<p>Provider the kubernetes provider (e.g. gke)</p>
</td>
</tr>
<tr>
<td>
<code>project</code></br>
<em>
string
</em>
</td>
<td>
<p>ProjectID the cloud project ID e.g. on GCP</p>
</td>
</tr>
<tr>
<td>
<code>clusterName</code></br>
<em>
string
</em>
</td>
<td>
<p>ClusterName the logical name of the cluster</p>
</td>
</tr>
<tr>
<td>
<code>region</code></br>
<em>
string
</em>
</td>
<td>
<p>Region the cloud region being used</p>
</td>
</tr>
<tr>
<td>
<code>zone</code></br>
<em>
string
</em>
</td>
<td>
<p>Zone the cloud zone being used</p>
</td>
</tr>
<tr>
<td>
<code>gitName</code></br>
<em>
string
</em>
</td>
<td>
<p>GitName is the name of the default git service</p>
</td>
</tr>
<tr>
<td>
<code>gitKind</code></br>
<em>
string
</em>
</td>
<td>
<p>GitKind is the kind of git server (github, bitbucketserver etc)</p>
</td>
</tr>
<tr>
<td>
<code>gitServer</code></br>
<em>
string
</em>
</td>
<td>
<p>GitServer is the URL of the git server</p>
</td>
</tr>
<tr>
<td>
<code>externalDNSSAName</code></br>
<em>
string
</em>
</td>
<td>
<p>ExternalDNSSAName the service account name for external dns</p>
</td>
</tr>
<tr>
<td>
<code>kanikoSAName</code></br>
<em>
string
</em>
</td>
<td>
<p>VaultSAName the service account name for vault
KanikoSAName the service account name for kaniko</p>
</td>
</tr>
<tr>
<td>
<code>devEnvApprovers</code></br>
<em>
[]string
</em>
</td>
<td>
<p>DevEnvApprovers contains an optional list of approvers to populate the initial OWNERS file in the dev env repo</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.DestinationConfig">DestinationConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.ClusterConfig">ClusterConfig</a>, 
<a href="#core.jenkins-x.io/v4beta1.SettingsConfig">SettingsConfig</a>)
</p>
<p>
<p>DestinationConfig the common cluster settings that can be specified in settings or requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>chartRepository</code></br>
<em>
string
</em>
</td>
<td>
<p>ChartRepository the repository URL to deploy charts to</p>
</td>
</tr>
<tr>
<td>
<code>chartKind</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.ChartRepositoryType">
ChartRepositoryType
</a>
</em>
</td>
<td>
<p>ChartKind the chart repository kind (e.g. normal, OCI or github pages)</p>
</td>
</tr>
<tr>
<td>
<code>chartSecret</code></br>
<em>
string
</em>
</td>
<td>
<p>ChartSecret an optional secret name used to be able to push to chart repositories</p>
</td>
</tr>
<tr>
<td>
<code>registry</code></br>
<em>
string
</em>
</td>
<td>
<p>Registry the host name of the container registry</p>
</td>
</tr>
<tr>
<td>
<code>dockerRegistryOrg</code></br>
<em>
string
</em>
</td>
<td>
<p>DockerRegistryOrg the default organisation used for container images</p>
</td>
</tr>
<tr>
<td>
<code>kanikoFlags</code></br>
<em>
string
</em>
</td>
<td>
<p>KanikoFlags allows global kaniko flags to be supplied such as to disable host verification</p>
</td>
</tr>
<tr>
<td>
<code>environmentGitOwner</code></br>
<em>
string
</em>
</td>
<td>
<p>EnvironmentGitOwner the default git owner for environment repositories if none is specified explicitly</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.EnvironmentConfig">EnvironmentConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>, 
<a href="#core.jenkins-x.io/v4beta1.SettingsConfig">SettingsConfig</a>)
</p>
<p>
<p>EnvironmentConfig configures the organisation and repository name of the git repositories for environments</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>key</code></br>
<em>
string
</em>
</td>
<td>
<p>Key is the key of the environment configuration</p>
</td>
</tr>
<tr>
<td>
<code>owner</code></br>
<em>
string
</em>
</td>
<td>
<p>Owner is the git user or organisation for the repository</p>
</td>
</tr>
<tr>
<td>
<code>repository</code></br>
<em>
string
</em>
</td>
<td>
<p>Repository is the name of the repository within the owner</p>
</td>
</tr>
<tr>
<td>
<code>gitServer</code></br>
<em>
string
</em>
</td>
<td>
<p>GitServer is the URL of the git server</p>
</td>
</tr>
<tr>
<td>
<code>gitKind</code></br>
<em>
string
</em>
</td>
<td>
<p>GitKind is the kind of git server (github, bitbucketserver etc)</p>
</td>
</tr>
<tr>
<td>
<code>gitUrl</code></br>
<em>
string
</em>
</td>
<td>
<p>GitURL optional git URL for the git repository for the environment. If its not specified its generated from the
git server, kind, owner and repository</p>
</td>
</tr>
<tr>
<td>
<code>ingress</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">
IngressConfig
</a>
</em>
</td>
<td>
<p>Ingress contains ingress specific requirements</p>
</td>
</tr>
<tr>
<td>
<code>remoteCluster</code></br>
<em>
bool
</em>
</td>
<td>
<p>RemoteCluster specifies this environment runs on a remote cluster to the development cluster</p>
</td>
</tr>
<tr>
<td>
<code>promotionStrategy</code></br>
<em>
github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1.PromotionStrategyType
</em>
</td>
<td>
<p>PromotionStrategy what kind of promotion strategy to use</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code></br>
<em>
string
</em>
</td>
<td>
<p>Namespace is the target namespace for deploying resources in this environment.  Will default to &ldquo;jx-{{ .Key }}&rdquo; if omitted</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.GKEConfig">GKEConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.ClusterConfig">ClusterConfig</a>)
</p>
<p>
<p>GKEConfig contains GKE specific requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>projectNumber</code></br>
<em>
string
</em>
</td>
<td>
<p>ProjectNumber the unique project number GKE assigns to a project (required for workload identity).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.IngressConfig">IngressConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.EnvironmentConfig">EnvironmentConfig</a>, 
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>IngressConfig contains dns specific requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>externalDNS</code></br>
<em>
bool
</em>
</td>
<td>
<p>DNS is enabled</p>
</td>
</tr>
<tr>
<td>
<code>cloud_dns_secret_name</code></br>
<em>
string
</em>
</td>
<td>
<p>CloudDNSSecretName secret name which contains the service account for external-dns and cert-manager issuer to
access the Cloud DNS service to resolve a DNS challenge</p>
</td>
</tr>
<tr>
<td>
<code>domain</code></br>
<em>
string
</em>
</td>
<td>
<p>Domain to expose ingress endpoints</p>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.IngressType">
IngressType
</a>
</em>
</td>
<td>
<p>Kind the kind of ingress used (ingress v1, ingress v2, istio etc)</p>
</td>
</tr>
<tr>
<td>
<code>ignoreLoadBalancer</code></br>
<em>
bool
</em>
</td>
<td>
<p>IgnoreLoadBalancer if the nginx-controller LoadBalancer service should not be used to detect and update the
domain if you are using a dynamic domain resolver like <code>.nip.io</code> rather than a real DNS configuration.
With this flag enabled the <code>Domain</code> value will be used and never re-created based on the current LoadBalancer IP address.</p>
</td>
</tr>
<tr>
<td>
<code>namespaceSubDomain</code></br>
<em>
string
</em>
</td>
<td>
<p>NamespaceSubDomain the sub domain expression to expose ingress. Defaults to &ldquo;.jx.&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>tls</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.TLSConfig">
TLSConfig
</a>
</em>
</td>
<td>
<p>TLS enable automated TLS using certmanager</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code></br>
<em>
map[string]string
</em>
</td>
<td>
<p>Annotations optional annotations added to ingresses</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.IngressType">IngressType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">IngressConfig</a>)
</p>
<p>
<p>IngressType is the type of a ingress strategy</p>
</p>
<h3 id="core.jenkins-x.io/v4beta1.LegacyStorageConfig">LegacyStorageConfig
</h3>
<p>
<p>Deprecated: migrate to top level Requirements object</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>logs</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.LegacyStorageEntryConfig">
LegacyStorageEntryConfig
</a>
</em>
</td>
<td>
<p>Logs for storing build logs</p>
</td>
</tr>
<tr>
<td>
<code>reports</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.LegacyStorageEntryConfig">
LegacyStorageEntryConfig
</a>
</em>
</td>
<td>
<p>Tests for storing test results, coverage + code quality reports</p>
</td>
</tr>
<tr>
<td>
<code>repository</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.LegacyStorageEntryConfig">
LegacyStorageEntryConfig
</a>
</em>
</td>
<td>
<p>Repository for storing repository artifacts</p>
</td>
</tr>
<tr>
<td>
<code>backup</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.LegacyStorageEntryConfig">
LegacyStorageEntryConfig
</a>
</em>
</td>
<td>
<p>Backup for backing up kubernetes resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.LegacyStorageEntryConfig">LegacyStorageEntryConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.LegacyStorageConfig">LegacyStorageConfig</a>)
</p>
<p>
<p>Deprecated: migrate to top level Requirements object</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Enabled if the storage is enabled</p>
</td>
</tr>
<tr>
<td>
<code>url</code></br>
<em>
string
</em>
</td>
<td>
<p>URL the cloud storage bucket URL such as &lsquo;gs://mybucket&rsquo; or &lsquo;s3://foo&rsquo; or `azblob://thingy&rsquo;
see <a href="https://jenkins-x.io/architecture/storage/">https://jenkins-x.io/architecture/storage/</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.RepositoryType">RepositoryType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>RepositoryType is the type of a repository we use to store artifacts (jars, tarballs, npm packages etc)</p>
</p>
<h3 id="core.jenkins-x.io/v4beta1.Requirements">Requirements
</h3>
<p>
<p>Requirements represents a collection installation requirements for Jenkins X</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">
RequirementsConfig
</a>
</em>
</td>
<td>
<p>Spec the definition of the secret mappings</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoUpdate</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.AutoUpdateConfig">
AutoUpdateConfig
</a>
</em>
</td>
<td>
<p>AutoUpdate contains auto update config</p>
</td>
</tr>
<tr>
<td>
<code>cluster</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.ClusterConfig">
ClusterConfig
</a>
</em>
</td>
<td>
<p>Cluster contains cluster specific requirements</p>
</td>
</tr>
<tr>
<td>
<code>environments</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.EnvironmentConfig">
[]EnvironmentConfig
</a>
</em>
</td>
<td>
<p>Environments the requirements for the environments</p>
</td>
</tr>
<tr>
<td>
<code>extraDomains</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">
[]IngressConfig
</a>
</em>
</td>
<td>
<p>ExtraDomains to expose alternate services with custom ingress for specific applications</p>
</td>
</tr>
<tr>
<td>
<code>ingress</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">
IngressConfig
</a>
</em>
</td>
<td>
<p>Ingress contains ingress specific requirements</p>
</td>
</tr>
<tr>
<td>
<code>kuberhealthy</code></br>
<em>
bool
</em>
</td>
<td>
<p>Kuberhealthy indicates if we have already installed Kuberhealthy upfront in the kubernetes cluster</p>
</td>
</tr>
<tr>
<td>
<code>pipelineUser</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.UserNameEmailConfig">
UserNameEmailConfig
</a>
</em>
</td>
<td>
<p>PipelineUser the user name and email used for running pipelines</p>
</td>
</tr>
<tr>
<td>
<code>repository</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.RepositoryType">
RepositoryType
</a>
</em>
</td>
<td>
<p>Repository specifies what kind of artifact repository you wish to use for storing artifacts (jars, tarballs, npm modules etc)</p>
</td>
</tr>
<tr>
<td>
<code>secretStorage</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.SecretStorageType">
SecretStorageType
</a>
</em>
</td>
<td>
<p>SecretStorage how should we store secrets for the cluster</p>
</td>
</tr>
<tr>
<td>
<code>storage</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.StorageConfig">
[]StorageConfig
</a>
</em>
</td>
<td>
<p>Storage contains storage requirements</p>
</td>
</tr>
<tr>
<td>
<code>terraform</code></br>
<em>
bool
</em>
</td>
<td>
<p>Terraform specifies if  we are managing the kubernetes cluster and cloud resources with Terraform</p>
</td>
</tr>
<tr>
<td>
<code>terraformVault</code></br>
<em>
bool
</em>
</td>
<td>
<p>TerraformVault indicates whether Vault has been installed upfront by Terraform</p>
</td>
</tr>
<tr>
<td>
<code>vault</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.VaultConfig">
VaultConfig
</a>
</em>
</td>
<td>
<p>Vault the configuration for vault</p>
</td>
</tr>
<tr>
<td>
<code>webhook</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.WebhookType">
WebhookType
</a>
</em>
</td>
<td>
<p>Webhook specifies what engine we should use for webhooks</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.Requirements">Requirements</a>, 
<a href="#core.jenkins-x.io/v4beta1.RequirementsValues">RequirementsValues</a>)
</p>
<p>
<p>RequirementsConfig contains the logical installation requirements in the <code>jx-requirements.yml</code> file when
installing, configuring or upgrading Jenkins X via <code>jx boot</code></p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoUpdate</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.AutoUpdateConfig">
AutoUpdateConfig
</a>
</em>
</td>
<td>
<p>AutoUpdate contains auto update config</p>
</td>
</tr>
<tr>
<td>
<code>cluster</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.ClusterConfig">
ClusterConfig
</a>
</em>
</td>
<td>
<p>Cluster contains cluster specific requirements</p>
</td>
</tr>
<tr>
<td>
<code>environments</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.EnvironmentConfig">
[]EnvironmentConfig
</a>
</em>
</td>
<td>
<p>Environments the requirements for the environments</p>
</td>
</tr>
<tr>
<td>
<code>extraDomains</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">
[]IngressConfig
</a>
</em>
</td>
<td>
<p>ExtraDomains to expose alternate services with custom ingress for specific applications</p>
</td>
</tr>
<tr>
<td>
<code>ingress</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">
IngressConfig
</a>
</em>
</td>
<td>
<p>Ingress contains ingress specific requirements</p>
</td>
</tr>
<tr>
<td>
<code>kuberhealthy</code></br>
<em>
bool
</em>
</td>
<td>
<p>Kuberhealthy indicates if we have already installed Kuberhealthy upfront in the kubernetes cluster</p>
</td>
</tr>
<tr>
<td>
<code>pipelineUser</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.UserNameEmailConfig">
UserNameEmailConfig
</a>
</em>
</td>
<td>
<p>PipelineUser the user name and email used for running pipelines</p>
</td>
</tr>
<tr>
<td>
<code>repository</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.RepositoryType">
RepositoryType
</a>
</em>
</td>
<td>
<p>Repository specifies what kind of artifact repository you wish to use for storing artifacts (jars, tarballs, npm modules etc)</p>
</td>
</tr>
<tr>
<td>
<code>secretStorage</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.SecretStorageType">
SecretStorageType
</a>
</em>
</td>
<td>
<p>SecretStorage how should we store secrets for the cluster</p>
</td>
</tr>
<tr>
<td>
<code>storage</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.StorageConfig">
[]StorageConfig
</a>
</em>
</td>
<td>
<p>Storage contains storage requirements</p>
</td>
</tr>
<tr>
<td>
<code>terraform</code></br>
<em>
bool
</em>
</td>
<td>
<p>Terraform specifies if  we are managing the kubernetes cluster and cloud resources with Terraform</p>
</td>
</tr>
<tr>
<td>
<code>terraformVault</code></br>
<em>
bool
</em>
</td>
<td>
<p>TerraformVault indicates whether Vault has been installed upfront by Terraform</p>
</td>
</tr>
<tr>
<td>
<code>vault</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.VaultConfig">
VaultConfig
</a>
</em>
</td>
<td>
<p>Vault the configuration for vault</p>
</td>
</tr>
<tr>
<td>
<code>webhook</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.WebhookType">
WebhookType
</a>
</em>
</td>
<td>
<p>Webhook specifies what engine we should use for webhooks</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.RequirementsValues">RequirementsValues
</h3>
<p>
<p>RequirementsValues contains the logical installation requirements in the <code>jx-requirements.yml</code> file as helm values</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>jxRequirements</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">
RequirementsConfig
</a>
</em>
</td>
<td>
<p>RequirementsConfig contains the logical installation requirements</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.ResourceReference">ResourceReference
</h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
<em>
string
</em>
</td>
<td>
<p>API version of the referent.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
<em>
string
</em>
</td>
<td>
<p>Kind of the referent.
More info: <a href="https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds</a></p>
</td>
</tr>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name of the referent.
More info: <a href="http://kubernetes.io/docs/user-guide/identifiers#names">http://kubernetes.io/docs/user-guide/identifiers#names</a></p>
</td>
</tr>
<tr>
<td>
<code>uid</code></br>
<em>
k8s.io/apimachinery/pkg/types.UID
</em>
</td>
<td>
<p>UID of the referent.
More info: <a href="http://kubernetes.io/docs/user-guide/identifiers#uids">http://kubernetes.io/docs/user-guide/identifiers#uids</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.SecretStorageType">SecretStorageType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>SecretStorageType is the type of storage used for secrets</p>
</p>
<h3 id="core.jenkins-x.io/v4beta1.Settings">Settings
</h3>
<p>
<p>Settings represents application specific settings for use inside a pipeline of an application</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.SettingsConfig">
SettingsConfig
</a>
</em>
</td>
<td>
<p>Spec the definition of the settings</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>gitUrl</code></br>
<em>
string
</em>
</td>
<td>
<p>GitURL the git URL for your development cluster where the default environments and cluster configuration are specified</p>
</td>
</tr>
<tr>
<td>
<code>destination</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.DestinationConfig">
DestinationConfig
</a>
</em>
</td>
<td>
<p>Destination settings to define where release artifacts go in terms of containers and charts</p>
</td>
</tr>
<tr>
<td>
<code>promoteEnvironments</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.EnvironmentConfig">
[]EnvironmentConfig
</a>
</em>
</td>
<td>
<p>PromoteEnvironments the environments for promotion</p>
</td>
</tr>
<tr>
<td>
<code>ignoreDevEnvironments</code></br>
<em>
bool
</em>
</td>
<td>
<p>IgnoreDevEnvironments if enabled do not inherit any environments from the</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.SettingsConfig">SettingsConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.Settings">Settings</a>)
</p>
<p>
<p>SettingsConfig contains the optional overrides you can specify on a per application basis</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>gitUrl</code></br>
<em>
string
</em>
</td>
<td>
<p>GitURL the git URL for your development cluster where the default environments and cluster configuration are specified</p>
</td>
</tr>
<tr>
<td>
<code>destination</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.DestinationConfig">
DestinationConfig
</a>
</em>
</td>
<td>
<p>Destination settings to define where release artifacts go in terms of containers and charts</p>
</td>
</tr>
<tr>
<td>
<code>promoteEnvironments</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.EnvironmentConfig">
[]EnvironmentConfig
</a>
</em>
</td>
<td>
<p>PromoteEnvironments the environments for promotion</p>
</td>
</tr>
<tr>
<td>
<code>ignoreDevEnvironments</code></br>
<em>
bool
</em>
</td>
<td>
<p>IgnoreDevEnvironments if enabled do not inherit any environments from the</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.StorageConfig">StorageConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>StorageConfig contains dns specific requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name of the bucket</p>
</td>
</tr>
<tr>
<td>
<code>url</code></br>
<em>
string
</em>
</td>
<td>
<p>URL the cloud storage bucket URL such as &lsquo;gs://mybucket&rsquo; or &lsquo;s3://foo&rsquo; or `azblob://thingy&rsquo;
see <a href="https://jenkins-x.io/architecture/storage/">https://jenkins-x.io/architecture/storage/</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.TLSConfig">TLSConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.IngressConfig">IngressConfig</a>)
</p>
<p>
<p>TLSConfig contains TLS specific requirements</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>TLS enabled</p>
</td>
</tr>
<tr>
<td>
<code>email</code></br>
<em>
string
</em>
</td>
<td>
<p>Email address to register with services like LetsEncrypt</p>
</td>
</tr>
<tr>
<td>
<code>production</code></br>
<em>
bool
</em>
</td>
<td>
<p>Production false uses self-signed certificates from the LetsEncrypt staging server, true enables the production
server which incurs higher rate limiting <a href="https://letsencrypt.org/docs/rate-limits/">https://letsencrypt.org/docs/rate-limits/</a></p>
</td>
</tr>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
<p>SecretName the name of the secret which contains the TLS certificate</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.UserNameEmailConfig">UserNameEmailConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>UserNameEmailConfig contains the user name and email of a user (e.g. pipeline user)</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>username</code></br>
<em>
string
</em>
</td>
<td>
<p>Username the username of the user</p>
</td>
</tr>
<tr>
<td>
<code>email</code></br>
<em>
string
</em>
</td>
<td>
<p>Email the email address of the user</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.VaultAWSConfig">VaultAWSConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.VaultConfig">VaultConfig</a>)
</p>
<p>
<p>VaultAWSConfig contains all the Vault configuration needed by Vault to be deployed in AWS</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>VaultAWSUnsealConfig</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.VaultAWSUnsealConfig">
VaultAWSUnsealConfig
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>autoCreate</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>dynamoDBTable</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>dynamoDBRegion</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>iamUserName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.VaultAWSUnsealConfig">VaultAWSUnsealConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.VaultAWSConfig">VaultAWSConfig</a>)
</p>
<p>
<p>VaultAWSUnsealConfig contains references to existing AWS resources that can be used to install Vault</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>kmsKeyId</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kmsRegion</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>s3Bucket</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>s3Prefix</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>s3Region</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.VaultAzureConfig">VaultAzureConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.VaultConfig">VaultConfig</a>)
</p>
<p>
<p>VaultAzureConfig contains all the Vault configuration needed by Vault to be deployed in Azure</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>tenantId</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>vaultName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>containerName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.VaultConfig">VaultConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>VaultConfig contains Vault configuration for Boot</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name the name of the Vault if using Jenkins X managed Vault instance.
Cannot be used in conjunction with the URL attribute</p>
</td>
</tr>
<tr>
<td>
<code>bucket</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>recreateBucket</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyring</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>key</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>disableURLDiscovery</code></br>
<em>
bool
</em>
</td>
<td>
<p>DisableURLDiscovery allows us to optionally override the default lookup of the Vault URL, could be incluster service or external ingress</p>
</td>
</tr>
<tr>
<td>
<code>aws</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.VaultAWSConfig">
VaultAWSConfig
</a>
</em>
</td>
<td>
<p>AWSConfig describes the AWS specific configuration needed for the Vault Operator.</p>
</td>
</tr>
<tr>
<td>
<code>azure</code></br>
<em>
<a href="#core.jenkins-x.io/v4beta1.VaultAzureConfig">
VaultAzureConfig
</a>
</em>
</td>
<td>
<p>AzureConfig describes the Azure specific configuration needed for the Vault Operator.</p>
</td>
</tr>
<tr>
<td>
<code>url</code></br>
<em>
string
</em>
</td>
<td>
<p>URL specifies the URL of an Vault instance to use for secret storage.
Needs to be specified together with the Service Account and namespace to use for connecting to Vault.
This cannot be used in conjunction with the Name attribute.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccount</code></br>
<em>
string
</em>
</td>
<td>
<p>ServiceAccount is the name of the Kubernetes service account allowed to authenticate against Vault.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code></br>
<em>
string
</em>
</td>
<td>
<p>Namespace of the Kubernetes service account allowed to authenticate against Vault.</p>
</td>
</tr>
<tr>
<td>
<code>secretEngineMountPoint</code></br>
<em>
string
</em>
</td>
<td>
<p>SecretEngineMountPoint is the secret engine mount point to be used for writing data into the KV engine of Vault.
If not specified the &lsquo;secret&rsquo; is used.</p>
</td>
</tr>
<tr>
<td>
<code>kubernetesAuthPath</code></br>
<em>
string
</em>
</td>
<td>
<p>KubernetesAuthPath is the auth path of used for this cluster
If not specified the &lsquo;kubernetes&rsquo; is used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="core.jenkins-x.io/v4beta1.WebhookType">WebhookType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#core.jenkins-x.io/v4beta1.RequirementsConfig">RequirementsConfig</a>)
</p>
<p>
<p>WebhookType is the type of a webhook strategy</p>
</p>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
on git commit <code>499503f</code>.
</em></p>
