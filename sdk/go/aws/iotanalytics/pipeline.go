// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotanalytics

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::IoTAnalytics::Pipeline
type Pipeline struct {
	pulumi.CustomResourceState

	PipelineActivities PipelineActivityArrayOutput `pulumi:"pipelineActivities"`
	PipelineName       pulumi.StringPtrOutput      `pulumi:"pipelineName"`
	Tags               PipelineTagArrayOutput      `pulumi:"tags"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PipelineActivities == nil {
		return nil, errors.New("invalid value for required argument 'PipelineActivities'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"pipelineName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pipeline
	err := ctx.RegisterResource("aws-native:iotanalytics:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("aws-native:iotanalytics:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
}

type PipelineState struct {
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	PipelineActivities []PipelineActivity `pulumi:"pipelineActivities"`
	PipelineName       *string            `pulumi:"pipelineName"`
	Tags               []PipelineTag      `pulumi:"tags"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	PipelineActivities PipelineActivityArrayInput
	PipelineName       pulumi.StringPtrInput
	Tags               PipelineTagArrayInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

func (i *Pipeline) ToOutput(ctx context.Context) pulumix.Output[*Pipeline] {
	return pulumix.Output[*Pipeline]{
		OutputState: i.ToPipelineOutputWithContext(ctx).OutputState,
	}
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func (o PipelineOutput) ToOutput(ctx context.Context) pulumix.Output[*Pipeline] {
	return pulumix.Output[*Pipeline]{
		OutputState: o.OutputState,
	}
}

func (o PipelineOutput) PipelineActivities() PipelineActivityArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineActivityArrayOutput { return v.PipelineActivities }).(PipelineActivityArrayOutput)
}

func (o PipelineOutput) PipelineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringPtrOutput { return v.PipelineName }).(pulumi.StringPtrOutput)
}

func (o PipelineOutput) Tags() PipelineTagArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineTagArrayOutput { return v.Tags }).(PipelineTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineInput)(nil)).Elem(), &Pipeline{})
	pulumi.RegisterOutputType(PipelineOutput{})
}
