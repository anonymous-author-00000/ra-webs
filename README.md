# RA-WEBs
[![Go](https://github.com/anonymous-author-00000/ra-webs/actions/workflows/go.yml/badge.svg)](https://github.com/anonymous-author-00000/ra-webs/actions/workflows/go.yml)

RA-WEBs is a protocol that enables browsers to verify proof of Remote Attestation while maintaining compatibility.

### Dependencies

- An Azure instance with Intel SGX (for running the example TA)
- Ubuntu 20.04
- Nix

### How to Deploy the Test Environment

The deployability was verified using DC1ds v3 (1 vCPU, 8 GiB memory).

#### 1. Clone the Repository

```bash
git clone https://github.com/anonymous-author-00000/ra-webs
cd ra-webs/test
```

#### 2. Run the Cloudflare Tunnel (Optional)


```sh
docker compose up tunnel
```

#### 3. Configure the Environment Files

Copy the templates and fill in each parameter.

```sh
cp env/common.env.template env/common.env
cp env/ta.env.template env/ta.env
cp env/monitor.env.template env/monitor.env
```


#### 4. Run the servers

```sh
docker compose up
```


### How to Conduct Formal Verification

```
nix-shell proverif.nix
proverif ra-webs.pv
```

### Demo

The demos are as follows:
- [TA](https://ra-webs-demo-2.eastus2.cloudapp.azure.com)
- [Monitor](https://dip-safe-launches-approved.trycloudflare.com/)

Note that the demo may be suspended due to budget constraints.
