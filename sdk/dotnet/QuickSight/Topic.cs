// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    /// <summary>
    /// Definition of the AWS::QuickSight::Topic Resource Type.
    /// </summary>
    [AwsNativeResourceType("aws-native:quicksight:Topic")]
    public partial class Topic : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("awsAccountId")]
        public Output<string?> AwsAccountId { get; private set; } = null!;

        [Output("dataSets")]
        public Output<ImmutableArray<Outputs.TopicDatasetMetadata>> DataSets { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("topicId")]
        public Output<string?> TopicId { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:quicksight:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:quicksight:Topic", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "awsAccountId",
                    "topicId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, options);
        }
    }

    public sealed class TopicArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        [Input("dataSets")]
        private InputList<Inputs.TopicDatasetMetadataArgs>? _dataSets;
        public InputList<Inputs.TopicDatasetMetadataArgs> DataSets
        {
            get => _dataSets ?? (_dataSets = new InputList<Inputs.TopicDatasetMetadataArgs>());
            set => _dataSets = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public TopicArgs()
        {
        }
        public static new TopicArgs Empty => new TopicArgs();
    }
}
