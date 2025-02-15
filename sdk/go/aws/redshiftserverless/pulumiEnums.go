// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshiftserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type NamespaceLogExport string

const (
	NamespaceLogExportUseractivitylog = NamespaceLogExport("useractivitylog")
	NamespaceLogExportUserlog         = NamespaceLogExport("userlog")
	NamespaceLogExportConnectionlog   = NamespaceLogExport("connectionlog")
)

func (NamespaceLogExport) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceLogExport)(nil)).Elem()
}

func (e NamespaceLogExport) ToNamespaceLogExportOutput() NamespaceLogExportOutput {
	return pulumi.ToOutput(e).(NamespaceLogExportOutput)
}

func (e NamespaceLogExport) ToNamespaceLogExportOutputWithContext(ctx context.Context) NamespaceLogExportOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NamespaceLogExportOutput)
}

func (e NamespaceLogExport) ToNamespaceLogExportPtrOutput() NamespaceLogExportPtrOutput {
	return e.ToNamespaceLogExportPtrOutputWithContext(context.Background())
}

func (e NamespaceLogExport) ToNamespaceLogExportPtrOutputWithContext(ctx context.Context) NamespaceLogExportPtrOutput {
	return NamespaceLogExport(e).ToNamespaceLogExportOutputWithContext(ctx).ToNamespaceLogExportPtrOutputWithContext(ctx)
}

func (e NamespaceLogExport) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NamespaceLogExport) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NamespaceLogExport) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NamespaceLogExport) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NamespaceLogExportOutput struct{ *pulumi.OutputState }

func (NamespaceLogExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceLogExport)(nil)).Elem()
}

func (o NamespaceLogExportOutput) ToNamespaceLogExportOutput() NamespaceLogExportOutput {
	return o
}

func (o NamespaceLogExportOutput) ToNamespaceLogExportOutputWithContext(ctx context.Context) NamespaceLogExportOutput {
	return o
}

func (o NamespaceLogExportOutput) ToNamespaceLogExportPtrOutput() NamespaceLogExportPtrOutput {
	return o.ToNamespaceLogExportPtrOutputWithContext(context.Background())
}

func (o NamespaceLogExportOutput) ToNamespaceLogExportPtrOutputWithContext(ctx context.Context) NamespaceLogExportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamespaceLogExport) *NamespaceLogExport {
		return &v
	}).(NamespaceLogExportPtrOutput)
}

func (o NamespaceLogExportOutput) ToOutput(ctx context.Context) pulumix.Output[NamespaceLogExport] {
	return pulumix.Output[NamespaceLogExport]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceLogExportOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NamespaceLogExportOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NamespaceLogExport) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NamespaceLogExportOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamespaceLogExportOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NamespaceLogExport) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NamespaceLogExportPtrOutput struct{ *pulumi.OutputState }

func (NamespaceLogExportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceLogExport)(nil)).Elem()
}

func (o NamespaceLogExportPtrOutput) ToNamespaceLogExportPtrOutput() NamespaceLogExportPtrOutput {
	return o
}

func (o NamespaceLogExportPtrOutput) ToNamespaceLogExportPtrOutputWithContext(ctx context.Context) NamespaceLogExportPtrOutput {
	return o
}

func (o NamespaceLogExportPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*NamespaceLogExport] {
	return pulumix.Output[*NamespaceLogExport]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceLogExportPtrOutput) Elem() NamespaceLogExportOutput {
	return o.ApplyT(func(v *NamespaceLogExport) NamespaceLogExport {
		if v != nil {
			return *v
		}
		var ret NamespaceLogExport
		return ret
	}).(NamespaceLogExportOutput)
}

func (o NamespaceLogExportPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamespaceLogExportPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NamespaceLogExport) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NamespaceLogExportInput is an input type that accepts NamespaceLogExportArgs and NamespaceLogExportOutput values.
// You can construct a concrete instance of `NamespaceLogExportInput` via:
//
//	NamespaceLogExportArgs{...}
type NamespaceLogExportInput interface {
	pulumi.Input

	ToNamespaceLogExportOutput() NamespaceLogExportOutput
	ToNamespaceLogExportOutputWithContext(context.Context) NamespaceLogExportOutput
}

var namespaceLogExportPtrType = reflect.TypeOf((**NamespaceLogExport)(nil)).Elem()

type NamespaceLogExportPtrInput interface {
	pulumi.Input

	ToNamespaceLogExportPtrOutput() NamespaceLogExportPtrOutput
	ToNamespaceLogExportPtrOutputWithContext(context.Context) NamespaceLogExportPtrOutput
}

