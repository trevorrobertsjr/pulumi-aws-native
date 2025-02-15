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
    public sealed class DashboardPivotTablePaginatedReportOptions
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? OverflowColumnHeaderVisibility;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? VerticalOverflowVisibility;

        [OutputConstructor]
        private DashboardPivotTablePaginatedReportOptions(
            Pulumi.AwsNative.QuickSight.DashboardVisibility? overflowColumnHeaderVisibility,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? verticalOverflowVisibility)
        {
            OverflowColumnHeaderVisibility = overflowColumnHeaderVisibility;
            VerticalOverflowVisibility = verticalOverflowVisibility;
        }
    }
}
