// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisVisualTitleLabelOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("formatText")]
        public Input<Inputs.AnalysisShortFormatTextArgs>? FormatText { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        public AnalysisVisualTitleLabelOptionsArgs()
        {
        }
        public static new AnalysisVisualTitleLabelOptionsArgs Empty => new AnalysisVisualTitleLabelOptionsArgs();
    }
}
