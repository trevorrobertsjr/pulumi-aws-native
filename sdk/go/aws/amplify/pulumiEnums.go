// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type AppAutoBranchCreationConfigStage string

const (
	AppAutoBranchCreationConfigStageExperimental = AppAutoBranchCreationConfigStage("EXPERIMENTAL")
	AppAutoBranchCreationConfigStageBeta         = AppAutoBranchCreationConfigStage("BETA")
	AppAutoBranchCreationConfigStagePullRequest  = AppAutoBranchCreationConfigStage("PULL_REQUEST")
	AppAutoBranchCreationConfigStageProduction   = AppAutoBranchCreationConfigStage("PRODUCTION")
	AppAutoBranchCreationConfigStageDevelopment  = AppAutoBranchCreationConfigStage("DEVELOPMENT")
)

func (AppAutoBranchCreationConfigStage) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAutoBranchCreationConfigStage)(nil)).Elem()
}

func (e AppAutoBranchCreationConfigStage) ToAppAutoBranchCreationConfigStageOutput() AppAutoBranchCreationConfigStageOutput {
	return pulumi.ToOutput(e).(AppAutoBranchCreationConfigStageOutput)
}

func (e AppAutoBranchCreationConfigStage) ToAppAutoBranchCreationConfigStageOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigStageOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppAutoBranchCreationConfigStageOutput)
}

func (e AppAutoBranchCreationConfigStage) ToAppAutoBranchCreationConfigStagePtrOutput() AppAutoBranchCreationConfigStagePtrOutput {
	return e.ToAppAutoBranchCreationConfigStagePtrOutputWithContext(context.Background())
}

func (e AppAutoBranchCreationConfigStage) ToAppAutoBranchCreationConfigStagePtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigStagePtrOutput {
	return AppAutoBranchCreationConfigStage(e).ToAppAutoBranchCreationConfigStageOutputWithContext(ctx).ToAppAutoBranchCreationConfigStagePtrOutputWithContext(ctx)
}

func (e AppAutoBranchCreationConfigStage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppAutoBranchCreationConfigStage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppAutoBranchCreationConfigStage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppAutoBranchCreationConfigStage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppAutoBranchCreationConfigStageOutput struct{ *pulumi.OutputState }

func (AppAutoBranchCreationConfigStageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAutoBranchCreationConfigStage)(nil)).Elem()
}

func (o AppAutoBranchCreationConfigStageOutput) ToAppAutoBranchCreationConfigStageOutput() AppAutoBranchCreationConfigStageOutput {
	return o
}

func (o AppAutoBranchCreationConfigStageOutput) ToAppAutoBranchCreationConfigStageOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigStageOutput {
	return o
}

func (o AppAutoBranchCreationConfigStageOutput) ToAppAutoBranchCreationConfigStagePtrOutput() AppAutoBranchCreationConfigStagePtrOutput {
	return o.ToAppAutoBranchCreationConfigStagePtrOutputWithContext(context.Background())
}

func (o AppAutoBranchCreationConfigStageOutput) ToAppAutoBranchCreationConfigStagePtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigStagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppAutoBranchCreationConfigStage) *AppAutoBranchCreationConfigStage {
		return &v
	}).(AppAutoBranchCreationConfigStagePtrOutput)
}

func (o AppAutoBranchCreationConfigStageOutput) ToOutput(ctx context.Context) pulumix.Output[AppAutoBranchCreationConfigStage] {
	return pulumix.Output[AppAutoBranchCreationConfigStage]{
		OutputState: o.OutputState,
	}
}

func (o AppAutoBranchCreationConfigStageOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppAutoBranchCreationConfigStageOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppAutoBranchCreationConfigStage) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppAutoBranchCreationConfigStageOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppAutoBranchCreationConfigStageOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppAutoBranchCreationConfigStage) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppAutoBranchCreationConfigStagePtrOutput struct{ *pulumi.OutputState }

func (AppAutoBranchCreationConfigStagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAutoBranchCreationConfigStage)(nil)).Elem()
}

func (o AppAutoBranchCreationConfigStagePtrOutput) ToAppAutoBranchCreationConfigStagePtrOutput() AppAutoBranchCreationConfigStagePtrOutput {
	return o
}

func (o AppAutoBranchCreationConfigStagePtrOutput) ToAppAutoBranchCreationConfigStagePtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigStagePtrOutput {
	return o
}

func (o AppAutoBranchCreationConfigStagePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AppAutoBranchCreationConfigStage] {
	return pulumix.Output[*AppAutoBranchCreationConfigStage]{
		OutputState: o.OutputState,
	}
}

