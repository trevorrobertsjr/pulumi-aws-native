// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedblockchain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type AccessorNetworkAccessorType string

const (
	AccessorNetworkAccessorTypeEthereumGoerli           = AccessorNetworkAccessorType("ETHEREUM_GOERLI")
	AccessorNetworkAccessorTypeEthereumMainnet          = AccessorNetworkAccessorType("ETHEREUM_MAINNET")
	AccessorNetworkAccessorTypeEthereumMainnetAndGoerli = AccessorNetworkAccessorType("ETHEREUM_MAINNET_AND_GOERLI")
	AccessorNetworkAccessorTypePolygonMainnet           = AccessorNetworkAccessorType("POLYGON_MAINNET")
	AccessorNetworkAccessorTypePolygonMumbai            = AccessorNetworkAccessorType("POLYGON_MUMBAI")
)

func (AccessorNetworkAccessorType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessorNetworkAccessorType)(nil)).Elem()
}

func (e AccessorNetworkAccessorType) ToAccessorNetworkAccessorTypeOutput() AccessorNetworkAccessorTypeOutput {
	return pulumi.ToOutput(e).(AccessorNetworkAccessorTypeOutput)
}

func (e AccessorNetworkAccessorType) ToAccessorNetworkAccessorTypeOutputWithContext(ctx context.Context) AccessorNetworkAccessorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessorNetworkAccessorTypeOutput)
}

func (e AccessorNetworkAccessorType) ToAccessorNetworkAccessorTypePtrOutput() AccessorNetworkAccessorTypePtrOutput {
	return e.ToAccessorNetworkAccessorTypePtrOutputWithContext(context.Background())
}

func (e AccessorNetworkAccessorType) ToAccessorNetworkAccessorTypePtrOutputWithContext(ctx context.Context) AccessorNetworkAccessorTypePtrOutput {
	return AccessorNetworkAccessorType(e).ToAccessorNetworkAccessorTypeOutputWithContext(ctx).ToAccessorNetworkAccessorTypePtrOutputWithContext(ctx)
}

func (e AccessorNetworkAccessorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessorNetworkAccessorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessorNetworkAccessorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessorNetworkAccessorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessorNetworkAccessorTypeOutput struct{ *pulumi.OutputState }

func (AccessorNetworkAccessorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessorNetworkAccessorType)(nil)).Elem()
}

func (o AccessorNetworkAccessorTypeOutput) ToAccessorNetworkAccessorTypeOutput() AccessorNetworkAccessorTypeOutput {
	return o
}

func (o AccessorNetworkAccessorTypeOutput) ToAccessorNetworkAccessorTypeOutputWithContext(ctx context.Context) AccessorNetworkAccessorTypeOutput {
	return o
}

func (o AccessorNetworkAccessorTypeOutput) ToAccessorNetworkAccessorTypePtrOutput() AccessorNetworkAccessorTypePtrOutput {
	return o.ToAccessorNetworkAccessorTypePtrOutputWithContext(context.Background())
}

func (o AccessorNetworkAccessorTypeOutput) ToAccessorNetworkAccessorTypePtrOutputWithContext(ctx context.Context) AccessorNetworkAccessorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessorNetworkAccessorType) *AccessorNetworkAccessorType {
		return &v
	}).(AccessorNetworkAccessorTypePtrOutput)
}

func (o AccessorNetworkAccessorTypeOutput) ToOutput(ctx context.Context) pulumix.Output[AccessorNetworkAccessorType] {
	return pulumix.Output[AccessorNetworkAccessorType]{
		OutputState: o.OutputState,
	}
}

func (o AccessorNetworkAccessorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessorNetworkAccessorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessorNetworkAccessorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessorNetworkAccessorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessorNetworkAccessorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessorNetworkAccessorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessorNetworkAccessorTypePtrOutput struct{ *pulumi.OutputState }

func (AccessorNetworkAccessorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessorNetworkAccessorType)(nil)).Elem()
}

func (o AccessorNetworkAccessorTypePtrOutput) ToAccessorNetworkAccessorTypePtrOutput() AccessorNetworkAccessorTypePtrOutput {
	return o
}

func (o AccessorNetworkAccessorTypePtrOutput) ToAccessorNetworkAccessorTypePtrOutputWithContext(ctx context.Context) AccessorNetworkAccessorTypePtrOutput {
	return o
}

func (o AccessorNetworkAccessorTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AccessorNetworkAccessorType] {
	return pulumix.Output[*AccessorNetworkAccessorType]{
		OutputState: o.OutputState,
	}
}

func (o AccessorNetworkAccessorTypePtrOutput) Elem() AccessorNetworkAccessorTypeOutput {
	return o.ApplyT(func(v *AccessorNetworkAccessorType) AccessorNetworkAccessorType {
		if v != nil {
			return *v
		}
		var ret AccessorNetworkAccessorType
		return ret
	}).(AccessorNetworkAccessorTypeOutput)
}

