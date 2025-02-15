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
    public sealed class DashboardSubtotalOptions
    {
        public readonly string? CustomLabel;
        public readonly Pulumi.AwsNative.QuickSight.DashboardPivotTableSubtotalLevel? FieldLevel;
        public readonly ImmutableArray<Outputs.DashboardPivotTableFieldSubtotalOptions> FieldLevelOptions;
        public readonly Outputs.DashboardTableCellStyle? MetricHeaderCellStyle;
        public readonly ImmutableArray<Outputs.DashboardTableStyleTarget> StyleTargets;
        public readonly Outputs.DashboardTableCellStyle? TotalCellStyle;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? TotalsVisibility;
        public readonly Outputs.DashboardTableCellStyle? ValueCellStyle;

        [OutputConstructor]
        private DashboardSubtotalOptions(
            string? customLabel,

            Pulumi.AwsNative.QuickSight.DashboardPivotTableSubtotalLevel? fieldLevel,

            ImmutableArray<Outputs.DashboardPivotTableFieldSubtotalOptions> fieldLevelOptions,

            Outputs.DashboardTableCellStyle? metricHeaderCellStyle,

            ImmutableArray<Outputs.DashboardTableStyleTarget> styleTargets,

            Outputs.DashboardTableCellStyle? totalCellStyle,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? totalsVisibility,

            Outputs.DashboardTableCellStyle? valueCellStyle)
        {
            CustomLabel = customLabel;
            FieldLevel = fieldLevel;
            FieldLevelOptions = fieldLevelOptions;
            MetricHeaderCellStyle = metricHeaderCellStyle;
            StyleTargets = styleTargets;
            TotalCellStyle = totalCellStyle;
            TotalsVisibility = totalsVisibility;
            ValueCellStyle = valueCellStyle;
        }
    }
}
