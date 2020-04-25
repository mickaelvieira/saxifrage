package lexer

import (
	"strings"
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
		uc := tc.input
		lc := strings.ToLower(uc)

		got := isKeyword(uc)
		assert.Equal(t, tc.want, got, "Test Case [Uppercase] %d %v", i, tc)

		got = isKeyword(lc)
		assert.Equal(t, tc.want, got, "Test Case [Lowercase] %d %v", i, tc)
	}
}

func TestGetKeyword(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"AddressFamily"},
		{"BatchMode"},
		{"BindAddress"},
		{"ChallengeResponseAuthentication"},
		{"CheckHostIP"},
		{"Cipher"},
		{"Ciphers"},
		{"ClearAllForwardings"},
		{"Compression"},
		{"CompressionLevel"},
		{"ConnectionAttempts"},
		{"ConnectTimeout"},
		{"ControlMaster"},
		{"ControlPath"},
		{"DynamicForward"},
		{"EscapeChar"},
		{"ExitOnForwardFailure"},
		{"ForwardAgent"},
		{"ForwardX11"},
		{"ForwardX11Trusted"},
		{"GatewayPorts"},
		{"GlobalKnownHostsFile"},
		{"GSSAPIAuthentication"},
		{"GSSAPIKeyExchange"},
		{"GSSAPIClientIdentity"},
		{"GSSAPIDelegateCredentials"},
		{"GSSAPIRenewalForcesRekey"},
		{"GSSAPITrustDns"},
		{"HashKnownHosts"},
		{"HostbasedAuthentication"},
		{"HostKeyAlgorithms"},
		{"HostKeyAlias"},
		{"HostName"},
		{"IdentitiesOnly"},
		{"IdentityFile"},
		{"KbdInteractiveAuthentication"},
		{"KbdInteractiveDevices"},
		{"LocalCommand"},
		{"LocalForward"},
		{"LogLevel"},
		{"MACs"},
		{"NoHostAuthenticationForLocalhost"},
		{"PasswordAuthentication"},
		{"PermitLocalCommand"},
		{"PreferredAuthentications"},
		{"Port"},
		{"Protocol"},
		{"ProxyCommand"},
		{"PubkeyAuthentication"},
		{"RekeyLimit"},
		{"RemoteForward"},
		{"RhostsRSAAuthentication"},
		{"RSAAuthentication"},
		{"SendEnv"},
		{"ServerAliveCountMax"},
		{"ServerAliveInterval"},
		{"SmartcardDevice"},
		{"StrictHostKeyChecking"},
		{"TCPKeepAlive"},
		{"Tunnel"},
		{"TunnelDevice"},
		{"User"},
		{"UsePrivilegedPort"},
		{"UserKnownHostsFile"},
		{"VerifyHostKeyDNS"},
		{"VisualHostKey"},
	}

	for i, tc := range cases {
		uc := tc.input
		lc := strings.ToLower(uc)

		got := getKeyword(tc.input)
		assert.Equal(t, lc, got.ID, "Test Case [Uppercase] %d %v", i, tc)
		assert.Equal(t, uc, got.Name, "Test Case [Uppercase] %d %v", i, tc)
	}
}
