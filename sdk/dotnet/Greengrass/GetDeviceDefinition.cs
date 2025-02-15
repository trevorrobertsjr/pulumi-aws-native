// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetDeviceDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::DeviceDefinition
        /// </summary>
        public static Task<GetDeviceDefinitionResult> InvokeAsync(GetDeviceDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceDefinitionResult>("aws-native:greengrass:getDeviceDefinition", args ?? new GetDeviceDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::DeviceDefinition
        /// </summary>
        public static Output<GetDeviceDefinitionResult> Invoke(GetDeviceDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceDefinitionResult>("aws-native:greengrass:getDeviceDefinition", args ?? new GetDeviceDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceDefinitionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeviceDefinitionArgs()
        {
        }
        public static new GetDeviceDefinitionArgs Empty => new GetDeviceDefinitionArgs();
    }

    public sealed class GetDeviceDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeviceDefinitionInvokeArgs()
        {
        }
        public static new GetDeviceDefinitionInvokeArgs Empty => new GetDeviceDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceDefinitionResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? LatestVersionArn;
        public readonly string? Name;
        public readonly object? Tags;

        [OutputConstructor]
        private GetDeviceDefinitionResult(
            string? arn,

            string? id,

            string? latestVersionArn,

            string? name,

            object? tags)
        {
            Arn = arn;
            Id = id;
            LatestVersionArn = latestVersionArn;
            Name = name;
            Tags = tags;
        }
    }
}
