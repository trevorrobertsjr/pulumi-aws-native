// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGradientStopArgs : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        [Input("dataValue")]
        public Input<double>? DataValue { get; set; }

        [Input("gradientOffset", required: true)]
        public Input<double> GradientOffset { get; set; } = null!;

        public DashboardGradientStopArgs()
        {
        }
        public static new DashboardGradientStopArgs Empty => new DashboardGradientStopArgs();
    }
}
