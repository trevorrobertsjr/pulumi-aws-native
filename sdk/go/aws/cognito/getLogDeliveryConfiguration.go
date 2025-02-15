// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Cognito::LogDeliveryConfiguration
func LookupLogDeliveryConfiguration(ctx *pulumi.Context, args *LookupLogDeliveryConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupLogDeliveryConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLogDeliveryConfigurationResult
	err := ctx.Invoke("aws-native:cognito:getLogDeliveryConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLogDeliveryConfigurationArgs struct {
	Id string `pulumi:"id"`
}

type LookupLogDeliveryConfigurationResult struct {
	Id                *string                                    `pulumi:"id"`
	LogConfigurations []LogDeliveryConfigurationLogConfiguration `pulumi:"logConfigurations"`
}

func LookupLogDeliveryConfigurationOutput(ctx *pulumi.Context, args LookupLogDeliveryConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupLogDeliveryConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogDeliveryConfigurationResult, error) {
			args := v.(LookupLogDeliveryConfigurationArgs)
			r, err := LookupLogDeliveryConfiguration(ctx, &args, opts...)
			var s LookupLogDeliveryConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogDeliveryConfigurationResultOutput)
}

type LookupLogDeliveryConfigurationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLogDeliveryConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogDeliveryConfigurationArgs)(nil)).Elem()
}

type LookupLogDeliveryConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupLogDeliveryConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogDeliveryConfigurationResult)(nil)).Elem()
}

func (o LookupLogDeliveryConfigurationResultOutput) ToLookupLogDeliveryConfigurationResultOutput() LookupLogDeliveryConfigurationResultOutput {
	return o
}

func (o LookupLogDeliveryConfigurationResultOutput) ToLookupLogDeliveryConfigurationResultOutputWithContext(ctx context.Context) LookupLogDeliveryConfigurationResultOutput {
	return o
}

func (o LookupLogDeliveryConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLogDeliveryConfigurationResult] {
	return pulumix.Output[LookupLogDeliveryConfigurationResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupLogDeliveryConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogDeliveryConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLogDeliveryConfigurationResultOutput) LogConfigurations() LogDeliveryConfigurationLogConfigurationArrayOutput {
	return o.ApplyT(func(v LookupLogDeliveryConfigurationResult) []LogDeliveryConfigurationLogConfiguration {
		return v.LogConfigurations
	}).(LogDeliveryConfigurationLogConfigurationArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogDeliveryConfigurationResultOutput{})
}
