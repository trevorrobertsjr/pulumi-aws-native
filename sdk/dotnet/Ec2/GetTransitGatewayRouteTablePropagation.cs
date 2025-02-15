// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetTransitGatewayRouteTablePropagation
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::TransitGatewayRouteTablePropagation
        /// </summary>
        public static Task<GetTransitGatewayRouteTablePropagationResult> InvokeAsync(GetTransitGatewayRouteTablePropagationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayRouteTablePropagationResult>("aws-native:ec2:getTransitGatewayRouteTablePropagation", args ?? new GetTransitGatewayRouteTablePropagationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::TransitGatewayRouteTablePropagation
        /// </summary>
        public static Output<GetTransitGatewayRouteTablePropagationResult> Invoke(GetTransitGatewayRouteTablePropagationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitGatewayRouteTablePropagationResult>("aws-native:ec2:getTransitGatewayRouteTablePropagation", args ?? new GetTransitGatewayRouteTablePropagationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitGatewayRouteTablePropagationArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTransitGatewayRouteTablePropagationArgs()
        {
        }
        public static new GetTransitGatewayRouteTablePropagationArgs Empty => new GetTransitGatewayRouteTablePropagationArgs();
    }

    public sealed class GetTransitGatewayRouteTablePropagationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTransitGatewayRouteTablePropagationInvokeArgs()
        {
        }
        public static new GetTransitGatewayRouteTablePropagationInvokeArgs Empty => new GetTransitGatewayRouteTablePropagationInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitGatewayRouteTablePropagationResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetTransitGatewayRouteTablePropagationResult(string? id)
        {
            Id = id;
        }
    }
}
