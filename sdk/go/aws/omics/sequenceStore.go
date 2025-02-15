// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::Omics::SequenceStore Resource Type
type SequenceStore struct {
	pulumi.CustomResourceState

	// The store's ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// When the store was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// A description for the store.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An S3 URI representing the bucket and folder to store failed read set uploads.
	FallbackLocation pulumi.StringPtrOutput `pulumi:"fallbackLocation"`
	// A name for the store.
	Name            pulumi.StringOutput             `pulumi:"name"`
	SequenceStoreId pulumi.StringOutput             `pulumi:"sequenceStoreId"`
	SseConfig       SequenceStoreSseConfigPtrOutput `pulumi:"sseConfig"`
	Tags            SequenceStoreTagMapPtrOutput    `pulumi:"tags"`
}

// NewSequenceStore registers a new resource with the given unique name, arguments, and options.
func NewSequenceStore(ctx *pulumi.Context,
	name string, args *SequenceStoreArgs, opts ...pulumi.ResourceOption) (*SequenceStore, error) {
	if args == nil {
		args = &SequenceStoreArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"fallbackLocation",
		"name",
		"sseConfig",
		"tags",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SequenceStore
	err := ctx.RegisterResource("aws-native:omics:SequenceStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSequenceStore gets an existing SequenceStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSequenceStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SequenceStoreState, opts ...pulumi.ResourceOption) (*SequenceStore, error) {
	var resource SequenceStore
	err := ctx.ReadResource("aws-native:omics:SequenceStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SequenceStore resources.
type sequenceStoreState struct {
}

type SequenceStoreState struct {
}

func (SequenceStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*sequenceStoreState)(nil)).Elem()
}

type sequenceStoreArgs struct {
	// A description for the store.
	Description *string `pulumi:"description"`
	// An S3 URI representing the bucket and folder to store failed read set uploads.
	FallbackLocation *string `pulumi:"fallbackLocation"`
	// A name for the store.
	Name      *string                 `pulumi:"name"`
	SseConfig *SequenceStoreSseConfig `pulumi:"sseConfig"`
	Tags      *SequenceStoreTagMap    `pulumi:"tags"`
}

// The set of arguments for constructing a SequenceStore resource.
type SequenceStoreArgs struct {
	// A description for the store.
	Description pulumi.StringPtrInput
	// An S3 URI representing the bucket and folder to store failed read set uploads.
	FallbackLocation pulumi.StringPtrInput
	// A name for the store.
	Name      pulumi.StringPtrInput
	SseConfig SequenceStoreSseConfigPtrInput
	Tags      SequenceStoreTagMapPtrInput
}

func (SequenceStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sequenceStoreArgs)(nil)).Elem()
}

type SequenceStoreInput interface {
	pulumi.Input

	ToSequenceStoreOutput() SequenceStoreOutput
	ToSequenceStoreOutputWithContext(ctx context.Context) SequenceStoreOutput
}

func (*SequenceStore) ElementType() reflect.Type {
	return reflect.TypeOf((**SequenceStore)(nil)).Elem()
}

func (i *SequenceStore) ToSequenceStoreOutput() SequenceStoreOutput {
	return i.ToSequenceStoreOutputWithContext(context.Background())
}

func (i *SequenceStore) ToSequenceStoreOutputWithContext(ctx context.Context) SequenceStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SequenceStoreOutput)
}

func (i *SequenceStore) ToOutput(ctx context.Context) pulumix.Output[*SequenceStore] {
	return pulumix.Output[*SequenceStore]{
		OutputState: i.ToSequenceStoreOutputWithContext(ctx).OutputState,
	}
}

type SequenceStoreOutput struct{ *pulumi.OutputState }

func (SequenceStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SequenceStore)(nil)).Elem()
}

func (o SequenceStoreOutput) ToSequenceStoreOutput() SequenceStoreOutput {
	return o
}

func (o SequenceStoreOutput) ToSequenceStoreOutputWithContext(ctx context.Context) SequenceStoreOutput {
	return o
}

func (o SequenceStoreOutput) ToOutput(ctx context.Context) pulumix.Output[*SequenceStore] {
	return pulumix.Output[*SequenceStore]{
		OutputState: o.OutputState,
	}
}

// The store's ARN.
func (o SequenceStoreOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SequenceStore) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// When the store was created.
func (o SequenceStoreOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SequenceStore) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// A description for the store.
func (o SequenceStoreOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SequenceStore) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An S3 URI representing the bucket and folder to store failed read set uploads.
func (o SequenceStoreOutput) FallbackLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SequenceStore) pulumi.StringPtrOutput { return v.FallbackLocation }).(pulumi.StringPtrOutput)
}

// A name for the store.
func (o SequenceStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SequenceStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SequenceStoreOutput) SequenceStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *SequenceStore) pulumi.StringOutput { return v.SequenceStoreId }).(pulumi.StringOutput)
}

func (o SequenceStoreOutput) SseConfig() SequenceStoreSseConfigPtrOutput {
	return o.ApplyT(func(v *SequenceStore) SequenceStoreSseConfigPtrOutput { return v.SseConfig }).(SequenceStoreSseConfigPtrOutput)
}

func (o SequenceStoreOutput) Tags() SequenceStoreTagMapPtrOutput {
	return o.ApplyT(func(v *SequenceStore) SequenceStoreTagMapPtrOutput { return v.Tags }).(SequenceStoreTagMapPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SequenceStoreInput)(nil)).Elem(), &SequenceStore{})
	pulumi.RegisterOutputType(SequenceStoreOutput{})
}
