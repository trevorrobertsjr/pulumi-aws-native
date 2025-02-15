// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rekognition
{
    /// <summary>
    /// The AWS::Rekognition::StreamProcessor type is used to create an Amazon Rekognition StreamProcessor that you can use to analyze streaming videos.
    /// </summary>
    [AwsNativeResourceType("aws-native:rekognition:StreamProcessor")]
    public partial class StreamProcessor : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The BoundingBoxRegionsOfInterest specifies an array of bounding boxes of interest in the video frames to analyze, as part of connected home feature. If an object is partially in a region of interest, Rekognition will tag it as detected if the overlap of the object with the region-of-interest is greater than 20%.
        /// </summary>
        [Output("boundingBoxRegionsOfInterest")]
        public Output<ImmutableArray<Outputs.StreamProcessorBoundingBox>> BoundingBoxRegionsOfInterest { get; private set; } = null!;

        [Output("connectedHomeSettings")]
        public Output<Outputs.StreamProcessorConnectedHomeSettings?> ConnectedHomeSettings { get; private set; } = null!;

        [Output("dataSharingPreference")]
        public Output<Outputs.StreamProcessorDataSharingPreference?> DataSharingPreference { get; private set; } = null!;

        [Output("faceSearchSettings")]
        public Output<Outputs.StreamProcessorFaceSearchSettings?> FaceSearchSettings { get; private set; } = null!;

        [Output("kinesisDataStream")]
        public Output<Outputs.StreamProcessorKinesisDataStream?> KinesisDataStream { get; private set; } = null!;

        [Output("kinesisVideoStream")]
        public Output<Outputs.StreamProcessorKinesisVideoStream> KinesisVideoStream { get; private set; } = null!;

        /// <summary>
        /// The KMS key that is used by Rekognition to encrypt any intermediate customer metadata and store in the customer's S3 bucket.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("notificationChannel")]
        public Output<Outputs.StreamProcessorNotificationChannel?> NotificationChannel { get; private set; } = null!;

        /// <summary>
        /// The PolygonRegionsOfInterest specifies a set of polygon areas of interest in the video frames to analyze, as part of connected home feature. Each polygon is in turn, an ordered list of Point
        /// </summary>
        [Output("polygonRegionsOfInterest")]
        public Output<ImmutableArray<ImmutableArray<Outputs.StreamProcessorPoint>>> PolygonRegionsOfInterest { get; private set; } = null!;

        /// <summary>
        /// ARN of the IAM role that allows access to the stream processor, and provides Rekognition read permissions for KVS stream and write permissions to S3 bucket and SNS topic.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("s3Destination")]
        public Output<Outputs.StreamProcessorS3Destination?> S3Destination { get; private set; } = null!;

        /// <summary>
        /// Current status of the stream processor.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Detailed status message about the stream processor.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.StreamProcessorTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a StreamProcessor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamProcessor(string name, StreamProcessorArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rekognition:StreamProcessor", name, args ?? new StreamProcessorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamProcessor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rekognition:StreamProcessor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "boundingBoxRegionsOfInterest[*]",
                    "connectedHomeSettings",
                    "dataSharingPreference",
                    "faceSearchSettings",
                    "kinesisDataStream",
                    "kinesisVideoStream",
                    "kmsKeyId",
                    "name",
                    "notificationChannel",
                    "polygonRegionsOfInterest[*]",
                    "roleArn",
                    "s3Destination",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamProcessor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamProcessor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StreamProcessor(name, id, options);
        }
    }

    public sealed class StreamProcessorArgs : global::Pulumi.ResourceArgs
    {
        [Input("boundingBoxRegionsOfInterest")]
        private InputList<Inputs.StreamProcessorBoundingBoxArgs>? _boundingBoxRegionsOfInterest;

        /// <summary>
        /// The BoundingBoxRegionsOfInterest specifies an array of bounding boxes of interest in the video frames to analyze, as part of connected home feature. If an object is partially in a region of interest, Rekognition will tag it as detected if the overlap of the object with the region-of-interest is greater than 20%.
        /// </summary>
        public InputList<Inputs.StreamProcessorBoundingBoxArgs> BoundingBoxRegionsOfInterest
        {
            get => _boundingBoxRegionsOfInterest ?? (_boundingBoxRegionsOfInterest = new InputList<Inputs.StreamProcessorBoundingBoxArgs>());
            set => _boundingBoxRegionsOfInterest = value;
        }

        [Input("connectedHomeSettings")]
        public Input<Inputs.StreamProcessorConnectedHomeSettingsArgs>? ConnectedHomeSettings { get; set; }

        [Input("dataSharingPreference")]
        public Input<Inputs.StreamProcessorDataSharingPreferenceArgs>? DataSharingPreference { get; set; }

        [Input("faceSearchSettings")]
        public Input<Inputs.StreamProcessorFaceSearchSettingsArgs>? FaceSearchSettings { get; set; }

        [Input("kinesisDataStream")]
        public Input<Inputs.StreamProcessorKinesisDataStreamArgs>? KinesisDataStream { get; set; }

        [Input("kinesisVideoStream", required: true)]
        public Input<Inputs.StreamProcessorKinesisVideoStreamArgs> KinesisVideoStream { get; set; } = null!;

        /// <summary>
        /// The KMS key that is used by Rekognition to encrypt any intermediate customer metadata and store in the customer's S3 bucket.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannel")]
        public Input<Inputs.StreamProcessorNotificationChannelArgs>? NotificationChannel { get; set; }

        [Input("polygonRegionsOfInterest")]
        private InputList<ImmutableArray<Inputs.StreamProcessorPointArgs>>? _polygonRegionsOfInterest;

        /// <summary>
        /// The PolygonRegionsOfInterest specifies a set of polygon areas of interest in the video frames to analyze, as part of connected home feature. Each polygon is in turn, an ordered list of Point
        /// </summary>
        public InputList<ImmutableArray<Inputs.StreamProcessorPointArgs>> PolygonRegionsOfInterest
        {
            get => _polygonRegionsOfInterest ?? (_polygonRegionsOfInterest = new InputList<ImmutableArray<Inputs.StreamProcessorPointArgs>>());
            set => _polygonRegionsOfInterest = value;
        }

        /// <summary>
        /// ARN of the IAM role that allows access to the stream processor, and provides Rekognition read permissions for KVS stream and write permissions to S3 bucket and SNS topic.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("s3Destination")]
        public Input<Inputs.StreamProcessorS3DestinationArgs>? S3Destination { get; set; }

        [Input("tags")]
        private InputList<Inputs.StreamProcessorTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.StreamProcessorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StreamProcessorTagArgs>());
            set => _tags = value;
        }

        public StreamProcessorArgs()
        {
        }
        public static new StreamProcessorArgs Empty => new StreamProcessorArgs();
    }
}
