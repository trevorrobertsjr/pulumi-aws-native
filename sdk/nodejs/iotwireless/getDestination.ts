// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Destination's resource schema demonstrating some basic constructs and validation rules.
 */
export function getDestination(args: GetDestinationArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotwireless:getDestination", {
        "name": args.name,
    }, opts);
}

export interface GetDestinationArgs {
    /**
     * Unique name of destination
     */
    name: string;
}

export interface GetDestinationResult {
    /**
     * Destination arn. Returned after successful create.
     */
    readonly arn?: string;
    /**
     * Destination description
     */
    readonly description?: string;
    /**
     * Destination expression
     */
    readonly expression?: string;
    /**
     * Must be RuleName
     */
    readonly expressionType?: enums.iotwireless.DestinationExpressionType;
    /**
     * AWS role ARN that grants access
     */
    readonly roleArn?: string;
    /**
     * A list of key-value pairs that contain metadata for the destination.
     */
    readonly tags?: outputs.iotwireless.DestinationTag[];
}
/**
 * Destination's resource schema demonstrating some basic constructs and validation rules.
 */
export function getDestinationOutput(args: GetDestinationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationResult> {
    return pulumi.output(args).apply((a: any) => getDestination(a, opts))
}

export interface GetDestinationOutputArgs {
    /**
     * Unique name of destination
     */
    name: pulumi.Input<string>;
}
