// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Schema for IAM Role Policy
 */
export class RolePolicy extends pulumi.CustomResource {
    /**
     * Get an existing RolePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RolePolicy {
        return new RolePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:RolePolicy';

    /**
     * Returns true if the given object is an instance of RolePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RolePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RolePolicy.__pulumiType;
    }

    /**
     * The policy document.
     */
    public readonly policyDocument!: pulumi.Output<any | undefined>;
    /**
     * The friendly name (not ARN) identifying the policy.
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * The name of the policy document.
     */
    public readonly roleName!: pulumi.Output<string>;

    /**
     * Create a RolePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RolePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
        } else {
            resourceInputs["policyDocument"] = undefined /*out*/;
            resourceInputs["policyName"] = undefined /*out*/;
            resourceInputs["roleName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["policyName", "roleName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RolePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RolePolicy resource.
 */
export interface RolePolicyArgs {
    /**
     * The policy document.
     */
    policyDocument?: any;
    /**
     * The friendly name (not ARN) identifying the policy.
     */
    policyName: pulumi.Input<string>;
    /**
     * The name of the policy document.
     */
    roleName: pulumi.Input<string>;
}
