package util

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

// ── helpers ─────────────────────────────────────────────────────────────────

func mustGetOutbound(t *testing.T, uri string) map[string]interface{} {
	t.Helper()
	out, _, err := GetOutbound(uri, 0)
	if err != nil {
		t.Fatalf("GetOutbound(%q) error: %v", uri, err)
	}
	if out == nil {
		t.Fatalf("GetOutbound(%q) returned nil outbound", uri)
	}
	return *out
}

func str(m map[string]interface{}, key string) string {
	v, _ := m[key].(string)
	return v
}

func mapVal(m map[string]interface{}, key string) map[string]interface{} {
	v, _ := m[key].(map[string]interface{})
	return v
}

// ── GetOutbound: unsupported scheme ─────────────────────────────────────────

func TestGetOutbound_UnsupportedScheme(t *testing.T) {
	_, _, err := GetOutbound("ftp://example.com", 0)
	if err == nil {
		t.Error("expected error for unsupported scheme, got nil")
	}
}

// ── vless ────────────────────────────────────────────────────────────────────

func TestGetOutbound_Vless_Basic(t *testing.T) {
	uri := "vless://some-uuid-1234@example.com:443?security=tls&sni=example.com&type=ws&path=/ws#MyTag"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "vless" {
		t.Errorf("type: got %q, want %q", str(out, "type"), "vless")
	}
	if str(out, "tag") != "MyTag" {
		t.Errorf("tag: got %q, want %q", str(out, "tag"), "MyTag")
	}
	if str(out, "server") != "example.com" {
		t.Errorf("server: got %q, want %q", str(out, "server"), "example.com")
	}
	if out["server_port"] != 443 {
		t.Errorf("server_port: got %v, want 443", out["server_port"])
	}
}

