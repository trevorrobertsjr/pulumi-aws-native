// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
type KeyOrigin string

const (
	KeyOriginAwsKms   = KeyOrigin("AWS_KMS")
	KeyOriginExternal = KeyOrigin("EXTERNAL")
)

func (KeyOrigin) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyOrigin)(nil)).Elem()
}

func (e KeyOrigin) ToKeyOriginOutput() KeyOriginOutput {
	return pulumi.ToOutput(e).(KeyOriginOutput)
}

func (e KeyOrigin) ToKeyOriginOutputWithContext(ctx context.Context) KeyOriginOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyOriginOutput)
}

func (e KeyOrigin) ToKeyOriginPtrOutput() KeyOriginPtrOutput {
	return e.ToKeyOriginPtrOutputWithContext(context.Background())
}

func (e KeyOrigin) ToKeyOriginPtrOutputWithContext(ctx context.Context) KeyOriginPtrOutput {
	return KeyOrigin(e).ToKeyOriginOutputWithContext(ctx).ToKeyOriginPtrOutputWithContext(ctx)
}

func (e KeyOrigin) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyOrigin) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyOrigin) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyOrigin) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyOriginOutput struct{ *pulumi.OutputState }

func (KeyOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyOrigin)(nil)).Elem()
}

func (o KeyOriginOutput) ToKeyOriginOutput() KeyOriginOutput {
	return o
}

func (o KeyOriginOutput) ToKeyOriginOutputWithContext(ctx context.Context) KeyOriginOutput {
	return o
}

func (o KeyOriginOutput) ToKeyOriginPtrOutput() KeyOriginPtrOutput {
	return o.ToKeyOriginPtrOutputWithContext(context.Background())
}

func (o KeyOriginOutput) ToKeyOriginPtrOutputWithContext(ctx context.Context) KeyOriginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyOrigin) *KeyOrigin {
		return &v
	}).(KeyOriginPtrOutput)
}

func (o KeyOriginOutput) ToOutput(ctx context.Context) pulumix.Output[KeyOrigin] {
	return pulumix.Output[KeyOrigin]{
		OutputState: o.OutputState,
	}
}

func (o KeyOriginOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyOriginOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyOrigin) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyOriginOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyOriginOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyOrigin) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyOriginPtrOutput struct{ *pulumi.OutputState }

func (KeyOriginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyOrigin)(nil)).Elem()
}

func (o KeyOriginPtrOutput) ToKeyOriginPtrOutput() KeyOriginPtrOutput {
	return o
}

func (o KeyOriginPtrOutput) ToKeyOriginPtrOutputWithContext(ctx context.Context) KeyOriginPtrOutput {
	return o
}

func (o KeyOriginPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*KeyOrigin] {
	return pulumix.Output[*KeyOrigin]{
		OutputState: o.OutputState,
	}
}

func (o KeyOriginPtrOutput) Elem() KeyOriginOutput {
	return o.ApplyT(func(v *KeyOrigin) KeyOrigin {
		if v != nil {
			return *v
		}
		var ret KeyOrigin
		return ret
	}).(KeyOriginOutput)
}

func (o KeyOriginPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyOriginPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyOrigin) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyOriginInput is an input type that accepts KeyOriginArgs and KeyOriginOutput values.
// You can construct a concrete instance of `KeyOriginInput` via:
//
//	KeyOriginArgs{...}
type KeyOriginInput interface {
	pulumi.Input

	ToKeyOriginOutput() KeyOriginOutput
	ToKeyOriginOutputWithContext(context.Context) KeyOriginOutput
}

var keyOriginPtrType = reflect.TypeOf((**KeyOrigin)(nil)).Elem()

type KeyOriginPtrInput interface {
	pulumi.Input

	ToKeyOriginPtrOutput() KeyOriginPtrOutput
	ToKeyOriginPtrOutputWithContext(context.Context) KeyOriginPtrOutput
}

type keyOriginPtr string

