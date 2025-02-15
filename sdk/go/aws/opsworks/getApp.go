// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::OpsWorks::App
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppResult
	err := ctx.Invoke("aws-native:opsworks:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	Id string `pulumi:"id"`
}

type LookupAppResult struct {
	AppSource        *AppSource               `pulumi:"appSource"`
	Attributes       interface{}              `pulumi:"attributes"`
	DataSources      []AppDataSource          `pulumi:"dataSources"`
	Description      *string                  `pulumi:"description"`
	Domains          []string                 `pulumi:"domains"`
	EnableSsl        *bool                    `pulumi:"enableSsl"`
	Environment      []AppEnvironmentVariable `pulumi:"environment"`
	Id               *string                  `pulumi:"id"`
	Name             *string                  `pulumi:"name"`
	SslConfiguration *AppSslConfiguration     `pulumi:"sslConfiguration"`
	Type             *string                  `pulumi:"type"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAppResult] {
	return pulumix.Output[LookupAppResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupAppResultOutput) AppSource() AppSourcePtrOutput {
	return o.ApplyT(func(v LookupAppResult) *AppSource { return v.AppSource }).(AppSourcePtrOutput)
}

func (o LookupAppResultOutput) Attributes() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAppResult) interface{} { return v.Attributes }).(pulumi.AnyOutput)
}

func (o LookupAppResultOutput) DataSources() AppDataSourceArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []AppDataSource { return v.DataSources }).(AppDataSourceArrayOutput)
}

func (o LookupAppResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []string { return v.Domains }).(pulumi.StringArrayOutput)
}

func (o LookupAppResultOutput) EnableSsl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *bool { return v.EnableSsl }).(pulumi.BoolPtrOutput)
}

func (o LookupAppResultOutput) Environment() AppEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []AppEnvironmentVariable { return v.Environment }).(AppEnvironmentVariableArrayOutput)
}

func (o LookupAppResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) SslConfiguration() AppSslConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *AppSslConfiguration { return v.SslConfiguration }).(AppSslConfigurationPtrOutput)
}

func (o LookupAppResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
