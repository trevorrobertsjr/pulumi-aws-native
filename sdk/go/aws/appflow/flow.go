// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appflow

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::AppFlow::Flow.
type Flow struct {
	pulumi.CustomResourceState

	// Description of the flow.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of Destination connectors of the flow.
	DestinationFlowConfigList FlowDestinationFlowConfigArrayOutput `pulumi:"destinationFlowConfigList"`
	// ARN identifier of the flow.
	FlowArn pulumi.StringOutput `pulumi:"flowArn"`
	// Name of the flow.
	FlowName pulumi.StringOutput `pulumi:"flowName"`
	// Flow activation status for Scheduled- and Event-triggered flows
	FlowStatus FlowStatusPtrOutput `pulumi:"flowStatus"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
	KmsArn pulumi.StringPtrOutput `pulumi:"kmsArn"`
	// Configurations of metadata catalog of the flow.
	MetadataCatalogConfig FlowMetadataCatalogConfigPtrOutput `pulumi:"metadataCatalogConfig"`
	// Configurations of Source connector of the flow.
	SourceFlowConfig FlowSourceFlowConfigOutput `pulumi:"sourceFlowConfig"`
	// List of Tags.
	Tags FlowTagArrayOutput `pulumi:"tags"`
	// List of tasks for the flow.
	Tasks FlowTaskArrayOutput `pulumi:"tasks"`
	// Trigger settings of the flow.
	TriggerConfig FlowTriggerConfigOutput `pulumi:"triggerConfig"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationFlowConfigList == nil {
		return nil, errors.New("invalid value for required argument 'DestinationFlowConfigList'")
	}
	if args.SourceFlowConfig == nil {
		return nil, errors.New("invalid value for required argument 'SourceFlowConfig'")
	}
	if args.Tasks == nil {
		return nil, errors.New("invalid value for required argument 'Tasks'")
	}
	if args.TriggerConfig == nil {
		return nil, errors.New("invalid value for required argument 'TriggerConfig'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"flowName",
		"kmsArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flow
	err := ctx.RegisterResource("aws-native:appflow:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("aws-native:appflow:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
}

type FlowState struct {
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	// Description of the flow.
	Description *string `pulumi:"description"`
	// List of Destination connectors of the flow.
	DestinationFlowConfigList []FlowDestinationFlowConfig `pulumi:"destinationFlowConfigList"`
	// Name of the flow.
	FlowName *string `pulumi:"flowName"`
	// Flow activation status for Scheduled- and Event-triggered flows
	FlowStatus *FlowStatus `pulumi:"flowStatus"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
	KmsArn *string `pulumi:"kmsArn"`
	// Configurations of metadata catalog of the flow.
	MetadataCatalogConfig *FlowMetadataCatalogConfig `pulumi:"metadataCatalogConfig"`
	// Configurations of Source connector of the flow.
	SourceFlowConfig FlowSourceFlowConfig `pulumi:"sourceFlowConfig"`
	// List of Tags.
	Tags []FlowTag `pulumi:"tags"`
	// List of tasks for the flow.
	Tasks []FlowTask `pulumi:"tasks"`
	// Trigger settings of the flow.
	TriggerConfig FlowTriggerConfig `pulumi:"triggerConfig"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	// Description of the flow.
	Description pulumi.StringPtrInput
	// List of Destination connectors of the flow.
	DestinationFlowConfigList FlowDestinationFlowConfigArrayInput
	// Name of the flow.
	FlowName pulumi.StringPtrInput
	// Flow activation status for Scheduled- and Event-triggered flows
	FlowStatus FlowStatusPtrInput
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
	KmsArn pulumi.StringPtrInput
	// Configurations of metadata catalog of the flow.
	MetadataCatalogConfig FlowMetadataCatalogConfigPtrInput
	// Configurations of Source connector of the flow.
	SourceFlowConfig FlowSourceFlowConfigInput
	// List of Tags.
	Tags FlowTagArrayInput
	// List of tasks for the flow.
	Tasks FlowTaskArrayInput
	// Trigger settings of the flow.
	TriggerConfig FlowTriggerConfigInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}

type FlowInput interface {
	pulumi.Input

	ToFlowOutput() FlowOutput
	ToFlowOutputWithContext(ctx context.Context) FlowOutput
}

func (*Flow) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil)).Elem()
}

