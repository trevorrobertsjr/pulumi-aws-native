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
    public sealed class AnalysisPeriodToDateComputation
    {
        public readonly string ComputationId;
        public readonly string? Name;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTimeGranularity? PeriodTimeGranularity;
        public readonly Outputs.AnalysisDimensionField? Time;
        public readonly Outputs.AnalysisMeasureField? Value;

        [OutputConstructor]
        private AnalysisPeriodToDateComputation(
            string computationId,

            string? name,

            Pulumi.AwsNative.QuickSight.AnalysisTimeGranularity? periodTimeGranularity,

            Outputs.AnalysisDimensionField? time,

            Outputs.AnalysisMeasureField? value)
        {
            ComputationId = computationId;
            Name = name;
            PeriodTimeGranularity = periodTimeGranularity;
            Time = time;
            Value = value;
        }
    }
}
