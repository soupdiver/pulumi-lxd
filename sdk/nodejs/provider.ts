// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The provider type for the lxd package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'lxd';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            inputs["LXDRemote"] = pulumi.output(args ? args.LXDRemote : undefined).apply(JSON.stringify);
            inputs["acceptRemoteCertificate"] = pulumi.output(args ? args.acceptRemoteCertificate : undefined).apply(JSON.stringify);
            inputs["configDir"] = args ? args.configDir : undefined;
            inputs["generateClientCertificates"] = pulumi.output(args ? args.generateClientCertificates : undefined).apply(JSON.stringify);
            inputs["refreshInterval"] = args ? args.refreshInterval : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    readonly LXDRemote?: pulumi.Input<pulumi.Input<inputs.ProviderLXDRemote>[]>;
    /**
     * Accept the server certificate
     */
    readonly acceptRemoteCertificate?: pulumi.Input<boolean>;
    /**
     * The directory to look for existing LXD configuration. default = $HOME/.config/lxc
     */
    readonly configDir?: pulumi.Input<string>;
    readonly generateClientCertificates?: pulumi.Input<boolean>;
    /**
     * How often to poll during state changes (default 10s)
     */
    readonly refreshInterval?: pulumi.Input<string>;
}
