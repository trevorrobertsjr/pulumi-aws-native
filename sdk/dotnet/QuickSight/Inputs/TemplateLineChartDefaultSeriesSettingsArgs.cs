// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateLineChartDefaultSeriesSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("axisBinding")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateAxisBinding>? AxisBinding { get; set; }

        [Input("lineStyleSettings")]
        public Input<Inputs.TemplateLineChartLineStyleSettingsArgs>? LineStyleSettings { get; set; }

        [Input("markerStyleSettings")]
        public Input<Inputs.TemplateLineChartMarkerStyleSettingsArgs>? MarkerStyleSettings { get; set; }

        public TemplateLineChartDefaultSeriesSettingsArgs()
        {
        }
        public static new TemplateLineChartDefaultSeriesSettingsArgs Empty => new TemplateLineChartDefaultSeriesSettingsArgs();
    }
}
