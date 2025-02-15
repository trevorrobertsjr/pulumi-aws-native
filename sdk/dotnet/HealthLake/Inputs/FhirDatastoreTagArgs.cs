// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.HealthLake.Inputs
{

    /// <summary>
    /// A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag values are both required, but tag values can be empty (null) strings.
    /// </summary>
    public sealed class FhirDatastoreTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the tag.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value of the tag.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public FhirDatastoreTagArgs()
        {
        }
        public static new FhirDatastoreTagArgs Empty => new FhirDatastoreTagArgs();
    }
}
