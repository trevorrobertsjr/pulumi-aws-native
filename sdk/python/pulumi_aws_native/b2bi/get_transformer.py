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
    'GetTransformerResult',
    'AwaitableGetTransformerResult',
    'get_transformer',
    'get_transformer_output',
]

@pulumi.output_type
class GetTransformerResult:
    def __init__(__self__, created_at=None, edi_type=None, file_format=None, mapping_template=None, modified_at=None, name=None, sample_document=None, status=None, tags=None, transformer_arn=None, transformer_id=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if edi_type and not isinstance(edi_type, dict):
            raise TypeError("Expected argument 'edi_type' to be a dict")
        pulumi.set(__self__, "edi_type", edi_type)
        if file_format and not isinstance(file_format, str):
            raise TypeError("Expected argument 'file_format' to be a str")
        pulumi.set(__self__, "file_format", file_format)
        if mapping_template and not isinstance(mapping_template, str):
            raise TypeError("Expected argument 'mapping_template' to be a str")
        pulumi.set(__self__, "mapping_template", mapping_template)
        if modified_at and not isinstance(modified_at, str):
            raise TypeError("Expected argument 'modified_at' to be a str")
        pulumi.set(__self__, "modified_at", modified_at)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sample_document and not isinstance(sample_document, str):
            raise TypeError("Expected argument 'sample_document' to be a str")
        pulumi.set(__self__, "sample_document", sample_document)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if transformer_arn and not isinstance(transformer_arn, str):
            raise TypeError("Expected argument 'transformer_arn' to be a str")
        pulumi.set(__self__, "transformer_arn", transformer_arn)
        if transformer_id and not isinstance(transformer_id, str):
            raise TypeError("Expected argument 'transformer_id' to be a str")
        pulumi.set(__self__, "transformer_id", transformer_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="ediType")
    def edi_type(self) -> Optional['outputs.TransformerEdiTypeProperties']:
        return pulumi.get(self, "edi_type")

    @property
    @pulumi.getter(name="fileFormat")
    def file_format(self) -> Optional['TransformerFileFormat']:
        return pulumi.get(self, "file_format")

    @property
    @pulumi.getter(name="mappingTemplate")
    def mapping_template(self) -> Optional[str]:
        return pulumi.get(self, "mapping_template")

    @property
    @pulumi.getter(name="modifiedAt")
    def modified_at(self) -> Optional[str]:
        return pulumi.get(self, "modified_at")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sampleDocument")
    def sample_document(self) -> Optional[str]:
        return pulumi.get(self, "sample_document")

    @property
    @pulumi.getter
    def status(self) -> Optional['TransformerStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TransformerTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transformerArn")
    def transformer_arn(self) -> Optional[str]:
        return pulumi.get(self, "transformer_arn")

    @property
    @pulumi.getter(name="transformerId")
    def transformer_id(self) -> Optional[str]:
        return pulumi.get(self, "transformer_id")


class AwaitableGetTransformerResult(GetTransformerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransformerResult(
            created_at=self.created_at,
            edi_type=self.edi_type,
            file_format=self.file_format,
            mapping_template=self.mapping_template,
            modified_at=self.modified_at,
            name=self.name,
            sample_document=self.sample_document,
            status=self.status,
            tags=self.tags,
            transformer_arn=self.transformer_arn,
            transformer_id=self.transformer_id)


def get_transformer(transformer_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransformerResult:
    """
    Definition of AWS::B2BI::Transformer Resource Type
    """
    __args__ = dict()
    __args__['transformerId'] = transformer_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:b2bi:getTransformer', __args__, opts=opts, typ=GetTransformerResult).value

    return AwaitableGetTransformerResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        edi_type=pulumi.get(__ret__, 'edi_type'),
        file_format=pulumi.get(__ret__, 'file_format'),
        mapping_template=pulumi.get(__ret__, 'mapping_template'),
        modified_at=pulumi.get(__ret__, 'modified_at'),
        name=pulumi.get(__ret__, 'name'),
        sample_document=pulumi.get(__ret__, 'sample_document'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        transformer_arn=pulumi.get(__ret__, 'transformer_arn'),
        transformer_id=pulumi.get(__ret__, 'transformer_id'))


@_utilities.lift_output_func(get_transformer)
def get_transformer_output(transformer_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransformerResult]:
    """
    Definition of AWS::B2BI::Transformer Resource Type
    """
    ...
