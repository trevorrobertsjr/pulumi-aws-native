// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A replication configuration that you later provide to configure and start a AWS DMS Serverless replication
type ReplicationConfig struct {
	pulumi.CustomResourceState

	ComputeConfig ReplicationConfigComputeConfigPtrOutput `pulumi:"computeConfig"`
	// The Amazon Resource Name (ARN) of the Replication Config
	ReplicationConfigArn pulumi.StringPtrOutput `pulumi:"replicationConfigArn"`
	// A unique identifier of replication configuration
	ReplicationConfigIdentifier pulumi.StringPtrOutput `pulumi:"replicationConfigIdentifier"`
	// JSON settings for Servereless replications that are provisioned using this replication configuration
	ReplicationSettings pulumi.AnyOutput `pulumi:"replicationSettings"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration
	ReplicationType ReplicationConfigReplicationTypePtrOutput `pulumi:"replicationType"`
	// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource
	ResourceIdentifier pulumi.StringPtrOutput `pulumi:"resourceIdentifier"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
	SourceEndpointArn pulumi.StringPtrOutput `pulumi:"sourceEndpointArn"`
	// JSON settings for specifying supplemental data
	SupplementalSettings pulumi.AnyOutput `pulumi:"supplementalSettings"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
	TableMappings pulumi.AnyOutput `pulumi:"tableMappings"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
	Tags ReplicationConfigTagArrayOutput `pulumi:"tags"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
	TargetEndpointArn pulumi.StringPtrOutput `pulumi:"targetEndpointArn"`
}

// NewReplicationConfig registers a new resource with the given unique name, arguments, and options.
func NewReplicationConfig(ctx *pulumi.Context,
	name string, args *ReplicationConfigArgs, opts ...pulumi.ResourceOption) (*ReplicationConfig, error) {
	if args == nil {
		args = &ReplicationConfigArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"resourceIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicationConfig
	err := ctx.RegisterResource("aws-native:dms:ReplicationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationConfig gets an existing ReplicationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationConfigState, opts ...pulumi.ResourceOption) (*ReplicationConfig, error) {
	var resource ReplicationConfig
	err := ctx.ReadResource("aws-native:dms:ReplicationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationConfig resources.
type replicationConfigState struct {
}

type ReplicationConfigState struct {
}

func (ReplicationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigState)(nil)).Elem()
}

type replicationConfigArgs struct {
	ComputeConfig *ReplicationConfigComputeConfig `pulumi:"computeConfig"`
	// The Amazon Resource Name (ARN) of the Replication Config
	ReplicationConfigArn *string `pulumi:"replicationConfigArn"`
	// A unique identifier of replication configuration
	ReplicationConfigIdentifier *string `pulumi:"replicationConfigIdentifier"`
	// JSON settings for Servereless replications that are provisioned using this replication configuration
	ReplicationSettings interface{} `pulumi:"replicationSettings"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration
	ReplicationType *ReplicationConfigReplicationType `pulumi:"replicationType"`
	// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource
	ResourceIdentifier *string `pulumi:"resourceIdentifier"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
	SourceEndpointArn *string `pulumi:"sourceEndpointArn"`
	// JSON settings for specifying supplemental data
	SupplementalSettings interface{} `pulumi:"supplementalSettings"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
	TableMappings interface{} `pulumi:"tableMappings"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
	Tags []ReplicationConfigTag `pulumi:"tags"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
	TargetEndpointArn *string `pulumi:"targetEndpointArn"`
}

// The set of arguments for constructing a ReplicationConfig resource.
type ReplicationConfigArgs struct {
	ComputeConfig ReplicationConfigComputeConfigPtrInput
	// The Amazon Resource Name (ARN) of the Replication Config
	ReplicationConfigArn pulumi.StringPtrInput
	// A unique identifier of replication configuration
	ReplicationConfigIdentifier pulumi.StringPtrInput
	// JSON settings for Servereless replications that are provisioned using this replication configuration
	ReplicationSettings pulumi.Input
	// The type of AWS DMS Serverless replication to provision using this replication configuration
	ReplicationType ReplicationConfigReplicationTypePtrInput
	// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource
	ResourceIdentifier pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
	SourceEndpointArn pulumi.StringPtrInput
	// JSON settings for specifying supplemental data
	SupplementalSettings pulumi.Input
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
	TableMappings pulumi.Input
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
	Tags ReplicationConfigTagArrayInput
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
	TargetEndpointArn pulumi.StringPtrInput
}

