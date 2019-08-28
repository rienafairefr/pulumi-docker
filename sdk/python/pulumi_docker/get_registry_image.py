# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetRegistryImageResult:
    """
    A collection of values returned by getRegistryImage.
    """
    def __init__(__self__, name=None, sha256_digest=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if sha256_digest and not isinstance(sha256_digest, str):
            raise TypeError("Expected argument 'sha256_digest' to be a str")
        __self__.sha256_digest = sha256_digest
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetRegistryImageResult(GetRegistryImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryImageResult(
            name=self.name,
            sha256_digest=self.sha256_digest,
            id=self.id)

def get_registry_image(name=None,opts=None):
    """
    Reads the image metadata from a Docker Registry. Used in conjunction with the
    [docker\_image](https://www.terraform.io/docs/providers/docker/r/image.html) resource to keep an image up
    to date on the latest available version of the tag.
    
    :param str name: The name of the Docker image, including any tags. e.g. `alpine:latest`

    > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/d/registry_image.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('docker:index/getRegistryImage:getRegistryImage', __args__, opts=opts).value

    return AwaitableGetRegistryImageResult(
        name=__ret__.get('name'),
        sha256_digest=__ret__.get('sha256Digest'),
        id=__ret__.get('id'))
