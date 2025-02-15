// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupStudioSessionMapping(ctx *pulumi.Context, args *LookupStudioSessionMappingArgs, opts ...pulumi.InvokeOption) (*LookupStudioSessionMappingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStudioSessionMappingResult
	err := ctx.Invoke("aws-native:emr:getStudioSessionMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudioSessionMappingArgs struct {
	// The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
	IdentityName string `pulumi:"identityName"`
	// Specifies whether the identity to map to the Studio is a user or a group.
	IdentityType StudioSessionMappingIdentityType `pulumi:"identityType"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId string `pulumi:"studioId"`
}

type LookupStudioSessionMappingResult struct {
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
	SessionPolicyArn *string `pulumi:"sessionPolicyArn"`
}

func LookupStudioSessionMappingOutput(ctx *pulumi.Context, args LookupStudioSessionMappingOutputArgs, opts ...pulumi.InvokeOption) LookupStudioSessionMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStudioSessionMappingResult, error) {
			args := v.(LookupStudioSessionMappingArgs)
			r, err := LookupStudioSessionMapping(ctx, &args, opts...)
			var s LookupStudioSessionMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStudioSessionMappingResultOutput)
}

type LookupStudioSessionMappingOutputArgs struct {
	// The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
	IdentityName pulumi.StringInput `pulumi:"identityName"`
	// Specifies whether the identity to map to the Studio is a user or a group.
	IdentityType StudioSessionMappingIdentityTypeInput `pulumi:"identityType"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId pulumi.StringInput `pulumi:"studioId"`
}

func (LookupStudioSessionMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioSessionMappingArgs)(nil)).Elem()
}

type LookupStudioSessionMappingResultOutput struct{ *pulumi.OutputState }

func (LookupStudioSessionMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioSessionMappingResult)(nil)).Elem()
}

func (o LookupStudioSessionMappingResultOutput) ToLookupStudioSessionMappingResultOutput() LookupStudioSessionMappingResultOutput {
	return o
}

func (o LookupStudioSessionMappingResultOutput) ToLookupStudioSessionMappingResultOutputWithContext(ctx context.Context) LookupStudioSessionMappingResultOutput {
	return o
}

func (o LookupStudioSessionMappingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStudioSessionMappingResult] {
	return pulumix.Output[LookupStudioSessionMappingResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
func (o LookupStudioSessionMappingResultOutput) SessionPolicyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioSessionMappingResult) *string { return v.SessionPolicyArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudioSessionMappingResultOutput{})
}
