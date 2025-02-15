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
    public sealed class TemplateLegendOptions
    {
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? Height;
        public readonly Pulumi.AwsNative.QuickSight.TemplateLegendPosition? Position;
        public readonly Outputs.TemplateLabelOptions? Title;
        public readonly Pulumi.AwsNative.QuickSight.TemplateVisibility? Visibility;
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? Width;

        [OutputConstructor]
        private TemplateLegendOptions(
            string? height,

            Pulumi.AwsNative.QuickSight.TemplateLegendPosition? position,

            Outputs.TemplateLabelOptions? title,

            Pulumi.AwsNative.QuickSight.TemplateVisibility? visibility,

            string? width)
        {
            Height = height;
            Position = position;
            Title = title;
            Visibility = visibility;
            Width = width;
        }
    }
}
