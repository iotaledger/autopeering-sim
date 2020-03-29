<h2 align="center">A simulator for testing the IOTA autopeering module</h2>

<p align="center">
  <a href="https://discord.iota.org/" style="text-decoration:none;"><img src="https://img.shields.io/badge/Discord-9cf.svg?logo=discord" alt="Discord"></a>
    <a href="https://iota.stackexchange.com/" style="text-decoration:none;"><img src="https://img.shields.io/badge/StackExchange-9cf.svg?logo=stackexchange" alt="StackExchange"></a>
    <a href="https://github.com/iotaledger/autopeering-sim/blob/master/LICENSE" style="text-decoration:none;"><img src="https://img.shields.io/github/license/iotaledger/autopeering-sim.svg" alt="Apache 2.0 license"></a>
    <a href="https://golang.org/doc/install" style="text-decoration:none;"><img src="https://img.shields.io/github/go-mod/go-version/iotaledger/autopeering-sim" alt="Go version"></a>
    <a href="https://travis-ci.org/iotaledger/autopeering-sim" style="text-decoration:none;"><img src="https://travis-ci.org/iotaledger/autopeering-sim.svg?branch=master" alt="Build status"></a>
    <a href="https://goreportcard.com/report/github.com/iotaledger/autopeering-sim" style="text-decoration:none;"><img src="https://goreportcard.com/badge/github.com/iotaledger/autopeering-sim" alt="Go Report Card"></a>
</p>
      
<p align="center">
  <a href="#about">About</a> ◈
  <a href="#design">Design</a> ◈
  <a href="#getting-started">Getting started</a> ◈
  <a href="#supporting-the-project">Supporting the project</a> ◈
  <a href="#joining-the-discussion">Joining the discussion</a> 
</p>

---

## About

This repository is where the IOTA Foundation's Research Team simulates tests on the [autopeering module](https://coordicide.iota.org/module2) to study and evaluate its performance.

By making this repository open source, the goal is to allow you to run your own simulations and get involved with development.

To find out more details about autopeering, see the following resources:

- [Coordicide White Paper](https://files.iota.org/papers/Coordicide_WP.pdf) by Coordicide Team, IOTA Foundation
- [Coordicide update - Autopeering: Part 1](https://blog.iota.org/coordicide-update-autopeering-part-1-fc72e21c7e11) by Dr. Angelo Capossele
- [How do we achieve a verifiably random network topology?](https://www.youtube.com/watch?v=-NZVwdZdZk4) by Dr. Hans Moog

You can also see a working example of autopeering in our prototype node software called [GoShimmer](https://github.com/iotaledger/goshimmer).

## Design

The autopeering module is divided into two submodules:

- Peer discovery: Responsible for operations such as discovering new peers and verifying their online status

- Neighbor selection: Responsible for finding and managing neighbors

![Autopeering design](images/autopeering.png)

## Prerequisites

To complete this guide, you need to have at least [version 1.13 of Go](https://golang.org/doc/install) installed on your device.

To check if you have Go installed, run the following command:

```bash
go version
```

## Getting started

To get started, follow these steps to build and run the simulator.

1. Clone this repository

    ```bash
    git clone https://github.com/iotaledger/autopeering-sim.git
    ```

2. Change into the `autopeering-sim` directory
    
    ```bash
    cd autopeering-sim
    ```

3. Build the executable file

    ```bash
    go build -o sim
    ```

4. If you're using Windows, append the `.exe` file extension to the `sim` file

5. Run the simulation

    ```
    ./sim
    ```

6. Open a web browser and go to `http://localhost:8844` to see the simulator

![visualize simulation](images/animation.gif)

The highlighted colors show the following:

**Blue line:** New connections between neighbors
**Red line:** A dropped connection between neighbors
**Blue circle:** A node that is accepting a peering request
**Green circle:** A node that is sending a peering request

### Examining the data

To analyse the results of the simulation, read the `.csv` files in the `data` directory:

- **comvAnalysis**: Proportion of nodes with a complete neighborhood and average number of neighbors as a function of time
- **linkAnalysis**: Probability Density Function (PDF) of the time a given link stays active
- **msgAnalysis**: Number of peering requests sent, accepted, rejected, received and the number of connections dropped of each peer, as well as their average

### Visualizing the data

To generate graphs of the data, run the provided Python script.

You must have Python and PIP installed to run this script. The script generates graphs in `.eps` files, so to view the graphs, you also need an application that can open these files.

1. Install the dependencies

    ```bash
    pip install numpy matplotlib
    ```

2. Run the script `plot.py` script from the `simulation` directory

    ```
    python plot.py
    ```

The graphs provide two figures:
- The proportion of nodes with a complete neighborhood and the average number of neighbors as a function of time
- The Probability Density Function (PDF) of the time a given link stays active

![Example graph](images/graph1.png)

### Parameters

These parameters affect how the simulation behaves. As a result, changing these parameters has an affect on how long the protocol takes to converge.

To change any of these parameters, edit them in the `config.json` file.

|   **Parameter**     |    **Type**   | **Description**    |
|---------------------|:-------------:|:--------------|    
|   `NumberNodes`     |   int         | Number of nodes that try to autopeer in the simulation |
|   `Duration`        |   int         | Duration of the simulation in seconds |
|   `SaltLifetime`    |   int         | How often the public salt changes for each node in seconds |
|   `VisualEnabled`   |   bool        | Whether the visualization server is enabled |
|   `DropOnUpdate`    |   bool        | Whether to drop all connections to neighbors each time the `SaltLifetime` parameter expires |

### Protobuf files

The messages exchanged during autopeering are serialized using [Protocol Buffers](https://developers.google.com/protocol-buffers/).
To generate the corresponding Go files after changing the protobuf files, use the following command:

```bash
make compile
```

## Supporting the project

If you want to contribute to the code, consider posting a [bug report](https://github.com/iotaledger/autopeering-sim/issues/new-issue), feature request or a [pull request](https://github.com/iotaledger/autopeering-sim/pulls/).

See the [contributing guidelines](.github/CONTRIBUTING.md) for more information.

## Joining the discussion

If you want to get involved in the community, need help getting started, have any issues related to the repository or just want to discuss blockchain, distributed ledgers, and IoT with other people, feel free to join our [Discord](https://discord.iota.org/).
