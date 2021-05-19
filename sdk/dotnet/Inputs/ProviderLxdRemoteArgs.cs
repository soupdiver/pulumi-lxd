// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Lxd.Inputs
{

    public sealed class ProviderLxdRemoteArgs : Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("default")]
        public Input<bool>? Default { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        public ProviderLxdRemoteArgs()
        {
        }
    }
}