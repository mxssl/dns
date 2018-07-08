dns is a simple CLI for [DNS-LG](http://www.dns-lg.com)

# Install

### Option 1

#### Compile

```
mkdir -p ~/go/src/github.com/mxssl/dns
cd ~/go/src/github.com/mxssl/dns
git clone https://github.com/mxssl/dns.git
dep ensure
go build -o dns
mv dns /bin/dns
chmod +x /bin/dns
```

### Option 2

#### Download compiled binary

Linux: https://github.com/mxssl/dns/releases/download/v0.0.1/dns-linux-amd64

Windows: https://github.com/mxssl/dns/releases/download/v0.0.1/dns-windows-amd64.exe

MacOS: https://github.com/mxssl/dns/releases/download/v0.0.1/dns-darwin-amd64

# Usage

```
# dns
dns is a CLI for DNS-LG API.

Usage:
  dns [command]

Available Commands:
  a           Get Host Address (A records)
  aaaa        Get IPv6 Host Address (AAAA records)
  cert        Get Certificate (CERT records)
  cname       Get Canonical Name (CNAME records)
  dhcid       Get DHCP Identifier (DHCID records)
  dlv         Get DNSSEC Lookaside Validation record (DLV records)
  dname       Get Delegation name (DNAME records)
  dnskey      Get DNS Key record (DNSKEY records)
  ds          Get Delegation Signer (DS records)
  help        Help about any command
  hinfo       Get Host Information (HINFO records)
  hip         Get Host Identity Protocol (HIP records)
  ipseckey    Get IPSec Key (IPSECKEY records)
  kx          Get Key eXchanger record (KX records)
  loc         Get Location record (LOC records)
  mx          Get Mail Exchange record (MX records)
  naptr       Get Name Authority Pointer (NAPTR records)
  ns          Get Name Servers (NS records)
  nsec        Get Next-Secure record (NSEC records)
  nsec3       Get NSEC record version 3 (NSEC3 records)
  nsec3param  Get NSEC3 parameters (NSEC3PARAM records)
  opt         Get Option record (OPT records)
  ptr         Get Pointer record (PTR records)
  rptr4       Get reverse (PTR) record from IPv4 addresses
  rptr6       Get reverse (PTR) record from IPv4 addresses
  rrsig       Get Resource Records Signature (RRSIG records)
  soa         Get Start of Authority (SOA record)
  spf         Get Sender Policy Framework (SPF records)
  srv         Get Service Locator (SRV records)
  sshfp       Get SSH Public Key Fingerprint (SSHFP records)
  ta          Get DNSSEC Trust Authorities (TA records)
  talink      Get Trust Anchor LINK (TALINK records)
  tlsa        Get TLSA records
  txt         Get Text record (TXT records)

Flags:
  -h, --help              help for dns
  -r, --resolver string   Choice dns resolver (default "google1")

Use "dns [command] --help" for more information about a command.

```

## Example

```
# dns a mxssl.github.com
{
    "question": [
        {
            "name": "mxssl.github.com.",
            "type": "A",
            "class": "IN"
        }
    ],
    "answer": [
        {
            "name": "mxssl.github.com.",
            "type": "CNAME",
            "class": "IN",
            "ttl": 3599,
            "rdlength": 18,
            "rdata": "github.github.io."
        },
        {
            "name": "github.github.io.",
            "type": "CNAME",
            "class": "IN",
            "ttl": 3599,
            "rdlength": 27,
            "rdata": "sni.github.map.fastly.net."
        },
        {
            "name": "sni.github.map.fastly.net.",
            "type": "A",
            "class": "IN",
            "ttl": 3599,
            "rdlength": 4,
            "rdata": "185.199.108.153"
        },
        {
            "name": "sni.github.map.fastly.net.",
            "type": "A",
            "class": "IN",
            "ttl": 3599,
            "rdlength": 4,
            "rdata": "185.199.109.153"
        },
        {
            "name": "sni.github.map.fastly.net.",
            "type": "A",
            "class": "IN",
            "ttl": 3599,
            "rdlength": 4,
            "rdata": "185.199.110.153"
        },
        {
            "name": "sni.github.map.fastly.net.",
            "type": "A",
            "class": "IN",
            "ttl": 3599,
            "rdlength": 4,
            "rdata": "185.199.111.153"
        }
    ]
}

```