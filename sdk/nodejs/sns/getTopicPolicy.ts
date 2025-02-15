// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Schema for AWS::SNS::TopicPolicy
 */
export function getTopicPolicy(args: GetTopicPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sns:getTopicPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetTopicPolicyArgs {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    id: string;
}

export interface GetTopicPolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id?: string;
    /**
     * A policy document that contains permissions to add to the specified SNS topics.
     */
    readonly policyDocument?: any;
    /**
     * The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` function to specify an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) resource.
     */
    readonly topics?: string[];
}
/**
 * Schema for AWS::SNS::TopicPolicy
 */
export function getTopicPolicyOutput(args: GetTopicPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicPolicyResult> {
    return pulumi.output(args).apply((a: any) => getTopicPolicy(a, opts))
}

export interface GetTopicPolicyOutputArgs {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    id: pulumi.Input<string>;
}
