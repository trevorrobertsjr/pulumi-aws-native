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

// Resource Type definition for AWS::Connect::UserHierarchyGroup
func LookupUserHierarchyGroup(ctx *pulumi.Context, args *LookupUserHierarchyGroupArgs, opts ...pulumi.InvokeOption) (*LookupUserHierarchyGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserHierarchyGroupResult
	err := ctx.Invoke("aws-native:connect:getUserHierarchyGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserHierarchyGroupArgs struct {
	// The Amazon Resource Name (ARN) for the user hierarchy group.
	UserHierarchyGroupArn string `pulumi:"userHierarchyGroupArn"`
}

type LookupUserHierarchyGroupResult struct {
	// The identifier of the Amazon Connect instance.
	InstanceArn *string `pulumi:"instanceArn"`
	// The name of the user hierarchy group.
	Name *string `pulumi:"name"`
	// One or more tags.
	Tags []UserHierarchyGroupTag `pulumi:"tags"`
	// The Amazon Resource Name (ARN) for the user hierarchy group.
	UserHierarchyGroupArn *string `pulumi:"userHierarchyGroupArn"`
}

func LookupUserHierarchyGroupOutput(ctx *pulumi.Context, args LookupUserHierarchyGroupOutputArgs, opts ...pulumi.InvokeOption) LookupUserHierarchyGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserHierarchyGroupResult, error) {
			args := v.(LookupUserHierarchyGroupArgs)
			r, err := LookupUserHierarchyGroup(ctx, &args, opts...)
			var s LookupUserHierarchyGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserHierarchyGroupResultOutput)
}

type LookupUserHierarchyGroupOutputArgs struct {
	// The Amazon Resource Name (ARN) for the user hierarchy group.
	UserHierarchyGroupArn pulumi.StringInput `pulumi:"userHierarchyGroupArn"`
}

func (LookupUserHierarchyGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserHierarchyGroupArgs)(nil)).Elem()
}

type LookupUserHierarchyGroupResultOutput struct{ *pulumi.OutputState }

func (LookupUserHierarchyGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserHierarchyGroupResult)(nil)).Elem()
}

func (o LookupUserHierarchyGroupResultOutput) ToLookupUserHierarchyGroupResultOutput() LookupUserHierarchyGroupResultOutput {
	return o
}

func (o LookupUserHierarchyGroupResultOutput) ToLookupUserHierarchyGroupResultOutputWithContext(ctx context.Context) LookupUserHierarchyGroupResultOutput {
	return o
}

func (o LookupUserHierarchyGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserHierarchyGroupResult] {
	return pulumix.Output[LookupUserHierarchyGroupResult]{
		OutputState: o.OutputState,
	}
}

// The identifier of the Amazon Connect instance.
func (o LookupUserHierarchyGroupResultOutput) InstanceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserHierarchyGroupResult) *string { return v.InstanceArn }).(pulumi.StringPtrOutput)
}

// The name of the user hierarchy group.
func (o LookupUserHierarchyGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserHierarchyGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// One or more tags.
func (o LookupUserHierarchyGroupResultOutput) Tags() UserHierarchyGroupTagArrayOutput {
	return o.ApplyT(func(v LookupUserHierarchyGroupResult) []UserHierarchyGroupTag { return v.Tags }).(UserHierarchyGroupTagArrayOutput)
}

// The Amazon Resource Name (ARN) for the user hierarchy group.
func (o LookupUserHierarchyGroupResultOutput) UserHierarchyGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserHierarchyGroupResult) *string { return v.UserHierarchyGroupArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserHierarchyGroupResultOutput{})
}