func TestGetOutbound_Vless_WithIndex(t *testing.T) {
	uri := "vless://uuid@host.com:8080?security=tls&type=tcp#Tag"
	_, tag, err := GetOutbound(uri, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(tag, "3.") {
		t.Errorf("expected tag to start with '3.', got %q", tag)
	}
}

// ── trojan ───────────────────────────────────────────────────────────────────

func TestGetOutbound_Trojan_Basic(t *testing.T) {
	uri := "trojan://mypassword@trojan.example.com:443?security=tls&sni=trojan.example.com#TrojanTag"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "trojan" {
		t.Errorf("type: got %q, want %q", str(out, "type"), "trojan")
	}
	if str(out, "password") != "mypassword" {
		t.Errorf("password: got %q, want %q", str(out, "password"), "mypassword")
	}
	if str(out, "server") != "trojan.example.com" {
		t.Errorf("server: got %q, want %q", str(out, "server"), "trojan.example.com")
	}
}

// ── hysteria2 ────────────────────────────────────────────────────────────────

func TestGetOutbound_Hysteria2_Basic(t *testing.T) {
	uri := "hy2://secret@hy2.example.com:443?sni=hy2.example.com#HY2Tag"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "hysteria2" {
		t.Errorf("type: got %q, want %q", str(out, "type"), "hysteria2")
	}
	if str(out, "password") != "secret" {
		t.Errorf("password: got %q, want %q", str(out, "password"), "secret")
	}
}

func TestGetOutbound_Hysteria2_WithObfs(t *testing.T) {
	uri := "hysteria2://secret@hy2.example.com:443?obfs=salamander&obfs-password=obfspass#HY2"
	out := mustGetOutbound(t, uri)
	obfs := mapVal(out, "obfs")
	if obfs == nil {
		t.Fatal("expected obfs map in hy2 result")
	}
	if obfs["type"] != "salamander" {
		t.Errorf("obfs.type: got %v, want salamander", obfs["type"])
	}
	if obfs["password"] != "obfspass" {
		t.Errorf("obfs.password: got %v, want obfspass", obfs["password"])
	}
}

// ── hysteria ─────────────────────────────────────────────────────────────────

func TestGetOutbound_Hysteria_Basic(t *testing.T) {
	uri := "hysteria://hy.example.com:4443?auth=myauth&upmbps=100&downmbps=200#HY1"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "hysteria" {
		t.Errorf("type: got %q, want hysteria", str(out, "type"))
	}
	if str(out, "server") != "hy.example.com" {
		t.Errorf("server: got %q, want hy.example.com", str(out, "server"))
	}
	if str(out, "auth_str") != "myauth" {
		t.Errorf("auth_str: got %q, want myauth", str(out, "auth_str"))
	}
}

// ── shadowsocks ───────────────────────────────────────────────────────────────

func TestGetOutbound_Shadowsocks_Basic(t *testing.T) {
	// ss://method:password@host:port#tag (plain userinfo)
	uri := "ss://chacha20-ietf-poly1305:mypassword@ss.example.com:8388#SSTag"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "shadowsocks" {
		t.Errorf("type: got %q, want shadowsocks", str(out, "type"))
	}
	if str(out, "server") != "ss.example.com" {
		t.Errorf("server: got %q, want ss.example.com", str(out, "server"))
	}
}

func TestGetOutbound_Shadowsocks_Base64Userinfo(t *testing.T) {
	// Some clients encode method:password in base64 as the username
	methodPass := "chacha20-ietf-poly1305:encodedpass"
	encoded := base64.StdEncoding.EncodeToString([]byte(methodPass))
	uri := "ss://" + encoded + "@ss.example.com:8388#SSBase64"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "shadowsocks" {
		t.Errorf("type: got %q, want shadowsocks", str(out, "type"))
	}
	if str(out, "method") != "chacha20-ietf-poly1305" {
		t.Errorf("method: got %q", str(out, "method"))
	}
}

// ── tuic ─────────────────────────────────────────────────────────────────────

func TestGetOutbound_Tuic_Basic(t *testing.T) {
	uri := "tuic://my-uuid:mypassword@tuic.example.com:443?sni=tuic.example.com#TuicTag"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "tuic" {
		t.Errorf("type: got %q, want tuic", str(out, "type"))
	}
	if str(out, "uuid") != "my-uuid" {
		t.Errorf("uuid: got %q, want my-uuid", str(out, "uuid"))
	}
}

// ── anytls ────────────────────────────────────────────────────────────────────

func TestGetOutbound_Anytls_Basic(t *testing.T) {
	uri := "anytls://secretpass@anytls.example.com:443?sni=anytls.example.com#AnyTLSTag"
	out := mustGetOutbound(t, uri)

	if str(out, "type") != "anytls" {
		t.Errorf("type: got %q, want anytls", str(out, "type"))
	}
	if str(out, "password") != "secretpass" {
		t.Errorf("password: got %q, want secretpass", str(out, "password"))
	}
}

// ── vmess ─────────────────────────────────────────────────────────────────────

func makeVmessURI(t *testing.T, data map[string]interface{}) string {
	t.Helper()
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	return "vmess://" + encoded
}

func TestGetOutbound_Vmess_Basic(t *testing.T) {
	data := map[string]interface{}{
		"v":    2,
		"ps":   "MyVmess",
		"add":  "vmess.example.com",
		"port": 443,
		"id":   "some-uuid",
		"aid":  0,
		"net":  "tcp",
		"type": "",
		"tls":  "tls",
		"sni":  "vmess.example.com",
	}
	out := mustGetOutbound(t, makeVmessURI(t, data))

	if str(out, "type") != "vmess" {
		t.Errorf("type: got %q, want vmess", str(out, "type"))
	}
	if str(out, "tag") != "MyVmess" {
		t.Errorf("tag: got %q, want MyVmess", str(out, "tag"))
	}
	if str(out, "uuid") != "some-uuid" {
		t.Errorf("uuid: got %q, want some-uuid", str(out, "uuid"))
	}
}

func TestGetOutbound_Vmess_WsTransport(t *testing.T) {
	data := map[string]interface{}{
		"v":    2,
		"ps":   "WsVmess",
		"add":  "vmess.example.com",
		"port": 443,
		"id":   "ws-uuid",
		"aid":  0,
		"net":  "ws",
		"path": "/ws",
		"host": "vmess.example.com",
		"tls":  "",
	}
	out := mustGetOutbound(t, makeVmessURI(t, data))
	transport := mapVal(out, "transport")
	if transport == nil {
		t.Fatal("expected transport map")
	}
	if transport["type"] != "ws" {
		t.Errorf("transport.type: got %v, want ws", transport["type"])
	}
	if transport["path"] != "/ws" {
		t.Errorf("transport.path: got %v, want /ws", transport["path"])
	}
}

func TestGetOutbound_Vmess_GrpcTransport(t *testing.T) {
	data := map[string]interface{}{
		"v":    2,
		"ps":   "GrpcVmess",
		"add":  "vmess.example.com",
		"port": 443,
		"id":   "grpc-uuid",
		"aid":  0,
		"net":  "grpc",
		"path": "myService",
		"tls":  "",
	}
	out := mustGetOutbound(t, makeVmessURI(t, data))
	transport := mapVal(out, "transport")
	if transport == nil {
		t.Fatal("expected transport map")
	}
	if transport["type"] != "grpc" {
		t.Errorf("transport.type: got %v, want grpc", transport["type"])
	}
	if transport["service_name"] != "myService" {
		t.Errorf("transport.service_name: got %v, want myService", transport["service_name"])
	}
}

// ── getTls ────────────────────────────────────────────────────────────────────

func TestGetTls_TlsSecurity(t *testing.T) {
	uri := "vless://uuid@host:443?security=tls&sni=myhost.com&fp=chrome&allowInsecure=1#T"
	out := mustGetOutbound(t, uri)
	tls := mapVal(out, "tls")
	if tls == nil {
		t.Fatal("expected tls map")
	}
	if tls["enabled"] != true {
		t.Error("expected tls.enabled = true")
	}
	if tls["server_name"] != "myhost.com" {
		t.Errorf("server_name: got %v", tls["server_name"])
	}
	if tls["insecure"] != true {
		t.Error("expected tls.insecure = true")
	}
	utls := mapVal(tls, "utls")
	if utls == nil {
		t.Fatal("expected utls map")
	}
	if utls["fingerprint"] != "chrome" {
		t.Errorf("fingerprint: got %v", utls["fingerprint"])
	}
}

func TestGetTls_RealitySecurity(t *testing.T) {
	uri := "vless://uuid@host:443?security=reality&pbk=mypubkey&sid=mysid&sni=real.com#R"
	out := mustGetOutbound(t, uri)
	tls := mapVal(out, "tls")
	if tls == nil {
		t.Fatal("expected tls map")
	}
	reality := mapVal(tls, "reality")
	if reality == nil {
		t.Fatal("expected reality map")
	}
	if reality["public_key"] != "mypubkey" {
		t.Errorf("public_key: got %v", reality["public_key"])
	}
	if reality["short_id"] != "mysid" {
		t.Errorf("short_id: got %v", reality["short_id"])
	}
}

// ── getTransport ──────────────────────────────────────────────────────────────

func TestGetTransport_Grpc(t *testing.T) {
	uri := "vless://uuid@host:443?security=tls&type=grpc&serviceName=myGrpcSvc#G"
	out := mustGetOutbound(t, uri)
	transport := mapVal(out, "transport")
	if transport == nil {
		t.Fatal("expected transport map")
	}
	if transport["type"] != "grpc" {
		t.Errorf("type: got %v, want grpc", transport["type"])
	}
	if transport["service_name"] != "myGrpcSvc" {
		t.Errorf("service_name: got %v, want myGrpcSvc", transport["service_name"])
	}
}

func TestGetTransport_Ws(t *testing.T) {
	uri := "vless://uuid@host:443?security=tls&type=ws&path=/mypath&host=cdn.example.com#W"
	out := mustGetOutbound(t, uri)
	transport := mapVal(out, "transport")
	if transport == nil {
		t.Fatal("expected transport map")
	}
	if transport["type"] != "ws" {
		t.Errorf("type: got %v, want ws", transport["type"])
	}
	if transport["path"] != "/mypath" {
		t.Errorf("path: got %v, want /mypath", transport["path"])
	}
	headers, _ := transport["headers"].(map[string]interface{})
	if headers == nil {
		t.Fatal("expected headers map for ws")
	}
	if headers["Host"] != "cdn.example.com" {
		t.Errorf("Host: got %v, want cdn.example.com", headers["Host"])
	}
}

func TestGetTransport_HttpUpgrade(t *testing.T) {
	uri := "vless://uuid@host:443?security=tls&type=httpupgrade&path=/upgrade&host=up.example.com#HU"
	out := mustGetOutbound(t, uri)
	transport := mapVal(out, "transport")
	if transport == nil {
		t.Fatal("expected transport map")
	}
	if transport["type"] != "httpupgrade" {
		t.Errorf("type: got %v, want httpupgrade", transport["type"])
	}
}
