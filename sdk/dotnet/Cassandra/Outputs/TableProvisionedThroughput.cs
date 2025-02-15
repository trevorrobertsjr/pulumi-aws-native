// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Outputs
{

    /// <summary>
    /// Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits
    /// </summary>
    [OutputType]
    public sealed class TableProvisionedThroughput
    {
        public readonly int ReadCapacityUnits;
        public readonly int WriteCapacityUnits;

        [OutputConstructor]
        private TableProvisionedThroughput(
            int readCapacityUnits,

            int writeCapacityUnits)
        {
            ReadCapacityUnits = readCapacityUnits;
            WriteCapacityUnits = writeCapacityUnits;
        }
    }
}
