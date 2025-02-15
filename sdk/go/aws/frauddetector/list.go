// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package frauddetector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A resource schema for a List in Amazon Fraud Detector.
type List struct {
	pulumi.CustomResourceState

	// The list ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The time when the list was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The description of the list.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The elements in this list.
	Elements pulumi.StringArrayOutput `pulumi:"elements"`
	// The time when the list was last updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// The name of the list.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags associated with this list.
	Tags ListTagArrayOutput `pulumi:"tags"`
	// The variable type of the list.
	VariableType pulumi.StringPtrOutput `pulumi:"variableType"`
}

// NewList registers a new resource with the given unique name, arguments, and options.
func NewList(ctx *pulumi.Context,
	name string, args *ListArgs, opts ...pulumi.ResourceOption) (*List, error) {
	if args == nil {
		args = &ListArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource List
	err := ctx.RegisterResource("aws-native:frauddetector:List", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetList gets an existing List resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListState, opts ...pulumi.ResourceOption) (*List, error) {
	var resource List
	err := ctx.ReadResource("aws-native:frauddetector:List", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering List resources.
type listState struct {
}

type ListState struct {
}

func (ListState) ElementType() reflect.Type {
	return reflect.TypeOf((*listState)(nil)).Elem()
}

type listArgs struct {
	// The description of the list.
	Description *string `pulumi:"description"`
	// The elements in this list.
	Elements []string `pulumi:"elements"`
	// The name of the list.
	Name *string `pulumi:"name"`
	// Tags associated with this list.
	Tags []ListTag `pulumi:"tags"`
	// The variable type of the list.
	VariableType *string `pulumi:"variableType"`
}

// The set of arguments for constructing a List resource.
type ListArgs struct {
	// The description of the list.
	Description pulumi.StringPtrInput
	// The elements in this list.
	Elements pulumi.StringArrayInput
	// The name of the list.
	Name pulumi.StringPtrInput
	// Tags associated with this list.
	Tags ListTagArrayInput
	// The variable type of the list.
	VariableType pulumi.StringPtrInput
}

func (ListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listArgs)(nil)).Elem()
}

type ListInput interface {
	pulumi.Input

	ToListOutput() ListOutput
	ToListOutputWithContext(ctx context.Context) ListOutput
}

func (*List) ElementType() reflect.Type {
	return reflect.TypeOf((**List)(nil)).Elem()
}

func (i *List) ToListOutput() ListOutput {
	return i.ToListOutputWithContext(context.Background())
}

func (i *List) ToListOutputWithContext(ctx context.Context) ListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListOutput)
}

func (i *List) ToOutput(ctx context.Context) pulumix.Output[*List] {
	return pulumix.Output[*List]{
		OutputState: i.ToListOutputWithContext(ctx).OutputState,
	}
}

type ListOutput struct{ *pulumi.OutputState }

func (ListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**List)(nil)).Elem()
}

func (o ListOutput) ToListOutput() ListOutput {
	return o
}

func (o ListOutput) ToListOutputWithContext(ctx context.Context) ListOutput {
	return o
}

func (o ListOutput) ToOutput(ctx context.Context) pulumix.Output[*List] {
	return pulumix.Output[*List]{
		OutputState: o.OutputState,
	}
}

// The list ARN.
func (o ListOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The time when the list was created.
func (o ListOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The description of the list.
func (o ListOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *List) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The elements in this list.
func (o ListOutput) Elements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *List) pulumi.StringArrayOutput { return v.Elements }).(pulumi.StringArrayOutput)
}

// The time when the list was last updated.
func (o ListOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// The name of the list.
func (o ListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags associated with this list.
func (o ListOutput) Tags() ListTagArrayOutput {
	return o.ApplyT(func(v *List) ListTagArrayOutput { return v.Tags }).(ListTagArrayOutput)
}

// The variable type of the list.
func (o ListOutput) VariableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *List) pulumi.StringPtrOutput { return v.VariableType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListInput)(nil)).Elem(), &List{})
	pulumi.RegisterOutputType(ListOutput{})
}
