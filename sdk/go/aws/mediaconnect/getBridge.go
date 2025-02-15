// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::MediaConnect::Bridge
func LookupBridge(ctx *pulumi.Context, args *LookupBridgeArgs, opts ...pulumi.InvokeOption) (*LookupBridgeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBridgeResult
	err := ctx.Invoke("aws-native:mediaconnect:getBridge", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBridgeArgs struct {
	// The Amazon Resource Number (ARN) of the bridge.
	BridgeArn string `pulumi:"bridgeArn"`
}

type LookupBridgeResult struct {
	// The Amazon Resource Number (ARN) of the bridge.
	BridgeArn            *string                     `pulumi:"bridgeArn"`
	BridgeState          *BridgeStateEnum            `pulumi:"bridgeState"`
	EgressGatewayBridge  *BridgeEgressGatewayBridge  `pulumi:"egressGatewayBridge"`
	IngressGatewayBridge *BridgeIngressGatewayBridge `pulumi:"ingressGatewayBridge"`
	// The name of the bridge.
	Name *string `pulumi:"name"`
	// The outputs on this bridge.
	Outputs []BridgeOutputType `pulumi:"outputs"`
	// The placement Amazon Resource Number (ARN) of the bridge.
	PlacementArn         *string               `pulumi:"placementArn"`
	SourceFailoverConfig *BridgeFailoverConfig `pulumi:"sourceFailoverConfig"`
	// The sources on this bridge.
	Sources []BridgeSourceType `pulumi:"sources"`
}

func LookupBridgeOutput(ctx *pulumi.Context, args LookupBridgeOutputArgs, opts ...pulumi.InvokeOption) LookupBridgeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBridgeResult, error) {
			args := v.(LookupBridgeArgs)
			r, err := LookupBridge(ctx, &args, opts...)
			var s LookupBridgeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBridgeResultOutput)
}

type LookupBridgeOutputArgs struct {
	// The Amazon Resource Number (ARN) of the bridge.
	BridgeArn pulumi.StringInput `pulumi:"bridgeArn"`
}

func (LookupBridgeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBridgeArgs)(nil)).Elem()
}

type LookupBridgeResultOutput struct{ *pulumi.OutputState }

func (LookupBridgeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBridgeResult)(nil)).Elem()
}

func (o LookupBridgeResultOutput) ToLookupBridgeResultOutput() LookupBridgeResultOutput {
	return o
}

func (o LookupBridgeResultOutput) ToLookupBridgeResultOutputWithContext(ctx context.Context) LookupBridgeResultOutput {
	return o
}

func (o LookupBridgeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBridgeResult] {
	return pulumix.Output[LookupBridgeResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Number (ARN) of the bridge.
func (o LookupBridgeResultOutput) BridgeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *string { return v.BridgeArn }).(pulumi.StringPtrOutput)
}

func (o LookupBridgeResultOutput) BridgeState() BridgeStateEnumPtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *BridgeStateEnum { return v.BridgeState }).(BridgeStateEnumPtrOutput)
}

func (o LookupBridgeResultOutput) EgressGatewayBridge() BridgeEgressGatewayBridgePtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *BridgeEgressGatewayBridge { return v.EgressGatewayBridge }).(BridgeEgressGatewayBridgePtrOutput)
}

func (o LookupBridgeResultOutput) IngressGatewayBridge() BridgeIngressGatewayBridgePtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *BridgeIngressGatewayBridge { return v.IngressGatewayBridge }).(BridgeIngressGatewayBridgePtrOutput)
}

// The name of the bridge.
func (o LookupBridgeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The outputs on this bridge.
func (o LookupBridgeResultOutput) Outputs() BridgeOutputTypeArrayOutput {
	return o.ApplyT(func(v LookupBridgeResult) []BridgeOutputType { return v.Outputs }).(BridgeOutputTypeArrayOutput)
}

// The placement Amazon Resource Number (ARN) of the bridge.
func (o LookupBridgeResultOutput) PlacementArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *string { return v.PlacementArn }).(pulumi.StringPtrOutput)
}

func (o LookupBridgeResultOutput) SourceFailoverConfig() BridgeFailoverConfigPtrOutput {
	return o.ApplyT(func(v LookupBridgeResult) *BridgeFailoverConfig { return v.SourceFailoverConfig }).(BridgeFailoverConfigPtrOutput)
}

// The sources on this bridge.
func (o LookupBridgeResultOutput) Sources() BridgeSourceTypeArrayOutput {
	return o.ApplyT(func(v LookupBridgeResult) []BridgeSourceType { return v.Sources }).(BridgeSourceTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBridgeResultOutput{})
}
