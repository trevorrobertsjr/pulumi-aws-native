// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatastorePartitions
    {
        public readonly ImmutableArray<Outputs.DatastorePartition> Partitions;

        [OutputConstructor]
        private DatastorePartitions(ImmutableArray<Outputs.DatastorePartition> partitions)
        {
            Partitions = partitions;
        }
    }
}
