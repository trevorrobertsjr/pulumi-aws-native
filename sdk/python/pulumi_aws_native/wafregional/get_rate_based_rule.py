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
    'GetRateBasedRuleResult',
    'AwaitableGetRateBasedRuleResult',
    'get_rate_based_rule',
    'get_rate_based_rule_output',
]

@pulumi.output_type
class GetRateBasedRuleResult:
    def __init__(__self__, id=None, match_predicates=None, rate_limit=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if match_predicates and not isinstance(match_predicates, list):
            raise TypeError("Expected argument 'match_predicates' to be a list")
        pulumi.set(__self__, "match_predicates", match_predicates)
        if rate_limit and not isinstance(rate_limit, int):
            raise TypeError("Expected argument 'rate_limit' to be a int")
        pulumi.set(__self__, "rate_limit", rate_limit)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="matchPredicates")
    def match_predicates(self) -> Optional[Sequence['outputs.RateBasedRulePredicate']]:
        return pulumi.get(self, "match_predicates")

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> Optional[int]:
        return pulumi.get(self, "rate_limit")


class AwaitableGetRateBasedRuleResult(GetRateBasedRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRateBasedRuleResult(
            id=self.id,
            match_predicates=self.match_predicates,
            rate_limit=self.rate_limit)


def get_rate_based_rule(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRateBasedRuleResult:
    """
    Resource Type definition for AWS::WAFRegional::RateBasedRule
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:wafregional:getRateBasedRule', __args__, opts=opts, typ=GetRateBasedRuleResult).value

    return AwaitableGetRateBasedRuleResult(
        id=pulumi.get(__ret__, 'id'),
        match_predicates=pulumi.get(__ret__, 'match_predicates'),
        rate_limit=pulumi.get(__ret__, 'rate_limit'))


@_utilities.lift_output_func(get_rate_based_rule)
def get_rate_based_rule_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRateBasedRuleResult]:
    """
    Resource Type definition for AWS::WAFRegional::RateBasedRule
    """
    ...
