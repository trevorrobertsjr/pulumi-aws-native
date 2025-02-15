// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT::TopicRuleDestination
 */
export class TopicRuleDestination extends pulumi.CustomResource {
    /**
     * Get an existing TopicRuleDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TopicRuleDestination {
        return new TopicRuleDestination(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iot:TopicRuleDestination';

    /**
     * Returns true if the given object is an instance of TopicRuleDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicRuleDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicRuleDestination.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * HTTP URL destination properties.
     */
    public readonly httpUrlProperties!: pulumi.Output<outputs.iot.TopicRuleDestinationHttpUrlDestinationSummary | undefined>;
    /**
     * The status of the TopicRuleDestination.
     */
    public readonly status!: pulumi.Output<enums.iot.TopicRuleDestinationStatus | undefined>;
    /**
     * The reasoning for the current status of the TopicRuleDestination.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;
    /**
     * VPC destination properties.
     */
    public readonly vpcProperties!: pulumi.Output<outputs.iot.TopicRuleDestinationVpcDestinationProperties | undefined>;

    /**
     * Create a TopicRuleDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TopicRuleDestinationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["httpUrlProperties"] = args ? args.httpUrlProperties : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vpcProperties"] = args ? args.vpcProperties : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["httpUrlProperties"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
            resourceInputs["vpcProperties"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["httpUrlProperties", "vpcProperties"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TopicRuleDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TopicRuleDestination resource.
 */
export interface TopicRuleDestinationArgs {
    /**
     * HTTP URL destination properties.
     */
    httpUrlProperties?: pulumi.Input<inputs.iot.TopicRuleDestinationHttpUrlDestinationSummaryArgs>;
    /**
     * The status of the TopicRuleDestination.
     */
    status?: pulumi.Input<enums.iot.TopicRuleDestinationStatus>;
    /**
     * VPC destination properties.
     */
    vpcProperties?: pulumi.Input<inputs.iot.TopicRuleDestinationVpcDestinationPropertiesArgs>;
}
