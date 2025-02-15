// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Detective.Inputs
{

    /// <summary>
    /// A key-value pair to associate with a resource.
    /// </summary>
    public sealed class GraphTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @ 
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @ 
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GraphTagArgs()
        {
        }
        public static new GraphTagArgs Empty => new GraphTagArgs();
    }
}
