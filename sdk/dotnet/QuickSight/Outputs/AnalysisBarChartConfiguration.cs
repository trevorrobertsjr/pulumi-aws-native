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
    public sealed class AnalysisBarChartConfiguration
    {
        public readonly Pulumi.AwsNative.QuickSight.AnalysisBarsArrangement? BarsArrangement;
        public readonly Outputs.AnalysisAxisDisplayOptions? CategoryAxis;
        public readonly Outputs.AnalysisChartAxisLabelOptions? CategoryLabelOptions;
        public readonly Outputs.AnalysisChartAxisLabelOptions? ColorLabelOptions;
        public readonly ImmutableArray<Outputs.AnalysisContributionAnalysisDefault> ContributionAnalysisDefaults;
        public readonly Outputs.AnalysisDataLabelOptions? DataLabels;
        public readonly Outputs.AnalysisBarChartFieldWells? FieldWells;
        public readonly Outputs.AnalysisLegendOptions? Legend;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisBarChartOrientation? Orientation;
        public readonly ImmutableArray<Outputs.AnalysisReferenceLine> ReferenceLines;
        public readonly Outputs.AnalysisSmallMultiplesOptions? SmallMultiplesOptions;
        public readonly Outputs.AnalysisBarChartSortConfiguration? SortConfiguration;
        public readonly Outputs.AnalysisTooltipOptions? Tooltip;
        public readonly Outputs.AnalysisAxisDisplayOptions? ValueAxis;
        public readonly Outputs.AnalysisChartAxisLabelOptions? ValueLabelOptions;
        public readonly Outputs.AnalysisVisualPalette? VisualPalette;

        [OutputConstructor]
        private AnalysisBarChartConfiguration(
            Pulumi.AwsNative.QuickSight.AnalysisBarsArrangement? barsArrangement,

            Outputs.AnalysisAxisDisplayOptions? categoryAxis,

            Outputs.AnalysisChartAxisLabelOptions? categoryLabelOptions,

            Outputs.AnalysisChartAxisLabelOptions? colorLabelOptions,

            ImmutableArray<Outputs.AnalysisContributionAnalysisDefault> contributionAnalysisDefaults,

            Outputs.AnalysisDataLabelOptions? dataLabels,

            Outputs.AnalysisBarChartFieldWells? fieldWells,

            Outputs.AnalysisLegendOptions? legend,

            Pulumi.AwsNative.QuickSight.AnalysisBarChartOrientation? orientation,

            ImmutableArray<Outputs.AnalysisReferenceLine> referenceLines,

            Outputs.AnalysisSmallMultiplesOptions? smallMultiplesOptions,

            Outputs.AnalysisBarChartSortConfiguration? sortConfiguration,

            Outputs.AnalysisTooltipOptions? tooltip,

            Outputs.AnalysisAxisDisplayOptions? valueAxis,

            Outputs.AnalysisChartAxisLabelOptions? valueLabelOptions,

            Outputs.AnalysisVisualPalette? visualPalette)
        {
            BarsArrangement = barsArrangement;
            CategoryAxis = categoryAxis;
            CategoryLabelOptions = categoryLabelOptions;
            ColorLabelOptions = colorLabelOptions;
            ContributionAnalysisDefaults = contributionAnalysisDefaults;
            DataLabels = dataLabels;
            FieldWells = fieldWells;
            Legend = legend;
            Orientation = orientation;
            ReferenceLines = referenceLines;
            SmallMultiplesOptions = smallMultiplesOptions;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
            ValueAxis = valueAxis;
            ValueLabelOptions = valueLabelOptions;
            VisualPalette = visualPalette;
        }
    }
}
