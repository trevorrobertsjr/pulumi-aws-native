// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const LinkResourceType = {
    AwsCloudWatchMetric: "AWS::CloudWatch::Metric",
    AwsLogsLogGroup: "AWS::Logs::LogGroup",
    AwsxRayTrace: "AWS::XRay::Trace",
    AwsApplicationInsightsApplication: "AWS::ApplicationInsights::Application",
} as const;

export type LinkResourceType = (typeof LinkResourceType)[keyof typeof LinkResourceType];
