// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This schema is for testing purpose only.
type RobotApplication struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The revision ID of robot application.
	CurrentRevisionId pulumi.StringPtrOutput `pulumi:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// The name of the robot application.
	Name               pulumi.StringPtrOutput                   `pulumi:"name"`
	RobotSoftwareSuite RobotApplicationRobotSoftwareSuiteOutput `pulumi:"robotSoftwareSuite"`
	// The sources of the robot application.
	Sources RobotApplicationSourceConfigArrayOutput `pulumi:"sources"`
	Tags    RobotApplicationTagsPtrOutput           `pulumi:"tags"`
}

// NewRobotApplication registers a new resource with the given unique name, arguments, and options.
func NewRobotApplication(ctx *pulumi.Context,
	name string, args *RobotApplicationArgs, opts ...pulumi.ResourceOption) (*RobotApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RobotSoftwareSuite == nil {
		return nil, errors.New("invalid value for required argument 'RobotSoftwareSuite'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RobotApplication
	err := ctx.RegisterResource("aws-native:robomaker:RobotApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRobotApplication gets an existing RobotApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRobotApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RobotApplicationState, opts ...pulumi.ResourceOption) (*RobotApplication, error) {
	var resource RobotApplication
	err := ctx.ReadResource("aws-native:robomaker:RobotApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RobotApplication resources.
type robotApplicationState struct {
}

type RobotApplicationState struct {
}

func (RobotApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationState)(nil)).Elem()
}

type robotApplicationArgs struct {
	// The revision ID of robot application.
	CurrentRevisionId *string `pulumi:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	Environment *string `pulumi:"environment"`
	// The name of the robot application.
	Name               *string                            `pulumi:"name"`
	RobotSoftwareSuite RobotApplicationRobotSoftwareSuite `pulumi:"robotSoftwareSuite"`
	// The sources of the robot application.
	Sources []RobotApplicationSourceConfig `pulumi:"sources"`
	Tags    *RobotApplicationTags          `pulumi:"tags"`
}

// The set of arguments for constructing a RobotApplication resource.
type RobotApplicationArgs struct {
	// The revision ID of robot application.
	CurrentRevisionId pulumi.StringPtrInput
	// The URI of the Docker image for the robot application.
	Environment pulumi.StringPtrInput
	// The name of the robot application.
	Name               pulumi.StringPtrInput
	RobotSoftwareSuite RobotApplicationRobotSoftwareSuiteInput
	// The sources of the robot application.
	Sources RobotApplicationSourceConfigArrayInput
	Tags    RobotApplicationTagsPtrInput
}

func (RobotApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationArgs)(nil)).Elem()
}

type RobotApplicationInput interface {
	pulumi.Input

	ToRobotApplicationOutput() RobotApplicationOutput
	ToRobotApplicationOutputWithContext(ctx context.Context) RobotApplicationOutput
}

func (*RobotApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**RobotApplication)(nil)).Elem()
}

func (i *RobotApplication) ToRobotApplicationOutput() RobotApplicationOutput {
	return i.ToRobotApplicationOutputWithContext(context.Background())
}

func (i *RobotApplication) ToRobotApplicationOutputWithContext(ctx context.Context) RobotApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotApplicationOutput)
}

func (i *RobotApplication) ToOutput(ctx context.Context) pulumix.Output[*RobotApplication] {
	return pulumix.Output[*RobotApplication]{
		OutputState: i.ToRobotApplicationOutputWithContext(ctx).OutputState,
	}
}

type RobotApplicationOutput struct{ *pulumi.OutputState }

func (RobotApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RobotApplication)(nil)).Elem()
}

func (o RobotApplicationOutput) ToRobotApplicationOutput() RobotApplicationOutput {
	return o
}

func (o RobotApplicationOutput) ToRobotApplicationOutputWithContext(ctx context.Context) RobotApplicationOutput {
	return o
}

func (o RobotApplicationOutput) ToOutput(ctx context.Context) pulumix.Output[*RobotApplication] {
	return pulumix.Output[*RobotApplication]{
		OutputState: o.OutputState,
	}
}

func (o RobotApplicationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RobotApplication) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The revision ID of robot application.
func (o RobotApplicationOutput) CurrentRevisionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RobotApplication) pulumi.StringPtrOutput { return v.CurrentRevisionId }).(pulumi.StringPtrOutput)
}

// The URI of the Docker image for the robot application.
func (o RobotApplicationOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RobotApplication) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// The name of the robot application.
func (o RobotApplicationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RobotApplication) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RobotApplicationOutput) RobotSoftwareSuite() RobotApplicationRobotSoftwareSuiteOutput {
	return o.ApplyT(func(v *RobotApplication) RobotApplicationRobotSoftwareSuiteOutput { return v.RobotSoftwareSuite }).(RobotApplicationRobotSoftwareSuiteOutput)
}

// The sources of the robot application.
func (o RobotApplicationOutput) Sources() RobotApplicationSourceConfigArrayOutput {
	return o.ApplyT(func(v *RobotApplication) RobotApplicationSourceConfigArrayOutput { return v.Sources }).(RobotApplicationSourceConfigArrayOutput)
}

func (o RobotApplicationOutput) Tags() RobotApplicationTagsPtrOutput {
	return o.ApplyT(func(v *RobotApplication) RobotApplicationTagsPtrOutput { return v.Tags }).(RobotApplicationTagsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RobotApplicationInput)(nil)).Elem(), &RobotApplication{})
	pulumi.RegisterOutputType(RobotApplicationOutput{})
}
