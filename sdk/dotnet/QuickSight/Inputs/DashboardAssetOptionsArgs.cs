// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardAssetOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        [Input("weekStart")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardDayOfTheWeek>? WeekStart { get; set; }

        public DashboardAssetOptionsArgs()
        {
        }
        public static new DashboardAssetOptionsArgs Empty => new DashboardAssetOptionsArgs();
    }
}
