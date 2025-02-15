// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2
{
    public static class GetApplicationCloudWatchLoggingOption
    {
        /// <summary>
        /// Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
        /// </summary>
        public static Task<GetApplicationCloudWatchLoggingOptionResult> InvokeAsync(GetApplicationCloudWatchLoggingOptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationCloudWatchLoggingOptionResult>("aws-native:kinesisanalyticsv2:getApplicationCloudWatchLoggingOption", args ?? new GetApplicationCloudWatchLoggingOptionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
        /// </summary>
        public static Output<GetApplicationCloudWatchLoggingOptionResult> Invoke(GetApplicationCloudWatchLoggingOptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationCloudWatchLoggingOptionResult>("aws-native:kinesisanalyticsv2:getApplicationCloudWatchLoggingOption", args ?? new GetApplicationCloudWatchLoggingOptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationCloudWatchLoggingOptionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetApplicationCloudWatchLoggingOptionArgs()
        {
        }
        public static new GetApplicationCloudWatchLoggingOptionArgs Empty => new GetApplicationCloudWatchLoggingOptionArgs();
    }

    public sealed class GetApplicationCloudWatchLoggingOptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetApplicationCloudWatchLoggingOptionInvokeArgs()
        {
        }
        public static new GetApplicationCloudWatchLoggingOptionInvokeArgs Empty => new GetApplicationCloudWatchLoggingOptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationCloudWatchLoggingOptionResult
    {
        public readonly Outputs.ApplicationCloudWatchLoggingOptionCloudWatchLoggingOption? CloudWatchLoggingOption;
        public readonly string? Id;

        [OutputConstructor]
        private GetApplicationCloudWatchLoggingOptionResult(
            Outputs.ApplicationCloudWatchLoggingOptionCloudWatchLoggingOption? cloudWatchLoggingOption,

            string? id)
        {
            CloudWatchLoggingOption = cloudWatchLoggingOption;
            Id = id;
        }
    }
}
