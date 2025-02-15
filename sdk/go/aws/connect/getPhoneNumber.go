// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Connect::PhoneNumber
func LookupPhoneNumber(ctx *pulumi.Context, args *LookupPhoneNumberArgs, opts ...pulumi.InvokeOption) (*LookupPhoneNumberResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPhoneNumberResult
	err := ctx.Invoke("aws-native:connect:getPhoneNumber", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPhoneNumberArgs struct {
	// The phone number ARN
	PhoneNumberArn string `pulumi:"phoneNumberArn"`
}

type LookupPhoneNumberResult struct {
	// The phone number e164 address.
	Address *string `pulumi:"address"`
	// The description of the phone number.
	Description *string `pulumi:"description"`
	// The phone number ARN
	PhoneNumberArn *string `pulumi:"phoneNumberArn"`
	// One or more tags.
	Tags []PhoneNumberTag `pulumi:"tags"`
	// The ARN of the target the phone number is claimed to.
	TargetArn *string `pulumi:"targetArn"`
}

func LookupPhoneNumberOutput(ctx *pulumi.Context, args LookupPhoneNumberOutputArgs, opts ...pulumi.InvokeOption) LookupPhoneNumberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPhoneNumberResult, error) {
			args := v.(LookupPhoneNumberArgs)
			r, err := LookupPhoneNumber(ctx, &args, opts...)
			var s LookupPhoneNumberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPhoneNumberResultOutput)
}

type LookupPhoneNumberOutputArgs struct {
	// The phone number ARN
	PhoneNumberArn pulumi.StringInput `pulumi:"phoneNumberArn"`
}

func (LookupPhoneNumberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPhoneNumberArgs)(nil)).Elem()
}

type LookupPhoneNumberResultOutput struct{ *pulumi.OutputState }

func (LookupPhoneNumberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPhoneNumberResult)(nil)).Elem()
}

func (o LookupPhoneNumberResultOutput) ToLookupPhoneNumberResultOutput() LookupPhoneNumberResultOutput {
	return o
}

func (o LookupPhoneNumberResultOutput) ToLookupPhoneNumberResultOutputWithContext(ctx context.Context) LookupPhoneNumberResultOutput {
	return o
}

func (o LookupPhoneNumberResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPhoneNumberResult] {
	return pulumix.Output[LookupPhoneNumberResult]{
		OutputState: o.OutputState,
	}
}

// The phone number e164 address.
func (o LookupPhoneNumberResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPhoneNumberResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The description of the phone number.
func (o LookupPhoneNumberResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPhoneNumberResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The phone number ARN
func (o LookupPhoneNumberResultOutput) PhoneNumberArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPhoneNumberResult) *string { return v.PhoneNumberArn }).(pulumi.StringPtrOutput)
}

// One or more tags.
func (o LookupPhoneNumberResultOutput) Tags() PhoneNumberTagArrayOutput {
	return o.ApplyT(func(v LookupPhoneNumberResult) []PhoneNumberTag { return v.Tags }).(PhoneNumberTagArrayOutput)
}

// The ARN of the target the phone number is claimed to.
func (o LookupPhoneNumberResultOutput) TargetArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPhoneNumberResult) *string { return v.TargetArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPhoneNumberResultOutput{})
}
