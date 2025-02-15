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
    public sealed class AnalysisSankeyDiagramAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.AnalysisDimensionField> Destination;
        public readonly ImmutableArray<Outputs.AnalysisDimensionField> Source;
        public readonly ImmutableArray<Outputs.AnalysisMeasureField> Weight;

        [OutputConstructor]
        private AnalysisSankeyDiagramAggregatedFieldWells(
            ImmutableArray<Outputs.AnalysisDimensionField> destination,

            ImmutableArray<Outputs.AnalysisDimensionField> source,

            ImmutableArray<Outputs.AnalysisMeasureField> weight)
        {
            Destination = destination;
            Source = source;
            Weight = weight;
        }
    }
}
