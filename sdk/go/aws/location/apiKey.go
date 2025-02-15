// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::Location::APIKey Resource Type
type ApiKey struct {
	pulumi.CustomResourceState

	Arn          pulumi.StringOutput      `pulumi:"arn"`
	CreateTime   pulumi.StringOutput      `pulumi:"createTime"`
	Description  pulumi.StringPtrOutput   `pulumi:"description"`
	ExpireTime   pulumi.StringPtrOutput   `pulumi:"expireTime"`
	ForceDelete  pulumi.BoolPtrOutput     `pulumi:"forceDelete"`
	ForceUpdate  pulumi.BoolPtrOutput     `pulumi:"forceUpdate"`
	KeyArn       pulumi.StringOutput      `pulumi:"keyArn"`
	KeyName      pulumi.StringOutput      `pulumi:"keyName"`
	NoExpiry     pulumi.BoolPtrOutput     `pulumi:"noExpiry"`
	Restrictions ApiKeyRestrictionsOutput `pulumi:"restrictions"`
	// An array of key-value pairs to apply to this resource.
	Tags       ApiKeyTagArrayOutput `pulumi:"tags"`
	UpdateTime pulumi.StringOutput  `pulumi:"updateTime"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.Restrictions == nil {
		return nil, errors.New("invalid value for required argument 'Restrictions'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"keyName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiKey
	err := ctx.RegisterResource("aws-native:location:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("aws-native:location:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
}

type ApiKeyState struct {
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	Description  *string            `pulumi:"description"`
	ExpireTime   *string            `pulumi:"expireTime"`
	ForceDelete  *bool              `pulumi:"forceDelete"`
	ForceUpdate  *bool              `pulumi:"forceUpdate"`
	KeyName      string             `pulumi:"keyName"`
	NoExpiry     *bool              `pulumi:"noExpiry"`
	Restrictions ApiKeyRestrictions `pulumi:"restrictions"`
	// An array of key-value pairs to apply to this resource.
	Tags []ApiKeyTag `pulumi:"tags"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	Description  pulumi.StringPtrInput
	ExpireTime   pulumi.StringPtrInput
	ForceDelete  pulumi.BoolPtrInput
	ForceUpdate  pulumi.BoolPtrInput
	KeyName      pulumi.StringInput
	NoExpiry     pulumi.BoolPtrInput
	Restrictions ApiKeyRestrictionsInput
	// An array of key-value pairs to apply to this resource.
	Tags ApiKeyTagArrayInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

func (*ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (i *ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i *ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

func (i *ApiKey) ToOutput(ctx context.Context) pulumix.Output[*ApiKey] {
	return pulumix.Output[*ApiKey]{
		OutputState: i.ToApiKeyOutputWithContext(ctx).OutputState,
	}
}

type ApiKeyOutput struct{ *pulumi.OutputState }

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*ApiKey] {
	return pulumix.Output[*ApiKey]{
		OutputState: o.OutputState,
	}
}

func (o ApiKeyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ApiKeyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o ApiKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiKeyOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.ExpireTime }).(pulumi.StringPtrOutput)
}

func (o ApiKeyOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

func (o ApiKeyOutput) ForceUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.BoolPtrOutput { return v.ForceUpdate }).(pulumi.BoolPtrOutput)
}

func (o ApiKeyOutput) KeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.KeyArn }).(pulumi.StringOutput)
}

func (o ApiKeyOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

func (o ApiKeyOutput) NoExpiry() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.BoolPtrOutput { return v.NoExpiry }).(pulumi.BoolPtrOutput)
}

func (o ApiKeyOutput) Restrictions() ApiKeyRestrictionsOutput {
	return o.ApplyT(func(v *ApiKey) ApiKeyRestrictionsOutput { return v.Restrictions }).(ApiKeyRestrictionsOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ApiKeyOutput) Tags() ApiKeyTagArrayOutput {
	return o.ApplyT(func(v *ApiKey) ApiKeyTagArrayOutput { return v.Tags }).(ApiKeyTagArrayOutput)
}

func (o ApiKeyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyInput)(nil)).Elem(), &ApiKey{})
	pulumi.RegisterOutputType(ApiKeyOutput{})
}
