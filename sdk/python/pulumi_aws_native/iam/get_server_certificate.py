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
    'GetServerCertificateResult',
    'AwaitableGetServerCertificateResult',
    'get_server_certificate',
    'get_server_certificate_output',
]

@pulumi.output_type
class GetServerCertificateResult:
    def __init__(__self__, arn=None, path=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Amazon Resource Name (ARN) of the server certificate
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ServerCertificateTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetServerCertificateResult(GetServerCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerCertificateResult(
            arn=self.arn,
            path=self.path,
            tags=self.tags)


def get_server_certificate(server_certificate_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerCertificateResult:
    """
    Resource Type definition for AWS::IAM::ServerCertificate
    """
    __args__ = dict()
    __args__['serverCertificateName'] = server_certificate_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iam:getServerCertificate', __args__, opts=opts, typ=GetServerCertificateResult).value

    return AwaitableGetServerCertificateResult(
        arn=pulumi.get(__ret__, 'arn'),
        path=pulumi.get(__ret__, 'path'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_server_certificate)
def get_server_certificate_output(server_certificate_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerCertificateResult]:
    """
    Resource Type definition for AWS::IAM::ServerCertificate
    """
    ...
