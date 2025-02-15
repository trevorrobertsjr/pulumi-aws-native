// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package proton

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type EnvironmentAccountConnectionStatus string

const (
	EnvironmentAccountConnectionStatusPending   = EnvironmentAccountConnectionStatus("PENDING")
	EnvironmentAccountConnectionStatusConnected = EnvironmentAccountConnectionStatus("CONNECTED")
	EnvironmentAccountConnectionStatusRejected  = EnvironmentAccountConnectionStatus("REJECTED")
)

type EnvironmentAccountConnectionStatusOutput struct{ *pulumi.OutputState }

func (EnvironmentAccountConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentAccountConnectionStatus)(nil)).Elem()
}

func (o EnvironmentAccountConnectionStatusOutput) ToEnvironmentAccountConnectionStatusOutput() EnvironmentAccountConnectionStatusOutput {
	return o
}

func (o EnvironmentAccountConnectionStatusOutput) ToEnvironmentAccountConnectionStatusOutputWithContext(ctx context.Context) EnvironmentAccountConnectionStatusOutput {
	return o
}

func (o EnvironmentAccountConnectionStatusOutput) ToEnvironmentAccountConnectionStatusPtrOutput() EnvironmentAccountConnectionStatusPtrOutput {
	return o.ToEnvironmentAccountConnectionStatusPtrOutputWithContext(context.Background())
}

func (o EnvironmentAccountConnectionStatusOutput) ToEnvironmentAccountConnectionStatusPtrOutputWithContext(ctx context.Context) EnvironmentAccountConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentAccountConnectionStatus) *EnvironmentAccountConnectionStatus {
		return &v
	}).(EnvironmentAccountConnectionStatusPtrOutput)
}

func (o EnvironmentAccountConnectionStatusOutput) ToOutput(ctx context.Context) pulumix.Output[EnvironmentAccountConnectionStatus] {
	return pulumix.Output[EnvironmentAccountConnectionStatus]{
		OutputState: o.OutputState,
	}
}

func (o EnvironmentAccountConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentAccountConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentAccountConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentAccountConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentAccountConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentAccountConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentAccountConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentAccountConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentAccountConnectionStatus)(nil)).Elem()
}

func (o EnvironmentAccountConnectionStatusPtrOutput) ToEnvironmentAccountConnectionStatusPtrOutput() EnvironmentAccountConnectionStatusPtrOutput {
	return o
}

func (o EnvironmentAccountConnectionStatusPtrOutput) ToEnvironmentAccountConnectionStatusPtrOutputWithContext(ctx context.Context) EnvironmentAccountConnectionStatusPtrOutput {
	return o
}

func (o EnvironmentAccountConnectionStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentAccountConnectionStatus] {
	return pulumix.Output[*EnvironmentAccountConnectionStatus]{
		OutputState: o.OutputState,
	}
}

func (o EnvironmentAccountConnectionStatusPtrOutput) Elem() EnvironmentAccountConnectionStatusOutput {
	return o.ApplyT(func(v *EnvironmentAccountConnectionStatus) EnvironmentAccountConnectionStatus {
		if v != nil {
			return *v
		}
		var ret EnvironmentAccountConnectionStatus
		return ret
	}).(EnvironmentAccountConnectionStatusOutput)
}

func (o EnvironmentAccountConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentAccountConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentAccountConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentTemplateProvisioning string

const (
	EnvironmentTemplateProvisioningCustomerManaged = EnvironmentTemplateProvisioning("CUSTOMER_MANAGED")
)

func (EnvironmentTemplateProvisioning) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentTemplateProvisioning)(nil)).Elem()
}

func (e EnvironmentTemplateProvisioning) ToEnvironmentTemplateProvisioningOutput() EnvironmentTemplateProvisioningOutput {
	return pulumi.ToOutput(e).(EnvironmentTemplateProvisioningOutput)
}

func (e EnvironmentTemplateProvisioning) ToEnvironmentTemplateProvisioningOutputWithContext(ctx context.Context) EnvironmentTemplateProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnvironmentTemplateProvisioningOutput)
}

func (e EnvironmentTemplateProvisioning) ToEnvironmentTemplateProvisioningPtrOutput() EnvironmentTemplateProvisioningPtrOutput {
	return e.ToEnvironmentTemplateProvisioningPtrOutputWithContext(context.Background())
}