func (o AppAutoBranchCreationConfigStagePtrOutput) Elem() AppAutoBranchCreationConfigStageOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfigStage) AppAutoBranchCreationConfigStage {
		if v != nil {
			return *v
		}
		var ret AppAutoBranchCreationConfigStage
		return ret
	}).(AppAutoBranchCreationConfigStageOutput)
}

func (o AppAutoBranchCreationConfigStagePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppAutoBranchCreationConfigStagePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppAutoBranchCreationConfigStage) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AppAutoBranchCreationConfigStageInput is an input type that accepts AppAutoBranchCreationConfigStageArgs and AppAutoBranchCreationConfigStageOutput values.
// You can construct a concrete instance of `AppAutoBranchCreationConfigStageInput` via:
//
//	AppAutoBranchCreationConfigStageArgs{...}
type AppAutoBranchCreationConfigStageInput interface {
	pulumi.Input

	ToAppAutoBranchCreationConfigStageOutput() AppAutoBranchCreationConfigStageOutput
	ToAppAutoBranchCreationConfigStageOutputWithContext(context.Context) AppAutoBranchCreationConfigStageOutput
}

var appAutoBranchCreationConfigStagePtrType = reflect.TypeOf((**AppAutoBranchCreationConfigStage)(nil)).Elem()

type AppAutoBranchCreationConfigStagePtrInput interface {
	pulumi.Input

	ToAppAutoBranchCreationConfigStagePtrOutput() AppAutoBranchCreationConfigStagePtrOutput
	ToAppAutoBranchCreationConfigStagePtrOutputWithContext(context.Context) AppAutoBranchCreationConfigStagePtrOutput
}

type appAutoBranchCreationConfigStagePtr string

func AppAutoBranchCreationConfigStagePtr(v string) AppAutoBranchCreationConfigStagePtrInput {
	return (*appAutoBranchCreationConfigStagePtr)(&v)
}

func (*appAutoBranchCreationConfigStagePtr) ElementType() reflect.Type {
	return appAutoBranchCreationConfigStagePtrType
}

func (in *appAutoBranchCreationConfigStagePtr) ToAppAutoBranchCreationConfigStagePtrOutput() AppAutoBranchCreationConfigStagePtrOutput {
	return pulumi.ToOutput(in).(AppAutoBranchCreationConfigStagePtrOutput)
}

func (in *appAutoBranchCreationConfigStagePtr) ToAppAutoBranchCreationConfigStagePtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigStagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppAutoBranchCreationConfigStagePtrOutput)
}

func (in *appAutoBranchCreationConfigStagePtr) ToOutput(ctx context.Context) pulumix.Output[*AppAutoBranchCreationConfigStage] {
	return pulumix.Output[*AppAutoBranchCreationConfigStage]{
		OutputState: in.ToAppAutoBranchCreationConfigStagePtrOutputWithContext(ctx).OutputState,
	}
}

type AppPlatform string

const (
	AppPlatformWeb        = AppPlatform("WEB")
	AppPlatformWebDynamic = AppPlatform("WEB_DYNAMIC")
	AppPlatformWebCompute = AppPlatform("WEB_COMPUTE")
)

func (AppPlatform) ElementType() reflect.Type {
	return reflect.TypeOf((*AppPlatform)(nil)).Elem()
}

func (e AppPlatform) ToAppPlatformOutput() AppPlatformOutput {
	return pulumi.ToOutput(e).(AppPlatformOutput)
}

func (e AppPlatform) ToAppPlatformOutputWithContext(ctx context.Context) AppPlatformOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppPlatformOutput)
}

func (e AppPlatform) ToAppPlatformPtrOutput() AppPlatformPtrOutput {
	return e.ToAppPlatformPtrOutputWithContext(context.Background())
}

func (e AppPlatform) ToAppPlatformPtrOutputWithContext(ctx context.Context) AppPlatformPtrOutput {
	return AppPlatform(e).ToAppPlatformOutputWithContext(ctx).ToAppPlatformPtrOutputWithContext(ctx)
}

func (e AppPlatform) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppPlatform) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppPlatform) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppPlatform) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppPlatformOutput struct{ *pulumi.OutputState }

func (AppPlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppPlatform)(nil)).Elem()
}

func (o AppPlatformOutput) ToAppPlatformOutput() AppPlatformOutput {
	return o
}

