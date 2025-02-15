// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object representing an Amazon EKS cluster.
type Cluster struct {
	pulumi.CustomResourceState

	AccessConfig ClusterAccessConfigPtrOutput `pulumi:"accessConfig"`
	// The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The certificate-authority-data for your cluster.
	CertificateAuthorityData pulumi.StringOutput `pulumi:"certificateAuthorityData"`
	// The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.
	ClusterSecurityGroupId pulumi.StringOutput                `pulumi:"clusterSecurityGroupId"`
	EncryptionConfig       ClusterEncryptionConfigArrayOutput `pulumi:"encryptionConfig"`
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	EncryptionConfigKeyArn pulumi.StringOutput `pulumi:"encryptionConfigKeyArn"`
	// The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.
	Endpoint                pulumi.StringOutput                     `pulumi:"endpoint"`
	KubernetesNetworkConfig ClusterKubernetesNetworkConfigPtrOutput `pulumi:"kubernetesNetworkConfig"`
	Logging                 LoggingPtrOutput                        `pulumi:"logging"`
	// The unique name to give to your cluster.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.
	OpenIdConnectIssuerUrl pulumi.StringOutput             `pulumi:"openIdConnectIssuerUrl"`
	OutpostConfig          ClusterOutpostConfigPtrOutput   `pulumi:"outpostConfig"`
	ResourcesVpcConfig     ClusterResourcesVpcConfigOutput `pulumi:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags ClusterTagArrayOutput `pulumi:"tags"`
	// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourcesVpcConfig == nil {
		return nil, errors.New("invalid value for required argument 'ResourcesVpcConfig'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"accessConfig.bootstrapClusterCreatorAdminPermissions",
		"encryptionConfig[*]",
		"kubernetesNetworkConfig",
		"name",
		"outpostConfig",
		"roleArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("aws-native:eks:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws-native:eks:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	AccessConfig            *ClusterAccessConfig            `pulumi:"accessConfig"`
	EncryptionConfig        []ClusterEncryptionConfig       `pulumi:"encryptionConfig"`
	KubernetesNetworkConfig *ClusterKubernetesNetworkConfig `pulumi:"kubernetesNetworkConfig"`
	Logging                 *Logging                        `pulumi:"logging"`
	// The unique name to give to your cluster.
	Name               *string                   `pulumi:"name"`
	OutpostConfig      *ClusterOutpostConfig     `pulumi:"outpostConfig"`
	ResourcesVpcConfig ClusterResourcesVpcConfig `pulumi:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags []ClusterTag `pulumi:"tags"`
	// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	AccessConfig            ClusterAccessConfigPtrInput
	EncryptionConfig        ClusterEncryptionConfigArrayInput
	KubernetesNetworkConfig ClusterKubernetesNetworkConfigPtrInput
	Logging                 LoggingPtrInput
	// The unique name to give to your cluster.
	Name               pulumi.StringPtrInput
	OutpostConfig      ClusterOutpostConfigPtrInput
	ResourcesVpcConfig ClusterResourcesVpcConfigInput
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	RoleArn pulumi.StringInput
	// An array of key-value pairs to apply to this resource.
	Tags ClusterTagArrayInput
	// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
	Version pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: i.ToClusterOutputWithContext(ctx).OutputState,
	}
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: o.OutputState,
	}
}

func (o ClusterOutput) AccessConfig() ClusterAccessConfigPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterAccessConfigPtrOutput { return v.AccessConfig }).(ClusterAccessConfigPtrOutput)
}

// The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
func (o ClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The certificate-authority-data for your cluster.
func (o ClusterOutput) CertificateAuthorityData() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CertificateAuthorityData }).(pulumi.StringOutput)
}

// The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.
func (o ClusterOutput) ClusterSecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterSecurityGroupId }).(pulumi.StringOutput)
}

func (o ClusterOutput) EncryptionConfig() ClusterEncryptionConfigArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterEncryptionConfigArrayOutput { return v.EncryptionConfig }).(ClusterEncryptionConfigArrayOutput)
}

// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
func (o ClusterOutput) EncryptionConfigKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.EncryptionConfigKeyArn }).(pulumi.StringOutput)
}

// The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.
func (o ClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o ClusterOutput) KubernetesNetworkConfig() ClusterKubernetesNetworkConfigPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterKubernetesNetworkConfigPtrOutput { return v.KubernetesNetworkConfig }).(ClusterKubernetesNetworkConfigPtrOutput)
}

func (o ClusterOutput) Logging() LoggingPtrOutput {
	return o.ApplyT(func(v *Cluster) LoggingPtrOutput { return v.Logging }).(LoggingPtrOutput)
}

// The unique name to give to your cluster.
func (o ClusterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.
func (o ClusterOutput) OpenIdConnectIssuerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.OpenIdConnectIssuerUrl }).(pulumi.StringOutput)
}

func (o ClusterOutput) OutpostConfig() ClusterOutpostConfigPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterOutpostConfigPtrOutput { return v.OutpostConfig }).(ClusterOutpostConfigPtrOutput)
}

func (o ClusterOutput) ResourcesVpcConfig() ClusterResourcesVpcConfigOutput {
	return o.ApplyT(func(v *Cluster) ClusterResourcesVpcConfigOutput { return v.ResourcesVpcConfig }).(ClusterResourcesVpcConfigOutput)
}

// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
func (o ClusterOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ClusterOutput) Tags() ClusterTagArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterTagArrayOutput { return v.Tags }).(ClusterTagArrayOutput)
}

// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
func (o ClusterOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterOutputType(ClusterOutput{})
}
