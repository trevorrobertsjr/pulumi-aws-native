// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.
type AutoScalingConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of this auto scaling configuration.
	AutoScalingConfigurationArn pulumi.StringOutput `pulumi:"autoScalingConfigurationArn"`
	// The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
	AutoScalingConfigurationName pulumi.StringPtrOutput `pulumi:"autoScalingConfigurationName"`
	// The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName.
	AutoScalingConfigurationRevision pulumi.IntOutput `pulumi:"autoScalingConfigurationRevision"`
	// It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
	Latest pulumi.BoolOutput `pulumi:"latest"`
	// The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
	MaxConcurrency pulumi.IntPtrOutput `pulumi:"maxConcurrency"`
	// The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
	MaxSize pulumi.IntPtrOutput `pulumi:"maxSize"`
	// The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
	MinSize pulumi.IntPtrOutput `pulumi:"minSize"`
	// A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
	Tags AutoScalingConfigurationTagArrayOutput `pulumi:"tags"`
}

// NewAutoScalingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAutoScalingConfiguration(ctx *pulumi.Context,
	name string, args *AutoScalingConfigurationArgs, opts ...pulumi.ResourceOption) (*AutoScalingConfiguration, error) {
	if args == nil {
		args = &AutoScalingConfigurationArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"autoScalingConfigurationName",
		"maxConcurrency",
		"maxSize",
		"minSize",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoScalingConfiguration
	err := ctx.RegisterResource("aws-native:apprunner:AutoScalingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoScalingConfiguration gets an existing AutoScalingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoScalingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoScalingConfigurationState, opts ...pulumi.ResourceOption) (*AutoScalingConfiguration, error) {
	var resource AutoScalingConfiguration
	err := ctx.ReadResource("aws-native:apprunner:AutoScalingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoScalingConfiguration resources.
type autoScalingConfigurationState struct {
}

type AutoScalingConfigurationState struct {
}

func (AutoScalingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScalingConfigurationState)(nil)).Elem()
}

type autoScalingConfigurationArgs struct {
	// The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
	AutoScalingConfigurationName *string `pulumi:"autoScalingConfigurationName"`
	// The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
	MaxConcurrency *int `pulumi:"maxConcurrency"`
	// The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
	MinSize *int `pulumi:"minSize"`
	// A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
	Tags []AutoScalingConfigurationTag `pulumi:"tags"`
}

// The set of arguments for constructing a AutoScalingConfiguration resource.
type AutoScalingConfigurationArgs struct {
	// The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
	AutoScalingConfigurationName pulumi.StringPtrInput
	// The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
	MaxConcurrency pulumi.IntPtrInput
	// The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
	MaxSize pulumi.IntPtrInput
	// The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
	MinSize pulumi.IntPtrInput
	// A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
	Tags AutoScalingConfigurationTagArrayInput
}

func (AutoScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScalingConfigurationArgs)(nil)).Elem()
}

type AutoScalingConfigurationInput interface {
	pulumi.Input

	ToAutoScalingConfigurationOutput() AutoScalingConfigurationOutput
	ToAutoScalingConfigurationOutputWithContext(ctx context.Context) AutoScalingConfigurationOutput
}

func (*AutoScalingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingConfiguration)(nil)).Elem()
}

func (i *AutoScalingConfiguration) ToAutoScalingConfigurationOutput() AutoScalingConfigurationOutput {
	return i.ToAutoScalingConfigurationOutputWithContext(context.Background())
}

func (i *AutoScalingConfiguration) ToAutoScalingConfigurationOutputWithContext(ctx context.Context) AutoScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalingConfigurationOutput)
}

func (i *AutoScalingConfiguration) ToOutput(ctx context.Context) pulumix.Output[*AutoScalingConfiguration] {
	return pulumix.Output[*AutoScalingConfiguration]{
		OutputState: i.ToAutoScalingConfigurationOutputWithContext(ctx).OutputState,
	}
}

type AutoScalingConfigurationOutput struct{ *pulumi.OutputState }

func (AutoScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingConfiguration)(nil)).Elem()
}

func (o AutoScalingConfigurationOutput) ToAutoScalingConfigurationOutput() AutoScalingConfigurationOutput {
	return o
}

func (o AutoScalingConfigurationOutput) ToAutoScalingConfigurationOutputWithContext(ctx context.Context) AutoScalingConfigurationOutput {
	return o
}

func (o AutoScalingConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*AutoScalingConfiguration] {
	return pulumix.Output[*AutoScalingConfiguration]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of this auto scaling configuration.
func (o AutoScalingConfigurationOutput) AutoScalingConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.StringOutput { return v.AutoScalingConfigurationArn }).(pulumi.StringOutput)
}

// The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
func (o AutoScalingConfigurationOutput) AutoScalingConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.StringPtrOutput { return v.AutoScalingConfigurationName }).(pulumi.StringPtrOutput)
}

// The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName.
func (o AutoScalingConfigurationOutput) AutoScalingConfigurationRevision() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.IntOutput { return v.AutoScalingConfigurationRevision }).(pulumi.IntOutput)
}

// It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
func (o AutoScalingConfigurationOutput) Latest() pulumi.BoolOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.BoolOutput { return v.Latest }).(pulumi.BoolOutput)
}

// The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
func (o AutoScalingConfigurationOutput) MaxConcurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.IntPtrOutput { return v.MaxConcurrency }).(pulumi.IntPtrOutput)
}

// The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.
func (o AutoScalingConfigurationOutput) MaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.IntPtrOutput { return v.MaxSize }).(pulumi.IntPtrOutput)
}

// The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
func (o AutoScalingConfigurationOutput) MinSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) pulumi.IntPtrOutput { return v.MinSize }).(pulumi.IntPtrOutput)
}

// A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.
func (o AutoScalingConfigurationOutput) Tags() AutoScalingConfigurationTagArrayOutput {
	return o.ApplyT(func(v *AutoScalingConfiguration) AutoScalingConfigurationTagArrayOutput { return v.Tags }).(AutoScalingConfigurationTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoScalingConfigurationInput)(nil)).Elem(), &AutoScalingConfiguration{})
	pulumi.RegisterOutputType(AutoScalingConfigurationOutput{})
}
