// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardWordCloudFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("wordCloudAggregatedFieldWells")]
        public Input<Inputs.DashboardWordCloudAggregatedFieldWellsArgs>? WordCloudAggregatedFieldWells { get; set; }

        public DashboardWordCloudFieldWellsArgs()
        {
        }
        public static new DashboardWordCloudFieldWellsArgs Empty => new DashboardWordCloudFieldWellsArgs();
    }
}
