// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    /// <summary>
    /// An object that contains SNS topic ARN associated event destination.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationSetEventDestinationSnsDestination
    {
        public readonly string TopicArn;

        [OutputConstructor]
        private ConfigurationSetEventDestinationSnsDestination(string topicArn)
        {
            TopicArn = topicArn;
        }
    }
}