func (o AccessorNetworkAccessorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessorNetworkAccessorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessorNetworkAccessorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessorNetworkAccessorTypeInput is an input type that accepts AccessorNetworkAccessorTypeArgs and AccessorNetworkAccessorTypeOutput values.
// You can construct a concrete instance of `AccessorNetworkAccessorTypeInput` via:
//
//	AccessorNetworkAccessorTypeArgs{...}
type AccessorNetworkAccessorTypeInput interface {
	pulumi.Input

	ToAccessorNetworkAccessorTypeOutput() AccessorNetworkAccessorTypeOutput
	ToAccessorNetworkAccessorTypeOutputWithContext(context.Context) AccessorNetworkAccessorTypeOutput
}

var accessorNetworkAccessorTypePtrType = reflect.TypeOf((**AccessorNetworkAccessorType)(nil)).Elem()

type AccessorNetworkAccessorTypePtrInput interface {
	pulumi.Input

	ToAccessorNetworkAccessorTypePtrOutput() AccessorNetworkAccessorTypePtrOutput
	ToAccessorNetworkAccessorTypePtrOutputWithContext(context.Context) AccessorNetworkAccessorTypePtrOutput
}

type accessorNetworkAccessorTypePtr string

func AccessorNetworkAccessorTypePtr(v string) AccessorNetworkAccessorTypePtrInput {
	return (*accessorNetworkAccessorTypePtr)(&v)
}

func (*accessorNetworkAccessorTypePtr) ElementType() reflect.Type {
	return accessorNetworkAccessorTypePtrType
}

func (in *accessorNetworkAccessorTypePtr) ToAccessorNetworkAccessorTypePtrOutput() AccessorNetworkAccessorTypePtrOutput {
	return pulumi.ToOutput(in).(AccessorNetworkAccessorTypePtrOutput)
}

func (in *accessorNetworkAccessorTypePtr) ToAccessorNetworkAccessorTypePtrOutputWithContext(ctx context.Context) AccessorNetworkAccessorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessorNetworkAccessorTypePtrOutput)
}

func (in *accessorNetworkAccessorTypePtr) ToOutput(ctx context.Context) pulumix.Output[*AccessorNetworkAccessorType] {
	return pulumix.Output[*AccessorNetworkAccessorType]{
		OutputState: in.ToAccessorNetworkAccessorTypePtrOutputWithContext(ctx).OutputState,
	}
}

type AccessorStatus string

const (
	AccessorStatusAvailable       = AccessorStatus("AVAILABLE")
	AccessorStatusPendingDeletion = AccessorStatus("PENDING_DELETION")
	AccessorStatusDeleted         = AccessorStatus("DELETED")
)

type AccessorStatusOutput struct{ *pulumi.OutputState }

func (AccessorStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessorStatus)(nil)).Elem()
}

func (o AccessorStatusOutput) ToAccessorStatusOutput() AccessorStatusOutput {
	return o
}

func (o AccessorStatusOutput) ToAccessorStatusOutputWithContext(ctx context.Context) AccessorStatusOutput {
	return o
}

func (o AccessorStatusOutput) ToAccessorStatusPtrOutput() AccessorStatusPtrOutput {
	return o.ToAccessorStatusPtrOutputWithContext(context.Background())
}

func (o AccessorStatusOutput) ToAccessorStatusPtrOutputWithContext(ctx context.Context) AccessorStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessorStatus) *AccessorStatus {
		return &v
	}).(AccessorStatusPtrOutput)
}

func (o AccessorStatusOutput) ToOutput(ctx context.Context) pulumix.Output[AccessorStatus] {
	return pulumix.Output[AccessorStatus]{
		OutputState: o.OutputState,
	}
}

func (o AccessorStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessorStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessorStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessorStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessorStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessorStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessorStatusPtrOutput struct{ *pulumi.OutputState }

func (AccessorStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessorStatus)(nil)).Elem()
}

func (o AccessorStatusPtrOutput) ToAccessorStatusPtrOutput() AccessorStatusPtrOutput {
	return o
}

func (o AccessorStatusPtrOutput) ToAccessorStatusPtrOutputWithContext(ctx context.Context) AccessorStatusPtrOutput {
	return o
}

func (o AccessorStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AccessorStatus] {
	return pulumix.Output[*AccessorStatus]{
		OutputState: o.OutputState,
	}
}

func (o AccessorStatusPtrOutput) Elem() AccessorStatusOutput {
	return o.ApplyT(func(v *AccessorStatus) AccessorStatus {
		if v != nil {
			return *v
		}
		var ret AccessorStatus
		return ret
	}).(AccessorStatusOutput)
}

func (o AccessorStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessorStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessorStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessorType string

const (
	AccessorTypeBillingToken = AccessorType("BILLING_TOKEN")
)

func (AccessorType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessorType)(nil)).Elem()
}

func (e AccessorType) ToAccessorTypeOutput() AccessorTypeOutput {
	return pulumi.ToOutput(e).(AccessorTypeOutput)
}

