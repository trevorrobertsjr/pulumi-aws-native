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

// Resource Type definition for AWS::SageMaker::NotebookInstanceLifecycleConfig
func LookupNotebookInstanceLifecycleConfig(ctx *pulumi.Context, args *LookupNotebookInstanceLifecycleConfigArgs, opts ...pulumi.InvokeOption) (*LookupNotebookInstanceLifecycleConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNotebookInstanceLifecycleConfigResult
	err := ctx.Invoke("aws-native:sagemaker:getNotebookInstanceLifecycleConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotebookInstanceLifecycleConfigArgs struct {
	Id string `pulumi:"id"`
}

type LookupNotebookInstanceLifecycleConfigResult struct {
	Id       *string                                                        `pulumi:"id"`
	OnCreate []NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook `pulumi:"onCreate"`
	OnStart  []NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook `pulumi:"onStart"`
}

func LookupNotebookInstanceLifecycleConfigOutput(ctx *pulumi.Context, args LookupNotebookInstanceLifecycleConfigOutputArgs, opts ...pulumi.InvokeOption) LookupNotebookInstanceLifecycleConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotebookInstanceLifecycleConfigResult, error) {
			args := v.(LookupNotebookInstanceLifecycleConfigArgs)
			r, err := LookupNotebookInstanceLifecycleConfig(ctx, &args, opts...)
			var s LookupNotebookInstanceLifecycleConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotebookInstanceLifecycleConfigResultOutput)
}

type LookupNotebookInstanceLifecycleConfigOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNotebookInstanceLifecycleConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookInstanceLifecycleConfigArgs)(nil)).Elem()
}

type LookupNotebookInstanceLifecycleConfigResultOutput struct{ *pulumi.OutputState }

func (LookupNotebookInstanceLifecycleConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookInstanceLifecycleConfigResult)(nil)).Elem()
}

func (o LookupNotebookInstanceLifecycleConfigResultOutput) ToLookupNotebookInstanceLifecycleConfigResultOutput() LookupNotebookInstanceLifecycleConfigResultOutput {
	return o
}

func (o LookupNotebookInstanceLifecycleConfigResultOutput) ToLookupNotebookInstanceLifecycleConfigResultOutputWithContext(ctx context.Context) LookupNotebookInstanceLifecycleConfigResultOutput {
	return o
}

func (o LookupNotebookInstanceLifecycleConfigResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNotebookInstanceLifecycleConfigResult] {
	return pulumix.Output[LookupNotebookInstanceLifecycleConfigResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupNotebookInstanceLifecycleConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotebookInstanceLifecycleConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNotebookInstanceLifecycleConfigResultOutput) OnCreate() NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookArrayOutput {
	return o.ApplyT(func(v LookupNotebookInstanceLifecycleConfigResult) []NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook {
		return v.OnCreate
	}).(NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookArrayOutput)
}

func (o LookupNotebookInstanceLifecycleConfigResultOutput) OnStart() NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookArrayOutput {
	return o.ApplyT(func(v LookupNotebookInstanceLifecycleConfigResult) []NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook {
		return v.OnStart
	}).(NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHookArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotebookInstanceLifecycleConfigResultOutput{})
}
