// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applicationautoscaling

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::ApplicationAutoScaling::ScalingPolicy
type ScalingPolicy struct {
	pulumi.CustomResourceState

	// ARN is a read only property for the resource.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing AWS::ApplicationAutoScaling::ScalingPolicy resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// TargetTrackingScaling Not supported for Amazon EMR
	//
	// StepScaling Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
	ScalableDimension pulumi.StringPtrOutput `pulumi:"scalableDimension"`
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the AWS::ApplicationAutoScaling::ScalableTarget resource.
	ScalingTargetId pulumi.StringPtrOutput `pulumi:"scalingTargetId"`
	// The namespace of the AWS service that provides the resource, or a custom-resource.
	ServiceNamespace pulumi.StringPtrOutput `pulumi:"serviceNamespace"`
	// A step scaling policy.
	StepScalingPolicyConfiguration ScalingPolicyStepScalingPolicyConfigurationPtrOutput `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// NewScalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewScalingPolicy(ctx *pulumi.Context,
	name string, args *ScalingPolicyArgs, opts ...pulumi.ResourceOption) (*ScalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"policyName",
		"resourceId",
		"scalableDimension",
		"scalingTargetId",
		"serviceNamespace",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScalingPolicy
	err := ctx.RegisterResource("aws-native:applicationautoscaling:ScalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPolicy gets an existing ScalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPolicyState, opts ...pulumi.ResourceOption) (*ScalingPolicy, error) {
	var resource ScalingPolicy
	err := ctx.ReadResource("aws-native:applicationautoscaling:ScalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPolicy resources.
type scalingPolicyState struct {
}

type ScalingPolicyState struct {
}

func (ScalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPolicyState)(nil)).Elem()
}

type scalingPolicyArgs struct {
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing AWS::ApplicationAutoScaling::ScalingPolicy resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	PolicyName string `pulumi:"policyName"`
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// TargetTrackingScaling Not supported for Amazon EMR
	//
	// StepScaling Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	PolicyType string `pulumi:"policyType"`
	// The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.
	ResourceId *string `pulumi:"resourceId"`
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
	ScalableDimension *string `pulumi:"scalableDimension"`
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the AWS::ApplicationAutoScaling::ScalableTarget resource.
	ScalingTargetId *string `pulumi:"scalingTargetId"`
	// The namespace of the AWS service that provides the resource, or a custom-resource.
	ServiceNamespace *string `pulumi:"serviceNamespace"`
	// A step scaling policy.
	StepScalingPolicyConfiguration *ScalingPolicyStepScalingPolicyConfiguration `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration *ScalingPolicyTargetTrackingScalingPolicyConfiguration `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// The set of arguments for constructing a ScalingPolicy resource.
type ScalingPolicyArgs struct {
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing AWS::ApplicationAutoScaling::ScalingPolicy resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	PolicyName pulumi.StringInput
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// TargetTrackingScaling Not supported for Amazon EMR
	//
	// StepScaling Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	PolicyType pulumi.StringInput
	// The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.
	ResourceId pulumi.StringPtrInput
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
	ScalableDimension pulumi.StringPtrInput
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the AWS::ApplicationAutoScaling::ScalableTarget resource.
	ScalingTargetId pulumi.StringPtrInput
	// The namespace of the AWS service that provides the resource, or a custom-resource.
	ServiceNamespace pulumi.StringPtrInput
	// A step scaling policy.
	StepScalingPolicyConfiguration ScalingPolicyStepScalingPolicyConfigurationPtrInput
	// A target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrInput
}

func (ScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPolicyArgs)(nil)).Elem()
}

type ScalingPolicyInput interface {
	pulumi.Input

	ToScalingPolicyOutput() ScalingPolicyOutput
	ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput
}

func (*ScalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPolicy)(nil)).Elem()
}

func (i *ScalingPolicy) ToScalingPolicyOutput() ScalingPolicyOutput {
	return i.ToScalingPolicyOutputWithContext(context.Background())
}

func (i *ScalingPolicy) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyOutput)
}

func (i *ScalingPolicy) ToOutput(ctx context.Context) pulumix.Output[*ScalingPolicy] {
	return pulumix.Output[*ScalingPolicy]{
		OutputState: i.ToScalingPolicyOutputWithContext(ctx).OutputState,
	}
}

type ScalingPolicyOutput struct{ *pulumi.OutputState }

func (ScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyOutput) ToScalingPolicyOutput() ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalingPolicy] {
	return pulumix.Output[*ScalingPolicy]{
		OutputState: o.OutputState,
	}
}

// ARN is a read only property for the resource.
func (o ScalingPolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the scaling policy.
//
// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing AWS::ApplicationAutoScaling::ScalingPolicy resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
func (o ScalingPolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// The scaling policy type.
//
// The following policy types are supported:
//
// # TargetTrackingScaling Not supported for Amazon EMR
//
// StepScaling Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
func (o ScalingPolicyOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.
func (o ScalingPolicyOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
func (o ScalingPolicyOutput) ScalableDimension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ScalableDimension }).(pulumi.StringPtrOutput)
}

// The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more information about the ID, see the Return Value section of the AWS::ApplicationAutoScaling::ScalableTarget resource.
func (o ScalingPolicyOutput) ScalingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ScalingTargetId }).(pulumi.StringPtrOutput)
}

// The namespace of the AWS service that provides the resource, or a custom-resource.
func (o ScalingPolicyOutput) ServiceNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ServiceNamespace }).(pulumi.StringPtrOutput)
}

// A step scaling policy.
func (o ScalingPolicyOutput) StepScalingPolicyConfiguration() ScalingPolicyStepScalingPolicyConfigurationPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) ScalingPolicyStepScalingPolicyConfigurationPtrOutput {
		return v.StepScalingPolicyConfiguration
	}).(ScalingPolicyStepScalingPolicyConfigurationPtrOutput)
}

// A target tracking scaling policy.
func (o ScalingPolicyOutput) TargetTrackingScalingPolicyConfiguration() ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput {
		return v.TargetTrackingScalingPolicyConfiguration
	}).(ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyInput)(nil)).Elem(), &ScalingPolicy{})
	pulumi.RegisterOutputType(ScalingPolicyOutput{})
}
