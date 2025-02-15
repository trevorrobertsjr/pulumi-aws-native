// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs
{
    public static class GetCapacityProvider
    {
        /// <summary>
        /// Resource Type definition for AWS::ECS::CapacityProvider.
        /// </summary>
        public static Task<GetCapacityProviderResult> InvokeAsync(GetCapacityProviderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCapacityProviderResult>("aws-native:ecs:getCapacityProvider", args ?? new GetCapacityProviderArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ECS::CapacityProvider.
        /// </summary>
        public static Output<GetCapacityProviderResult> Invoke(GetCapacityProviderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCapacityProviderResult>("aws-native:ecs:getCapacityProvider", args ?? new GetCapacityProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCapacityProviderArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCapacityProviderArgs()
        {
        }
        public static new GetCapacityProviderArgs Empty => new GetCapacityProviderArgs();
    }

    public sealed class GetCapacityProviderInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetCapacityProviderInvokeArgs()
        {
        }
        public static new GetCapacityProviderInvokeArgs Empty => new GetCapacityProviderInvokeArgs();
    }


    [OutputType]
    public sealed class GetCapacityProviderResult
    {
        public readonly Outputs.CapacityProviderAutoScalingGroupProvider? AutoScalingGroupProvider;
        public readonly ImmutableArray<Outputs.CapacityProviderTag> Tags;

        [OutputConstructor]
        private GetCapacityProviderResult(
            Outputs.CapacityProviderAutoScalingGroupProvider? autoScalingGroupProvider,

            ImmutableArray<Outputs.CapacityProviderTag> tags)
        {
            AutoScalingGroupProvider = autoScalingGroupProvider;
            Tags = tags;
        }
    }
}
