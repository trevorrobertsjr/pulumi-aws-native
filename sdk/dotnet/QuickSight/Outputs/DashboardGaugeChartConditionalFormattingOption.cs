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
    public sealed class DashboardGaugeChartConditionalFormattingOption
    {
        public readonly Outputs.DashboardGaugeChartArcConditionalFormatting? Arc;
        public readonly Outputs.DashboardGaugeChartPrimaryValueConditionalFormatting? PrimaryValue;

        [OutputConstructor]
        private DashboardGaugeChartConditionalFormattingOption(
            Outputs.DashboardGaugeChartArcConditionalFormatting? arc,

            Outputs.DashboardGaugeChartPrimaryValueConditionalFormatting? primaryValue)
        {
            Arc = arc;
            PrimaryValue = primaryValue;
        }
    }
}
