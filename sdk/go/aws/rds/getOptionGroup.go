// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AWS::RDS::OptionGroup resource creates an option group, to enable and configure features that are specific to a particular DB engine.
func LookupOptionGroup(ctx *pulumi.Context, args *LookupOptionGroupArgs, opts ...pulumi.InvokeOption) (*LookupOptionGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOptionGroupResult
	err := ctx.Invoke("aws-native:rds:getOptionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOptionGroupArgs struct {
	// Specifies the name of the option group.
	OptionGroupName string `pulumi:"optionGroupName"`
}

type LookupOptionGroupResult struct {
	// Indicates what options are available in the option group.
	OptionConfigurations []OptionGroupOptionConfiguration `pulumi:"optionConfigurations"`
	// An array of key-value pairs to apply to this resource.
	Tags []OptionGroupTag `pulumi:"tags"`
}

func LookupOptionGroupOutput(ctx *pulumi.Context, args LookupOptionGroupOutputArgs, opts ...pulumi.InvokeOption) LookupOptionGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOptionGroupResult, error) {
			args := v.(LookupOptionGroupArgs)
			r, err := LookupOptionGroup(ctx, &args, opts...)
			var s LookupOptionGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOptionGroupResultOutput)
}

type LookupOptionGroupOutputArgs struct {
	// Specifies the name of the option group.
	OptionGroupName pulumi.StringInput `pulumi:"optionGroupName"`
}

func (LookupOptionGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOptionGroupArgs)(nil)).Elem()
}

type LookupOptionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupOptionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOptionGroupResult)(nil)).Elem()
}

func (o LookupOptionGroupResultOutput) ToLookupOptionGroupResultOutput() LookupOptionGroupResultOutput {
	return o
}

func (o LookupOptionGroupResultOutput) ToLookupOptionGroupResultOutputWithContext(ctx context.Context) LookupOptionGroupResultOutput {
	return o
}

func (o LookupOptionGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOptionGroupResult] {
	return pulumix.Output[LookupOptionGroupResult]{
		OutputState: o.OutputState,
	}
}

// Indicates what options are available in the option group.
func (o LookupOptionGroupResultOutput) OptionConfigurations() OptionGroupOptionConfigurationArrayOutput {
	return o.ApplyT(func(v LookupOptionGroupResult) []OptionGroupOptionConfiguration { return v.OptionConfigurations }).(OptionGroupOptionConfigurationArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupOptionGroupResultOutput) Tags() OptionGroupTagArrayOutput {
	return o.ApplyT(func(v LookupOptionGroupResult) []OptionGroupTag { return v.Tags }).(OptionGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOptionGroupResultOutput{})
}