func KeyOriginPtr(v string) KeyOriginPtrInput {
	return (*keyOriginPtr)(&v)
}

func (*keyOriginPtr) ElementType() reflect.Type {
	return keyOriginPtrType
}

func (in *keyOriginPtr) ToKeyOriginPtrOutput() KeyOriginPtrOutput {
	return pulumi.ToOutput(in).(KeyOriginPtrOutput)
}

func (in *keyOriginPtr) ToKeyOriginPtrOutputWithContext(ctx context.Context) KeyOriginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyOriginPtrOutput)
}

func (in *keyOriginPtr) ToOutput(ctx context.Context) pulumix.Output[*KeyOrigin] {
	return pulumix.Output[*KeyOrigin]{
		OutputState: in.ToKeyOriginPtrOutputWithContext(ctx).OutputState,
	}
}

// Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
type KeySpec string

const (
	KeySpecSymmetricDefault = KeySpec("SYMMETRIC_DEFAULT")
	KeySpecRsa2048          = KeySpec("RSA_2048")
	KeySpecRsa3072          = KeySpec("RSA_3072")
	KeySpecRsa4096          = KeySpec("RSA_4096")
	KeySpecEccNistP256      = KeySpec("ECC_NIST_P256")
	KeySpecEccNistP384      = KeySpec("ECC_NIST_P384")
	KeySpecEccNistP521      = KeySpec("ECC_NIST_P521")
	KeySpecEccSecgP256k1    = KeySpec("ECC_SECG_P256K1")
	KeySpecHmac224          = KeySpec("HMAC_224")
	KeySpecHmac256          = KeySpec("HMAC_256")
	KeySpecHmac384          = KeySpec("HMAC_384")
	KeySpecHmac512          = KeySpec("HMAC_512")
	KeySpecSm2              = KeySpec("SM2")
)

func (KeySpec) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySpec)(nil)).Elem()
}

func (e KeySpec) ToKeySpecOutput() KeySpecOutput {
	return pulumi.ToOutput(e).(KeySpecOutput)
}

func (e KeySpec) ToKeySpecOutputWithContext(ctx context.Context) KeySpecOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeySpecOutput)
}

func (e KeySpec) ToKeySpecPtrOutput() KeySpecPtrOutput {
	return e.ToKeySpecPtrOutputWithContext(context.Background())
}

func (e KeySpec) ToKeySpecPtrOutputWithContext(ctx context.Context) KeySpecPtrOutput {
	return KeySpec(e).ToKeySpecOutputWithContext(ctx).ToKeySpecPtrOutputWithContext(ctx)
}

func (e KeySpec) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySpec) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySpec) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeySpec) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeySpecOutput struct{ *pulumi.OutputState }

func (KeySpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySpec)(nil)).Elem()
}

func (o KeySpecOutput) ToKeySpecOutput() KeySpecOutput {
	return o
}

func (o KeySpecOutput) ToKeySpecOutputWithContext(ctx context.Context) KeySpecOutput {
	return o
}

func (o KeySpecOutput) ToKeySpecPtrOutput() KeySpecPtrOutput {
	return o.ToKeySpecPtrOutputWithContext(context.Background())
}

func (o KeySpecOutput) ToKeySpecPtrOutputWithContext(ctx context.Context) KeySpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeySpec) *KeySpec {
		return &v
	}).(KeySpecPtrOutput)
}

func (o KeySpecOutput) ToOutput(ctx context.Context) pulumix.Output[KeySpec] {
	return pulumix.Output[KeySpec]{
		OutputState: o.OutputState,
	}
}

func (o KeySpecOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeySpecOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySpec) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeySpecOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySpecOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySpec) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeySpecPtrOutput struct{ *pulumi.OutputState }

func (KeySpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySpec)(nil)).Elem()
}

func (o KeySpecPtrOutput) ToKeySpecPtrOutput() KeySpecPtrOutput {
	return o
}

