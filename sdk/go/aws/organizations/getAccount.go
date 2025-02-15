// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// You can use AWS::Organizations::Account to manage accounts in organization.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("aws-native:organizations:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	// If the account was created successfully, the unique identifier (ID) of the new account.
	AccountId string `pulumi:"accountId"`
}

type LookupAccountResult struct {
	// If the account was created successfully, the unique identifier (ID) of the new account.
	AccountId *string `pulumi:"accountId"`
	// The friendly name of the member account.
	AccountName *string `pulumi:"accountName"`
	// The Amazon Resource Name (ARN) of the account.
	Arn *string `pulumi:"arn"`
	// The email address of the owner to assign to the new member account.
	Email *string `pulumi:"email"`
	// The method by which the account joined the organization.
	JoinedMethod *AccountJoinedMethod `pulumi:"joinedMethod"`
	// The date the account became a part of the organization.
	JoinedTimestamp *string `pulumi:"joinedTimestamp"`
	// List of parent nodes for the member account. Currently only one parent at a time is supported. Default is root.
	ParentIds []string `pulumi:"parentIds"`
	// The status of the account in the organization.
	Status *AccountStatus `pulumi:"status"`
	// A list of tags that you want to attach to the newly created account. For each tag in the list, you must specify both a tag key and a value.
	Tags []AccountTag `pulumi:"tags"`
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	// If the account was created successfully, the unique identifier (ID) of the new account.
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccountResult] {
	return pulumix.Output[LookupAccountResult]{
		OutputState: o.OutputState,
	}
}

// If the account was created successfully, the unique identifier (ID) of the new account.
func (o LookupAccountResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The friendly name of the member account.
func (o LookupAccountResultOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the account.
func (o LookupAccountResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The email address of the owner to assign to the new member account.
func (o LookupAccountResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The method by which the account joined the organization.
func (o LookupAccountResultOutput) JoinedMethod() AccountJoinedMethodPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *AccountJoinedMethod { return v.JoinedMethod }).(AccountJoinedMethodPtrOutput)
}

// The date the account became a part of the organization.
func (o LookupAccountResultOutput) JoinedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.JoinedTimestamp }).(pulumi.StringPtrOutput)
}

// List of parent nodes for the member account. Currently only one parent at a time is supported. Default is root.
func (o LookupAccountResultOutput) ParentIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []string { return v.ParentIds }).(pulumi.StringArrayOutput)
}

// The status of the account in the organization.
func (o LookupAccountResultOutput) Status() AccountStatusPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *AccountStatus { return v.Status }).(AccountStatusPtrOutput)
}

// A list of tags that you want to attach to the newly created account. For each tag in the list, you must specify both a tag key and a value.
func (o LookupAccountResultOutput) Tags() AccountTagArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []AccountTag { return v.Tags }).(AccountTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
