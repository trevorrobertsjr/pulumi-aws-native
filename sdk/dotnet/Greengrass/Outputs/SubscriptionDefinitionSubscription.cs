// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Outputs
{

    [OutputType]
    public sealed class SubscriptionDefinitionSubscription
    {
        public readonly string Id;
        public readonly string Source;
        public readonly string Subject;
        public readonly string Target;

        [OutputConstructor]
        private SubscriptionDefinitionSubscription(
            string id,

            string source,

            string subject,

            string target)
        {
            Id = id;
            Source = source;
            Subject = subject;
            Target = target;
        }
    }
}
