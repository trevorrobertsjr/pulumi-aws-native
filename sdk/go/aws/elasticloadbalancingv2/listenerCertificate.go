// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::ListenerCertificate
//
// Deprecated: ListenerCertificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ListenerCertificate struct {
	pulumi.CustomResourceState

	Certificates ListenerCertificateCertificateArrayOutput `pulumi:"certificates"`
	ListenerArn  pulumi.StringOutput                       `pulumi:"listenerArn"`
}

// NewListenerCertificate registers a new resource with the given unique name, arguments, and options.
func NewListenerCertificate(ctx *pulumi.Context,
	name string, args *ListenerCertificateArgs, opts ...pulumi.ResourceOption) (*ListenerCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificates == nil {
		return nil, errors.New("invalid value for required argument 'Certificates'")
	}
	if args.ListenerArn == nil {
		return nil, errors.New("invalid value for required argument 'ListenerArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"listenerArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ListenerCertificate
	err := ctx.RegisterResource("aws-native:elasticloadbalancingv2:ListenerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerCertificate gets an existing ListenerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerCertificateState, opts ...pulumi.ResourceOption) (*ListenerCertificate, error) {
	var resource ListenerCertificate
	err := ctx.ReadResource("aws-native:elasticloadbalancingv2:ListenerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerCertificate resources.
type listenerCertificateState struct {
}

type ListenerCertificateState struct {
}

func (ListenerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerCertificateState)(nil)).Elem()
}

type listenerCertificateArgs struct {
	Certificates []ListenerCertificateCertificate `pulumi:"certificates"`
	ListenerArn  string                           `pulumi:"listenerArn"`
}

// The set of arguments for constructing a ListenerCertificate resource.
type ListenerCertificateArgs struct {
	Certificates ListenerCertificateCertificateArrayInput
	ListenerArn  pulumi.StringInput
}

func (ListenerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerCertificateArgs)(nil)).Elem()
}

type ListenerCertificateInput interface {
	pulumi.Input

	ToListenerCertificateOutput() ListenerCertificateOutput
	ToListenerCertificateOutputWithContext(ctx context.Context) ListenerCertificateOutput
}

func (*ListenerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerCertificate)(nil)).Elem()
}

func (i *ListenerCertificate) ToListenerCertificateOutput() ListenerCertificateOutput {
	return i.ToListenerCertificateOutputWithContext(context.Background())
}

func (i *ListenerCertificate) ToListenerCertificateOutputWithContext(ctx context.Context) ListenerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerCertificateOutput)
}

func (i *ListenerCertificate) ToOutput(ctx context.Context) pulumix.Output[*ListenerCertificate] {
	return pulumix.Output[*ListenerCertificate]{
		OutputState: i.ToListenerCertificateOutputWithContext(ctx).OutputState,
	}
}

type ListenerCertificateOutput struct{ *pulumi.OutputState }

func (ListenerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerCertificate)(nil)).Elem()
}

func (o ListenerCertificateOutput) ToListenerCertificateOutput() ListenerCertificateOutput {
	return o
}

func (o ListenerCertificateOutput) ToListenerCertificateOutputWithContext(ctx context.Context) ListenerCertificateOutput {
	return o
}

func (o ListenerCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[*ListenerCertificate] {
	return pulumix.Output[*ListenerCertificate]{
		OutputState: o.OutputState,
	}
}

func (o ListenerCertificateOutput) Certificates() ListenerCertificateCertificateArrayOutput {
	return o.ApplyT(func(v *ListenerCertificate) ListenerCertificateCertificateArrayOutput { return v.Certificates }).(ListenerCertificateCertificateArrayOutput)
}

func (o ListenerCertificateOutput) ListenerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerCertificate) pulumi.StringOutput { return v.ListenerArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerCertificateInput)(nil)).Elem(), &ListenerCertificate{})
	pulumi.RegisterOutputType(ListenerCertificateOutput{})
}
