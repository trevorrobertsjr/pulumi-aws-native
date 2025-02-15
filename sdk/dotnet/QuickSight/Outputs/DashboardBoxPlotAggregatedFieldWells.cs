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
    public sealed class DashboardBoxPlotAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.DashboardDimensionField> GroupBy;
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Values;

        [OutputConstructor]
        private DashboardBoxPlotAggregatedFieldWells(
            ImmutableArray<Outputs.DashboardDimensionField> groupBy,

            ImmutableArray<Outputs.DashboardMeasureField> values)
        {
            GroupBy = groupBy;
            Values = values;
        }
    }
}
