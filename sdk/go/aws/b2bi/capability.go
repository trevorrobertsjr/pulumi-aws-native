// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package b2bi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::B2BI::Capability Resource Type
type Capability struct {
	pulumi.CustomResourceState

	CapabilityArn         pulumi.StringOutput                     `pulumi:"capabilityArn"`
	CapabilityId          pulumi.StringOutput                     `pulumi:"capabilityId"`
	Configuration         CapabilityConfigurationPropertiesOutput `pulumi:"configuration"`
	CreatedAt             pulumi.StringOutput                     `pulumi:"createdAt"`
	InstructionsDocuments CapabilityS3LocationArrayOutput         `pulumi:"instructionsDocuments"`
	ModifiedAt            pulumi.StringOutput                     `pulumi:"modifiedAt"`
	Name                  pulumi.StringOutput                     `pulumi:"name"`
	Tags                  CapabilityTagArrayOutput                `pulumi:"tags"`
	Type                  CapabilityTypeOutput                    `pulumi:"type"`
}

// NewCapability registers a new resource with the given unique name, arguments, and options.
func NewCapability(ctx *pulumi.Context,
	name string, args *CapabilityArgs, opts ...pulumi.ResourceOption) (*Capability, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Capability
	err := ctx.RegisterResource("aws-native:b2bi:Capability", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapability gets an existing Capability resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapability(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapabilityState, opts ...pulumi.ResourceOption) (*Capability, error) {
	var resource Capability
	err := ctx.ReadResource("aws-native:b2bi:Capability", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Capability resources.
type capabilityState struct {
}

type CapabilityState struct {
}

func (CapabilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*capabilityState)(nil)).Elem()
}

type capabilityArgs struct {
	Configuration         CapabilityConfigurationProperties `pulumi:"configuration"`
	InstructionsDocuments []CapabilityS3Location            `pulumi:"instructionsDocuments"`
	Name                  *string                           `pulumi:"name"`
	Tags                  []CapabilityTag                   `pulumi:"tags"`
	Type                  CapabilityType                    `pulumi:"type"`
}

// The set of arguments for constructing a Capability resource.
type CapabilityArgs struct {
	Configuration         CapabilityConfigurationPropertiesInput
	InstructionsDocuments CapabilityS3LocationArrayInput
	Name                  pulumi.StringPtrInput
	Tags                  CapabilityTagArrayInput
	Type                  CapabilityTypeInput
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capabilityArgs)(nil)).Elem()
}

type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput
}

func (*Capability) ElementType() reflect.Type {
	return reflect.TypeOf((**Capability)(nil)).Elem()
}

func (i *Capability) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i *Capability) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}

func (i *Capability) ToOutput(ctx context.Context) pulumix.Output[*Capability] {
	return pulumix.Output[*Capability]{
		OutputState: i.ToCapabilityOutputWithContext(ctx).OutputState,
	}
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToOutput(ctx context.Context) pulumix.Output[*Capability] {
	return pulumix.Output[*Capability]{
		OutputState: o.OutputState,
	}
}

func (o CapabilityOutput) CapabilityArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.CapabilityArn }).(pulumi.StringOutput)
}

func (o CapabilityOutput) CapabilityId() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.CapabilityId }).(pulumi.StringOutput)
}

func (o CapabilityOutput) Configuration() CapabilityConfigurationPropertiesOutput {
	return o.ApplyT(func(v *Capability) CapabilityConfigurationPropertiesOutput { return v.Configuration }).(CapabilityConfigurationPropertiesOutput)
}

func (o CapabilityOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o CapabilityOutput) InstructionsDocuments() CapabilityS3LocationArrayOutput {
	return o.ApplyT(func(v *Capability) CapabilityS3LocationArrayOutput { return v.InstructionsDocuments }).(CapabilityS3LocationArrayOutput)
}

func (o CapabilityOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o CapabilityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CapabilityOutput) Tags() CapabilityTagArrayOutput {
	return o.ApplyT(func(v *Capability) CapabilityTagArrayOutput { return v.Tags }).(CapabilityTagArrayOutput)
}

func (o CapabilityOutput) Type() CapabilityTypeOutput {
	return o.ApplyT(func(v *Capability) CapabilityTypeOutput { return v.Type }).(CapabilityTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapabilityInput)(nil)).Elem(), &Capability{})
	pulumi.RegisterOutputType(CapabilityOutput{})
}
