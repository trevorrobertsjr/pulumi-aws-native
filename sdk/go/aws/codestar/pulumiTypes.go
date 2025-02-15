// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestar

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type GitHubRepositoryCode struct {
	S3 GitHubRepositoryS3 `pulumi:"s3"`
}

// GitHubRepositoryCodeInput is an input type that accepts GitHubRepositoryCodeArgs and GitHubRepositoryCodeOutput values.
// You can construct a concrete instance of `GitHubRepositoryCodeInput` via:
//
//	GitHubRepositoryCodeArgs{...}
type GitHubRepositoryCodeInput interface {
	pulumi.Input

	ToGitHubRepositoryCodeOutput() GitHubRepositoryCodeOutput
	ToGitHubRepositoryCodeOutputWithContext(context.Context) GitHubRepositoryCodeOutput
}

type GitHubRepositoryCodeArgs struct {
	S3 GitHubRepositoryS3Input `pulumi:"s3"`
}

func (GitHubRepositoryCodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubRepositoryCode)(nil)).Elem()
}

func (i GitHubRepositoryCodeArgs) ToGitHubRepositoryCodeOutput() GitHubRepositoryCodeOutput {
	return i.ToGitHubRepositoryCodeOutputWithContext(context.Background())
}

func (i GitHubRepositoryCodeArgs) ToGitHubRepositoryCodeOutputWithContext(ctx context.Context) GitHubRepositoryCodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryCodeOutput)
}

func (i GitHubRepositoryCodeArgs) ToOutput(ctx context.Context) pulumix.Output[GitHubRepositoryCode] {
	return pulumix.Output[GitHubRepositoryCode]{
		OutputState: i.ToGitHubRepositoryCodeOutputWithContext(ctx).OutputState,
	}
}

func (i GitHubRepositoryCodeArgs) ToGitHubRepositoryCodePtrOutput() GitHubRepositoryCodePtrOutput {
	return i.ToGitHubRepositoryCodePtrOutputWithContext(context.Background())
}

func (i GitHubRepositoryCodeArgs) ToGitHubRepositoryCodePtrOutputWithContext(ctx context.Context) GitHubRepositoryCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryCodeOutput).ToGitHubRepositoryCodePtrOutputWithContext(ctx)
}

// GitHubRepositoryCodePtrInput is an input type that accepts GitHubRepositoryCodeArgs, GitHubRepositoryCodePtr and GitHubRepositoryCodePtrOutput values.
// You can construct a concrete instance of `GitHubRepositoryCodePtrInput` via:
//
//	        GitHubRepositoryCodeArgs{...}
//
//	or:
//
//	        nil
type GitHubRepositoryCodePtrInput interface {
	pulumi.Input

	ToGitHubRepositoryCodePtrOutput() GitHubRepositoryCodePtrOutput
	ToGitHubRepositoryCodePtrOutputWithContext(context.Context) GitHubRepositoryCodePtrOutput
}

type gitHubRepositoryCodePtrType GitHubRepositoryCodeArgs

func GitHubRepositoryCodePtr(v *GitHubRepositoryCodeArgs) GitHubRepositoryCodePtrInput {
	return (*gitHubRepositoryCodePtrType)(v)
}

func (*gitHubRepositoryCodePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubRepositoryCode)(nil)).Elem()
}

func (i *gitHubRepositoryCodePtrType) ToGitHubRepositoryCodePtrOutput() GitHubRepositoryCodePtrOutput {
	return i.ToGitHubRepositoryCodePtrOutputWithContext(context.Background())
}

func (i *gitHubRepositoryCodePtrType) ToGitHubRepositoryCodePtrOutputWithContext(ctx context.Context) GitHubRepositoryCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryCodePtrOutput)
}

func (i *gitHubRepositoryCodePtrType) ToOutput(ctx context.Context) pulumix.Output[*GitHubRepositoryCode] {
	return pulumix.Output[*GitHubRepositoryCode]{
		OutputState: i.ToGitHubRepositoryCodePtrOutputWithContext(ctx).OutputState,
	}
}

type GitHubRepositoryCodeOutput struct{ *pulumi.OutputState }

func (GitHubRepositoryCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubRepositoryCode)(nil)).Elem()
}

func (o GitHubRepositoryCodeOutput) ToGitHubRepositoryCodeOutput() GitHubRepositoryCodeOutput {
	return o
}

func (o GitHubRepositoryCodeOutput) ToGitHubRepositoryCodeOutputWithContext(ctx context.Context) GitHubRepositoryCodeOutput {
	return o
}

func (o GitHubRepositoryCodeOutput) ToGitHubRepositoryCodePtrOutput() GitHubRepositoryCodePtrOutput {
	return o.ToGitHubRepositoryCodePtrOutputWithContext(context.Background())
}

func (o GitHubRepositoryCodeOutput) ToGitHubRepositoryCodePtrOutputWithContext(ctx context.Context) GitHubRepositoryCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubRepositoryCode) *GitHubRepositoryCode {
		return &v
	}).(GitHubRepositoryCodePtrOutput)
}