func (o AppPlatformOutput) ToAppPlatformOutputWithContext(ctx context.Context) AppPlatformOutput {
	return o
}

func (o AppPlatformOutput) ToAppPlatformPtrOutput() AppPlatformPtrOutput {
	return o.ToAppPlatformPtrOutputWithContext(context.Background())
}

func (o AppPlatformOutput) ToAppPlatformPtrOutputWithContext(ctx context.Context) AppPlatformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppPlatform) *AppPlatform {
		return &v
	}).(AppPlatformPtrOutput)
}

func (o AppPlatformOutput) ToOutput(ctx context.Context) pulumix.Output[AppPlatform] {
	return pulumix.Output[AppPlatform]{
		OutputState: o.OutputState,
	}
}

func (o AppPlatformOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppPlatformOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppPlatform) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppPlatformOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppPlatformOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppPlatform) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppPlatformPtrOutput struct{ *pulumi.OutputState }

func (AppPlatformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppPlatform)(nil)).Elem()
}

func (o AppPlatformPtrOutput) ToAppPlatformPtrOutput() AppPlatformPtrOutput {
	return o
}

func (o AppPlatformPtrOutput) ToAppPlatformPtrOutputWithContext(ctx context.Context) AppPlatformPtrOutput {
	return o
}

func (o AppPlatformPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AppPlatform] {
	return pulumix.Output[*AppPlatform]{
		OutputState: o.OutputState,
	}
}

func (o AppPlatformPtrOutput) Elem() AppPlatformOutput {
	return o.ApplyT(func(v *AppPlatform) AppPlatform {
		if v != nil {
			return *v
		}
		var ret AppPlatform
		return ret
	}).(AppPlatformOutput)
}

func (o AppPlatformPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppPlatformPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppPlatform) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AppPlatformInput is an input type that accepts AppPlatformArgs and AppPlatformOutput values.
// You can construct a concrete instance of `AppPlatformInput` via:
//
//	AppPlatformArgs{...}
type AppPlatformInput interface {
	pulumi.Input

	ToAppPlatformOutput() AppPlatformOutput
	ToAppPlatformOutputWithContext(context.Context) AppPlatformOutput
}

var appPlatformPtrType = reflect.TypeOf((**AppPlatform)(nil)).Elem()

type AppPlatformPtrInput interface {
	pulumi.Input

	ToAppPlatformPtrOutput() AppPlatformPtrOutput
	ToAppPlatformPtrOutputWithContext(context.Context) AppPlatformPtrOutput
}

type appPlatformPtr string

func AppPlatformPtr(v string) AppPlatformPtrInput {
	return (*appPlatformPtr)(&v)
}

func (*appPlatformPtr) ElementType() reflect.Type {
	return appPlatformPtrType
}

func (in *appPlatformPtr) ToAppPlatformPtrOutput() AppPlatformPtrOutput {
	return pulumi.ToOutput(in).(AppPlatformPtrOutput)
}

func (in *appPlatformPtr) ToAppPlatformPtrOutputWithContext(ctx context.Context) AppPlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppPlatformPtrOutput)
}

func (in *appPlatformPtr) ToOutput(ctx context.Context) pulumix.Output[*AppPlatform] {
	return pulumix.Output[*AppPlatform]{
		OutputState: in.ToAppPlatformPtrOutputWithContext(ctx).OutputState,
	}
}

type BranchStage string

const (
	BranchStageExperimental = BranchStage("EXPERIMENTAL")
	BranchStageBeta         = BranchStage("BETA")
	BranchStagePullRequest  = BranchStage("PULL_REQUEST")
	BranchStageProduction   = BranchStage("PRODUCTION")
	BranchStageDevelopment  = BranchStage("DEVELOPMENT")
)

func (BranchStage) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchStage)(nil)).Elem()
}

func (e BranchStage) ToBranchStageOutput() BranchStageOutput {
	return pulumi.ToOutput(e).(BranchStageOutput)
}

func (e BranchStage) ToBranchStageOutputWithContext(ctx context.Context) BranchStageOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BranchStageOutput)
}

func (e BranchStage) ToBranchStagePtrOutput() BranchStagePtrOutput {
	return e.ToBranchStagePtrOutputWithContext(context.Background())
}

func (e BranchStage) ToBranchStagePtrOutputWithContext(ctx context.Context) BranchStagePtrOutput {
	return BranchStage(e).ToBranchStageOutputWithContext(ctx).ToBranchStagePtrOutputWithContext(ctx)
}

