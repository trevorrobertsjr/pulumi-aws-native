// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::Policy
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetPolicyArgs {
    /**
     * The provider-assigned unique ID for this resource
     */
    id: string;
}

export interface GetPolicyResult {
    /**
     * The name of the group to associate the policy with. This parameter allows (through its regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
     */
    readonly groups?: string[];
    /**
     * The provider-assigned unique ID for this resource
     */
    readonly id?: string;
    /**
     * The policy document. You must provide policies in JSON format in IAM. However, for AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.
     */
    readonly policyDocument?: any;
    /**
     * The name of the policy document. This parameter allows (through its regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
     */
    readonly policyName?: string;
    /**
     * The name of the role to associate the policy with. This parameter allows (per its regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
     */
    readonly roles?: string[];
    /**
     * The name of the user to associate the policy with. This parameter allows (through its regex pattern) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
     */
    readonly users?: string[];
}
/**
 * Resource Type definition for AWS::IAM::Policy
 */
export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyResult> {
    return pulumi.output(args).apply((a: any) => getPolicy(a, opts))
}

export interface GetPolicyOutputArgs {
    /**
     * The provider-assigned unique ID for this resource
     */
    id: pulumi.Input<string>;
}
