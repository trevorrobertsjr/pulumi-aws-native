// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisWaterfallChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryAxisDisplayOptions")]
        public Input<Inputs.AnalysisAxisDisplayOptionsArgs>? CategoryAxisDisplayOptions { get; set; }

        [Input("categoryAxisLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? CategoryAxisLabelOptions { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.AnalysisDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisWaterfallChartFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.AnalysisLegendOptionsArgs>? Legend { get; set; }

        [Input("primaryYAxisDisplayOptions")]
        public Input<Inputs.AnalysisAxisDisplayOptionsArgs>? PrimaryYAxisDisplayOptions { get; set; }

        [Input("primaryYAxisLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? PrimaryYAxisLabelOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisWaterfallChartSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.AnalysisVisualPaletteArgs>? VisualPalette { get; set; }

        [Input("waterfallChartOptions")]
        public Input<Inputs.AnalysisWaterfallChartOptionsArgs>? WaterfallChartOptions { get; set; }

        public AnalysisWaterfallChartConfigurationArgs()
        {
        }
        public static new AnalysisWaterfallChartConfigurationArgs Empty => new AnalysisWaterfallChartConfigurationArgs();
    }
}
