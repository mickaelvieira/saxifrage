package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsKeyword(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"foo", false},
		{"Host", false},
		{"Match", false},
		{"AddressFamily", true},
		{"BatchMode", true},
		{"BindAddress", true},
		{"ChallengeResponseAuthentication", true},
		{"CheckHostIP", true},
		{"Cipher", true},
		{"Ciphers", true},
		{"ClearAllForwardings", true},
		{"Compression", true},
		{"CompressionLevel", true},
		{"ConnectionAttempts", true},
		{"ConnectTimeout", true},
		{"ControlMaster", true},
		{"ControlPath", true},
		{"DynamicForward", true},
		{"EscapeChar", true},
		{"ExitOnForwardFailure", true},
		{"ForwardAgent", true},
		{"ForwardX11", true},
		{"ForwardX11Trusted", true},
		{"GatewayPorts", true},
		{"GlobalKnownHostsFile", true},
		{"GSSAPIAuthentication", true},
		{"GSSAPIKeyExchange", true},
		{"GSSAPIClientIdentity", true},
		{"GSSAPIDelegateCredentials", true},
		{"GSSAPIRenewalForcesRekey", true},
		{"GSSAPITrustDns", true},
		{"HashKnownHosts", true},
		{"HostbasedAuthentication", true},
		{"HostKeyAlgorithms", true},
		{"HostKeyAlias", true},
		{"HostName", true},
		{"IdentitiesOnly", true},
		{"IdentityFile", true},
		{"KbdInteractiveAuthentication", true},
		{"KbdInteractiveDevices", true},
		{"LocalCommand", true},
		{"LocalForward", true},
		{"LogLevel", true},
		{"MACs", true},
		{"NoHostAuthenticationForLocalhost", true},
		{"PasswordAuthentication", true},
		{"PermitLocalCommand", true},
		{"PreferredAuthentications", true},
		{"Port", true},
		{"Protocol", true},
		{"ProxyCommand", true},
		{"PubkeyAuthentication", true},
		{"RekeyLimit", true},
		{"RemoteForward", true},
		{"RhostsRSAAuthentication", true},
		{"RSAAuthentication", true},
		{"SendEnv", true},
		{"ServerAliveCountMax", true},
		{"ServerAliveInterval", true},
		{"SmartcardDevice", true},
		{"StrictHostKeyChecking", true},
		{"TCPKeepAlive", true},
		{"Tunnel", true},
		{"TunnelDevice", true},
		{"User", true},
		{"UsePrivilegedPort", true},
		{"UserKnownHostsFile", true},
		{"VerifyHostKeyDNS", true},
		{"VisualHostKey", true},
	}

	for i, tc := range cases {
		got := isKeyword(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}
