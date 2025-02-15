// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Schema for AWS::SNS::TopicInlinePolicy
 */
export class TopicInlinePolicy extends pulumi.CustomResource {
    /**
     * Get an existing TopicInlinePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TopicInlinePolicy {
        return new TopicInlinePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sns:TopicInlinePolicy';

    /**
     * Returns true if the given object is an instance of TopicInlinePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicInlinePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicInlinePolicy.__pulumiType;
    }

    /**
     * A policy document that contains permissions to add to the specified SNS topics.
     */
    public readonly policyDocument!: pulumi.Output<any>;
    /**
     * The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
     */
    public readonly topicArn!: pulumi.Output<string>;

    /**
     * Create a TopicInlinePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicInlinePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.policyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyDocument'");
            }
            if ((!args || args.topicArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicArn'");
            }
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["topicArn"] = args ? args.topicArn : undefined;
        } else {
            resourceInputs["policyDocument"] = undefined /*out*/;
            resourceInputs["topicArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["topicArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TopicInlinePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TopicInlinePolicy resource.
 */
export interface TopicInlinePolicyArgs {
    /**
     * A policy document that contains permissions to add to the specified SNS topics.
     */
    policyDocument: any;
    /**
     * The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
     */
    topicArn: pulumi.Input<string>;
}
