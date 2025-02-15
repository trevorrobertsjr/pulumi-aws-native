// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::InstanceProfile
 */
export function getInstanceProfile(args: GetInstanceProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceProfileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getInstanceProfile", {
        "instanceProfileName": args.instanceProfileName,
    }, opts);
}

export interface GetInstanceProfileArgs {
    /**
     * The name of the instance profile to create.
     */
    instanceProfileName: string;
}

export interface GetInstanceProfileResult {
    /**
     * The Amazon Resource Name (ARN) of the instance profile.
     */
    readonly arn?: string;
    /**
     * The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
     */
    readonly roles?: string[];
}
/**
 * Resource Type definition for AWS::IAM::InstanceProfile
 */
export function getInstanceProfileOutput(args: GetInstanceProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceProfileResult> {
    return pulumi.output(args).apply((a: any) => getInstanceProfile(a, opts))
}

export interface GetInstanceProfileOutputArgs {
    /**
     * The name of the instance profile to create.
     */
    instanceProfileName: pulumi.Input<string>;
}
