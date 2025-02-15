// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wisdom

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::Wisdom::KnowledgeBase Resource Type
type KnowledgeBase struct {
	pulumi.CustomResourceState

	Description                       pulumi.StringPtrOutput                                  `pulumi:"description"`
	KnowledgeBaseArn                  pulumi.StringOutput                                     `pulumi:"knowledgeBaseArn"`
	KnowledgeBaseId                   pulumi.StringOutput                                     `pulumi:"knowledgeBaseId"`
	KnowledgeBaseType                 KnowledgeBaseTypeOutput                                 `pulumi:"knowledgeBaseType"`
	Name                              pulumi.StringOutput                                     `pulumi:"name"`
	RenderingConfiguration            KnowledgeBaseRenderingConfigurationPtrOutput            `pulumi:"renderingConfiguration"`
	ServerSideEncryptionConfiguration KnowledgeBaseServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	SourceConfiguration               KnowledgeBaseSourceConfigurationPtrOutput               `pulumi:"sourceConfiguration"`
	Tags                              KnowledgeBaseTagArrayOutput                             `pulumi:"tags"`
}

// NewKnowledgeBase registers a new resource with the given unique name, arguments, and options.
func NewKnowledgeBase(ctx *pulumi.Context,
	name string, args *KnowledgeBaseArgs, opts ...pulumi.ResourceOption) (*KnowledgeBase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KnowledgeBaseType == nil {
		return nil, errors.New("invalid value for required argument 'KnowledgeBaseType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"knowledgeBaseType",
		"name",
		"serverSideEncryptionConfiguration",
		"sourceConfiguration",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KnowledgeBase
	err := ctx.RegisterResource("aws-native:wisdom:KnowledgeBase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKnowledgeBase gets an existing KnowledgeBase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKnowledgeBase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KnowledgeBaseState, opts ...pulumi.ResourceOption) (*KnowledgeBase, error) {
	var resource KnowledgeBase
	err := ctx.ReadResource("aws-native:wisdom:KnowledgeBase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KnowledgeBase resources.
type knowledgeBaseState struct {
}

type KnowledgeBaseState struct {
}

func (KnowledgeBaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*knowledgeBaseState)(nil)).Elem()
}

type knowledgeBaseArgs struct {
	Description                       *string                                         `pulumi:"description"`
	KnowledgeBaseType                 KnowledgeBaseType                               `pulumi:"knowledgeBaseType"`
	Name                              *string                                         `pulumi:"name"`
	RenderingConfiguration            *KnowledgeBaseRenderingConfiguration            `pulumi:"renderingConfiguration"`
	ServerSideEncryptionConfiguration *KnowledgeBaseServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	SourceConfiguration               *KnowledgeBaseSourceConfiguration               `pulumi:"sourceConfiguration"`
	Tags                              []KnowledgeBaseTag                              `pulumi:"tags"`
}

// The set of arguments for constructing a KnowledgeBase resource.
type KnowledgeBaseArgs struct {
	Description                       pulumi.StringPtrInput
	KnowledgeBaseType                 KnowledgeBaseTypeInput
	Name                              pulumi.StringPtrInput
	RenderingConfiguration            KnowledgeBaseRenderingConfigurationPtrInput
	ServerSideEncryptionConfiguration KnowledgeBaseServerSideEncryptionConfigurationPtrInput
	SourceConfiguration               KnowledgeBaseSourceConfigurationPtrInput
	Tags                              KnowledgeBaseTagArrayInput
}

func (KnowledgeBaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*knowledgeBaseArgs)(nil)).Elem()
}

type KnowledgeBaseInput interface {
	pulumi.Input

	ToKnowledgeBaseOutput() KnowledgeBaseOutput
	ToKnowledgeBaseOutputWithContext(ctx context.Context) KnowledgeBaseOutput
}

func (*KnowledgeBase) ElementType() reflect.Type {
	return reflect.TypeOf((**KnowledgeBase)(nil)).Elem()
}

func (i *KnowledgeBase) ToKnowledgeBaseOutput() KnowledgeBaseOutput {
	return i.ToKnowledgeBaseOutputWithContext(context.Background())
}

func (i *KnowledgeBase) ToKnowledgeBaseOutputWithContext(ctx context.Context) KnowledgeBaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KnowledgeBaseOutput)
}

func (i *KnowledgeBase) ToOutput(ctx context.Context) pulumix.Output[*KnowledgeBase] {
	return pulumix.Output[*KnowledgeBase]{
		OutputState: i.ToKnowledgeBaseOutputWithContext(ctx).OutputState,
	}
}

type KnowledgeBaseOutput struct{ *pulumi.OutputState }

func (KnowledgeBaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KnowledgeBase)(nil)).Elem()
}

func (o KnowledgeBaseOutput) ToKnowledgeBaseOutput() KnowledgeBaseOutput {
	return o
}

func (o KnowledgeBaseOutput) ToKnowledgeBaseOutputWithContext(ctx context.Context) KnowledgeBaseOutput {
	return o
}

func (o KnowledgeBaseOutput) ToOutput(ctx context.Context) pulumix.Output[*KnowledgeBase] {
	return pulumix.Output[*KnowledgeBase]{
		OutputState: o.OutputState,
	}
}

func (o KnowledgeBaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KnowledgeBase) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o KnowledgeBaseOutput) KnowledgeBaseArn() pulumi.StringOutput {
	return o.ApplyT(func(v *KnowledgeBase) pulumi.StringOutput { return v.KnowledgeBaseArn }).(pulumi.StringOutput)
}

func (o KnowledgeBaseOutput) KnowledgeBaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *KnowledgeBase) pulumi.StringOutput { return v.KnowledgeBaseId }).(pulumi.StringOutput)
}

func (o KnowledgeBaseOutput) KnowledgeBaseType() KnowledgeBaseTypeOutput {
	return o.ApplyT(func(v *KnowledgeBase) KnowledgeBaseTypeOutput { return v.KnowledgeBaseType }).(KnowledgeBaseTypeOutput)
}

func (o KnowledgeBaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KnowledgeBase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KnowledgeBaseOutput) RenderingConfiguration() KnowledgeBaseRenderingConfigurationPtrOutput {
	return o.ApplyT(func(v *KnowledgeBase) KnowledgeBaseRenderingConfigurationPtrOutput { return v.RenderingConfiguration }).(KnowledgeBaseRenderingConfigurationPtrOutput)
}

func (o KnowledgeBaseOutput) ServerSideEncryptionConfiguration() KnowledgeBaseServerSideEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v *KnowledgeBase) KnowledgeBaseServerSideEncryptionConfigurationPtrOutput {
		return v.ServerSideEncryptionConfiguration
	}).(KnowledgeBaseServerSideEncryptionConfigurationPtrOutput)
}

func (o KnowledgeBaseOutput) SourceConfiguration() KnowledgeBaseSourceConfigurationPtrOutput {
	return o.ApplyT(func(v *KnowledgeBase) KnowledgeBaseSourceConfigurationPtrOutput { return v.SourceConfiguration }).(KnowledgeBaseSourceConfigurationPtrOutput)
}

func (o KnowledgeBaseOutput) Tags() KnowledgeBaseTagArrayOutput {
	return o.ApplyT(func(v *KnowledgeBase) KnowledgeBaseTagArrayOutput { return v.Tags }).(KnowledgeBaseTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KnowledgeBaseInput)(nil)).Elem(), &KnowledgeBase{})
	pulumi.RegisterOutputType(KnowledgeBaseOutput{})
}
