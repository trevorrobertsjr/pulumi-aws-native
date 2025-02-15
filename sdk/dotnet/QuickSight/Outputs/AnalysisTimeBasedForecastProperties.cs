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
    public sealed class AnalysisTimeBasedForecastProperties
    {
        public readonly double? LowerBoundary;
        public readonly double? PeriodsBackward;
        public readonly double? PeriodsForward;
        public readonly double? PredictionInterval;
        public readonly double? Seasonality;
        public readonly double? UpperBoundary;

        [OutputConstructor]
        private AnalysisTimeBasedForecastProperties(
            double? lowerBoundary,

            double? periodsBackward,

            double? periodsForward,

            double? predictionInterval,

            double? seasonality,

            double? upperBoundary)
        {
            LowerBoundary = lowerBoundary;
            PeriodsBackward = periodsBackward;
            PeriodsForward = periodsForward;
            PredictionInterval = predictionInterval;
            Seasonality = seasonality;
            UpperBoundary = upperBoundary;
        }
    }
}
