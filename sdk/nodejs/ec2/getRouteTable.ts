// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::RouteTable
 */
export function getRouteTable(args: GetRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTableResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getRouteTable", {
        "routeTableId": args.routeTableId,
    }, opts);
}

export interface GetRouteTableArgs {
    /**
     * The route table ID.
     */
    routeTableId: string;
}

export interface GetRouteTableResult {
    /**
     * The route table ID.
     */
    readonly routeTableId?: string;
    /**
     * Any tags assigned to the route table.
     */
    readonly tags?: outputs.ec2.RouteTableTag[];
}
/**
 * Resource Type definition for AWS::EC2::RouteTable
 */
export function getRouteTableOutput(args: GetRouteTableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteTableResult> {
    return pulumi.output(args).apply((a: any) => getRouteTable(a, opts))
}

export interface GetRouteTableOutputArgs {
    /**
     * The route table ID.
     */
    routeTableId: pulumi.Input<string>;
}
