// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::Image
 */
export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:imagebuilder:getImage", {
        "arn": args.arn,
    }, opts);
}

export interface GetImageArgs {
    /**
     * The Amazon Resource Name (ARN) of the image.
     */
    arn: string;
}

export interface GetImageResult {
    /**
     * The Amazon Resource Name (ARN) of the image.
     */
    readonly arn?: string;
    /**
     * The execution role name/ARN for the image build, if provided
     */
    readonly executionRole?: string;
    /**
     * The AMI ID of the EC2 AMI in current region.
     */
    readonly imageId?: string;
    /**
     * URI for containers created in current Region with default ECR image tag
     */
    readonly imageUri?: string;
    /**
     * The name of the image.
     */
    readonly name?: string;
}
/**
 * Resource schema for AWS::ImageBuilder::Image
 */
export function getImageOutput(args: GetImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageResult> {
    return pulumi.output(args).apply((a: any) => getImage(a, opts))
}

export interface GetImageOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the image.
     */
    arn: pulumi.Input<string>;
}
