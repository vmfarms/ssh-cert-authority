package main

import (
	"github.com/cloudtools/ssh-cert-authority/util"
	"golang.org/x/crypto/ssh"
	"testing"
)

// 768 bit key to try to keep the paste as small as possible
const signerPublicKeyString = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAYQDQ0aYhDW5wP8rVBIpXiCTiKw33/nI1mAti3EoX/zRYWtP2XMBePqnf2w2V+z5FEJm8hFSDpb8DtRBNonGEceuTn5CNjfNpDIJ7H3OOa8mP0CBiJY9A9Oopwn/hsRLy0vk= user-key"
const userPublicKeyString = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwE= signer-key"
const caPublicKeyString = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAYQDBdHEBPftP3CYZms38s0/1OgczR0UWSBmsbkoT1wL0PFG3OWkbiLE8zUue15EmajA/1VlbZPChbAF8Ub/9ZpdM/lzTv5v0cF8xSlw5oXOZEnIpSYD6M9OboTs3f5HyXfk= ca-key"
const caPrivateKeyString = `-----BEGIN RSA PRIVATE KEY-----
MIIBygIBAAJhAMF0cQE9+0/cJhmazfyzT/U6BzNHRRZIGaxuShPXAvQ8Ubc5aRuI
sTzNS57XkSZqMD/VWVtk8KFsAXxRv/1ml0z+XNO/m/RwXzFKXDmhc5kScilJgPoz
05uhOzd/kfJd+QIDAQABAmBQExmfcP9wO+jNWmV+/t3O3JkUYaC4K1ntJK2m7q27
WKheVfYqvnbWewedFQ9wvizH5USTOnt1TQXeMvIt3ysQCEfs36slJNhiYcndZtQg
CHebcmfaVwz6bV7+oZxxuoECMQDuWAZInNwgy2H+Vl3ElkqqPWG13RwZf0DI5BHQ
bmOMWPL1wCEJyCPs9sYh8QqQj8kCMQDPySM90tdhZUMZ0+7dWJmTEO6vxs+SdxBa
lQB/chG2QaFdSUQJO08IUwmfv3DJVLECMCS+DyHshIbNs6qYt9XRcWszETgPAQDx
PBR8DD78dX4yTCoUV0OBxgAGvt6GoSFN+QIwGZOVne+NCXUQfGZk+aQFS2ADMWnU
dR/oyG2c4RMmcPvFJBl3oXdGdCzce2hyNqYRAjEApZH6zZCu/mLz0+jiuxiSKR1i
USCW6VgmTJhf7XHWkp4GsykjNY/exQGdO+T3CPF+
-----END RSA PRIVATE KEY-----`
const userPrivateKeyString = `-----BEGIN RSA PRIVATE KEY-----
MIIBywIBAAJhANDRpiENbnA/ytUEileIJOIrDff+cjWYC2LcShf/NFha0/ZcwF4+
qd/bDZX7PkUQmbyEVIOlvwO1EE2icYRx65OfkI2N82kMgnsfc45ryY/QIGIlj0D0
6inCf+GxEvLS+QIDAQABAmEArs4x8g1aXCEq3LPWU3wm1CYSpX2dgfvr3DBo3jnH
SgeO1PfEGaD/d+PaNamC8TH43KjYnnkz9d2vZvHiuqrKlAz/DzxRb1SIQumece86
wbc3Z/kwkVSNThpnV/g6r8pZAjEA9Xyl6951daRsUPJGGjPTsQOiqyiJK88x/51V
OkT/qUoyqsSbsSIax7nr81vria2fAjEA2cL/hKqUGRfPtPiejpEgzZqGMJY4i5YH
9k5TeFgxFenr0Pc2Kl1UW9jwIAVZPwhnAjAXrVMPgeBQXXB5CjUKt+72BsS8v2cj
i5Nl9RXQTfFesaJbaCUgG4r7son4aeg42j8CMBPHu7AQUo2I9SwKHVTz59flPmUx
cAd15FlCOiDHWgYUjoAXxIrKmXwSU5WFBttL5wIxAMUNKEElbLlK6/u/LknYWxUI
wD1uIS8C/S226HDgbJI5fcI3sWlgr65NqMnW/6PnWA==
-----END RSA PRIVATE KEY-----`
const signerPrivateKeyString = `-----BEGIN RSA PRIVATE KEY-----
MIIBzAIBAAJhALpKGXmRS5MaqRJzwPa8/aSWzYagH94aWAhvQGcz/OwixvDlxaw+
i6udk3YjCfAyEIUtP/noLNP1nNV+fbqHGfzooymDvo1x8/kad3rp2xceGIxecF5L
2voo9B6rNUifAQIDAQABAmEAp89wO05jIdR2USTswldktQsTgR5lFpHsk0yEW3M9
dwms4/xXoN2Gu8VqvJS7sx+kr9QNr8N2tgnLz6UB1zU4Zusw3PVEb0qKTXAnYF9R
iMXRV43sCT3yUvnQAz6Nj2FxAjEA6MuOnyd9vNsnFoVbXq9NqGFqhh5koAxEKZ30
t3xxf2Ai+hUS+5hGINHfQUMVs9MHAjEAzNvQmCID63pEg34QAaa2Z7JTTCO6UCGi
UN88qyZJzF7hVz8HWzMfqOlOozyVQtO3AjBi1YlHqMyJUcHWneec23Bs/G7tYhn2
mT6XLKio/fxxx68R3cChcJTVekT+wCyGnCECMQCujByPg2wTl3oJD8BTp9iDQk32
8fotjHrgrVTj/xuiJrWZwPpjmou/QArgyx3icsECMQC7SWUafU8e1udo22uY8sVT
c2wCh1KaBOotbpXX0zISFvWsDtAAb8/o2A43eRlTA2k=
-----END RSA PRIVATE KEY-----`
const boringUserCertString = "ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAgD1NjpHjT+p/i1qgvl9D/kNlm/sbK30K3yQQJrKbI694AAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwEAAAAAAAAAAQAAAAEAAAAIYnZ6LXRlc3QAAAAWAAAABnVidW50dQAAAAhlYzItdXNlcgAAAAAAAAAA//////////8AAAAAAAAAggAAABVwZXJtaXQtWDExLWZvcndhcmRpbmcAAAAAAAAAF3Blcm1pdC1hZ2VudC1mb3J3YXJkaW5nAAAAAAAAABZwZXJtaXQtcG9ydC1mb3J3YXJkaW5nAAAAAAAAAApwZXJtaXQtcHR5AAAAAAAAAA5wZXJtaXQtdXNlci1yYwAAAAAAAAAAAAAAdwAAAAdzc2gtcnNhAAAAAwEAAQAAAGEAukoZeZFLkxqpEnPA9rz9pJbNhqAf3hpYCG9AZzP87CLG8OXFrD6Lq52TdiMJ8DIQhS0/+egs0/Wc1X59uocZ/OijKYO+jXHz+Rp3eunbFx4YjF5wXkva+ij0Hqs1SJ8BAAAAbwAAAAdzc2gtcnNhAAAAYEUiGq8/zcN4UglbHsbU4PwHoAfzQQDkXk41oxYya5Ig858Y4EDf+BzhOOLNoAoTawKyWkTCKRayXpNdRfSFjiraO3a57XRGEJWF4ybgP1leYQ7QlERmmViAHMSxrNpL3Q==  bvanzant@bvz-air.local"
const signedByInvalidUserCertString = "ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAg9KP0im2yxWQuEQkSGKK9Ym4tb17ZyWs9xSwN0lsU3o4AAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwEAAAAAAAAAAQAAAAEAAAAIYnZ6LXRlc3QAAAAWAAAABnVidW50dQAAAAhlYzItdXNlcgAAAAAAAAAA//////////8AAAAAAAAAggAAABVwZXJtaXQtWDExLWZvcndhcmRpbmcAAAAAAAAAF3Blcm1pdC1hZ2VudC1mb3J3YXJkaW5nAAAAAAAAABZwZXJtaXQtcG9ydC1mb3J3YXJkaW5nAAAAAAAAAApwZXJtaXQtcHR5AAAAAAAAAA5wZXJtaXQtdXNlci1yYwAAAAAAAAAAAAAAdwAAAAdzc2gtcnNhAAAAAwEAAQAAAGEAwXRxAT37T9wmGZrN/LNP9ToHM0dFFkgZrG5KE9cC9DxRtzlpG4ixPM1LnteRJmowP9VZW2TwoWwBfFG//WaXTP5c07+b9HBfMUpcOaFzmRJyKUmA+jPTm6E7N3+R8l35AAAAbwAAAAdzc2gtcnNhAAAAYGuA1rEVeDacEoKSFHlxpD85O7ueBWHm8zRdmqc/txt88FnFKlb/cBFkL//8P5mB5WbjWGSnOXokUqj9Xf5gaIb1jfWzWql3HoYomlkyVGcI+g+lMq1w6/rI7lTcpEAjaw==  bvanzant@bvz-air.local"
const foreverUserCertString = "ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAgOQ95vOpIs7uJPiJ1VikwGbmf7LZWFCkD4SYzP9XEvrkAAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwEAAAAAAAAAAAAAAAEAAAANdGVzdGZhcmZ1dHVyZQAAAAoAAAAGdWJ1bnR1AAAAAAAAAAD//////////wAAAAAAAACCAAAAFXBlcm1pdC1YMTEtZm9yd2FyZGluZwAAAAAAAAAXcGVybWl0LWFnZW50LWZvcndhcmRpbmcAAAAAAAAAFnBlcm1pdC1wb3J0LWZvcndhcmRpbmcAAAAAAAAACnBlcm1pdC1wdHkAAAAAAAAADnBlcm1pdC11c2VyLXJjAAAAAAAAAAAAAAB3AAAAB3NzaC1yc2EAAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwEAAABvAAAAB3NzaC1yc2EAAABgrLrgF1y6t1s0APpKi2JcBDI9OhtsIa7SLaTvRrhAmXc+0OQPlgtmqp581Mwnpv+cuJpoPurO3cQpfzT+XDnRZdoEOBz9wTzkbb6tqQ/O+Ltf5ss5cyKd0GHgGCzUVYB7 user-key"
const twentyTwentyFiveCertString = "ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAgFJu+SQd43RC/DZjYXtsiSAZslugnFX0LQV44bQL5XlgAAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwEAAAAAAAAAAAAAAAEAAAANdGVzdGZhcmZ1dHVyZQAAAAoAAAAGdWJ1bnR1AAAAAFYW52AAAAAAaOLqwgAAAAAAAACCAAAAFXBlcm1pdC1YMTEtZm9yd2FyZGluZwAAAAAAAAAXcGVybWl0LWFnZW50LWZvcndhcmRpbmcAAAAAAAAAFnBlcm1pdC1wb3J0LWZvcndhcmRpbmcAAAAAAAAACnBlcm1pdC1wdHkAAAAAAAAADnBlcm1pdC11c2VyLXJjAAAAAAAAAAAAAAB3AAAAB3NzaC1yc2EAAAADAQABAAAAYQC6Shl5kUuTGqkSc8D2vP2kls2GoB/eGlgIb0BnM/zsIsbw5cWsPournZN2IwnwMhCFLT/56CzT9ZzVfn26hxn86KMpg76NcfP5Gnd66dsXHhiMXnBeS9r6KPQeqzVInwEAAABvAAAAB3NzaC1yc2EAAABgHPStnlKu6jr8bI0LG3LwwWczyLU56rMozAngOo6ovCAMqXtWZ5GvfUZf+Ok+B+6usJrMZQqrE+OhdC+GVWqzee9nH9Wy1n8ACbGVWJtN63iJxqrMqOnLVvSrd7s7tPqv user-key"