func (e BranchStage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BranchStage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BranchStage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BranchStage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BranchStageOutput struct{ *pulumi.OutputState }

func (BranchStageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchStage)(nil)).Elem()
}

func (o BranchStageOutput) ToBranchStageOutput() BranchStageOutput {
	return o
}

func (o BranchStageOutput) ToBranchStageOutputWithContext(ctx context.Context) BranchStageOutput {
	return o
}

func (o BranchStageOutput) ToBranchStagePtrOutput() BranchStagePtrOutput {
	return o.ToBranchStagePtrOutputWithContext(context.Background())
}

func (o BranchStageOutput) ToBranchStagePtrOutputWithContext(ctx context.Context) BranchStagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BranchStage) *BranchStage {
		return &v
	}).(BranchStagePtrOutput)
}

func (o BranchStageOutput) ToOutput(ctx context.Context) pulumix.Output[BranchStage] {
	return pulumix.Output[BranchStage]{
		OutputState: o.OutputState,
	}
}

func (o BranchStageOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BranchStageOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BranchStage) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BranchStageOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BranchStageOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BranchStage) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BranchStagePtrOutput struct{ *pulumi.OutputState }

func (BranchStagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchStage)(nil)).Elem()
}

func (o BranchStagePtrOutput) ToBranchStagePtrOutput() BranchStagePtrOutput {
	return o
}

func (o BranchStagePtrOutput) ToBranchStagePtrOutputWithContext(ctx context.Context) BranchStagePtrOutput {
	return o
}

func (o BranchStagePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*BranchStage] {
	return pulumix.Output[*BranchStage]{
		OutputState: o.OutputState,
	}
}

func (o BranchStagePtrOutput) Elem() BranchStageOutput {
	return o.ApplyT(func(v *BranchStage) BranchStage {
		if v != nil {
			return *v
		}
		var ret BranchStage
		return ret
	}).(BranchStageOutput)
}

func (o BranchStagePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BranchStagePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BranchStage) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BranchStageInput is an input type that accepts BranchStageArgs and BranchStageOutput values.
// You can construct a concrete instance of `BranchStageInput` via:
//
//	BranchStageArgs{...}
type BranchStageInput interface {
	pulumi.Input

	ToBranchStageOutput() BranchStageOutput
	ToBranchStageOutputWithContext(context.Context) BranchStageOutput
}

var branchStagePtrType = reflect.TypeOf((**BranchStage)(nil)).Elem()

type BranchStagePtrInput interface {
	pulumi.Input

	ToBranchStagePtrOutput() BranchStagePtrOutput
	ToBranchStagePtrOutputWithContext(context.Context) BranchStagePtrOutput
}

type branchStagePtr string

func BranchStagePtr(v string) BranchStagePtrInput {
	return (*branchStagePtr)(&v)
}

func (*branchStagePtr) ElementType() reflect.Type {
	return branchStagePtrType
}

func (in *branchStagePtr) ToBranchStagePtrOutput() BranchStagePtrOutput {
	return pulumi.ToOutput(in).(BranchStagePtrOutput)
}

func (in *branchStagePtr) ToBranchStagePtrOutputWithContext(ctx context.Context) BranchStagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BranchStagePtrOutput)
}

func (in *branchStagePtr) ToOutput(ctx context.Context) pulumix.Output[*BranchStage] {
	return pulumix.Output[*BranchStage]{
		OutputState: in.ToBranchStagePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppAutoBranchCreationConfigStageInput)(nil)).Elem(), AppAutoBranchCreationConfigStage("EXPERIMENTAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppAutoBranchCreationConfigStagePtrInput)(nil)).Elem(), AppAutoBranchCreationConfigStage("EXPERIMENTAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppPlatformInput)(nil)).Elem(), AppPlatform("WEB"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppPlatformPtrInput)(nil)).Elem(), AppPlatform("WEB"))
	pulumi.RegisterInputType(reflect.TypeOf((*BranchStageInput)(nil)).Elem(), BranchStage("EXPERIMENTAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*BranchStagePtrInput)(nil)).Elem(), BranchStage("EXPERIMENTAL"))
	pulumi.RegisterOutputType(AppAutoBranchCreationConfigStageOutput{})
	pulumi.RegisterOutputType(AppAutoBranchCreationConfigStagePtrOutput{})
	pulumi.RegisterOutputType(AppPlatformOutput{})
	pulumi.RegisterOutputType(AppPlatformPtrOutput{})
	pulumi.RegisterOutputType(BranchStageOutput{})
	pulumi.RegisterOutputType(BranchStagePtrOutput{})
}
