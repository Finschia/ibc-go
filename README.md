# ibc-go
<div align="center">
  <a href="https://github.com/Finschia/ibc-go/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/Finschia/ibc-go.svg" />
  </a>
  <a href="https://github.com/Finschia/ibc-go/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/Finschia/ibc-go.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/Finschia/ibc-go?tab=doc">
    <img alt="GoDoc" src="https://godoc.org/github.com/Finschia/ibc-go?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/Finschia/ibc-go">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/Finschia/ibc-go" />
  </a>
  <a href="https://codecov.io/gh/Finschia/ibc-go">
    <img alt="Code Coverage" src="https://codecov.io/gh/Finschia/ibc-go/branch/main/graph/badge.svg" />
  </a>
</div>
<div align="center">
  <a href="https://github.com/Finschia/ibc-go">
    <img alt="Lines Of Code" src="https://tokei.rs/b1/github/Finschia/ibc-go" />
  </a>
  <a href="https://discord.gg/AzefAFd">
    <img alt="Discord" src="https://img.shields.io/discord/669268347736686612.svg" />
  </a>
  <a href="https://sourcegraph.com/github.com/Finschia/ibc-go?badge">
    <img alt="Imported by" src="https://sourcegraph.com/github.com/Finschia/ibc-go/-/badge.svg" />
  </a>
    <img alt="Lint Status" src="https://github.com/Finschia/ibc-go/workflows/Lint/badge.svg" />
</div>

The Inter-Blockchain Communication protocol (IBC) allows blockchains to talk to each other. IBC handles transport across different sovereign blockchains. This end-to-end, connection-oriented, stateful protocol provides reliable, ordered, and authenticated communication between heterogeneous blockchains. This IBC implementation in Golang is built as a Cosmos SDK module.

## Contents

1. **[Core IBC Implementation](https://github.com/Finschia/ibc-go/tree/main/modules/core)**

    1.1 [ICS 02 Client](https://github.com/Finschia/ibc-go/tree/main/modules/core/02-client)

    1.2 [ICS 03 Connection](https://github.com/Finschia/ibc-go/tree/main/modules/core/03-connection)

    1.3 [ICS 04 Channel](https://github.com/Finschia/ibc-go/tree/main/modules/core/04-channel)

    1.4 [ICS 05 Port](https://github.com/Finschia/ibc-go/tree/main/modules/core/05-port)

    1.5 [ICS 23 Commitment](https://github.com/Finschia/ibc-go/tree/main/modules/core/23-commitment/types)

    1.6 [ICS 24 Host](https://github.com/Finschia/ibc-go/tree/main/modules/core/24-host)

2. **Applications**

    2.1 [ICS 20 Fungible Token Transfers](https://github.com/Finschia/ibc-go/tree/main/modules/apps/transfer)

    2.2 [ICS 27 Interchain Accounts](https://github.com/Finschia/ibc-go/tree/main/modules/apps/27-interchain-accounts)

3. **Light Clients**

    3.1 [ICS 06 Solo Machine](https://github.com/Finschia/ibc-go/tree/main/modules/light-clients/06-solomachine)

    3.2 [ICS 07 Tendermint](https://github.com/Finschia/ibc-go/tree/main/modules/light-clients/07-tendermint)

Note: The localhost client is currently non-functional. 

## Roadmap

For an overview of upcoming changes to ibc-go take a look at the [roadmap](./docs/roadmap/roadmap.md).

## Resources

- [IBC Website](https://ibcprotocol.org/)
- [IBC Specification](https://github.com/cosmos/ibc)
- [Documentation](https://ibc.cosmos.network/main/ibc/overview.html)
