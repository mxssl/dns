# dns

`dns` is a simple CLI tool for [DNS-LG API](http://www.dns-lg.com)

```sh
dns a github.com
{
    "question": [
        {
            "name": "github.com.",
            "type": "A",
            "class": "IN"
        }
    ],
    "answer": [
        {
            "name": "github.com.",
            "type": "A",
            "class": "IN",
            "ttl": 60,
            "rdlength": 4,
            "rdata": "140.82.113.4"
        }
    ]
}
```

## Install

### Download compiled binary

[Linux amd64](https://github.com/mxssl/dns/releases/download/1.0.7/dns_Linux_x86_64.tar.gz)

[MacOS amd64](https://github.com/mxssl/dns/releases/download/1.0.7/dns_darwin_amd64.tar.gz)

[MacOS arm64](https://github.com/mxssl/dns/releases/download/1.0.7/dns_darwin_arm64.tar.gz)

### Examples

Linux amd64:

```bash
wget https://github.com/mxssl/dns/releases/download/1.0.7/dns_linux_amd64.tar.gz
tar zvxf dns_linux_amd64.tar.gz
mv dns /usr/local/bin/dns
chmod +x /usr/local/bin/dns
rm dns_linux_amd64.tar.gz
```

MacOS amd64:

```bash
wget https://github.com/mxssl/dns/releases/download/1.0.7/dns_darwin_amd64.tar.gz
tar zvxf dns_darwin_amd64.tar.gz
mv dns /usr/local/bin/dns
chmod +x /usr/local/bin/dns
rm dns_darwin_amd64.tar.gz
```

MacOS arm64:

```bash
wget https://github.com/mxssl/dns/releases/download/1.0.7/dns_darwin_arm64.tar.gz
tar zvxf dns_darwin_arm64.tar.gz
mv dns /usr/local/bin/dns
chmod +x /usr/local/bin/dns
rm dns_darwin_arm64.tar.gz
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
dns a go.dev
{
    "question": [
        {
            "name": "go.dev.",
            "type": "A",
            "class": "IN"
        }
    ],
    "answer": [
        {
            "name": "go.dev.",
            "type": "A",
            "class": "IN",
            "ttl": 300,
            "rdlength": 4,
            "rdata": "216.239.34.21"
        },
        {
            "name": "go.dev.",
            "type": "A",
            "class": "IN",
            "ttl": 300,
            "rdlength": 4,
            "rdata": "216.239.32.21"
        },
        {
            "name": "go.dev.",
            "type": "A",
            "class": "IN",
            "ttl": 300,
            "rdlength": 4,
            "rdata": "216.239.36.21"
        },
        {
            "name": "go.dev.",
            "type": "A",
            "class": "IN",
            "ttl": 300,
            "rdlength": 4,
            "rdata": "216.239.38.21"
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

## Docker

```sh
docker container \
  run \
  --rm \
  mxssl/dns:1.0.5 \
  dns a golang.com
```

## Development

[Mage](https://github.com/magefile/mage) is used as an alternative to `make`.

```sh
mage -l
Targets:
  build                  the app
  clean                  delete compiled binary
  dockerBuild            build container with latest tag
  dockerRelease          build and push container to the registry
  dockerTestRun          test run latest container
  gitHubRelease          run goreleaser
  gitHubReleaseDryRun    goreleaser dry run
  lint                   the app
```
