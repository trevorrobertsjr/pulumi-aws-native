// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// The automation property name of the question.
    /// </summary>
    [OutputType]
    public sealed class EvaluationFormNumericQuestionPropertyValueAutomation
    {
        /// <summary>
        /// The automation property label.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.EvaluationFormNumericQuestionPropertyValueAutomationLabel Label;

        [OutputConstructor]
        private EvaluationFormNumericQuestionPropertyValueAutomation(Pulumi.AwsNative.Connect.EvaluationFormNumericQuestionPropertyValueAutomationLabel label)
        {
            Label = label;
        }
    }
}
