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
    public sealed class AnalysisDataPathSort
    {
        public readonly Pulumi.AwsNative.QuickSight.AnalysisSortDirection Direction;
        public readonly ImmutableArray<Outputs.AnalysisDataPathValue> SortPaths;

        [OutputConstructor]
        private AnalysisDataPathSort(
            Pulumi.AwsNative.QuickSight.AnalysisSortDirection direction,

            ImmutableArray<Outputs.AnalysisDataPathValue> sortPaths)
        {
            Direction = direction;
            SortPaths = sortPaths;
        }
    }
}
