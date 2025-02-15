// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketLifecycleConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.BucketRuleArgs>? _rules;

        /// <summary>
        /// A lifecycle rule for individual objects in an Amazon S3 bucket.
        /// </summary>
        public InputList<Inputs.BucketRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketRuleArgs>());
            set => _rules = value;
        }

        public BucketLifecycleConfigurationArgs()
        {
        }
        public static new BucketLifecycleConfigurationArgs Empty => new BucketLifecycleConfigurationArgs();
    }
}