func (i *Flow) ToFlowOutput() FlowOutput {
	return i.ToFlowOutputWithContext(context.Background())
}

func (i *Flow) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowOutput)
}

func (i *Flow) ToOutput(ctx context.Context) pulumix.Output[*Flow] {
	return pulumix.Output[*Flow]{
		OutputState: i.ToFlowOutputWithContext(ctx).OutputState,
	}
}

type FlowOutput struct{ *pulumi.OutputState }

func (FlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil)).Elem()
}

func (o FlowOutput) ToFlowOutput() FlowOutput {
	return o
}

func (o FlowOutput) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return o
}

func (o FlowOutput) ToOutput(ctx context.Context) pulumix.Output[*Flow] {
	return pulumix.Output[*Flow]{
		OutputState: o.OutputState,
	}
}

// Description of the flow.
func (o FlowOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of Destination connectors of the flow.
func (o FlowOutput) DestinationFlowConfigList() FlowDestinationFlowConfigArrayOutput {
	return o.ApplyT(func(v *Flow) FlowDestinationFlowConfigArrayOutput { return v.DestinationFlowConfigList }).(FlowDestinationFlowConfigArrayOutput)
}

// ARN identifier of the flow.
func (o FlowOutput) FlowArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.FlowArn }).(pulumi.StringOutput)
}

// Name of the flow.
func (o FlowOutput) FlowName() pulumi.StringOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringOutput { return v.FlowName }).(pulumi.StringOutput)
}

// Flow activation status for Scheduled- and Event-triggered flows
func (o FlowOutput) FlowStatus() FlowStatusPtrOutput {
	return o.ApplyT(func(v *Flow) FlowStatusPtrOutput { return v.FlowStatus }).(FlowStatusPtrOutput)
}

// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
func (o FlowOutput) KmsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flow) pulumi.StringPtrOutput { return v.KmsArn }).(pulumi.StringPtrOutput)
}

// Configurations of metadata catalog of the flow.
func (o FlowOutput) MetadataCatalogConfig() FlowMetadataCatalogConfigPtrOutput {
	return o.ApplyT(func(v *Flow) FlowMetadataCatalogConfigPtrOutput { return v.MetadataCatalogConfig }).(FlowMetadataCatalogConfigPtrOutput)
}

// Configurations of Source connector of the flow.
func (o FlowOutput) SourceFlowConfig() FlowSourceFlowConfigOutput {
	return o.ApplyT(func(v *Flow) FlowSourceFlowConfigOutput { return v.SourceFlowConfig }).(FlowSourceFlowConfigOutput)
}

// List of Tags.
func (o FlowOutput) Tags() FlowTagArrayOutput {
	return o.ApplyT(func(v *Flow) FlowTagArrayOutput { return v.Tags }).(FlowTagArrayOutput)
}

// List of tasks for the flow.
func (o FlowOutput) Tasks() FlowTaskArrayOutput {
	return o.ApplyT(func(v *Flow) FlowTaskArrayOutput { return v.Tasks }).(FlowTaskArrayOutput)
}

// Trigger settings of the flow.
func (o FlowOutput) TriggerConfig() FlowTriggerConfigOutput {
	return o.ApplyT(func(v *Flow) FlowTriggerConfigOutput { return v.TriggerConfig }).(FlowTriggerConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowInput)(nil)).Elem(), &Flow{})
	pulumi.RegisterOutputType(FlowOutput{})
}
