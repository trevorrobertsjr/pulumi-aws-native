// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardWordCloudAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.DashboardDimensionField> GroupBy;
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Size;

        [OutputConstructor]
        private DashboardWordCloudAggregatedFieldWells(
            ImmutableArray<Outputs.DashboardDimensionField> groupBy,

            ImmutableArray<Outputs.DashboardMeasureField> size)
        {
            GroupBy = groupBy;
            Size = size;
        }
    }
}