func (e EnvironmentTemplateProvisioning) ToEnvironmentTemplateProvisioningPtrOutputWithContext(ctx context.Context) EnvironmentTemplateProvisioningPtrOutput {
	return EnvironmentTemplateProvisioning(e).ToEnvironmentTemplateProvisioningOutputWithContext(ctx).ToEnvironmentTemplateProvisioningPtrOutputWithContext(ctx)
}

func (e EnvironmentTemplateProvisioning) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentTemplateProvisioning) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentTemplateProvisioning) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentTemplateProvisioning) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnvironmentTemplateProvisioningOutput struct{ *pulumi.OutputState }

func (EnvironmentTemplateProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentTemplateProvisioning)(nil)).Elem()
}

func (o EnvironmentTemplateProvisioningOutput) ToEnvironmentTemplateProvisioningOutput() EnvironmentTemplateProvisioningOutput {
	return o
}

func (o EnvironmentTemplateProvisioningOutput) ToEnvironmentTemplateProvisioningOutputWithContext(ctx context.Context) EnvironmentTemplateProvisioningOutput {
	return o
}

func (o EnvironmentTemplateProvisioningOutput) ToEnvironmentTemplateProvisioningPtrOutput() EnvironmentTemplateProvisioningPtrOutput {
	return o.ToEnvironmentTemplateProvisioningPtrOutputWithContext(context.Background())
}

func (o EnvironmentTemplateProvisioningOutput) ToEnvironmentTemplateProvisioningPtrOutputWithContext(ctx context.Context) EnvironmentTemplateProvisioningPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentTemplateProvisioning) *EnvironmentTemplateProvisioning {
		return &v
	}).(EnvironmentTemplateProvisioningPtrOutput)
}

func (o EnvironmentTemplateProvisioningOutput) ToOutput(ctx context.Context) pulumix.Output[EnvironmentTemplateProvisioning] {
	return pulumix.Output[EnvironmentTemplateProvisioning]{
		OutputState: o.OutputState,
	}
}

func (o EnvironmentTemplateProvisioningOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentTemplateProvisioningOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentTemplateProvisioning) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentTemplateProvisioningOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentTemplateProvisioningOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentTemplateProvisioning) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentTemplateProvisioningPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentTemplateProvisioningPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentTemplateProvisioning)(nil)).Elem()
}

func (o EnvironmentTemplateProvisioningPtrOutput) ToEnvironmentTemplateProvisioningPtrOutput() EnvironmentTemplateProvisioningPtrOutput {
	return o
}

func (o EnvironmentTemplateProvisioningPtrOutput) ToEnvironmentTemplateProvisioningPtrOutputWithContext(ctx context.Context) EnvironmentTemplateProvisioningPtrOutput {
	return o
}

func (o EnvironmentTemplateProvisioningPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentTemplateProvisioning] {
	return pulumix.Output[*EnvironmentTemplateProvisioning]{
		OutputState: o.OutputState,
	}
}

func (o EnvironmentTemplateProvisioningPtrOutput) Elem() EnvironmentTemplateProvisioningOutput {
	return o.ApplyT(func(v *EnvironmentTemplateProvisioning) EnvironmentTemplateProvisioning {
		if v != nil {
			return *v
		}
		var ret EnvironmentTemplateProvisioning
		return ret
	}).(EnvironmentTemplateProvisioningOutput)
}

func (o EnvironmentTemplateProvisioningPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentTemplateProvisioningPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentTemplateProvisioning) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnvironmentTemplateProvisioningInput is an input type that accepts EnvironmentTemplateProvisioningArgs and EnvironmentTemplateProvisioningOutput values.
// You can construct a concrete instance of `EnvironmentTemplateProvisioningInput` via:
//
//	EnvironmentTemplateProvisioningArgs{...}
type EnvironmentTemplateProvisioningInput interface {
	pulumi.Input

	ToEnvironmentTemplateProvisioningOutput() EnvironmentTemplateProvisioningOutput
	ToEnvironmentTemplateProvisioningOutputWithContext(context.Context) EnvironmentTemplateProvisioningOutput
}

var environmentTemplateProvisioningPtrType = reflect.TypeOf((**EnvironmentTemplateProvisioning)(nil)).Elem()

