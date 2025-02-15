// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sns.Inputs
{

    public sealed class TopicLoggingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch
        /// </summary>
        [Input("failureFeedbackRoleArn")]
        public Input<string>? FailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// Indicates one of the supported protocols for the SNS topic
        /// </summary>
        [Input("protocol", required: true)]
        public Input<Pulumi.AwsNative.Sns.TopicLoggingConfigProtocol> Protocol { get; set; } = null!;

        /// <summary>
        /// The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch
        /// </summary>
        [Input("successFeedbackRoleArn")]
        public Input<string>? SuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid percentage values range from 0 to 100
        /// </summary>
        [Input("successFeedbackSampleRate")]
        public Input<string>? SuccessFeedbackSampleRate { get; set; }

        public TopicLoggingConfigArgs()
        {
        }
        public static new TopicLoggingConfigArgs Empty => new TopicLoggingConfigArgs();
    }
}
