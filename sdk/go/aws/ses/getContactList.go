// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::SES::ContactList.
func LookupContactList(ctx *pulumi.Context, args *LookupContactListArgs, opts ...pulumi.InvokeOption) (*LookupContactListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContactListResult
	err := ctx.Invoke("aws-native:ses:getContactList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactListArgs struct {
	// The name of the contact list.
	ContactListName string `pulumi:"contactListName"`
}

type LookupContactListResult struct {
	// The description of the contact list.
	Description *string `pulumi:"description"`
	// The tags (keys and values) associated with the contact list.
	Tags []ContactListTag `pulumi:"tags"`
	// The topics associated with the contact list.
	Topics []ContactListTopic `pulumi:"topics"`
}

func LookupContactListOutput(ctx *pulumi.Context, args LookupContactListOutputArgs, opts ...pulumi.InvokeOption) LookupContactListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactListResult, error) {
			args := v.(LookupContactListArgs)
			r, err := LookupContactList(ctx, &args, opts...)
			var s LookupContactListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactListResultOutput)
}

type LookupContactListOutputArgs struct {
	// The name of the contact list.
	ContactListName pulumi.StringInput `pulumi:"contactListName"`
}

func (LookupContactListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactListArgs)(nil)).Elem()
}

type LookupContactListResultOutput struct{ *pulumi.OutputState }

func (LookupContactListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactListResult)(nil)).Elem()
}

func (o LookupContactListResultOutput) ToLookupContactListResultOutput() LookupContactListResultOutput {
	return o
}

func (o LookupContactListResultOutput) ToLookupContactListResultOutputWithContext(ctx context.Context) LookupContactListResultOutput {
	return o
}

func (o LookupContactListResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupContactListResult] {
	return pulumix.Output[LookupContactListResult]{
		OutputState: o.OutputState,
	}
}

// The description of the contact list.
func (o LookupContactListResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactListResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The tags (keys and values) associated with the contact list.
func (o LookupContactListResultOutput) Tags() ContactListTagArrayOutput {
	return o.ApplyT(func(v LookupContactListResult) []ContactListTag { return v.Tags }).(ContactListTagArrayOutput)
}

// The topics associated with the contact list.
func (o LookupContactListResultOutput) Topics() ContactListTopicArrayOutput {
	return o.ApplyT(func(v LookupContactListResult) []ContactListTopic { return v.Topics }).(ContactListTopicArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactListResultOutput{})
}
