// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::ElastiCache::ParameterGroup
func LookupParameterGroup(ctx *pulumi.Context, args *LookupParameterGroupArgs, opts ...pulumi.InvokeOption) (*LookupParameterGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupParameterGroupResult
	err := ctx.Invoke("aws-native:elasticache:getParameterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupParameterGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupParameterGroupResult struct {
	Description *string             `pulumi:"description"`
	Id          *string             `pulumi:"id"`
	Properties  interface{}         `pulumi:"properties"`
	Tags        []ParameterGroupTag `pulumi:"tags"`
}

func LookupParameterGroupOutput(ctx *pulumi.Context, args LookupParameterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupParameterGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupParameterGroupResult, error) {
			args := v.(LookupParameterGroupArgs)
			r, err := LookupParameterGroup(ctx, &args, opts...)
			var s LookupParameterGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupParameterGroupResultOutput)
}

type LookupParameterGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupParameterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterGroupArgs)(nil)).Elem()
}

type LookupParameterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupParameterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterGroupResult)(nil)).Elem()
}

func (o LookupParameterGroupResultOutput) ToLookupParameterGroupResultOutput() LookupParameterGroupResultOutput {
	return o
}

func (o LookupParameterGroupResultOutput) ToLookupParameterGroupResultOutputWithContext(ctx context.Context) LookupParameterGroupResultOutput {
	return o
}

func (o LookupParameterGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupParameterGroupResult] {
	return pulumix.Output[LookupParameterGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupParameterGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupParameterGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupParameterGroupResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupParameterGroupResultOutput) Tags() ParameterGroupTagArrayOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) []ParameterGroupTag { return v.Tags }).(ParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterGroupResultOutput{})
}
