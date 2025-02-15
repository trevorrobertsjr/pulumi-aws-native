// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::AppBlock
 */
export function getAppBlock(args: GetAppBlockArgs, opts?: pulumi.InvokeOptions): Promise<GetAppBlockResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appstream:getAppBlock", {
        "arn": args.arn,
    }, opts);
}

export interface GetAppBlockArgs {
    arn: string;
}

export interface GetAppBlockResult {
    readonly arn?: string;
    readonly createdTime?: string;
}
/**
 * Resource Type definition for AWS::AppStream::AppBlock
 */
export function getAppBlockOutput(args: GetAppBlockOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppBlockResult> {
    return pulumi.output(args).apply((a: any) => getAppBlock(a, opts))
}

export interface GetAppBlockOutputArgs {
    arn: pulumi.Input<string>;
}
