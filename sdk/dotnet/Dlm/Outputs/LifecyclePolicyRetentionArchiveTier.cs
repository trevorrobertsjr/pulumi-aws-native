// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dlm.Outputs
{

    [OutputType]
    public sealed class LifecyclePolicyRetentionArchiveTier
    {
        public readonly int? Count;
        public readonly int? Interval;
        public readonly string? IntervalUnit;

        [OutputConstructor]
        private LifecyclePolicyRetentionArchiveTier(
            int? count,

            int? interval,

            string? intervalUnit)
        {
            Count = count;
            Interval = interval;
            IntervalUnit = intervalUnit;
        }
    }
}
