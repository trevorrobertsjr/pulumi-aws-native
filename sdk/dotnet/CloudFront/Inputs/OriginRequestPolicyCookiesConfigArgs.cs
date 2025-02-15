// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class OriginRequestPolicyCookiesConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("cookieBehavior", required: true)]
        public Input<string> CookieBehavior { get; set; } = null!;

        [Input("cookies")]
        private InputList<string>? _cookies;
        public InputList<string> Cookies
        {
            get => _cookies ?? (_cookies = new InputList<string>());
            set => _cookies = value;
        }

        public OriginRequestPolicyCookiesConfigArgs()
        {
        }
        public static new OriginRequestPolicyCookiesConfigArgs Empty => new OriginRequestPolicyCookiesConfigArgs();
    }
}
