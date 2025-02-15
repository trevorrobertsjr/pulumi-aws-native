// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Greengrass::ConnectorDefinition
func LookupConnectorDefinition(ctx *pulumi.Context, args *LookupConnectorDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupConnectorDefinitionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorDefinitionResult
	err := ctx.Invoke("aws-native:greengrass:getConnectorDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorDefinitionArgs struct {
	Id string `pulumi:"id"`
}

type LookupConnectorDefinitionResult struct {
	Arn              *string     `pulumi:"arn"`
	Id               *string     `pulumi:"id"`
	LatestVersionArn *string     `pulumi:"latestVersionArn"`
	Name             *string     `pulumi:"name"`
	Tags             interface{} `pulumi:"tags"`
}

func LookupConnectorDefinitionOutput(ctx *pulumi.Context, args LookupConnectorDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorDefinitionResult, error) {
			args := v.(LookupConnectorDefinitionArgs)
			r, err := LookupConnectorDefinition(ctx, &args, opts...)
			var s LookupConnectorDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorDefinitionResultOutput)
}

type LookupConnectorDefinitionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupConnectorDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorDefinitionArgs)(nil)).Elem()
}

type LookupConnectorDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorDefinitionResult)(nil)).Elem()
}

func (o LookupConnectorDefinitionResultOutput) ToLookupConnectorDefinitionResultOutput() LookupConnectorDefinitionResultOutput {
	return o
}

func (o LookupConnectorDefinitionResultOutput) ToLookupConnectorDefinitionResultOutputWithContext(ctx context.Context) LookupConnectorDefinitionResultOutput {
	return o
}

func (o LookupConnectorDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectorDefinitionResult] {
	return pulumix.Output[LookupConnectorDefinitionResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupConnectorDefinitionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorDefinitionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorDefinitionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorDefinitionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorDefinitionResultOutput) LatestVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorDefinitionResult) *string { return v.LatestVersionArn }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorDefinitionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorDefinitionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorDefinitionResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorDefinitionResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorDefinitionResultOutput{})
}