type EnvironmentTemplateProvisioningPtrInput interface {
	pulumi.Input

	ToEnvironmentTemplateProvisioningPtrOutput() EnvironmentTemplateProvisioningPtrOutput
	ToEnvironmentTemplateProvisioningPtrOutputWithContext(context.Context) EnvironmentTemplateProvisioningPtrOutput
}

type environmentTemplateProvisioningPtr string

func EnvironmentTemplateProvisioningPtr(v string) EnvironmentTemplateProvisioningPtrInput {
	return (*environmentTemplateProvisioningPtr)(&v)
}

func (*environmentTemplateProvisioningPtr) ElementType() reflect.Type {
	return environmentTemplateProvisioningPtrType
}

func (in *environmentTemplateProvisioningPtr) ToEnvironmentTemplateProvisioningPtrOutput() EnvironmentTemplateProvisioningPtrOutput {
	return pulumi.ToOutput(in).(EnvironmentTemplateProvisioningPtrOutput)
}

func (in *environmentTemplateProvisioningPtr) ToEnvironmentTemplateProvisioningPtrOutputWithContext(ctx context.Context) EnvironmentTemplateProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnvironmentTemplateProvisioningPtrOutput)
}

func (in *environmentTemplateProvisioningPtr) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentTemplateProvisioning] {
	return pulumix.Output[*EnvironmentTemplateProvisioning]{
		OutputState: in.ToEnvironmentTemplateProvisioningPtrOutputWithContext(ctx).OutputState,
	}
}

type ServiceTemplateProvisioning string

const (
	ServiceTemplateProvisioningCustomerManaged = ServiceTemplateProvisioning("CUSTOMER_MANAGED")
)

func (ServiceTemplateProvisioning) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateProvisioning)(nil)).Elem()
}

func (e ServiceTemplateProvisioning) ToServiceTemplateProvisioningOutput() ServiceTemplateProvisioningOutput {
	return pulumi.ToOutput(e).(ServiceTemplateProvisioningOutput)
}

func (e ServiceTemplateProvisioning) ToServiceTemplateProvisioningOutputWithContext(ctx context.Context) ServiceTemplateProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceTemplateProvisioningOutput)
}

func (e ServiceTemplateProvisioning) ToServiceTemplateProvisioningPtrOutput() ServiceTemplateProvisioningPtrOutput {
	return e.ToServiceTemplateProvisioningPtrOutputWithContext(context.Background())
}

func (e ServiceTemplateProvisioning) ToServiceTemplateProvisioningPtrOutputWithContext(ctx context.Context) ServiceTemplateProvisioningPtrOutput {
	return ServiceTemplateProvisioning(e).ToServiceTemplateProvisioningOutputWithContext(ctx).ToServiceTemplateProvisioningPtrOutputWithContext(ctx)
}

func (e ServiceTemplateProvisioning) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceTemplateProvisioning) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceTemplateProvisioning) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceTemplateProvisioning) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceTemplateProvisioningOutput struct{ *pulumi.OutputState }

func (ServiceTemplateProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateProvisioning)(nil)).Elem()
}

func (o ServiceTemplateProvisioningOutput) ToServiceTemplateProvisioningOutput() ServiceTemplateProvisioningOutput {
	return o
}

func (o ServiceTemplateProvisioningOutput) ToServiceTemplateProvisioningOutputWithContext(ctx context.Context) ServiceTemplateProvisioningOutput {
	return o
}

func (o ServiceTemplateProvisioningOutput) ToServiceTemplateProvisioningPtrOutput() ServiceTemplateProvisioningPtrOutput {
	return o.ToServiceTemplateProvisioningPtrOutputWithContext(context.Background())
}

func (o ServiceTemplateProvisioningOutput) ToServiceTemplateProvisioningPtrOutputWithContext(ctx context.Context) ServiceTemplateProvisioningPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTemplateProvisioning) *ServiceTemplateProvisioning {
		return &v
	}).(ServiceTemplateProvisioningPtrOutput)
}

func (o ServiceTemplateProvisioningOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceTemplateProvisioning] {
	return pulumix.Output[ServiceTemplateProvisioning]{
		OutputState: o.OutputState,
	}
}

func (o ServiceTemplateProvisioningOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceTemplateProvisioningOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceTemplateProvisioning) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceTemplateProvisioningOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceTemplateProvisioningOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceTemplateProvisioning) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceTemplateProvisioningPtrOutput struct{ *pulumi.OutputState }

func (ServiceTemplateProvisioningPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTemplateProvisioning)(nil)).Elem()
}

func (o ServiceTemplateProvisioningPtrOutput) ToServiceTemplateProvisioningPtrOutput() ServiceTemplateProvisioningPtrOutput {
	return o
}

func (o ServiceTemplateProvisioningPtrOutput) ToServiceTemplateProvisioningPtrOutputWithContext(ctx context.Context) ServiceTemplateProvisioningPtrOutput {
	return o
}

func (o ServiceTemplateProvisioningPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceTemplateProvisioning] {
	return pulumix.Output[*ServiceTemplateProvisioning]{
		OutputState: o.OutputState,
	}
}

func (o ServiceTemplateProvisioningPtrOutput) Elem() ServiceTemplateProvisioningOutput {
	return o.ApplyT(func(v *ServiceTemplateProvisioning) ServiceTemplateProvisioning {
		if v != nil {
			return *v
		}
		var ret ServiceTemplateProvisioning
		return ret
	}).(ServiceTemplateProvisioningOutput)
}

func (o ServiceTemplateProvisioningPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceTemplateProvisioningPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceTemplateProvisioning) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServiceTemplateProvisioningInput is an input type that accepts ServiceTemplateProvisioningArgs and ServiceTemplateProvisioningOutput values.
// You can construct a concrete instance of `ServiceTemplateProvisioningInput` via:
//
//	ServiceTemplateProvisioningArgs{...}
type ServiceTemplateProvisioningInput interface {
	pulumi.Input

	ToServiceTemplateProvisioningOutput() ServiceTemplateProvisioningOutput
	ToServiceTemplateProvisioningOutputWithContext(context.Context) ServiceTemplateProvisioningOutput
}

var serviceTemplateProvisioningPtrType = reflect.TypeOf((**ServiceTemplateProvisioning)(nil)).Elem()

type ServiceTemplateProvisioningPtrInput interface {
	pulumi.Input

	ToServiceTemplateProvisioningPtrOutput() ServiceTemplateProvisioningPtrOutput
	ToServiceTemplateProvisioningPtrOutputWithContext(context.Context) ServiceTemplateProvisioningPtrOutput
}

type serviceTemplateProvisioningPtr string

func ServiceTemplateProvisioningPtr(v string) ServiceTemplateProvisioningPtrInput {
	return (*serviceTemplateProvisioningPtr)(&v)
}

func (*serviceTemplateProvisioningPtr) ElementType() reflect.Type {
	return serviceTemplateProvisioningPtrType
}

func (in *serviceTemplateProvisioningPtr) ToServiceTemplateProvisioningPtrOutput() ServiceTemplateProvisioningPtrOutput {
	return pulumi.ToOutput(in).(ServiceTemplateProvisioningPtrOutput)
}

func (in *serviceTemplateProvisioningPtr) ToServiceTemplateProvisioningPtrOutputWithContext(ctx context.Context) ServiceTemplateProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceTemplateProvisioningPtrOutput)
}

func (in *serviceTemplateProvisioningPtr) ToOutput(ctx context.Context) pulumix.Output[*ServiceTemplateProvisioning] {
	return pulumix.Output[*ServiceTemplateProvisioning]{
		OutputState: in.ToServiceTemplateProvisioningPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentTemplateProvisioningInput)(nil)).Elem(), EnvironmentTemplateProvisioning("CUSTOMER_MANAGED"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentTemplateProvisioningPtrInput)(nil)).Elem(), EnvironmentTemplateProvisioning("CUSTOMER_MANAGED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTemplateProvisioningInput)(nil)).Elem(), ServiceTemplateProvisioning("CUSTOMER_MANAGED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTemplateProvisioningPtrInput)(nil)).Elem(), ServiceTemplateProvisioning("CUSTOMER_MANAGED"))
	pulumi.RegisterOutputType(EnvironmentAccountConnectionStatusOutput{})
	pulumi.RegisterOutputType(EnvironmentAccountConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentTemplateProvisioningOutput{})
	pulumi.RegisterOutputType(EnvironmentTemplateProvisioningPtrOutput{})
	pulumi.RegisterOutputType(ServiceTemplateProvisioningOutput{})
	pulumi.RegisterOutputType(ServiceTemplateProvisioningPtrOutput{})
}
