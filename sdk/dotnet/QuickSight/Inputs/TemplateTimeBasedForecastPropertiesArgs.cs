// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTimeBasedForecastPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("lowerBoundary")]
        public Input<double>? LowerBoundary { get; set; }

        [Input("periodsBackward")]
        public Input<double>? PeriodsBackward { get; set; }

        [Input("periodsForward")]
        public Input<double>? PeriodsForward { get; set; }

        [Input("predictionInterval")]
        public Input<double>? PredictionInterval { get; set; }

        [Input("seasonality")]
        public Input<double>? Seasonality { get; set; }

        [Input("upperBoundary")]
        public Input<double>? UpperBoundary { get; set; }

        public TemplateTimeBasedForecastPropertiesArgs()
        {
        }
        public static new TemplateTimeBasedForecastPropertiesArgs Empty => new TemplateTimeBasedForecastPropertiesArgs();
    }
}