func (ReplicationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigArgs)(nil)).Elem()
}

type ReplicationConfigInput interface {
	pulumi.Input

	ToReplicationConfigOutput() ReplicationConfigOutput
	ToReplicationConfigOutputWithContext(ctx context.Context) ReplicationConfigOutput
}

func (*ReplicationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfig)(nil)).Elem()
}

func (i *ReplicationConfig) ToReplicationConfigOutput() ReplicationConfigOutput {
	return i.ToReplicationConfigOutputWithContext(context.Background())
}

func (i *ReplicationConfig) ToReplicationConfigOutputWithContext(ctx context.Context) ReplicationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigOutput)
}

func (i *ReplicationConfig) ToOutput(ctx context.Context) pulumix.Output[*ReplicationConfig] {
	return pulumix.Output[*ReplicationConfig]{
		OutputState: i.ToReplicationConfigOutputWithContext(ctx).OutputState,
	}
}

type ReplicationConfigOutput struct{ *pulumi.OutputState }

func (ReplicationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfig)(nil)).Elem()
}

func (o ReplicationConfigOutput) ToReplicationConfigOutput() ReplicationConfigOutput {
	return o
}

func (o ReplicationConfigOutput) ToReplicationConfigOutputWithContext(ctx context.Context) ReplicationConfigOutput {
	return o
}

func (o ReplicationConfigOutput) ToOutput(ctx context.Context) pulumix.Output[*ReplicationConfig] {
	return pulumix.Output[*ReplicationConfig]{
		OutputState: o.OutputState,
	}
}

func (o ReplicationConfigOutput) ComputeConfig() ReplicationConfigComputeConfigPtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) ReplicationConfigComputeConfigPtrOutput { return v.ComputeConfig }).(ReplicationConfigComputeConfigPtrOutput)
}

// The Amazon Resource Name (ARN) of the Replication Config
func (o ReplicationConfigOutput) ReplicationConfigArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringPtrOutput { return v.ReplicationConfigArn }).(pulumi.StringPtrOutput)
}

// A unique identifier of replication configuration
func (o ReplicationConfigOutput) ReplicationConfigIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringPtrOutput { return v.ReplicationConfigIdentifier }).(pulumi.StringPtrOutput)
}

// JSON settings for Servereless replications that are provisioned using this replication configuration
func (o ReplicationConfigOutput) ReplicationSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.AnyOutput { return v.ReplicationSettings }).(pulumi.AnyOutput)
}

// The type of AWS DMS Serverless replication to provision using this replication configuration
func (o ReplicationConfigOutput) ReplicationType() ReplicationConfigReplicationTypePtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) ReplicationConfigReplicationTypePtrOutput { return v.ReplicationType }).(ReplicationConfigReplicationTypePtrOutput)
}

// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource
func (o ReplicationConfigOutput) ResourceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringPtrOutput { return v.ResourceIdentifier }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration
func (o ReplicationConfigOutput) SourceEndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringPtrOutput { return v.SourceEndpointArn }).(pulumi.StringPtrOutput)
}

// JSON settings for specifying supplemental data
func (o ReplicationConfigOutput) SupplementalSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.AnyOutput { return v.SupplementalSettings }).(pulumi.AnyOutput)
}

// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration
func (o ReplicationConfigOutput) TableMappings() pulumi.AnyOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.AnyOutput { return v.TableMappings }).(pulumi.AnyOutput)
}

// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>
func (o ReplicationConfigOutput) Tags() ReplicationConfigTagArrayOutput {
	return o.ApplyT(func(v *ReplicationConfig) ReplicationConfigTagArrayOutput { return v.Tags }).(ReplicationConfigTagArrayOutput)
}

// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration
func (o ReplicationConfigOutput) TargetEndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringPtrOutput { return v.TargetEndpointArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigInput)(nil)).Elem(), &ReplicationConfig{})
	pulumi.RegisterOutputType(ReplicationConfigOutput{})
}
