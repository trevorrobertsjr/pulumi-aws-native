// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.
 */
export class Destination extends pulumi.CustomResource {
    /**
     * Get an existing Destination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Destination {
        return new Destination(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:logs:Destination';

    /**
     * Returns true if the given object is an instance of Destination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Destination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Destination.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the destination resource
     */
    public readonly destinationName!: pulumi.Output<string>;
    /**
     * An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
     */
    public readonly destinationPolicy!: pulumi.Output<string | undefined>;
    /**
     * The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)
     */
    public readonly targetArn!: pulumi.Output<string>;

    /**
     * Create a Destination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DestinationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.targetArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetArn'");
            }
            resourceInputs["destinationName"] = args ? args.destinationName : undefined;
            resourceInputs["destinationPolicy"] = args ? args.destinationPolicy : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["targetArn"] = args ? args.targetArn : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["destinationName"] = undefined /*out*/;
            resourceInputs["destinationPolicy"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["targetArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["destinationName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Destination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Destination resource.
 */
export interface DestinationArgs {
    /**
     * The name of the destination resource
     */
    destinationName?: pulumi.Input<string>;
    /**
     * An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
     */
    destinationPolicy?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
     */
    roleArn: pulumi.Input<string>;
    /**
     * The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)
     */
    targetArn: pulumi.Input<string>;
}