func (o GitHubRepositoryCodeOutput) ToOutput(ctx context.Context) pulumix.Output[GitHubRepositoryCode] {
	return pulumix.Output[GitHubRepositoryCode]{
		OutputState: o.OutputState,
	}
}

func (o GitHubRepositoryCodeOutput) S3() GitHubRepositoryS3Output {
	return o.ApplyT(func(v GitHubRepositoryCode) GitHubRepositoryS3 { return v.S3 }).(GitHubRepositoryS3Output)
}

type GitHubRepositoryCodePtrOutput struct{ *pulumi.OutputState }

func (GitHubRepositoryCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubRepositoryCode)(nil)).Elem()
}

func (o GitHubRepositoryCodePtrOutput) ToGitHubRepositoryCodePtrOutput() GitHubRepositoryCodePtrOutput {
	return o
}

func (o GitHubRepositoryCodePtrOutput) ToGitHubRepositoryCodePtrOutputWithContext(ctx context.Context) GitHubRepositoryCodePtrOutput {
	return o
}

func (o GitHubRepositoryCodePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*GitHubRepositoryCode] {
	return pulumix.Output[*GitHubRepositoryCode]{
		OutputState: o.OutputState,
	}
}

func (o GitHubRepositoryCodePtrOutput) Elem() GitHubRepositoryCodeOutput {
	return o.ApplyT(func(v *GitHubRepositoryCode) GitHubRepositoryCode {
		if v != nil {
			return *v
		}
		var ret GitHubRepositoryCode
		return ret
	}).(GitHubRepositoryCodeOutput)
}

func (o GitHubRepositoryCodePtrOutput) S3() GitHubRepositoryS3PtrOutput {
	return o.ApplyT(func(v *GitHubRepositoryCode) *GitHubRepositoryS3 {
		if v == nil {
			return nil
		}
		return &v.S3
	}).(GitHubRepositoryS3PtrOutput)
}

type GitHubRepositoryS3 struct {
	Bucket        string  `pulumi:"bucket"`
	Key           string  `pulumi:"key"`
	ObjectVersion *string `pulumi:"objectVersion"`
}

// GitHubRepositoryS3Input is an input type that accepts GitHubRepositoryS3Args and GitHubRepositoryS3Output values.
// You can construct a concrete instance of `GitHubRepositoryS3Input` via:
//
//	GitHubRepositoryS3Args{...}
type GitHubRepositoryS3Input interface {
	pulumi.Input

	ToGitHubRepositoryS3Output() GitHubRepositoryS3Output
	ToGitHubRepositoryS3OutputWithContext(context.Context) GitHubRepositoryS3Output
}

type GitHubRepositoryS3Args struct {
	Bucket        pulumi.StringInput    `pulumi:"bucket"`
	Key           pulumi.StringInput    `pulumi:"key"`
	ObjectVersion pulumi.StringPtrInput `pulumi:"objectVersion"`
}

func (GitHubRepositoryS3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubRepositoryS3)(nil)).Elem()
}

func (i GitHubRepositoryS3Args) ToGitHubRepositoryS3Output() GitHubRepositoryS3Output {
	return i.ToGitHubRepositoryS3OutputWithContext(context.Background())
}

func (i GitHubRepositoryS3Args) ToGitHubRepositoryS3OutputWithContext(ctx context.Context) GitHubRepositoryS3Output {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryS3Output)
}

func (i GitHubRepositoryS3Args) ToOutput(ctx context.Context) pulumix.Output[GitHubRepositoryS3] {
	return pulumix.Output[GitHubRepositoryS3]{
		OutputState: i.ToGitHubRepositoryS3OutputWithContext(ctx).OutputState,
	}
}

func (i GitHubRepositoryS3Args) ToGitHubRepositoryS3PtrOutput() GitHubRepositoryS3PtrOutput {
	return i.ToGitHubRepositoryS3PtrOutputWithContext(context.Background())
}

func (i GitHubRepositoryS3Args) ToGitHubRepositoryS3PtrOutputWithContext(ctx context.Context) GitHubRepositoryS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryS3Output).ToGitHubRepositoryS3PtrOutputWithContext(ctx)
}

// GitHubRepositoryS3PtrInput is an input type that accepts GitHubRepositoryS3Args, GitHubRepositoryS3Ptr and GitHubRepositoryS3PtrOutput values.
// You can construct a concrete instance of `GitHubRepositoryS3PtrInput` via:
//
//	        GitHubRepositoryS3Args{...}
//
//	or:
//
//	        nil
type GitHubRepositoryS3PtrInput interface {
	pulumi.Input

	ToGitHubRepositoryS3PtrOutput() GitHubRepositoryS3PtrOutput
	ToGitHubRepositoryS3PtrOutputWithContext(context.Context) GitHubRepositoryS3PtrOutput
}

type gitHubRepositoryS3PtrType GitHubRepositoryS3Args

func GitHubRepositoryS3Ptr(v *GitHubRepositoryS3Args) GitHubRepositoryS3PtrInput {
	return (*gitHubRepositoryS3PtrType)(v)
}

