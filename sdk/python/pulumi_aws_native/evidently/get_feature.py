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
    'GetFeatureResult',
    'AwaitableGetFeatureResult',
    'get_feature',
    'get_feature_output',
]

@pulumi.output_type
class GetFeatureResult:
    def __init__(__self__, arn=None, default_variation=None, description=None, entity_overrides=None, evaluation_strategy=None, tags=None, variations=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if default_variation and not isinstance(default_variation, str):
            raise TypeError("Expected argument 'default_variation' to be a str")
        pulumi.set(__self__, "default_variation", default_variation)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if entity_overrides and not isinstance(entity_overrides, list):
            raise TypeError("Expected argument 'entity_overrides' to be a list")
        pulumi.set(__self__, "entity_overrides", entity_overrides)
        if evaluation_strategy and not isinstance(evaluation_strategy, str):
            raise TypeError("Expected argument 'evaluation_strategy' to be a str")
        pulumi.set(__self__, "evaluation_strategy", evaluation_strategy)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if variations and not isinstance(variations, list):
            raise TypeError("Expected argument 'variations' to be a list")
        pulumi.set(__self__, "variations", variations)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="defaultVariation")
    def default_variation(self) -> Optional[str]:
        return pulumi.get(self, "default_variation")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entityOverrides")
    def entity_overrides(self) -> Optional[Sequence['outputs.FeatureEntityOverride']]:
        return pulumi.get(self, "entity_overrides")

    @property
    @pulumi.getter(name="evaluationStrategy")
    def evaluation_strategy(self) -> Optional['FeatureEvaluationStrategy']:
        return pulumi.get(self, "evaluation_strategy")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.FeatureTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def variations(self) -> Optional[Sequence['outputs.FeatureVariationObject']]:
        return pulumi.get(self, "variations")


class AwaitableGetFeatureResult(GetFeatureResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeatureResult(
            arn=self.arn,
            default_variation=self.default_variation,
            description=self.description,
            entity_overrides=self.entity_overrides,
            evaluation_strategy=self.evaluation_strategy,
            tags=self.tags,
            variations=self.variations)


def get_feature(arn: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeatureResult:
    """
    Resource Type definition for AWS::Evidently::Feature.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:evidently:getFeature', __args__, opts=opts, typ=GetFeatureResult).value

    return AwaitableGetFeatureResult(
        arn=pulumi.get(__ret__, 'arn'),
        default_variation=pulumi.get(__ret__, 'default_variation'),
        description=pulumi.get(__ret__, 'description'),
        entity_overrides=pulumi.get(__ret__, 'entity_overrides'),
        evaluation_strategy=pulumi.get(__ret__, 'evaluation_strategy'),
        tags=pulumi.get(__ret__, 'tags'),
        variations=pulumi.get(__ret__, 'variations'))


@_utilities.lift_output_func(get_feature)
def get_feature_output(arn: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeatureResult]:
    """
    Resource Type definition for AWS::Evidently::Feature.
    """
    ...
