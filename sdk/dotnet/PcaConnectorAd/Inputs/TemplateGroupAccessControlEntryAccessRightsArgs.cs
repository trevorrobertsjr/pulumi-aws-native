// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Inputs
{

    public sealed class TemplateGroupAccessControlEntryAccessRightsArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoEnroll")]
        public Input<Pulumi.AwsNative.PcaConnectorAd.TemplateGroupAccessControlEntryAccessRight>? AutoEnroll { get; set; }

        [Input("enroll")]
        public Input<Pulumi.AwsNative.PcaConnectorAd.TemplateGroupAccessControlEntryAccessRight>? Enroll { get; set; }

        public TemplateGroupAccessControlEntryAccessRightsArgs()
        {
        }
        public static new TemplateGroupAccessControlEntryAccessRightsArgs Empty => new TemplateGroupAccessControlEntryAccessRightsArgs();
    }
}
