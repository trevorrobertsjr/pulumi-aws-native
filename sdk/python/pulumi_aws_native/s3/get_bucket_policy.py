# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetBucketPolicyResult',
    'AwaitableGetBucketPolicyResult',
    'get_bucket_policy',
    'get_bucket_policy_output',
]

@pulumi.output_type
class GetBucketPolicyResult:
    def __init__(__self__, policy_document=None):
        if policy_document and not isinstance(policy_document, dict):
            raise TypeError("Expected argument 'policy_document' to be a dict")
        pulumi.set(__self__, "policy_document", policy_document)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> Optional[Any]:
        """
        A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
        """
        return pulumi.get(self, "policy_document")


class AwaitableGetBucketPolicyResult(GetBucketPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketPolicyResult(
            policy_document=self.policy_document)


def get_bucket_policy(bucket: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBucketPolicyResult:
    """
    Resource Type definition for AWS::S3::BucketPolicy


    :param str bucket: The name of the Amazon S3 bucket to which the policy applies.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:s3:getBucketPolicy', __args__, opts=opts, typ=GetBucketPolicyResult).value

    return AwaitableGetBucketPolicyResult(
        policy_document=pulumi.get(__ret__, 'policy_document'))


@_utilities.lift_output_func(get_bucket_policy)
def get_bucket_policy_output(bucket: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBucketPolicyResult]:
    """
    Resource Type definition for AWS::S3::BucketPolicy


    :param str bucket: The name of the Amazon S3 bucket to which the policy applies.
    """
    ...
