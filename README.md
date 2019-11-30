# dns

`dns` is a simple CLI tool for [DNS-LG API](http://www.dns-lg.com)

![screen](./screen.png)

## Install

### Download compiled binary

[Linux](https://github.com/mxssl/dns/releases/download/0.0.3/dns_Linux_x86_64.tar.gz)

[Windows](https://github.com/mxssl/dns/releases/download/0.0.3/dns_Windows_x86_64.tar.gz)

[MacOS](https://github.com/mxssl/dns/releases/download/0.0.3/dns_Darwin_x86_64.tar.gz)

### Examples

Linux:

```bash
wget https://github.com/mxssl/dns/releases/download/0.0.3/dns_Linux_x86_64.tar.gz
tar zvxf dns_Linux_x86_64.tar.gz
cp dns /usr/local/bin/dns
chmod +x /usr/local/bin/dns
```

MacOS

```bash
wget https://github.com/mxssl/dns/releases/download/0.0.3/dns_Darwin_x86_64.tar.gz
tar zvxf dns_Darwin_x86_64.tar.gz
cp dns /usr/local/bin/dns
chmod +x /usr/local/bin/dns
```

## Usage

```bash
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
      --raw               Raw output without color
  -r, --resolver string   Choice dns resolver (default "google1")

Use "dns [command] --help" for more information about a command
```

## Example

```sh
dns a golang.com
{
    "question": [
        {
            "name": "golang.com.",
            "type": "A",
            "class": "IN"
        }
    ],
    "answer": [
        {
            "name": "golang.com.",
            "type": "A",
            "class": "IN",
            "ttl": 299,
            "rdlength": 4,
            "rdata": "216.58.198.81"
        }
    ]
}
```

## Resolver

`google1` is used by default

You can use these resolvers:

| Name | IP |
|---|---|
| cloudflare | 1.1.1.1 |
| google1 | 8.8.8.8 |
| google2 | 8.8.4.4 |
| he | 74.82.42.42 |
| opendns1 | 208.67.222.222 |
| opendns2 | 208.67.220.220 |
| quad9 | 9.9.9.9 |

```sh
dns -r cloudflare a golang.com
{
    "question": [
        {
            "name": "golang.com.",
            "type": "A",
            "class": "IN"
        }
    ],
    "answer": [
        {
            "name": "golang.com.",
            "type": "A",
            "class": "IN",
            "ttl": 47,
            "rdlength": 4,
            "rdata": "172.217.168.49"
        }
    ]
}
```

## Docker

```sh
docker container \
  run \
  --rm \
  mxssl/dns:v0.0.3 \
  dns a golang.com
```
