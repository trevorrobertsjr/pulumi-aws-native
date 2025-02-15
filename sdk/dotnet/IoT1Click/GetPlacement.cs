// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT1Click
{
    public static class GetPlacement
    {
        /// <summary>
        /// Resource Type definition for AWS::IoT1Click::Placement
        /// </summary>
        public static Task<GetPlacementResult> InvokeAsync(GetPlacementArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlacementResult>("aws-native:iot1click:getPlacement", args ?? new GetPlacementArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoT1Click::Placement
        /// </summary>
        public static Output<GetPlacementResult> Invoke(GetPlacementInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlacementResult>("aws-native:iot1click:getPlacement", args ?? new GetPlacementInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlacementArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPlacementArgs()
        {
        }
        public static new GetPlacementArgs Empty => new GetPlacementArgs();
    }

    public sealed class GetPlacementInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPlacementInvokeArgs()
        {
        }
        public static new GetPlacementInvokeArgs Empty => new GetPlacementInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlacementResult
    {
        public readonly object? Attributes;
        public readonly string? Id;

        [OutputConstructor]
        private GetPlacementResult(
            object? attributes,

            string? id)
        {
            Attributes = attributes;
            Id = id;
        }
    }
}
