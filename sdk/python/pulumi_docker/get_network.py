# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, driver=None, id=None, internal=None, ipam_configs=None, name=None, options=None, scope=None):
        if driver and not isinstance(driver, str):
            raise TypeError("Expected argument 'driver' to be a str")
        __self__.driver = driver
        """
        (Optional, string) The driver of the Docker network. 
        Possible values are `bridge`, `host`, `overlay`, `macvlan`.
        See [docker docs][networkdocs] for more details.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if internal and not isinstance(internal, bool):
            raise TypeError("Expected argument 'internal' to be a bool")
        __self__.internal = internal
        if ipam_configs and not isinstance(ipam_configs, list):
            raise TypeError("Expected argument 'ipam_configs' to be a list")
        __self__.ipam_configs = ipam_configs
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        __self__.options = options
        """
        (Optional, map) Only available with bridge networks. See
        [docker docs][bridgeoptionsdocs] for more details.
        * `internal` (Optional, bool) Boolean flag for whether the network is internal.
        * `ipam_config` (Optional, map) See IPAM below for details.
        * `scope` (Optional, string) Scope of the network. One of `swarm`, `global`, or `local`.
        """
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        __self__.scope = scope
class AwaitableGetNetworkResult(GetNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkResult(
            driver=self.driver,
            id=self.id,
            internal=self.internal,
            ipam_configs=self.ipam_configs,
            name=self.name,
            options=self.options,
            scope=self.scope)

def get_network(id=None,name=None,opts=None):
    """
    Finds a specific docker network and returns information about it.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/d/docker_network.html.markdown.


    :param str id: The id of the Docker network.
    :param str name: The name of the Docker network.
    """
    __args__ = dict()


    __args__['id'] = id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('docker:index/getNetwork:getNetwork', __args__, opts=opts).value

    return AwaitableGetNetworkResult(
        driver=__ret__.get('driver'),
        id=__ret__.get('id'),
        internal=__ret__.get('internal'),
        ipam_configs=__ret__.get('ipamConfigs'),
        name=__ret__.get('name'),
        options=__ret__.get('options'),
        scope=__ret__.get('scope'))
