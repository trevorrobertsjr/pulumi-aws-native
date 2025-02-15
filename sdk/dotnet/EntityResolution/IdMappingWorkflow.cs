// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution
{
    /// <summary>
    /// IdMappingWorkflow defined in AWS Entity Resolution service
    /// </summary>
    [AwsNativeResourceType("aws-native:entityresolution:IdMappingWorkflow")]
    public partial class IdMappingWorkflow : global::Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the IdMappingWorkflow
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("idMappingTechniques")]
        public Output<Outputs.IdMappingWorkflowIdMappingTechniques> IdMappingTechniques { get; private set; } = null!;

        [Output("inputSourceConfig")]
        public Output<ImmutableArray<Outputs.IdMappingWorkflowInputSource>> InputSourceConfig { get; private set; } = null!;

        [Output("outputSourceConfig")]
        public Output<ImmutableArray<Outputs.IdMappingWorkflowOutputSource>> OutputSourceConfig { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.IdMappingWorkflowTag>> Tags { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        [Output("workflowArn")]
        public Output<string> WorkflowArn { get; private set; } = null!;

        /// <summary>
        /// The name of the IdMappingWorkflow
        /// </summary>
        [Output("workflowName")]
        public Output<string> WorkflowName { get; private set; } = null!;


        /// <summary>
        /// Create a IdMappingWorkflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdMappingWorkflow(string name, IdMappingWorkflowArgs args, CustomResourceOptions? options = null)
            : base("aws-native:entityresolution:IdMappingWorkflow", name, args ?? new IdMappingWorkflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdMappingWorkflow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:entityresolution:IdMappingWorkflow", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "workflowName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdMappingWorkflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdMappingWorkflow Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IdMappingWorkflow(name, id, options);
        }
    }

    public sealed class IdMappingWorkflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the IdMappingWorkflow
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("idMappingTechniques", required: true)]
        public Input<Inputs.IdMappingWorkflowIdMappingTechniquesArgs> IdMappingTechniques { get; set; } = null!;

        [Input("inputSourceConfig", required: true)]
        private InputList<Inputs.IdMappingWorkflowInputSourceArgs>? _inputSourceConfig;
        public InputList<Inputs.IdMappingWorkflowInputSourceArgs> InputSourceConfig
        {
            get => _inputSourceConfig ?? (_inputSourceConfig = new InputList<Inputs.IdMappingWorkflowInputSourceArgs>());
            set => _inputSourceConfig = value;
        }

        [Input("outputSourceConfig", required: true)]
        private InputList<Inputs.IdMappingWorkflowOutputSourceArgs>? _outputSourceConfig;
        public InputList<Inputs.IdMappingWorkflowOutputSourceArgs> OutputSourceConfig
        {
            get => _outputSourceConfig ?? (_outputSourceConfig = new InputList<Inputs.IdMappingWorkflowOutputSourceArgs>());
            set => _outputSourceConfig = value;
        }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.IdMappingWorkflowTagArgs>? _tags;
        public InputList<Inputs.IdMappingWorkflowTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.IdMappingWorkflowTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the IdMappingWorkflow
        /// </summary>
        [Input("workflowName", required: true)]
        public Input<string> WorkflowName { get; set; } = null!;

        public IdMappingWorkflowArgs()
        {
        }
        public static new IdMappingWorkflowArgs Empty => new IdMappingWorkflowArgs();
    }
}
