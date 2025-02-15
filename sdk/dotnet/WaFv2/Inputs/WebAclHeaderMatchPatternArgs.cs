// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// The pattern to look for in the request headers.
    /// </summary>
    public sealed class WebAclHeaderMatchPatternArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Inspect all parts of the web request headers.
        /// </summary>
        [Input("all")]
        public Input<object>? All { get; set; }

        [Input("excludedHeaders")]
        private InputList<string>? _excludedHeaders;
        public InputList<string> ExcludedHeaders
        {
            get => _excludedHeaders ?? (_excludedHeaders = new InputList<string>());
            set => _excludedHeaders = value;
        }

        [Input("includedHeaders")]
        private InputList<string>? _includedHeaders;
        public InputList<string> IncludedHeaders
        {
            get => _includedHeaders ?? (_includedHeaders = new InputList<string>());
            set => _includedHeaders = value;
        }

        public WebAclHeaderMatchPatternArgs()
        {
        }
        public static new WebAclHeaderMatchPatternArgs Empty => new WebAclHeaderMatchPatternArgs();
    }
}