func SetupSignerdConfig(numSignersRequired, maxCertLifetime int) map[string]ssh_ca_util.SignerdConfig {
	config := make(map[string]ssh_ca_util.SignerdConfig)
	config["testing"] = ssh_ca_util.SignerdConfig{
		SigningKeyFingerprint: "4c:c6:1e:31:ed:7b:7c:33:ff:7d:51:9e:59:da:68:f5",
		AuthorizedUsers: map[string]string{
			"f6:e3:42:5e:72:85:ce:26:e8:45:1f:79:2d:dc:0d:52": "test-user",
		},
		AuthorizedSigners: map[string]string{
			"23:10:8d:d0:54:90:d5:d1:2e:4d:05:fe:4b:54:29:e4": "test-signer",
		},
		NumberSignersRequired: numSignersRequired,
		MaxCertLifetime:       maxCertLifetime,
	}
	return config
}

func TestRejectRequest(t *testing.T) {
	allConfig := SetupSignerdConfig(1, 0)
	environment := "testing"
	envConfig := allConfig[environment]
	requestHandler := makeCertRequestHandler(allConfig)

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(boringUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert := pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err != nil {
		t.Fatalf("Should have succeeded. Failed with: %v", err)
	}

	err = requestHandler.rejectRequest("DEADBEEFDEADBEEF", "23:10:8d:d0:54:90:d5:d1:2e:4d:05:fe:4b:54:29:e4", envConfig)
	if err != nil {
		t.Fatalf("Should have succeeded. Failed with: %v", err)
	}
	err = requestHandler.addConfirmation("DEADBEEFDEADBEEF", "23:10:8d:d0:54:90:d5:d1:2e:4d:05:fe:4b:54:29:e4", envConfig, false)
	if err == nil {
		t.Fatalf("Sign after reject should fail.")
	}
}

func TestRejectRequestAfterSigning(t *testing.T) {
	allConfig := SetupSignerdConfig(2, 0)
	environment := "testing"
	envConfig := allConfig[environment]
	requestHandler := makeCertRequestHandler(allConfig)

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(boringUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert := pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err != nil {
		t.Fatalf("Should have succeeded. Failed with: %v", err)
	}

	err = requestHandler.addConfirmation("DEADBEEFDEADBEEF", "23:10:8d:d0:54:90:d5:d1:2e:4d:05:fe:4b:54:29:e4", envConfig, false)
	if err != nil {
		t.Fatalf("Sign should have worked. It failed: %v", err)
	}

	err = requestHandler.rejectRequest("DEADBEEFDEADBEEF", "23:10:8d:d0:54:90:d5:d1:2e:4d:05:fe:4b:54:29:e4", envConfig)
	if err != nil {
		t.Fatalf("Should have succeeded. Failed with: %v", err)
	}

	err = requestHandler.addConfirmation("DEADBEEFDEADBEEF", "23:10:8d:d0:54:90:d5:d1:2e:4d:05:fe:4b:54:29:e4", envConfig, false)
	if err == nil {
		t.Fatalf("Sign after reject should fail.")
	}
}

func TestSaveForeverCertDisallowed(t *testing.T) {
	allConfig := SetupSignerdConfig(1, 14400)
	environment := "testing"
	envConfig := allConfig[environment]
	requestHandler := makeCertRequestHandler(allConfig)

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(foreverUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert := pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed because cert never expires.")
	}

	pubKey, _, _, _, err = ssh.ParseAuthorizedKey([]byte(twentyTwentyFiveCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert = pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed because cert expires in 2025.")
	}
}

func TestSaveForeverCertAllowed(t *testing.T) {
	allConfig := SetupSignerdConfig(1, 0)
	environment := "testing"
	envConfig := allConfig[environment]
	requestHandler := makeCertRequestHandler(allConfig)

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(foreverUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert := pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err != nil {
		t.Fatalf("Should have worked, failed with: %v", err)
	}

	pubKey, _, _, _, err = ssh.ParseAuthorizedKey([]byte(twentyTwentyFiveCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert = pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF2", 1, cert)
	if err != nil {
		t.Fatalf("Should have worked, failed with: %v", err)
	}
}

func TestSaveRequestValidCert(t *testing.T) {
	allConfig := SetupSignerdConfig(1, 0)
	environment := "testing"
	envConfig := allConfig[environment]
	requestHandler := makeCertRequestHandler(allConfig)

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(boringUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert := pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err != nil {
		t.Fatalf("Should have succeeded. Failed with: %v", err)
	}

	err = requestHandler.saveSigningRequest(envConfig, environment, "", "DEADBEEFDEAD1111", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed, reason was missing.")
	}
	err = requestHandler.saveSigningRequest(envConfig, "", "reason: testing", "DEADBEEFDEAD2222", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed, environment was missing.")
	}
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "EEEFDF", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed with invalid request id (too short).")
	}

	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "IAM_A_DUPLICATE_ID", 1, cert)
	if err != nil {
		t.Fatalf("Should have succeeded. Failed with: %v", err)
	}
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "IAM_A_DUPLICATE_ID", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed with duplicate error")
	}
}

func TestSaveRequestInvalidCert(t *testing.T) {
	allConfig := SetupSignerdConfig(1, 0)
	environment := "testing"
	envConfig := allConfig[environment]
	requestHandler := makeCertRequestHandler(allConfig)

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(signedByInvalidUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	cert := pubKey.(*ssh.Certificate)
	err = requestHandler.saveSigningRequest(envConfig, environment, "reason: testing", "DEADBEEFDEADBEEF", 1, cert)
	if err == nil {
		t.Fatalf("Should have failed with fingerprint not in list error.")
	}
}

func getTwoBoringCerts(t *testing.T) (*ssh.Certificate, *ssh.Certificate) {
	pubKeyOne, _, _, _, err := ssh.ParseAuthorizedKey([]byte(boringUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	boringCertOne := pubKeyOne.(*ssh.Certificate)

	pubKeyTwo, _, _, _, err := ssh.ParseAuthorizedKey([]byte(boringUserCertString))
	if err != nil {
		t.Fatalf("Parsing canned cert failed: %v", err)
	}
	boringCertTwo := pubKeyTwo.(*ssh.Certificate)

	return boringCertOne, boringCertTwo
}

func TestCompareCertsEasy(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	if !compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should have compared equal.")
	}
}

func TestCompareCertsExtendedTime(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.ValidBefore += 1024
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsSerialMismatch(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.Serial = 0
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsTypeMismatch(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.CertType = ssh.HostCert
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsKeyId(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.KeyId = "I'm a liar!"
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsAddInSomeRoot(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.ValidPrincipals = append(boringCertTwo.ValidPrincipals, "root")
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsReplaceWithRoot(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.ValidPrincipals[0] = "root"
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsCriticalOptions(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.CriticalOptions["foobar"] = "we don't use these"
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsExtensions(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.CriticalOptions["environment@cloudtools.github.io"] = "prod"
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsNonce(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertOne.Nonce = []byte("this is nonce one")
	boringCertTwo.Nonce = []byte("this is nonce two")
	if !compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should have been equal, we don't compare Nonce.")
	}
}

func TestCompareCertsReserved(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.Reserved = []byte("some random garbage")
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}

func TestCompareCertsPublicKey(t *testing.T) {
	boringCertOne, boringCertTwo := getTwoBoringCerts(t)
	boringCertTwo.Key, _, _, _, _ = ssh.ParseAuthorizedKey([]byte(signerPublicKeyString))
	if compareCerts(boringCertOne, boringCertTwo) {
		t.Fatalf("Certs should not have been equal.")
	}
}
