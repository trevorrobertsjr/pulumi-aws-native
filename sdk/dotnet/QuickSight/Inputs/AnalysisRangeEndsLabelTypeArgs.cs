// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisRangeEndsLabelTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        public AnalysisRangeEndsLabelTypeArgs()
        {
        }
        public static new AnalysisRangeEndsLabelTypeArgs Empty => new AnalysisRangeEndsLabelTypeArgs();
    }
}
