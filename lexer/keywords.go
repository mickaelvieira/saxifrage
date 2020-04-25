package lexer

import (
	"strings"
)

// https://linux.die.net/man/5/ssh_config
// Keywords contains the list of ssh configuration mapping
var mapping = map[string]string{
	"AddressFamily":                    "any",
	"BatchMode":                        "no",
	"BindAddress":                      "",
	"ChallengeResponseAuthentication":  "yes",
	"CheckHostIP":                      "yes",
	"Cipher":                           "3des",
	"Ciphers":                          "aes128-ctr,aes192-ctr,aes256-ctr,aes128-cbc,3des-cbc",
	"ClearAllForwardings":              "no",
	"Compression":                      "no",
	"CompressionLevel":                 "",
	"ConnectionAttempts":               "",
	"ConnectTimeout":                   "0",
	"ControlMaster":                    "no",
	"ControlPath":                      "",
	"DynamicForward":                   "",
	"EscapeChar":                       "~",
	"ExitOnForwardFailure":             "",
	"ForwardAgent":                     "no",
	"ForwardX11":                       "no",
	"ForwardX11Trusted":                "",
	"GatewayPorts":                     "",
	"GlobalKnownHostsFile":             "",
	"GSSAPIAuthentication":             "no",
	"GSSAPIKeyExchange":                "",
	"GSSAPIClientIdentity":             "",
	"GSSAPIDelegateCredentials":        "no",
	"GSSAPIRenewalForcesRekey":         "",
	"GSSAPITrustDns":                   "",
	"HashKnownHosts":                   "",
	"HostbasedAuthentication":          "no",
	"HostKeyAlgorithms":                "",
	"HostKeyAlias":                     "",
	"HostName":                         "",
	"IdentitiesOnly":                   "",
	"IdentityFile":                     "~/.ssh/id_rsa",
	"KbdInteractiveAuthentication":     "",
	"KbdInteractiveDevices":            "",
	"LocalCommand":                     "",
	"LocalForward":                     "",
	"LogLevel":                         "",
	"MACs":                             "hmac-md5,hmac-sha1,umac-64@openssh.com",
	"NoHostAuthenticationForLocalhost": "",
	"PasswordAuthentication":           "yes",
	"PermitLocalCommand":               "no",
	"PreferredAuthentications":         "",
	"Port":                             "22",
	"Protocol":                         "",
	"ProxyCommand":                     "ssh -q -W %h:%p gateway.example.com",
	"PubkeyAuthentication":             "",
	"RekeyLimit":                       "1G 1h",
	"RemoteForward":                    "",
	"RhostsRSAAuthentication":          "",
	"RSAAuthentication":                "",
	"SendEnv":                          "",
	"ServerAliveCountMax":              "",
	"ServerAliveInterval":              "",
	"SmartcardDevice":                  "",
	"StrictHostKeyChecking":            "ask",
	"TCPKeepAlive":                     "",
	"Tunnel":                           "no",
	"TunnelDevice":                     "any:any",
	"UsePrivilegedPort":                "",
	"User":                             "",
	"UserKnownHostsFile":               "",
	"VerifyHostKeyDNS":                 "",
	"VisualHostKey":                    "no",
}

type keyword struct {
	ID      string
	Name    string
	Default string
}

var keywords = make(map[string]*keyword, len(mapping))

func init() {
	for k, v := range mapping {
		id := strings.ToLower(k)
		keywords[id] = &keyword{
			Name:    k,
			Default: v,
		}
	}
}

func isKeyword(i string) bool {
	k := strings.ToLower(i)
	_, ok := keywords[k]
	return ok
}

func getKeyword(i string) *keyword {
	k := strings.ToLower(i)
	kw, ok := keywords[k]
	if ok {
		return kw
	}
	return nil
}
