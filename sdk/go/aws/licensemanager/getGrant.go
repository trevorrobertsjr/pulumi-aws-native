// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupGrant(ctx *pulumi.Context, args *LookupGrantArgs, opts ...pulumi.InvokeOption) (*LookupGrantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGrantResult
	err := ctx.Invoke("aws-native:licensemanager:getGrant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGrantArgs struct {
	// Arn of the grant.
	GrantArn string `pulumi:"grantArn"`
}

type LookupGrantResult struct {
	// Arn of the grant.
	GrantArn *string `pulumi:"grantArn"`
	// Name for the created Grant.
	GrantName *string `pulumi:"grantName"`
	// Home region for the created grant.
	HomeRegion *string `pulumi:"homeRegion"`
	// License Arn for the grant.
	LicenseArn *string `pulumi:"licenseArn"`
	// The version of the grant.
	Version *string `pulumi:"version"`
}

func LookupGrantOutput(ctx *pulumi.Context, args LookupGrantOutputArgs, opts ...pulumi.InvokeOption) LookupGrantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGrantResult, error) {
			args := v.(LookupGrantArgs)
			r, err := LookupGrant(ctx, &args, opts...)
			var s LookupGrantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGrantResultOutput)
}

type LookupGrantOutputArgs struct {
	// Arn of the grant.
	GrantArn pulumi.StringInput `pulumi:"grantArn"`
}

func (LookupGrantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrantArgs)(nil)).Elem()
}

type LookupGrantResultOutput struct{ *pulumi.OutputState }

func (LookupGrantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrantResult)(nil)).Elem()
}

func (o LookupGrantResultOutput) ToLookupGrantResultOutput() LookupGrantResultOutput {
	return o
}

func (o LookupGrantResultOutput) ToLookupGrantResultOutputWithContext(ctx context.Context) LookupGrantResultOutput {
	return o
}

func (o LookupGrantResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGrantResult] {
	return pulumix.Output[LookupGrantResult]{
		OutputState: o.OutputState,
	}
}

// Arn of the grant.
func (o LookupGrantResultOutput) GrantArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrantResult) *string { return v.GrantArn }).(pulumi.StringPtrOutput)
}

// Name for the created Grant.
func (o LookupGrantResultOutput) GrantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrantResult) *string { return v.GrantName }).(pulumi.StringPtrOutput)
}

// Home region for the created grant.
func (o LookupGrantResultOutput) HomeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrantResult) *string { return v.HomeRegion }).(pulumi.StringPtrOutput)
}

// License Arn for the grant.
func (o LookupGrantResultOutput) LicenseArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrantResult) *string { return v.LicenseArn }).(pulumi.StringPtrOutput)
}

// The version of the grant.
func (o LookupGrantResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrantResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGrantResultOutput{})
}