func (*gitHubRepositoryS3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubRepositoryS3)(nil)).Elem()
}

func (i *gitHubRepositoryS3PtrType) ToGitHubRepositoryS3PtrOutput() GitHubRepositoryS3PtrOutput {
	return i.ToGitHubRepositoryS3PtrOutputWithContext(context.Background())
}

func (i *gitHubRepositoryS3PtrType) ToGitHubRepositoryS3PtrOutputWithContext(ctx context.Context) GitHubRepositoryS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryS3PtrOutput)
}

func (i *gitHubRepositoryS3PtrType) ToOutput(ctx context.Context) pulumix.Output[*GitHubRepositoryS3] {
	return pulumix.Output[*GitHubRepositoryS3]{
		OutputState: i.ToGitHubRepositoryS3PtrOutputWithContext(ctx).OutputState,
	}
}

type GitHubRepositoryS3Output struct{ *pulumi.OutputState }

func (GitHubRepositoryS3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubRepositoryS3)(nil)).Elem()
}

func (o GitHubRepositoryS3Output) ToGitHubRepositoryS3Output() GitHubRepositoryS3Output {
	return o
}

func (o GitHubRepositoryS3Output) ToGitHubRepositoryS3OutputWithContext(ctx context.Context) GitHubRepositoryS3Output {
	return o
}

func (o GitHubRepositoryS3Output) ToGitHubRepositoryS3PtrOutput() GitHubRepositoryS3PtrOutput {
	return o.ToGitHubRepositoryS3PtrOutputWithContext(context.Background())
}

func (o GitHubRepositoryS3Output) ToGitHubRepositoryS3PtrOutputWithContext(ctx context.Context) GitHubRepositoryS3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubRepositoryS3) *GitHubRepositoryS3 {
		return &v
	}).(GitHubRepositoryS3PtrOutput)
}

func (o GitHubRepositoryS3Output) ToOutput(ctx context.Context) pulumix.Output[GitHubRepositoryS3] {
	return pulumix.Output[GitHubRepositoryS3]{
		OutputState: o.OutputState,
	}
}

func (o GitHubRepositoryS3Output) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GitHubRepositoryS3) string { return v.Bucket }).(pulumi.StringOutput)
}

func (o GitHubRepositoryS3Output) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GitHubRepositoryS3) string { return v.Key }).(pulumi.StringOutput)
}

func (o GitHubRepositoryS3Output) ObjectVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubRepositoryS3) *string { return v.ObjectVersion }).(pulumi.StringPtrOutput)
}

type GitHubRepositoryS3PtrOutput struct{ *pulumi.OutputState }

func (GitHubRepositoryS3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubRepositoryS3)(nil)).Elem()
}

func (o GitHubRepositoryS3PtrOutput) ToGitHubRepositoryS3PtrOutput() GitHubRepositoryS3PtrOutput {
	return o
}

func (o GitHubRepositoryS3PtrOutput) ToGitHubRepositoryS3PtrOutputWithContext(ctx context.Context) GitHubRepositoryS3PtrOutput {
	return o
}

func (o GitHubRepositoryS3PtrOutput) ToOutput(ctx context.Context) pulumix.Output[*GitHubRepositoryS3] {
	return pulumix.Output[*GitHubRepositoryS3]{
		OutputState: o.OutputState,
	}
}

func (o GitHubRepositoryS3PtrOutput) Elem() GitHubRepositoryS3Output {
	return o.ApplyT(func(v *GitHubRepositoryS3) GitHubRepositoryS3 {
		if v != nil {
			return *v
		}
		var ret GitHubRepositoryS3
		return ret
	}).(GitHubRepositoryS3Output)
}

func (o GitHubRepositoryS3PtrOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubRepositoryS3) *string {
		if v == nil {
			return nil
		}
		return &v.Bucket
	}).(pulumi.StringPtrOutput)
}

func (o GitHubRepositoryS3PtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubRepositoryS3) *string {
		if v == nil {
			return nil
		}
		return &v.Key
	}).(pulumi.StringPtrOutput)
}

func (o GitHubRepositoryS3PtrOutput) ObjectVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubRepositoryS3) *string {
		if v == nil {
			return nil
		}
		return v.ObjectVersion
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitHubRepositoryCodeInput)(nil)).Elem(), GitHubRepositoryCodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitHubRepositoryCodePtrInput)(nil)).Elem(), GitHubRepositoryCodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitHubRepositoryS3Input)(nil)).Elem(), GitHubRepositoryS3Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitHubRepositoryS3PtrInput)(nil)).Elem(), GitHubRepositoryS3Args{})
	pulumi.RegisterOutputType(GitHubRepositoryCodeOutput{})
	pulumi.RegisterOutputType(GitHubRepositoryCodePtrOutput{})
	pulumi.RegisterOutputType(GitHubRepositoryS3Output{})
	pulumi.RegisterOutputType(GitHubRepositoryS3PtrOutput{})
}
