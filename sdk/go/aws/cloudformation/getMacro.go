// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::CloudFormation::Macro
func LookupMacro(ctx *pulumi.Context, args *LookupMacroArgs, opts ...pulumi.InvokeOption) (*LookupMacroResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMacroResult
	err := ctx.Invoke("aws-native:cloudformation:getMacro", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMacroArgs struct {
	Id string `pulumi:"id"`
}

type LookupMacroResult struct {
	Description  *string `pulumi:"description"`
	FunctionName *string `pulumi:"functionName"`
	Id           *string `pulumi:"id"`
	LogGroupName *string `pulumi:"logGroupName"`
	LogRoleArn   *string `pulumi:"logRoleArn"`
}

func LookupMacroOutput(ctx *pulumi.Context, args LookupMacroOutputArgs, opts ...pulumi.InvokeOption) LookupMacroResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMacroResult, error) {
			args := v.(LookupMacroArgs)
			r, err := LookupMacro(ctx, &args, opts...)
			var s LookupMacroResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMacroResultOutput)
}

type LookupMacroOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMacroOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMacroArgs)(nil)).Elem()
}

type LookupMacroResultOutput struct{ *pulumi.OutputState }

func (LookupMacroResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMacroResult)(nil)).Elem()
}

func (o LookupMacroResultOutput) ToLookupMacroResultOutput() LookupMacroResultOutput {
	return o
}

func (o LookupMacroResultOutput) ToLookupMacroResultOutputWithContext(ctx context.Context) LookupMacroResultOutput {
	return o
}

func (o LookupMacroResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMacroResult] {
	return pulumix.Output[LookupMacroResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupMacroResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMacroResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMacroResultOutput) FunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMacroResult) *string { return v.FunctionName }).(pulumi.StringPtrOutput)
}

func (o LookupMacroResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMacroResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMacroResultOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMacroResult) *string { return v.LogGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupMacroResultOutput) LogRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMacroResult) *string { return v.LogRoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMacroResultOutput{})
}
