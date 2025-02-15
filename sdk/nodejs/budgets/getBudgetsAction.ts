// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getBudgetsAction(args: GetBudgetsActionArgs, opts?: pulumi.InvokeOptions): Promise<GetBudgetsActionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:budgets:getBudgetsAction", {
        "actionId": args.actionId,
        "budgetName": args.budgetName,
    }, opts);
}

export interface GetBudgetsActionArgs {
    actionId: string;
    budgetName: string;
}

export interface GetBudgetsActionResult {
    readonly actionId?: string;
    readonly actionThreshold?: outputs.budgets.BudgetsActionActionThreshold;
    readonly approvalModel?: enums.budgets.BudgetsActionApprovalModel;
    readonly definition?: outputs.budgets.BudgetsActionDefinition;
    readonly executionRoleArn?: string;
    readonly notificationType?: enums.budgets.BudgetsActionNotificationType;
    readonly subscribers?: outputs.budgets.BudgetsActionSubscriber[];
}
/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getBudgetsActionOutput(args: GetBudgetsActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBudgetsActionResult> {
    return pulumi.output(args).apply((a: any) => getBudgetsAction(a, opts))
}

export interface GetBudgetsActionOutputArgs {
    actionId: pulumi.Input<string>;
    budgetName: pulumi.Input<string>;
}
