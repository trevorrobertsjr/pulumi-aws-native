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

// Resource Type definition for AWS::OpsWorks::UserProfile
func LookupUserProfile(ctx *pulumi.Context, args *LookupUserProfileArgs, opts ...pulumi.InvokeOption) (*LookupUserProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserProfileResult
	err := ctx.Invoke("aws-native:opsworks:getUserProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserProfileArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserProfileResult struct {
	AllowSelfManagement *bool   `pulumi:"allowSelfManagement"`
	Id                  *string `pulumi:"id"`
	SshPublicKey        *string `pulumi:"sshPublicKey"`
	SshUsername         *string `pulumi:"sshUsername"`
}

func LookupUserProfileOutput(ctx *pulumi.Context, args LookupUserProfileOutputArgs, opts ...pulumi.InvokeOption) LookupUserProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserProfileResult, error) {
			args := v.(LookupUserProfileArgs)
			r, err := LookupUserProfile(ctx, &args, opts...)
			var s LookupUserProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserProfileResultOutput)
}

type LookupUserProfileOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserProfileArgs)(nil)).Elem()
}

type LookupUserProfileResultOutput struct{ *pulumi.OutputState }

func (LookupUserProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserProfileResult)(nil)).Elem()
}

func (o LookupUserProfileResultOutput) ToLookupUserProfileResultOutput() LookupUserProfileResultOutput {
	return o
}

func (o LookupUserProfileResultOutput) ToLookupUserProfileResultOutputWithContext(ctx context.Context) LookupUserProfileResultOutput {
	return o
}

func (o LookupUserProfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserProfileResult] {
	return pulumix.Output[LookupUserProfileResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupUserProfileResultOutput) AllowSelfManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserProfileResult) *bool { return v.AllowSelfManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupUserProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupUserProfileResultOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserProfileResult) *string { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupUserProfileResultOutput) SshUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserProfileResult) *string { return v.SshUsername }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserProfileResultOutput{})
}
