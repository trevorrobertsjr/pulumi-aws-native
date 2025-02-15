// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::SageMaker::Pipeline
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPipelineResult
	err := ctx.Invoke("aws-native:sagemaker:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	// The name of the Pipeline.
	PipelineName string `pulumi:"pipelineName"`
}

type LookupPipelineResult struct {
	ParallelismConfiguration *ParallelismConfigurationProperties `pulumi:"parallelismConfiguration"`
	PipelineDefinition       interface{}                         `pulumi:"pipelineDefinition"`
	// The description of the Pipeline.
	PipelineDescription *string `pulumi:"pipelineDescription"`
	// The display name of the Pipeline.
	PipelineDisplayName *string `pulumi:"pipelineDisplayName"`
	// Role Arn
	RoleArn *string       `pulumi:"roleArn"`
	Tags    []PipelineTag `pulumi:"tags"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			var s LookupPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	// The name of the Pipeline.
	PipelineName pulumi.StringInput `pulumi:"pipelineName"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}

type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPipelineResult] {
	return pulumix.Output[LookupPipelineResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupPipelineResultOutput) ParallelismConfiguration() ParallelismConfigurationPropertiesPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *ParallelismConfigurationProperties { return v.ParallelismConfiguration }).(ParallelismConfigurationPropertiesPtrOutput)
}

func (o LookupPipelineResultOutput) PipelineDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPipelineResult) interface{} { return v.PipelineDefinition }).(pulumi.AnyOutput)
}

// The description of the Pipeline.
func (o LookupPipelineResultOutput) PipelineDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.PipelineDescription }).(pulumi.StringPtrOutput)
}

// The display name of the Pipeline.
func (o LookupPipelineResultOutput) PipelineDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.PipelineDisplayName }).(pulumi.StringPtrOutput)
}

// Role Arn
func (o LookupPipelineResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineResultOutput) Tags() PipelineTagArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineTag { return v.Tags }).(PipelineTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
