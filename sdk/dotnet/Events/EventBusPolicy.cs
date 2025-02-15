// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events
{
    /// <summary>
    /// Resource Type definition for AWS::Events::EventBusPolicy
    /// </summary>
    [Obsolete(@"EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:events:EventBusPolicy")]
    public partial class EventBusPolicy : global::Pulumi.CustomResource
    {
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.EventBusPolicyCondition?> Condition { get; private set; } = null!;

        [Output("eventBusName")]
        public Output<string?> EventBusName { get; private set; } = null!;

        [Output("principal")]
        public Output<string?> Principal { get; private set; } = null!;

        [Output("statement")]
        public Output<object?> Statement { get; private set; } = null!;

        [Output("statementId")]
        public Output<string> StatementId { get; private set; } = null!;


        /// <summary>
        /// Create a EventBusPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventBusPolicy(string name, EventBusPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:events:EventBusPolicy", name, args ?? new EventBusPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventBusPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:events:EventBusPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "eventBusName",
                    "statementId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventBusPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventBusPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventBusPolicy(name, id, options);
        }
    }

    public sealed class EventBusPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("condition")]
        public Input<Inputs.EventBusPolicyConditionArgs>? Condition { get; set; }

        [Input("eventBusName")]
        public Input<string>? EventBusName { get; set; }

        [Input("principal")]
        public Input<string>? Principal { get; set; }

        [Input("statement")]
        public Input<object>? Statement { get; set; }

        [Input("statementId", required: true)]
        public Input<string> StatementId { get; set; } = null!;

        public EventBusPolicyArgs()
        {
        }
        public static new EventBusPolicyArgs Empty => new EventBusPolicyArgs();
    }
}
