// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryFilter")]
        public Input<Inputs.AnalysisCategoryDrillDownFilterArgs>? CategoryFilter { get; set; }

        [Input("numericEqualityFilter")]
        public Input<Inputs.AnalysisNumericEqualityDrillDownFilterArgs>? NumericEqualityFilter { get; set; }

        [Input("timeRangeFilter")]
        public Input<Inputs.AnalysisTimeRangeDrillDownFilterArgs>? TimeRangeFilter { get; set; }

        public AnalysisDrillDownFilterArgs()
        {
        }
        public static new AnalysisDrillDownFilterArgs Empty => new AnalysisDrillDownFilterArgs();
    }
}
