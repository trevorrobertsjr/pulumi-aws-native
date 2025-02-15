// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::MediaLive::InputSecurityGroup
func LookupInputSecurityGroup(ctx *pulumi.Context, args *LookupInputSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupInputSecurityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInputSecurityGroupResult
	err := ctx.Invoke("aws-native:medialive:getInputSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInputSecurityGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupInputSecurityGroupResult struct {
	Arn            *string                                    `pulumi:"arn"`
	Id             *string                                    `pulumi:"id"`
	Tags           interface{}                                `pulumi:"tags"`
	WhitelistRules []InputSecurityGroupInputWhitelistRuleCidr `pulumi:"whitelistRules"`
}

func LookupInputSecurityGroupOutput(ctx *pulumi.Context, args LookupInputSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupInputSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInputSecurityGroupResult, error) {
			args := v.(LookupInputSecurityGroupArgs)
			r, err := LookupInputSecurityGroup(ctx, &args, opts...)
			var s LookupInputSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInputSecurityGroupResultOutput)
}

type LookupInputSecurityGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupInputSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputSecurityGroupArgs)(nil)).Elem()
}

type LookupInputSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupInputSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputSecurityGroupResult)(nil)).Elem()
}

func (o LookupInputSecurityGroupResultOutput) ToLookupInputSecurityGroupResultOutput() LookupInputSecurityGroupResultOutput {
	return o
}

func (o LookupInputSecurityGroupResultOutput) ToLookupInputSecurityGroupResultOutputWithContext(ctx context.Context) LookupInputSecurityGroupResultOutput {
	return o
}

func (o LookupInputSecurityGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInputSecurityGroupResult] {
	return pulumix.Output[LookupInputSecurityGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupInputSecurityGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputSecurityGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupInputSecurityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputSecurityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInputSecurityGroupResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupInputSecurityGroupResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupInputSecurityGroupResultOutput) WhitelistRules() InputSecurityGroupInputWhitelistRuleCidrArrayOutput {
	return o.ApplyT(func(v LookupInputSecurityGroupResult) []InputSecurityGroupInputWhitelistRuleCidr {
		return v.WhitelistRules
	}).(InputSecurityGroupInputWhitelistRuleCidrArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInputSecurityGroupResultOutput{})
}
