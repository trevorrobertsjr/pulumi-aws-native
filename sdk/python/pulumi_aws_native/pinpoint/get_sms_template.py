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
    'GetSmsTemplateResult',
    'AwaitableGetSmsTemplateResult',
    'get_sms_template',
    'get_sms_template_output',
]

@pulumi.output_type
class GetSmsTemplateResult:
    def __init__(__self__, arn=None, body=None, default_substitutions=None, id=None, tags=None, template_description=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if default_substitutions and not isinstance(default_substitutions, str):
            raise TypeError("Expected argument 'default_substitutions' to be a str")
        pulumi.set(__self__, "default_substitutions", default_substitutions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if template_description and not isinstance(template_description, str):
            raise TypeError("Expected argument 'template_description' to be a str")
        pulumi.set(__self__, "template_description", template_description)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def body(self) -> Optional[str]:
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="defaultSubstitutions")
    def default_substitutions(self) -> Optional[str]:
        return pulumi.get(self, "default_substitutions")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateDescription")
    def template_description(self) -> Optional[str]:
        return pulumi.get(self, "template_description")


class AwaitableGetSmsTemplateResult(GetSmsTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSmsTemplateResult(
            arn=self.arn,
            body=self.body,
            default_substitutions=self.default_substitutions,
            id=self.id,
            tags=self.tags,
            template_description=self.template_description)


def get_sms_template(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSmsTemplateResult:
    """
    Resource Type definition for AWS::Pinpoint::SmsTemplate
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:pinpoint:getSmsTemplate', __args__, opts=opts, typ=GetSmsTemplateResult).value

    return AwaitableGetSmsTemplateResult(
        arn=pulumi.get(__ret__, 'arn'),
        body=pulumi.get(__ret__, 'body'),
        default_substitutions=pulumi.get(__ret__, 'default_substitutions'),
        id=pulumi.get(__ret__, 'id'),
        tags=pulumi.get(__ret__, 'tags'),
        template_description=pulumi.get(__ret__, 'template_description'))


@_utilities.lift_output_func(get_sms_template)
def get_sms_template_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSmsTemplateResult]:
    """
    Resource Type definition for AWS::Pinpoint::SmsTemplate
    """
    ...
