// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A billing group is a set of linked account which belong to the same end customer. It can be seen as a virtual consolidated billing family.
 */
export function getBillingGroup(args: GetBillingGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:billingconductor:getBillingGroup", {
        "arn": args.arn,
    }, opts);
}

export interface GetBillingGroupArgs {
    /**
     * Billing Group ARN
     */
    arn: string;
}

export interface GetBillingGroupResult {
    readonly accountGrouping?: outputs.billingconductor.BillingGroupAccountGrouping;
    /**
     * Billing Group ARN
     */
    readonly arn?: string;
    readonly computationPreference?: outputs.billingconductor.BillingGroupComputationPreference;
    /**
     * Creation timestamp in UNIX epoch time format
     */
    readonly creationTime?: number;
    readonly description?: string;
    /**
     * Latest modified timestamp in UNIX epoch time format
     */
    readonly lastModifiedTime?: number;
    readonly name?: string;
    /**
     * Number of accounts in the billing group
     */
    readonly size?: number;
    readonly status?: enums.billingconductor.BillingGroupStatus;
    readonly statusReason?: string;
    readonly tags?: outputs.billingconductor.BillingGroupTag[];
}
/**
 * A billing group is a set of linked account which belong to the same end customer. It can be seen as a virtual consolidated billing family.
 */
export function getBillingGroupOutput(args: GetBillingGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBillingGroupResult> {
    return pulumi.output(args).apply((a: any) => getBillingGroup(a, opts))
}

export interface GetBillingGroupOutputArgs {
    /**
     * Billing Group ARN
     */
    arn: pulumi.Input<string>;
}