func (o KeySpecPtrOutput) ToKeySpecPtrOutputWithContext(ctx context.Context) KeySpecPtrOutput {
	return o
}

func (o KeySpecPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*KeySpec] {
	return pulumix.Output[*KeySpec]{
		OutputState: o.OutputState,
	}
}

func (o KeySpecPtrOutput) Elem() KeySpecOutput {
	return o.ApplyT(func(v *KeySpec) KeySpec {
		if v != nil {
			return *v
		}
		var ret KeySpec
		return ret
	}).(KeySpecOutput)
}

func (o KeySpecPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySpecPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeySpec) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeySpecInput is an input type that accepts KeySpecArgs and KeySpecOutput values.
// You can construct a concrete instance of `KeySpecInput` via:
//
//	KeySpecArgs{...}
type KeySpecInput interface {
	pulumi.Input

	ToKeySpecOutput() KeySpecOutput
	ToKeySpecOutputWithContext(context.Context) KeySpecOutput
}

var keySpecPtrType = reflect.TypeOf((**KeySpec)(nil)).Elem()

type KeySpecPtrInput interface {
	pulumi.Input

	ToKeySpecPtrOutput() KeySpecPtrOutput
	ToKeySpecPtrOutputWithContext(context.Context) KeySpecPtrOutput
}

type keySpecPtr string

func KeySpecPtr(v string) KeySpecPtrInput {
	return (*keySpecPtr)(&v)
}

func (*keySpecPtr) ElementType() reflect.Type {
	return keySpecPtrType
}

func (in *keySpecPtr) ToKeySpecPtrOutput() KeySpecPtrOutput {
	return pulumi.ToOutput(in).(KeySpecPtrOutput)
}

func (in *keySpecPtr) ToKeySpecPtrOutputWithContext(ctx context.Context) KeySpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeySpecPtrOutput)
}

func (in *keySpecPtr) ToOutput(ctx context.Context) pulumix.Output[*KeySpec] {
	return pulumix.Output[*KeySpec]{
		OutputState: in.ToKeySpecPtrOutputWithContext(ctx).OutputState,
	}
}

// Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
type KeyUsage string

const (
	KeyUsageEncryptDecrypt    = KeyUsage("ENCRYPT_DECRYPT")
	KeyUsageSignVerify        = KeyUsage("SIGN_VERIFY")
	KeyUsageGenerateVerifyMac = KeyUsage("GENERATE_VERIFY_MAC")
)

func (KeyUsage) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyUsage)(nil)).Elem()
}

func (e KeyUsage) ToKeyUsageOutput() KeyUsageOutput {
	return pulumi.ToOutput(e).(KeyUsageOutput)
}

func (e KeyUsage) ToKeyUsageOutputWithContext(ctx context.Context) KeyUsageOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyUsageOutput)
}

func (e KeyUsage) ToKeyUsagePtrOutput() KeyUsagePtrOutput {
	return e.ToKeyUsagePtrOutputWithContext(context.Background())
}

func (e KeyUsage) ToKeyUsagePtrOutputWithContext(ctx context.Context) KeyUsagePtrOutput {
	return KeyUsage(e).ToKeyUsageOutputWithContext(ctx).ToKeyUsagePtrOutputWithContext(ctx)
}

func (e KeyUsage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyUsage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyUsage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyUsage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyUsageOutput struct{ *pulumi.OutputState }

func (KeyUsageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyUsage)(nil)).Elem()
}

func (o KeyUsageOutput) ToKeyUsageOutput() KeyUsageOutput {
	return o
}

func (o KeyUsageOutput) ToKeyUsageOutputWithContext(ctx context.Context) KeyUsageOutput {
	return o
}

func (o KeyUsageOutput) ToKeyUsagePtrOutput() KeyUsagePtrOutput {
	return o.ToKeyUsagePtrOutputWithContext(context.Background())
}

