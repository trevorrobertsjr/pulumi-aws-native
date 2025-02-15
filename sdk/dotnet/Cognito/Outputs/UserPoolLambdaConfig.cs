// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolLambdaConfig
    {
        public readonly string? CreateAuthChallenge;
        public readonly Outputs.UserPoolCustomEmailSender? CustomEmailSender;
        public readonly string? CustomMessage;
        public readonly Outputs.UserPoolCustomSmsSender? CustomSmsSender;
        public readonly string? DefineAuthChallenge;
        public readonly string? KmsKeyId;
        public readonly string? PostAuthentication;
        public readonly string? PostConfirmation;
        public readonly string? PreAuthentication;
        public readonly string? PreSignUp;
        public readonly string? PreTokenGeneration;
        public readonly string? UserMigration;
        public readonly string? VerifyAuthChallengeResponse;

        [OutputConstructor]
        private UserPoolLambdaConfig(
            string? createAuthChallenge,

            Outputs.UserPoolCustomEmailSender? customEmailSender,

            string? customMessage,

            Outputs.UserPoolCustomSmsSender? customSmsSender,

            string? defineAuthChallenge,

            string? kmsKeyId,

            string? postAuthentication,

            string? postConfirmation,

            string? preAuthentication,

            string? preSignUp,

            string? preTokenGeneration,

            string? userMigration,

            string? verifyAuthChallengeResponse)
        {
            CreateAuthChallenge = createAuthChallenge;
            CustomEmailSender = customEmailSender;
            CustomMessage = customMessage;
            CustomSmsSender = customSmsSender;
            DefineAuthChallenge = defineAuthChallenge;
            KmsKeyId = kmsKeyId;
            PostAuthentication = postAuthentication;
            PostConfirmation = postConfirmation;
            PreAuthentication = preAuthentication;
            PreSignUp = preSignUp;
            PreTokenGeneration = preTokenGeneration;
            UserMigration = userMigration;
            VerifyAuthChallengeResponse = verifyAuthChallengeResponse;
        }
    }
}
