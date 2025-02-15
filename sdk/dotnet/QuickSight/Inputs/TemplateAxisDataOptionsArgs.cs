// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAxisDataOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateAxisOptions")]
        public Input<Inputs.TemplateDateAxisOptionsArgs>? DateAxisOptions { get; set; }

        [Input("numericAxisOptions")]
        public Input<Inputs.TemplateNumericAxisOptionsArgs>? NumericAxisOptions { get; set; }

        public TemplateAxisDataOptionsArgs()
        {
        }
        public static new TemplateAxisDataOptionsArgs Empty => new TemplateAxisDataOptionsArgs();
    }
}
