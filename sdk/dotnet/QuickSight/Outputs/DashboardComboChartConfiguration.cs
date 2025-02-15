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
    public sealed class DashboardComboChartConfiguration
    {
        public readonly Outputs.DashboardDataLabelOptions? BarDataLabels;
        public readonly Pulumi.AwsNative.QuickSight.DashboardBarsArrangement? BarsArrangement;
        public readonly Outputs.DashboardAxisDisplayOptions? CategoryAxis;
        public readonly Outputs.DashboardChartAxisLabelOptions? CategoryLabelOptions;
        public readonly Outputs.DashboardChartAxisLabelOptions? ColorLabelOptions;
        public readonly Outputs.DashboardComboChartFieldWells? FieldWells;
        public readonly Outputs.DashboardLegendOptions? Legend;
        public readonly Outputs.DashboardDataLabelOptions? LineDataLabels;
        public readonly Outputs.DashboardAxisDisplayOptions? PrimaryYAxisDisplayOptions;
        public readonly Outputs.DashboardChartAxisLabelOptions? PrimaryYAxisLabelOptions;
        public readonly ImmutableArray<Outputs.DashboardReferenceLine> ReferenceLines;
        public readonly Outputs.DashboardAxisDisplayOptions? SecondaryYAxisDisplayOptions;
        public readonly Outputs.DashboardChartAxisLabelOptions? SecondaryYAxisLabelOptions;
        public readonly Outputs.DashboardComboChartSortConfiguration? SortConfiguration;
        public readonly Outputs.DashboardTooltipOptions? Tooltip;
        public readonly Outputs.DashboardVisualPalette? VisualPalette;

        [OutputConstructor]
        private DashboardComboChartConfiguration(
            Outputs.DashboardDataLabelOptions? barDataLabels,

            Pulumi.AwsNative.QuickSight.DashboardBarsArrangement? barsArrangement,

            Outputs.DashboardAxisDisplayOptions? categoryAxis,

            Outputs.DashboardChartAxisLabelOptions? categoryLabelOptions,

            Outputs.DashboardChartAxisLabelOptions? colorLabelOptions,

            Outputs.DashboardComboChartFieldWells? fieldWells,

            Outputs.DashboardLegendOptions? legend,

            Outputs.DashboardDataLabelOptions? lineDataLabels,

            Outputs.DashboardAxisDisplayOptions? primaryYAxisDisplayOptions,

            Outputs.DashboardChartAxisLabelOptions? primaryYAxisLabelOptions,

            ImmutableArray<Outputs.DashboardReferenceLine> referenceLines,

            Outputs.DashboardAxisDisplayOptions? secondaryYAxisDisplayOptions,

            Outputs.DashboardChartAxisLabelOptions? secondaryYAxisLabelOptions,

            Outputs.DashboardComboChartSortConfiguration? sortConfiguration,

            Outputs.DashboardTooltipOptions? tooltip,

            Outputs.DashboardVisualPalette? visualPalette)
        {
            BarDataLabels = barDataLabels;
            BarsArrangement = barsArrangement;
            CategoryAxis = categoryAxis;
            CategoryLabelOptions = categoryLabelOptions;
            ColorLabelOptions = colorLabelOptions;
            FieldWells = fieldWells;
            Legend = legend;
            LineDataLabels = lineDataLabels;
            PrimaryYAxisDisplayOptions = primaryYAxisDisplayOptions;
            PrimaryYAxisLabelOptions = primaryYAxisLabelOptions;
            ReferenceLines = referenceLines;
            SecondaryYAxisDisplayOptions = secondaryYAxisDisplayOptions;
            SecondaryYAxisLabelOptions = secondaryYAxisLabelOptions;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
            VisualPalette = visualPalette;
        }
    }
}