type namespaceLogExportPtr string

func NamespaceLogExportPtr(v string) NamespaceLogExportPtrInput {
	return (*namespaceLogExportPtr)(&v)
}

func (*namespaceLogExportPtr) ElementType() reflect.Type {
	return namespaceLogExportPtrType
}

func (in *namespaceLogExportPtr) ToNamespaceLogExportPtrOutput() NamespaceLogExportPtrOutput {
	return pulumi.ToOutput(in).(NamespaceLogExportPtrOutput)
}

func (in *namespaceLogExportPtr) ToNamespaceLogExportPtrOutputWithContext(ctx context.Context) NamespaceLogExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NamespaceLogExportPtrOutput)
}

func (in *namespaceLogExportPtr) ToOutput(ctx context.Context) pulumix.Output[*NamespaceLogExport] {
	return pulumix.Output[*NamespaceLogExport]{
		OutputState: in.ToNamespaceLogExportPtrOutputWithContext(ctx).OutputState,
	}
}

// NamespaceLogExportArrayInput is an input type that accepts NamespaceLogExportArray and NamespaceLogExportArrayOutput values.
// You can construct a concrete instance of `NamespaceLogExportArrayInput` via:
//
//	NamespaceLogExportArray{ NamespaceLogExportArgs{...} }
type NamespaceLogExportArrayInput interface {
	pulumi.Input

	ToNamespaceLogExportArrayOutput() NamespaceLogExportArrayOutput
	ToNamespaceLogExportArrayOutputWithContext(context.Context) NamespaceLogExportArrayOutput
}

type NamespaceLogExportArray []NamespaceLogExport

func (NamespaceLogExportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceLogExport)(nil)).Elem()
}

func (i NamespaceLogExportArray) ToNamespaceLogExportArrayOutput() NamespaceLogExportArrayOutput {
	return i.ToNamespaceLogExportArrayOutputWithContext(context.Background())
}

func (i NamespaceLogExportArray) ToNamespaceLogExportArrayOutputWithContext(ctx context.Context) NamespaceLogExportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceLogExportArrayOutput)
}

func (i NamespaceLogExportArray) ToOutput(ctx context.Context) pulumix.Output[[]NamespaceLogExport] {
	return pulumix.Output[[]NamespaceLogExport]{
		OutputState: i.ToNamespaceLogExportArrayOutputWithContext(ctx).OutputState,
	}
}

type NamespaceLogExportArrayOutput struct{ *pulumi.OutputState }

func (NamespaceLogExportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceLogExport)(nil)).Elem()
}

func (o NamespaceLogExportArrayOutput) ToNamespaceLogExportArrayOutput() NamespaceLogExportArrayOutput {
	return o
}

func (o NamespaceLogExportArrayOutput) ToNamespaceLogExportArrayOutputWithContext(ctx context.Context) NamespaceLogExportArrayOutput {
	return o
}

func (o NamespaceLogExportArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]NamespaceLogExport] {
	return pulumix.Output[[]NamespaceLogExport]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceLogExportArrayOutput) Index(i pulumi.IntInput) NamespaceLogExportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceLogExport {
		return vs[0].([]NamespaceLogExport)[vs[1].(int)]
	}).(NamespaceLogExportOutput)
}

type NamespaceStatus string

const (
	NamespaceStatusAvailable = NamespaceStatus("AVAILABLE")
	NamespaceStatusModifying = NamespaceStatus("MODIFYING")
	NamespaceStatusDeleting  = NamespaceStatus("DELETING")
)

type NamespaceStatusOutput struct{ *pulumi.OutputState }

func (NamespaceStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceStatus)(nil)).Elem()
}

func (o NamespaceStatusOutput) ToNamespaceStatusOutput() NamespaceStatusOutput {
	return o
}

func (o NamespaceStatusOutput) ToNamespaceStatusOutputWithContext(ctx context.Context) NamespaceStatusOutput {
	return o
}

func (o NamespaceStatusOutput) ToNamespaceStatusPtrOutput() NamespaceStatusPtrOutput {
	return o.ToNamespaceStatusPtrOutputWithContext(context.Background())
}

func (o NamespaceStatusOutput) ToNamespaceStatusPtrOutputWithContext(ctx context.Context) NamespaceStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamespaceStatus) *NamespaceStatus {
		return &v
	}).(NamespaceStatusPtrOutput)
}

