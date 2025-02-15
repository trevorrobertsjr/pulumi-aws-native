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
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    def __init__(__self__, arn=None, kms_key_id=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        The KMS key for the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DatabaseTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            arn=self.arn,
            kms_key_id=self.kms_key_id,
            tags=self.tags)


def get_database(database_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    The AWS::Timestream::Database resource creates a Timestream database.


    :param str database_name: The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:timestream:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        arn=pulumi.get(__ret__, 'arn'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_database)
def get_database_output(database_name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    The AWS::Timestream::Database resource creates a Timestream database.


    :param str database_name: The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.
    """
    ...
