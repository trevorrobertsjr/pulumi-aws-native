// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisAnchorDateConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("anchorOption")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisAnchorOption>? AnchorOption { get; set; }

        [Input("parameterName")]
        public Input<string>? ParameterName { get; set; }

        public AnalysisAnchorDateConfigurationArgs()
        {
        }
        public static new AnalysisAnchorDateConfigurationArgs Empty => new AnalysisAnchorDateConfigurationArgs();
    }
}
