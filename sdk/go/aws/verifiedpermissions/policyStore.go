// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedpermissions

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::VerifiedPermissions::PolicyStore Resource Type
type PolicyStore struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput                  `pulumi:"arn"`
	PolicyStoreId      pulumi.StringOutput                  `pulumi:"policyStoreId"`
	Schema             PolicyStoreSchemaDefinitionPtrOutput `pulumi:"schema"`
	ValidationSettings PolicyStoreValidationSettingsOutput  `pulumi:"validationSettings"`
}

// NewPolicyStore registers a new resource with the given unique name, arguments, and options.
func NewPolicyStore(ctx *pulumi.Context,
	name string, args *PolicyStoreArgs, opts ...pulumi.ResourceOption) (*PolicyStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ValidationSettings == nil {
		return nil, errors.New("invalid value for required argument 'ValidationSettings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyStore
	err := ctx.RegisterResource("aws-native:verifiedpermissions:PolicyStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyStore gets an existing PolicyStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyStoreState, opts ...pulumi.ResourceOption) (*PolicyStore, error) {
	var resource PolicyStore
	err := ctx.ReadResource("aws-native:verifiedpermissions:PolicyStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyStore resources.
type policyStoreState struct {
}

type PolicyStoreState struct {
}

func (PolicyStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyStoreState)(nil)).Elem()
}

type policyStoreArgs struct {
	Schema             *PolicyStoreSchemaDefinition  `pulumi:"schema"`
	ValidationSettings PolicyStoreValidationSettings `pulumi:"validationSettings"`
}

// The set of arguments for constructing a PolicyStore resource.
type PolicyStoreArgs struct {
	Schema             PolicyStoreSchemaDefinitionPtrInput
	ValidationSettings PolicyStoreValidationSettingsInput
}

func (PolicyStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyStoreArgs)(nil)).Elem()
}

type PolicyStoreInput interface {
	pulumi.Input

	ToPolicyStoreOutput() PolicyStoreOutput
	ToPolicyStoreOutputWithContext(ctx context.Context) PolicyStoreOutput
}

func (*PolicyStore) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStore)(nil)).Elem()
}

func (i *PolicyStore) ToPolicyStoreOutput() PolicyStoreOutput {
	return i.ToPolicyStoreOutputWithContext(context.Background())
}

func (i *PolicyStore) ToPolicyStoreOutputWithContext(ctx context.Context) PolicyStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyStoreOutput)
}

func (i *PolicyStore) ToOutput(ctx context.Context) pulumix.Output[*PolicyStore] {
	return pulumix.Output[*PolicyStore]{
		OutputState: i.ToPolicyStoreOutputWithContext(ctx).OutputState,
	}
}

type PolicyStoreOutput struct{ *pulumi.OutputState }

func (PolicyStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStore)(nil)).Elem()
}

func (o PolicyStoreOutput) ToPolicyStoreOutput() PolicyStoreOutput {
	return o
}

func (o PolicyStoreOutput) ToPolicyStoreOutputWithContext(ctx context.Context) PolicyStoreOutput {
	return o
}

func (o PolicyStoreOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyStore] {
	return pulumix.Output[*PolicyStore]{
		OutputState: o.OutputState,
	}
}

func (o PolicyStoreOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyStore) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o PolicyStoreOutput) PolicyStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyStore) pulumi.StringOutput { return v.PolicyStoreId }).(pulumi.StringOutput)
}

func (o PolicyStoreOutput) Schema() PolicyStoreSchemaDefinitionPtrOutput {
	return o.ApplyT(func(v *PolicyStore) PolicyStoreSchemaDefinitionPtrOutput { return v.Schema }).(PolicyStoreSchemaDefinitionPtrOutput)
}

func (o PolicyStoreOutput) ValidationSettings() PolicyStoreValidationSettingsOutput {
	return o.ApplyT(func(v *PolicyStore) PolicyStoreValidationSettingsOutput { return v.ValidationSettings }).(PolicyStoreValidationSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStoreInput)(nil)).Elem(), &PolicyStore{})
	pulumi.RegisterOutputType(PolicyStoreOutput{})
}
