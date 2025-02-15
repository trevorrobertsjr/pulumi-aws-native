// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT1Click::Project
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot1click:getProject", {
        "id": args.id,
    }, opts);
}

export interface GetProjectArgs {
    id: string;
}

export interface GetProjectResult {
    readonly arn?: string;
    readonly description?: string;
    readonly id?: string;
    readonly placementTemplate?: outputs.iot1click.ProjectPlacementTemplate;
}
/**
 * Resource Type definition for AWS::IoT1Click::Project
 */
export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply((a: any) => getProject(a, opts))
}

export interface GetProjectOutputArgs {
    id: pulumi.Input<string>;
}
