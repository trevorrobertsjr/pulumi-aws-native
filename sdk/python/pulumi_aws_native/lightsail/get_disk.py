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
from ._enums import *

__all__ = [
    'GetDiskResult',
    'AwaitableGetDiskResult',
    'get_disk',
    'get_disk_output',
]

@pulumi.output_type
class GetDiskResult:
    def __init__(__self__, add_ons=None, attached_to=None, attachment_state=None, disk_arn=None, iops=None, is_attached=None, location=None, path=None, resource_type=None, state=None, support_code=None, tags=None):
        if add_ons and not isinstance(add_ons, list):
            raise TypeError("Expected argument 'add_ons' to be a list")
        pulumi.set(__self__, "add_ons", add_ons)
        if attached_to and not isinstance(attached_to, str):
            raise TypeError("Expected argument 'attached_to' to be a str")
        pulumi.set(__self__, "attached_to", attached_to)
        if attachment_state and not isinstance(attachment_state, str):
            raise TypeError("Expected argument 'attachment_state' to be a str")
        pulumi.set(__self__, "attachment_state", attachment_state)
        if disk_arn and not isinstance(disk_arn, str):
            raise TypeError("Expected argument 'disk_arn' to be a str")
        pulumi.set(__self__, "disk_arn", disk_arn)
        if iops and not isinstance(iops, int):
            raise TypeError("Expected argument 'iops' to be a int")
        pulumi.set(__self__, "iops", iops)
        if is_attached and not isinstance(is_attached, bool):
            raise TypeError("Expected argument 'is_attached' to be a bool")
        pulumi.set(__self__, "is_attached", is_attached)
        if location and not isinstance(location, dict):
            raise TypeError("Expected argument 'location' to be a dict")
        pulumi.set(__self__, "location", location)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if support_code and not isinstance(support_code, str):
            raise TypeError("Expected argument 'support_code' to be a str")
        pulumi.set(__self__, "support_code", support_code)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="addOns")
    def add_ons(self) -> Optional[Sequence['outputs.DiskAddOn']]:
        """
        An array of objects representing the add-ons to enable for the new instance.
        """
        return pulumi.get(self, "add_ons")

    @property
    @pulumi.getter(name="attachedTo")
    def attached_to(self) -> Optional[str]:
        """
        Name of the attached Lightsail Instance
        """
        return pulumi.get(self, "attached_to")

    @property
    @pulumi.getter(name="attachmentState")
    def attachment_state(self) -> Optional[str]:
        """
        Attachment State of the Lightsail disk
        """
        return pulumi.get(self, "attachment_state")

    @property
    @pulumi.getter(name="diskArn")
    def disk_arn(self) -> Optional[str]:
        return pulumi.get(self, "disk_arn")

    @property
    @pulumi.getter
    def iops(self) -> Optional[int]:
        """
        Iops of the Lightsail disk
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="isAttached")
    def is_attached(self) -> Optional[bool]:
        """
        Check is Disk is attached state
        """
        return pulumi.get(self, "is_attached")

    @property
    @pulumi.getter
    def location(self) -> Optional['outputs.DiskLocation']:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        Path of the  attached Disk
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        Resource type of Lightsail instance.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        State of the Lightsail disk
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="supportCode")
    def support_code(self) -> Optional[str]:
        """
        Support code to help identify any issues
        """
        return pulumi.get(self, "support_code")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DiskTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDiskResult(GetDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiskResult(
            add_ons=self.add_ons,
            attached_to=self.attached_to,
            attachment_state=self.attachment_state,
            disk_arn=self.disk_arn,
            iops=self.iops,
            is_attached=self.is_attached,
            location=self.location,
            path=self.path,
            resource_type=self.resource_type,
            state=self.state,
            support_code=self.support_code,
            tags=self.tags)


def get_disk(disk_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiskResult:
    """
    Resource Type definition for AWS::Lightsail::Disk


    :param str disk_name: The names to use for your new Lightsail disk.
    """
    __args__ = dict()
    __args__['diskName'] = disk_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:lightsail:getDisk', __args__, opts=opts, typ=GetDiskResult).value

    return AwaitableGetDiskResult(
        add_ons=pulumi.get(__ret__, 'add_ons'),
        attached_to=pulumi.get(__ret__, 'attached_to'),
        attachment_state=pulumi.get(__ret__, 'attachment_state'),
        disk_arn=pulumi.get(__ret__, 'disk_arn'),
        iops=pulumi.get(__ret__, 'iops'),
        is_attached=pulumi.get(__ret__, 'is_attached'),
        location=pulumi.get(__ret__, 'location'),
        path=pulumi.get(__ret__, 'path'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        state=pulumi.get(__ret__, 'state'),
        support_code=pulumi.get(__ret__, 'support_code'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_disk)
def get_disk_output(disk_name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDiskResult]:
    """
    Resource Type definition for AWS::Lightsail::Disk


    :param str disk_name: The names to use for your new Lightsail disk.
    """
    ...
