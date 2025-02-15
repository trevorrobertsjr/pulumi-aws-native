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
    'GetReplicationSetResult',
    'AwaitableGetReplicationSetResult',
    'get_replication_set',
    'get_replication_set_output',
]

@pulumi.output_type
class GetReplicationSetResult:
    def __init__(__self__, arn=None, deletion_protected=None, regions=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if deletion_protected and not isinstance(deletion_protected, bool):
            raise TypeError("Expected argument 'deletion_protected' to be a bool")
        pulumi.set(__self__, "deletion_protected", deletion_protected)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The ARN of the ReplicationSet.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deletionProtected")
    def deletion_protected(self) -> Optional[bool]:
        return pulumi.get(self, "deletion_protected")

    @property
    @pulumi.getter
    def regions(self) -> Optional[Sequence['outputs.ReplicationSetReplicationRegion']]:
        """
        The ReplicationSet configuration.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ReplicationSetTag']]:
        """
        The tags to apply to the replication set.
        """
        return pulumi.get(self, "tags")


class AwaitableGetReplicationSetResult(GetReplicationSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReplicationSetResult(
            arn=self.arn,
            deletion_protected=self.deletion_protected,
            regions=self.regions,
            tags=self.tags)


def get_replication_set(arn: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReplicationSetResult:
    """
    Resource type definition for AWS::SSMIncidents::ReplicationSet


    :param str arn: The ARN of the ReplicationSet.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ssmincidents:getReplicationSet', __args__, opts=opts, typ=GetReplicationSetResult).value

    return AwaitableGetReplicationSetResult(
        arn=pulumi.get(__ret__, 'arn'),
        deletion_protected=pulumi.get(__ret__, 'deletion_protected'),
        regions=pulumi.get(__ret__, 'regions'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_replication_set)
def get_replication_set_output(arn: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReplicationSetResult]:
    """
    Resource type definition for AWS::SSMIncidents::ReplicationSet


    :param str arn: The ARN of the ReplicationSet.
    """
    ...
