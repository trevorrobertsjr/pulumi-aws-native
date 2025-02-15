// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms
{
    /// <summary>
    /// Represents a stored analysis within a collaboration
    /// </summary>
    [AwsNativeResourceType("aws-native:cleanrooms:AnalysisTemplate")]
    public partial class AnalysisTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The member who can query can provide this placeholder for a literal data value in an analysis template
        /// </summary>
        [Output("analysisParameters")]
        public Output<ImmutableArray<Outputs.AnalysisTemplateAnalysisParameter>> AnalysisParameters { get; private set; } = null!;

        [Output("analysisTemplateIdentifier")]
        public Output<string> AnalysisTemplateIdentifier { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("collaborationArn")]
        public Output<string> CollaborationArn { get; private set; } = null!;

        [Output("collaborationIdentifier")]
        public Output<string> CollaborationIdentifier { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("format")]
        public Output<Pulumi.AwsNative.CleanRooms.AnalysisTemplateFormat> Format { get; private set; } = null!;

        [Output("membershipArn")]
        public Output<string> MembershipArn { get; private set; } = null!;

        [Output("membershipIdentifier")]
        public Output<string> MembershipIdentifier { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("schema")]
        public Output<Outputs.AnalysisTemplateAnalysisSchema> Schema { get; private set; } = null!;

        [Output("source")]
        public Output<Outputs.AnalysisTemplateAnalysisSource> Source { get; private set; } = null!;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.AnalysisTemplateTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AnalysisTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AnalysisTemplate(string name, AnalysisTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cleanrooms:AnalysisTemplate", name, args ?? new AnalysisTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AnalysisTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cleanrooms:AnalysisTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "analysisParameters[*]",
                    "format",
                    "membershipIdentifier",
                    "name",
                    "source",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AnalysisTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AnalysisTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AnalysisTemplate(name, id, options);
        }
    }

    public sealed class AnalysisTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("analysisParameters")]
        private InputList<Inputs.AnalysisTemplateAnalysisParameterArgs>? _analysisParameters;

        /// <summary>
        /// The member who can query can provide this placeholder for a literal data value in an analysis template
        /// </summary>
        public InputList<Inputs.AnalysisTemplateAnalysisParameterArgs> AnalysisParameters
        {
            get => _analysisParameters ?? (_analysisParameters = new InputList<Inputs.AnalysisTemplateAnalysisParameterArgs>());
            set => _analysisParameters = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("format", required: true)]
        public Input<Pulumi.AwsNative.CleanRooms.AnalysisTemplateFormat> Format { get; set; } = null!;

        [Input("membershipIdentifier", required: true)]
        public Input<string> MembershipIdentifier { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("source", required: true)]
        public Input<Inputs.AnalysisTemplateAnalysisSourceArgs> Source { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.AnalysisTemplateTagArgs>? _tags;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
        /// </summary>
        public InputList<Inputs.AnalysisTemplateTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AnalysisTemplateTagArgs>());
            set => _tags = value;
        }

        public AnalysisTemplateArgs()
        {
        }
        public static new AnalysisTemplateArgs Empty => new AnalysisTemplateArgs();
    }
}
