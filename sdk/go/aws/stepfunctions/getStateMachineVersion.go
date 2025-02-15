// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stepfunctions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for StateMachineVersion
func LookupStateMachineVersion(ctx *pulumi.Context, args *LookupStateMachineVersionArgs, opts ...pulumi.InvokeOption) (*LookupStateMachineVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStateMachineVersionResult
	err := ctx.Invoke("aws-native:stepfunctions:getStateMachineVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStateMachineVersionArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupStateMachineVersionResult struct {
	Arn         *string `pulumi:"arn"`
	Description *string `pulumi:"description"`
}

func LookupStateMachineVersionOutput(ctx *pulumi.Context, args LookupStateMachineVersionOutputArgs, opts ...pulumi.InvokeOption) LookupStateMachineVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStateMachineVersionResult, error) {
			args := v.(LookupStateMachineVersionArgs)
			r, err := LookupStateMachineVersion(ctx, &args, opts...)
			var s LookupStateMachineVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStateMachineVersionResultOutput)
}

type LookupStateMachineVersionOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupStateMachineVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStateMachineVersionArgs)(nil)).Elem()
}

type LookupStateMachineVersionResultOutput struct{ *pulumi.OutputState }

func (LookupStateMachineVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStateMachineVersionResult)(nil)).Elem()
}

func (o LookupStateMachineVersionResultOutput) ToLookupStateMachineVersionResultOutput() LookupStateMachineVersionResultOutput {
	return o
}

func (o LookupStateMachineVersionResultOutput) ToLookupStateMachineVersionResultOutputWithContext(ctx context.Context) LookupStateMachineVersionResultOutput {
	return o
}

func (o LookupStateMachineVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStateMachineVersionResult] {
	return pulumix.Output[LookupStateMachineVersionResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupStateMachineVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStateMachineVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupStateMachineVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStateMachineVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStateMachineVersionResultOutput{})
}
