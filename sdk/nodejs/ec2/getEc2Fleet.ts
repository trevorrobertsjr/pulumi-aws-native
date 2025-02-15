// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::EC2Fleet
 */
export function getEc2Fleet(args: GetEc2FleetArgs, opts?: pulumi.InvokeOptions): Promise<GetEc2FleetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getEc2Fleet", {
        "fleetId": args.fleetId,
    }, opts);
}

export interface GetEc2FleetArgs {
    fleetId: string;
}

export interface GetEc2FleetResult {
    readonly context?: string;
    readonly excessCapacityTerminationPolicy?: enums.ec2.Ec2FleetExcessCapacityTerminationPolicy;
    readonly fleetId?: string;
    readonly targetCapacitySpecification?: outputs.ec2.Ec2FleetTargetCapacitySpecificationRequest;
}
/**
 * Resource Type definition for AWS::EC2::EC2Fleet
 */
export function getEc2FleetOutput(args: GetEc2FleetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEc2FleetResult> {
    return pulumi.output(args).apply((a: any) => getEc2Fleet(a, opts))
}

export interface GetEc2FleetOutputArgs {
    fleetId: pulumi.Input<string>;
}