func (e AccessorType) ToAccessorTypeOutputWithContext(ctx context.Context) AccessorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessorTypeOutput)
}

func (e AccessorType) ToAccessorTypePtrOutput() AccessorTypePtrOutput {
	return e.ToAccessorTypePtrOutputWithContext(context.Background())
}

func (e AccessorType) ToAccessorTypePtrOutputWithContext(ctx context.Context) AccessorTypePtrOutput {
	return AccessorType(e).ToAccessorTypeOutputWithContext(ctx).ToAccessorTypePtrOutputWithContext(ctx)
}

func (e AccessorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessorTypeOutput struct{ *pulumi.OutputState }

func (AccessorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessorType)(nil)).Elem()
}

func (o AccessorTypeOutput) ToAccessorTypeOutput() AccessorTypeOutput {
	return o
}

func (o AccessorTypeOutput) ToAccessorTypeOutputWithContext(ctx context.Context) AccessorTypeOutput {
	return o
}

func (o AccessorTypeOutput) ToAccessorTypePtrOutput() AccessorTypePtrOutput {
	return o.ToAccessorTypePtrOutputWithContext(context.Background())
}

func (o AccessorTypeOutput) ToAccessorTypePtrOutputWithContext(ctx context.Context) AccessorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessorType) *AccessorType {
		return &v
	}).(AccessorTypePtrOutput)
}

func (o AccessorTypeOutput) ToOutput(ctx context.Context) pulumix.Output[AccessorType] {
	return pulumix.Output[AccessorType]{
		OutputState: o.OutputState,
	}
}

func (o AccessorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessorTypePtrOutput struct{ *pulumi.OutputState }

func (AccessorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessorType)(nil)).Elem()
}

func (o AccessorTypePtrOutput) ToAccessorTypePtrOutput() AccessorTypePtrOutput {
	return o
}

func (o AccessorTypePtrOutput) ToAccessorTypePtrOutputWithContext(ctx context.Context) AccessorTypePtrOutput {
	return o
}

func (o AccessorTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AccessorType] {
	return pulumix.Output[*AccessorType]{
		OutputState: o.OutputState,
	}
}

func (o AccessorTypePtrOutput) Elem() AccessorTypeOutput {
	return o.ApplyT(func(v *AccessorType) AccessorType {
		if v != nil {
			return *v
		}
		var ret AccessorType
		return ret
	}).(AccessorTypeOutput)
}

func (o AccessorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessorTypeInput is an input type that accepts AccessorTypeArgs and AccessorTypeOutput values.
// You can construct a concrete instance of `AccessorTypeInput` via:
//
//	AccessorTypeArgs{...}
type AccessorTypeInput interface {
	pulumi.Input

	ToAccessorTypeOutput() AccessorTypeOutput
	ToAccessorTypeOutputWithContext(context.Context) AccessorTypeOutput
}

var accessorTypePtrType = reflect.TypeOf((**AccessorType)(nil)).Elem()

type AccessorTypePtrInput interface {
	pulumi.Input

	ToAccessorTypePtrOutput() AccessorTypePtrOutput
	ToAccessorTypePtrOutputWithContext(context.Context) AccessorTypePtrOutput
}

type accessorTypePtr string

func AccessorTypePtr(v string) AccessorTypePtrInput {
	return (*accessorTypePtr)(&v)
}

func (*accessorTypePtr) ElementType() reflect.Type {
	return accessorTypePtrType
}

func (in *accessorTypePtr) ToAccessorTypePtrOutput() AccessorTypePtrOutput {
	return pulumi.ToOutput(in).(AccessorTypePtrOutput)
}

func (in *accessorTypePtr) ToAccessorTypePtrOutputWithContext(ctx context.Context) AccessorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessorTypePtrOutput)
}

func (in *accessorTypePtr) ToOutput(ctx context.Context) pulumix.Output[*AccessorType] {
	return pulumix.Output[*AccessorType]{
		OutputState: in.ToAccessorTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessorNetworkAccessorTypeInput)(nil)).Elem(), AccessorNetworkAccessorType("ETHEREUM_GOERLI"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessorNetworkAccessorTypePtrInput)(nil)).Elem(), AccessorNetworkAccessorType("ETHEREUM_GOERLI"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessorTypeInput)(nil)).Elem(), AccessorType("BILLING_TOKEN"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessorTypePtrInput)(nil)).Elem(), AccessorType("BILLING_TOKEN"))
	pulumi.RegisterOutputType(AccessorNetworkAccessorTypeOutput{})
	pulumi.RegisterOutputType(AccessorNetworkAccessorTypePtrOutput{})
	pulumi.RegisterOutputType(AccessorStatusOutput{})
	pulumi.RegisterOutputType(AccessorStatusPtrOutput{})
	pulumi.RegisterOutputType(AccessorTypeOutput{})
	pulumi.RegisterOutputType(AccessorTypePtrOutput{})
}
