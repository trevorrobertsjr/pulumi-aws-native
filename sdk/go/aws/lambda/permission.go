// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Lambda::Permission
type Permission struct {
	pulumi.CustomResourceState

	// The action that the principal can use on the function.
	Action pulumi.StringOutput `pulumi:"action"`
	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	EventSourceToken pulumi.StringPtrOutput `pulumi:"eventSourceToken"`
	// The name of the Lambda function, version, or alias.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType PermissionFunctionUrlAuthTypePtrOutput `pulumi:"functionUrlAuthType"`
	// The AWS service or account that invokes the function. If you specify a service, use SourceArn or SourceAccount to limit who can invoke the function through that service.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The identifier for your organization in AWS Organizations. Use this to grant permissions to all the AWS accounts under this organization.
	PrincipalOrgId pulumi.StringPtrOutput `pulumi:"principalOrgId"`
	// For Amazon S3, the ID of the account that owns the resource. Use this together with SourceArn to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount pulumi.StringPtrOutput `pulumi:"sourceAccount"`
	// For AWS services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
	SourceArn pulumi.StringPtrOutput `pulumi:"sourceArn"`
}

// NewPermission registers a new resource with the given unique name, arguments, and options.
func NewPermission(ctx *pulumi.Context,
	name string, args *PermissionArgs, opts ...pulumi.ResourceOption) (*Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"action",
		"eventSourceToken",
		"functionName",
		"functionUrlAuthType",
		"principal",
		"principalOrgId",
		"sourceAccount",
		"sourceArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Permission
	err := ctx.RegisterResource("aws-native:lambda:Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermission gets an existing Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionState, opts ...pulumi.ResourceOption) (*Permission, error) {
	var resource Permission
	err := ctx.ReadResource("aws-native:lambda:Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permission resources.
type permissionState struct {
}

type PermissionState struct {
}

func (PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionState)(nil)).Elem()
}

type permissionArgs struct {
	// The action that the principal can use on the function.
	Action string `pulumi:"action"`
	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	EventSourceToken *string `pulumi:"eventSourceToken"`
	// The name of the Lambda function, version, or alias.
	FunctionName string `pulumi:"functionName"`
	// The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType *PermissionFunctionUrlAuthType `pulumi:"functionUrlAuthType"`
	// The AWS service or account that invokes the function. If you specify a service, use SourceArn or SourceAccount to limit who can invoke the function through that service.
	Principal string `pulumi:"principal"`
	// The identifier for your organization in AWS Organizations. Use this to grant permissions to all the AWS accounts under this organization.
	PrincipalOrgId *string `pulumi:"principalOrgId"`
	// For Amazon S3, the ID of the account that owns the resource. Use this together with SourceArn to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount *string `pulumi:"sourceAccount"`
	// For AWS services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
	SourceArn *string `pulumi:"sourceArn"`
}

// The set of arguments for constructing a Permission resource.
type PermissionArgs struct {
	// The action that the principal can use on the function.
	Action pulumi.StringInput
	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	EventSourceToken pulumi.StringPtrInput
	// The name of the Lambda function, version, or alias.
	FunctionName pulumi.StringInput
	// The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType PermissionFunctionUrlAuthTypePtrInput
	// The AWS service or account that invokes the function. If you specify a service, use SourceArn or SourceAccount to limit who can invoke the function through that service.
	Principal pulumi.StringInput
	// The identifier for your organization in AWS Organizations. Use this to grant permissions to all the AWS accounts under this organization.
	PrincipalOrgId pulumi.StringPtrInput
	// For Amazon S3, the ID of the account that owns the resource. Use this together with SourceArn to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount pulumi.StringPtrInput
	// For AWS services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
	SourceArn pulumi.StringPtrInput
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionArgs)(nil)).Elem()
}

type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(ctx context.Context) PermissionOutput
}

func (*Permission) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (i *Permission) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i *Permission) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}

func (i *Permission) ToOutput(ctx context.Context) pulumix.Output[*Permission] {
	return pulumix.Output[*Permission]{
		OutputState: i.ToPermissionOutputWithContext(ctx).OutputState,
	}
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func (o PermissionOutput) ToOutput(ctx context.Context) pulumix.Output[*Permission] {
	return pulumix.Output[*Permission]{
		OutputState: o.OutputState,
	}
}

// The action that the principal can use on the function.
func (o PermissionOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// For Alexa Smart Home functions, a token that must be supplied by the invoker.
func (o PermissionOutput) EventSourceToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.EventSourceToken }).(pulumi.StringPtrOutput)
}

// The name of the Lambda function, version, or alias.
func (o PermissionOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.FunctionName }).(pulumi.StringOutput)
}

// The type of authentication that your function URL uses. Set to AWS_IAM if you want to restrict access to authenticated users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint.
func (o PermissionOutput) FunctionUrlAuthType() PermissionFunctionUrlAuthTypePtrOutput {
	return o.ApplyT(func(v *Permission) PermissionFunctionUrlAuthTypePtrOutput { return v.FunctionUrlAuthType }).(PermissionFunctionUrlAuthTypePtrOutput)
}

// The AWS service or account that invokes the function. If you specify a service, use SourceArn or SourceAccount to limit who can invoke the function through that service.
func (o PermissionOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The identifier for your organization in AWS Organizations. Use this to grant permissions to all the AWS accounts under this organization.
func (o PermissionOutput) PrincipalOrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.PrincipalOrgId }).(pulumi.StringPtrOutput)
}

// For Amazon S3, the ID of the account that owns the resource. Use this together with SourceArn to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
func (o PermissionOutput) SourceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.SourceAccount }).(pulumi.StringPtrOutput)
}

// For AWS services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
func (o PermissionOutput) SourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.SourceArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionInput)(nil)).Elem(), &Permission{})
	pulumi.RegisterOutputType(PermissionOutput{})
}