func (o NamespaceStatusOutput) ToOutput(ctx context.Context) pulumix.Output[NamespaceStatus] {
	return pulumix.Output[NamespaceStatus]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NamespaceStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NamespaceStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NamespaceStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamespaceStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NamespaceStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NamespaceStatusPtrOutput struct{ *pulumi.OutputState }

func (NamespaceStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceStatus)(nil)).Elem()
}

func (o NamespaceStatusPtrOutput) ToNamespaceStatusPtrOutput() NamespaceStatusPtrOutput {
	return o
}

func (o NamespaceStatusPtrOutput) ToNamespaceStatusPtrOutputWithContext(ctx context.Context) NamespaceStatusPtrOutput {
	return o
}

func (o NamespaceStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*NamespaceStatus] {
	return pulumix.Output[*NamespaceStatus]{
		OutputState: o.OutputState,
	}
}

func (o NamespaceStatusPtrOutput) Elem() NamespaceStatusOutput {
	return o.ApplyT(func(v *NamespaceStatus) NamespaceStatus {
		if v != nil {
			return *v
		}
		var ret NamespaceStatus
		return ret
	}).(NamespaceStatusOutput)
}

func (o NamespaceStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamespaceStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NamespaceStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkgroupStatus string

const (
	WorkgroupStatusCreating  = WorkgroupStatus("CREATING")
	WorkgroupStatusAvailable = WorkgroupStatus("AVAILABLE")
	WorkgroupStatusModifying = WorkgroupStatus("MODIFYING")
	WorkgroupStatusDeleting  = WorkgroupStatus("DELETING")
)

type WorkgroupStatusOutput struct{ *pulumi.OutputState }

func (WorkgroupStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkgroupStatus)(nil)).Elem()
}

func (o WorkgroupStatusOutput) ToWorkgroupStatusOutput() WorkgroupStatusOutput {
	return o
}

func (o WorkgroupStatusOutput) ToWorkgroupStatusOutputWithContext(ctx context.Context) WorkgroupStatusOutput {
	return o
}

func (o WorkgroupStatusOutput) ToWorkgroupStatusPtrOutput() WorkgroupStatusPtrOutput {
	return o.ToWorkgroupStatusPtrOutputWithContext(context.Background())
}

func (o WorkgroupStatusOutput) ToWorkgroupStatusPtrOutputWithContext(ctx context.Context) WorkgroupStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkgroupStatus) *WorkgroupStatus {
		return &v
	}).(WorkgroupStatusPtrOutput)
}

func (o WorkgroupStatusOutput) ToOutput(ctx context.Context) pulumix.Output[WorkgroupStatus] {
	return pulumix.Output[WorkgroupStatus]{
		OutputState: o.OutputState,
	}
}

func (o WorkgroupStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkgroupStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkgroupStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkgroupStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkgroupStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkgroupStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkgroupStatusPtrOutput struct{ *pulumi.OutputState }

func (WorkgroupStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkgroupStatus)(nil)).Elem()
}

func (o WorkgroupStatusPtrOutput) ToWorkgroupStatusPtrOutput() WorkgroupStatusPtrOutput {
	return o
}

func (o WorkgroupStatusPtrOutput) ToWorkgroupStatusPtrOutputWithContext(ctx context.Context) WorkgroupStatusPtrOutput {
	return o
}

func (o WorkgroupStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkgroupStatus] {
	return pulumix.Output[*WorkgroupStatus]{
		OutputState: o.OutputState,
	}
}

func (o WorkgroupStatusPtrOutput) Elem() WorkgroupStatusOutput {
	return o.ApplyT(func(v *WorkgroupStatus) WorkgroupStatus {
		if v != nil {
			return *v
		}
		var ret WorkgroupStatus
		return ret
	}).(WorkgroupStatusOutput)
}

func (o WorkgroupStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkgroupStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkgroupStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceLogExportInput)(nil)).Elem(), NamespaceLogExport("useractivitylog"))
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceLogExportPtrInput)(nil)).Elem(), NamespaceLogExport("useractivitylog"))
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceLogExportArrayInput)(nil)).Elem(), NamespaceLogExportArray{})
	pulumi.RegisterOutputType(NamespaceLogExportOutput{})
	pulumi.RegisterOutputType(NamespaceLogExportPtrOutput{})
	pulumi.RegisterOutputType(NamespaceLogExportArrayOutput{})
	pulumi.RegisterOutputType(NamespaceStatusOutput{})
	pulumi.RegisterOutputType(NamespaceStatusPtrOutput{})
	pulumi.RegisterOutputType(WorkgroupStatusOutput{})
	pulumi.RegisterOutputType(WorkgroupStatusPtrOutput{})
}
