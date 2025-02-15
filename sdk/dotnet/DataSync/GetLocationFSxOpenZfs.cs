// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetLocationFSxOpenZfs
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::LocationFSxOpenZFS.
        /// </summary>
        public static Task<GetLocationFSxOpenZfsResult> InvokeAsync(GetLocationFSxOpenZfsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocationFSxOpenZfsResult>("aws-native:datasync:getLocationFSxOpenZfs", args ?? new GetLocationFSxOpenZfsArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::LocationFSxOpenZFS.
        /// </summary>
        public static Output<GetLocationFSxOpenZfsResult> Invoke(GetLocationFSxOpenZfsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocationFSxOpenZfsResult>("aws-native:datasync:getLocationFSxOpenZfs", args ?? new GetLocationFSxOpenZfsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationFSxOpenZfsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public string LocationArn { get; set; } = null!;

        public GetLocationFSxOpenZfsArgs()
        {
        }
        public static new GetLocationFSxOpenZfsArgs Empty => new GetLocationFSxOpenZfsArgs();
    }

    public sealed class GetLocationFSxOpenZfsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public Input<string> LocationArn { get; set; } = null!;

        public GetLocationFSxOpenZfsInvokeArgs()
        {
        }
        public static new GetLocationFSxOpenZfsInvokeArgs Empty => new GetLocationFSxOpenZfsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocationFSxOpenZfsResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
        /// </summary>
        public readonly string? LocationArn;
        /// <summary>
        /// The URL of the FSx OpenZFS that was described.
        /// </summary>
        public readonly string? LocationUri;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationFSxOpenZfsTag> Tags;

        [OutputConstructor]
        private GetLocationFSxOpenZfsResult(
            string? locationArn,

            string? locationUri,

            ImmutableArray<Outputs.LocationFSxOpenZfsTag> tags)
        {
            LocationArn = locationArn;
            LocationUri = locationUri;
            Tags = tags;
        }
    }
}
