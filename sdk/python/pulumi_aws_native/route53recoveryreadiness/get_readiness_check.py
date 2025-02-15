# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetReadinessCheckResult',
    'AwaitableGetReadinessCheckResult',
    'get_readiness_check',
    'get_readiness_check_output',
]

@pulumi.output_type
class GetReadinessCheckResult:
    def __init__(__self__, readiness_check_arn=None, resource_set_name=None, tags=None):
        if readiness_check_arn and not isinstance(readiness_check_arn, str):
            raise TypeError("Expected argument 'readiness_check_arn' to be a str")
        pulumi.set(__self__, "readiness_check_arn", readiness_check_arn)
        if resource_set_name and not isinstance(resource_set_name, str):
            raise TypeError("Expected argument 'resource_set_name' to be a str")
        pulumi.set(__self__, "resource_set_name", resource_set_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="readinessCheckArn")
    def readiness_check_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the readiness check.
        """
        return pulumi.get(self, "readiness_check_arn")

    @property
    @pulumi.getter(name="resourceSetName")
    def resource_set_name(self) -> Optional[str]:
        """
        The name of the resource set to check.
        """
        return pulumi.get(self, "resource_set_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ReadinessCheckTag']]:
        """
        A collection of tags associated with a resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetReadinessCheckResult(GetReadinessCheckResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReadinessCheckResult(
            readiness_check_arn=self.readiness_check_arn,
            resource_set_name=self.resource_set_name,
            tags=self.tags)


def get_readiness_check(readiness_check_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReadinessCheckResult:
    """
    Aws Route53 Recovery Readiness Check Schema and API specification.


    :param str readiness_check_name: Name of the ReadinessCheck to create.
    """
    __args__ = dict()
    __args__['readinessCheckName'] = readiness_check_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:route53recoveryreadiness:getReadinessCheck', __args__, opts=opts, typ=GetReadinessCheckResult).value

    return AwaitableGetReadinessCheckResult(
        readiness_check_arn=pulumi.get(__ret__, 'readiness_check_arn'),
        resource_set_name=pulumi.get(__ret__, 'resource_set_name'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_readiness_check)
def get_readiness_check_output(readiness_check_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReadinessCheckResult]:
    """
    Aws Route53 Recovery Readiness Check Schema and API specification.


    :param str readiness_check_name: Name of the ReadinessCheck to create.
    """
    ...
