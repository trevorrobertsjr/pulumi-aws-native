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
    public sealed class TemplateSetParameterValueConfiguration
    {
        public readonly string DestinationParameterName;
        public readonly Outputs.TemplateDestinationParameterValueConfiguration Value;

        [OutputConstructor]
        private TemplateSetParameterValueConfiguration(
            string destinationParameterName,

            Outputs.TemplateDestinationParameterValueConfiguration value)
        {
            DestinationParameterName = destinationParameterName;
            Value = value;
        }
    }
}