func (o KeyUsageOutput) ToKeyUsagePtrOutputWithContext(ctx context.Context) KeyUsagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyUsage) *KeyUsage {
		return &v
	}).(KeyUsagePtrOutput)
}

func (o KeyUsageOutput) ToOutput(ctx context.Context) pulumix.Output[KeyUsage] {
	return pulumix.Output[KeyUsage]{
		OutputState: o.OutputState,
	}
}

func (o KeyUsageOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyUsageOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyUsage) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyUsageOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyUsageOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyUsage) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyUsagePtrOutput struct{ *pulumi.OutputState }

func (KeyUsagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyUsage)(nil)).Elem()
}

func (o KeyUsagePtrOutput) ToKeyUsagePtrOutput() KeyUsagePtrOutput {
	return o
}

func (o KeyUsagePtrOutput) ToKeyUsagePtrOutputWithContext(ctx context.Context) KeyUsagePtrOutput {
	return o
}

func (o KeyUsagePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*KeyUsage] {
	return pulumix.Output[*KeyUsage]{
		OutputState: o.OutputState,
	}
}

func (o KeyUsagePtrOutput) Elem() KeyUsageOutput {
	return o.ApplyT(func(v *KeyUsage) KeyUsage {
		if v != nil {
			return *v
		}
		var ret KeyUsage
		return ret
	}).(KeyUsageOutput)
}

func (o KeyUsagePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyUsagePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyUsage) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyUsageInput is an input type that accepts KeyUsageArgs and KeyUsageOutput values.
// You can construct a concrete instance of `KeyUsageInput` via:
//
//	KeyUsageArgs{...}
type KeyUsageInput interface {
	pulumi.Input

	ToKeyUsageOutput() KeyUsageOutput
	ToKeyUsageOutputWithContext(context.Context) KeyUsageOutput
}

var keyUsagePtrType = reflect.TypeOf((**KeyUsage)(nil)).Elem()

type KeyUsagePtrInput interface {
	pulumi.Input

	ToKeyUsagePtrOutput() KeyUsagePtrOutput
	ToKeyUsagePtrOutputWithContext(context.Context) KeyUsagePtrOutput
}

type keyUsagePtr string

func KeyUsagePtr(v string) KeyUsagePtrInput {
	return (*keyUsagePtr)(&v)
}

func (*keyUsagePtr) ElementType() reflect.Type {
	return keyUsagePtrType
}

func (in *keyUsagePtr) ToKeyUsagePtrOutput() KeyUsagePtrOutput {
	return pulumi.ToOutput(in).(KeyUsagePtrOutput)
}

func (in *keyUsagePtr) ToKeyUsagePtrOutputWithContext(ctx context.Context) KeyUsagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyUsagePtrOutput)
}

func (in *keyUsagePtr) ToOutput(ctx context.Context) pulumix.Output[*KeyUsage] {
	return pulumix.Output[*KeyUsage]{
		OutputState: in.ToKeyUsagePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyOriginInput)(nil)).Elem(), KeyOrigin("AWS_KMS"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyOriginPtrInput)(nil)).Elem(), KeyOrigin("AWS_KMS"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySpecInput)(nil)).Elem(), KeySpec("SYMMETRIC_DEFAULT"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySpecPtrInput)(nil)).Elem(), KeySpec("SYMMETRIC_DEFAULT"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyUsageInput)(nil)).Elem(), KeyUsage("ENCRYPT_DECRYPT"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyUsagePtrInput)(nil)).Elem(), KeyUsage("ENCRYPT_DECRYPT"))
	pulumi.RegisterOutputType(KeyOriginOutput{})
	pulumi.RegisterOutputType(KeyOriginPtrOutput{})
	pulumi.RegisterOutputType(KeySpecOutput{})
	pulumi.RegisterOutputType(KeySpecPtrOutput{})
	pulumi.RegisterOutputType(KeyUsageOutput{})
	pulumi.RegisterOutputType(KeyUsagePtrOutput{})
}
