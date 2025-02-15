// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// Resource Type definition for AWS::SageMaker::Image
    /// </summary>
    [AwsNativeResourceType("aws-native:sagemaker:Image")]
    public partial class Image : global::Pulumi.CustomResource
    {
        [Output("imageArn")]
        public Output<string> ImageArn { get; private set; } = null!;

        [Output("imageDescription")]
        public Output<string?> ImageDescription { get; private set; } = null!;

        [Output("imageDisplayName")]
        public Output<string?> ImageDisplayName { get; private set; } = null!;

        [Output("imageName")]
        public Output<string> ImageName { get; private set; } = null!;

        [Output("imageRoleArn")]
        public Output<string> ImageRoleArn { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ImageTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:Image", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "imageName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Image(name, id, options);
        }
    }

    public sealed class ImageArgs : global::Pulumi.ResourceArgs
    {
        [Input("imageDescription")]
        public Input<string>? ImageDescription { get; set; }

        [Input("imageDisplayName")]
        public Input<string>? ImageDisplayName { get; set; }

        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        [Input("imageRoleArn", required: true)]
        public Input<string> ImageRoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ImageTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.ImageTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ImageTagArgs>());
            set => _tags = value;
        }

        public ImageArgs()
        {
        }
        public static new ImageArgs Empty => new ImageArgs();
    }
}
