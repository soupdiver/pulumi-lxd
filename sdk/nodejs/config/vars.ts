// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

let __config = new pulumi.Config("lxd");

export let LXDRemote: outputs.config.LXDRemote[] | undefined = __config.getObject<outputs.config.LXDRemote[]>("LXDRemote");
/**
 * Accept the server certificate
 */
export let acceptRemoteCertificate: boolean | undefined = __config.getObject<boolean>("acceptRemoteCertificate");
export let address: string | undefined = __config.get("address");
/**
 * The directory to look for existing LXD configuration. default = $HOME/.config/lxc
 */
export let configDir: string | undefined = __config.get("configDir");
export let generateClientCertificates: boolean | undefined = __config.getObject<boolean>("generateClientCertificates");
export let port: string | undefined = __config.get("port");
/**
 * How often to poll during state changes (default 10s)
 */
export let refreshInterval: string | undefined = __config.get("refreshInterval");
export let remote: string | undefined = __config.get("remote");
export let remotePassword: string | undefined = __config.get("remotePassword");
export let scheme: string | undefined = __config.get("scheme");
