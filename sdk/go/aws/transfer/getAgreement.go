// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Transfer::Agreement
func LookupAgreement(ctx *pulumi.Context, args *LookupAgreementArgs, opts ...pulumi.InvokeOption) (*LookupAgreementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAgreementResult
	err := ctx.Invoke("aws-native:transfer:getAgreement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgreementArgs struct {
	// A unique identifier for the agreement.
	AgreementId string `pulumi:"agreementId"`
	// A unique identifier for the server.
	ServerId string `pulumi:"serverId"`
}

type LookupAgreementResult struct {
	// Specifies the access role for the agreement.
	AccessRole *string `pulumi:"accessRole"`
	// A unique identifier for the agreement.
	AgreementId *string `pulumi:"agreementId"`
	// Specifies the unique Amazon Resource Name (ARN) for the agreement.
	Arn *string `pulumi:"arn"`
	// Specifies the base directory for the agreement.
	BaseDirectory *string `pulumi:"baseDirectory"`
	// A textual description for the agreement.
	Description *string `pulumi:"description"`
	// A unique identifier for the local profile.
	LocalProfileId *string `pulumi:"localProfileId"`
	// A unique identifier for the partner profile.
	PartnerProfileId *string `pulumi:"partnerProfileId"`
	// Specifies the status of the agreement.
	Status *AgreementStatus `pulumi:"status"`
	// Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.
	Tags []AgreementTag `pulumi:"tags"`
}

func LookupAgreementOutput(ctx *pulumi.Context, args LookupAgreementOutputArgs, opts ...pulumi.InvokeOption) LookupAgreementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgreementResult, error) {
			args := v.(LookupAgreementArgs)
			r, err := LookupAgreement(ctx, &args, opts...)
			var s LookupAgreementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgreementResultOutput)
}

type LookupAgreementOutputArgs struct {
	// A unique identifier for the agreement.
	AgreementId pulumi.StringInput `pulumi:"agreementId"`
	// A unique identifier for the server.
	ServerId pulumi.StringInput `pulumi:"serverId"`
}

func (LookupAgreementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgreementArgs)(nil)).Elem()
}

type LookupAgreementResultOutput struct{ *pulumi.OutputState }

func (LookupAgreementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgreementResult)(nil)).Elem()
}

func (o LookupAgreementResultOutput) ToLookupAgreementResultOutput() LookupAgreementResultOutput {
	return o
}

func (o LookupAgreementResultOutput) ToLookupAgreementResultOutputWithContext(ctx context.Context) LookupAgreementResultOutput {
	return o
}

func (o LookupAgreementResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAgreementResult] {
	return pulumix.Output[LookupAgreementResult]{
		OutputState: o.OutputState,
	}
}

// Specifies the access role for the agreement.
func (o LookupAgreementResultOutput) AccessRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.AccessRole }).(pulumi.StringPtrOutput)
}

// A unique identifier for the agreement.
func (o LookupAgreementResultOutput) AgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.AgreementId }).(pulumi.StringPtrOutput)
}

// Specifies the unique Amazon Resource Name (ARN) for the agreement.
func (o LookupAgreementResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Specifies the base directory for the agreement.
func (o LookupAgreementResultOutput) BaseDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.BaseDirectory }).(pulumi.StringPtrOutput)
}

// A textual description for the agreement.
func (o LookupAgreementResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique identifier for the local profile.
func (o LookupAgreementResultOutput) LocalProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.LocalProfileId }).(pulumi.StringPtrOutput)
}

// A unique identifier for the partner profile.
func (o LookupAgreementResultOutput) PartnerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.PartnerProfileId }).(pulumi.StringPtrOutput)
}

// Specifies the status of the agreement.
func (o LookupAgreementResultOutput) Status() AgreementStatusPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *AgreementStatus { return v.Status }).(AgreementStatusPtrOutput)
}

// Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.
func (o LookupAgreementResultOutput) Tags() AgreementTagArrayOutput {
	return o.ApplyT(func(v LookupAgreementResult) []AgreementTag { return v.Tags }).(AgreementTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgreementResultOutput{})
}
