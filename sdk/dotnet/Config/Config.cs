// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Lxd
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("lxd");
        /// <summary>
        /// Accept the server certificate
        /// </summary>
        public static bool? AcceptRemoteCertificate { get; set; } = __config.GetBoolean("acceptRemoteCertificate");

        public static string? Address { get; set; } = __config.Get("address");

        /// <summary>
        /// The directory to look for existing LXD configuration. default = $HOME/.config/lxc
        /// </summary>
        public static string? ConfigDir { get; set; } = __config.Get("configDir");

        public static bool? GenerateClientCertificates { get; set; } = __config.GetBoolean("generateClientCertificates");

        public static ImmutableArray<Pulumi.Lxd.Config.Types.LxdRemotes> LxdRemotes { get; set; } = __config.GetObject<ImmutableArray<Pulumi.Lxd.Config.Types.LxdRemotes>>("lxdRemotes");

        public static string? Port { get; set; } = __config.Get("port");

        /// <summary>
        /// How often to poll during state changes (default 10s)
        /// </summary>
        public static string? RefreshInterval { get; set; } = __config.Get("refreshInterval");

        public static string? Remote { get; set; } = __config.Get("remote");

        public static string? RemotePassword { get; set; } = __config.Get("remotePassword");

        public static string? Scheme { get; set; } = __config.Get("scheme");

        public static class Types
        {

             public class LxdRemotes
             {
                public string? Address { get; set; } = null!;
                public bool? Default { get; set; }
                public string Name { get; set; }
                public string? Password { get; set; } = null!;
                public string? Port { get; set; } = null!;
                public string? Scheme { get; set; } = null!;
            }
        }
    }
}
