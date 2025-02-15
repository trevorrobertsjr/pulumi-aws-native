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
    'GetHostedConfigurationVersionResult',
    'AwaitableGetHostedConfigurationVersionResult',
    'get_hosted_configuration_version',
    'get_hosted_configuration_version_output',
]

@pulumi.output_type
class GetHostedConfigurationVersionResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetHostedConfigurationVersionResult(GetHostedConfigurationVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostedConfigurationVersionResult(
            id=self.id)


def get_hosted_configuration_version(id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostedConfigurationVersionResult:
    """
    Resource Type definition for AWS::AppConfig::HostedConfigurationVersion
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appconfig:getHostedConfigurationVersion', __args__, opts=opts, typ=GetHostedConfigurationVersionResult).value

    return AwaitableGetHostedConfigurationVersionResult(
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_hosted_configuration_version)
def get_hosted_configuration_version_output(id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHostedConfigurationVersionResult]:
    """
    Resource Type definition for AWS::AppConfig::HostedConfigurationVersion
    """
    ...
