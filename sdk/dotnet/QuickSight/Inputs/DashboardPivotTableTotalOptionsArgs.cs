// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableTotalOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnSubtotalOptions")]
        public Input<Inputs.DashboardSubtotalOptionsArgs>? ColumnSubtotalOptions { get; set; }

        [Input("columnTotalOptions")]
        public Input<Inputs.DashboardPivotTotalOptionsArgs>? ColumnTotalOptions { get; set; }

        [Input("rowSubtotalOptions")]
        public Input<Inputs.DashboardSubtotalOptionsArgs>? RowSubtotalOptions { get; set; }

        [Input("rowTotalOptions")]
        public Input<Inputs.DashboardPivotTotalOptionsArgs>? RowTotalOptions { get; set; }

        public DashboardPivotTableTotalOptionsArgs()
        {
        }
        public static new DashboardPivotTableTotalOptionsArgs Empty => new DashboardPivotTableTotalOptionsArgs();
    }
}
