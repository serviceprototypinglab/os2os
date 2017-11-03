# os2os

`os2os` is a command line interface tool which helps `openshift` users to move their apps
 between different `openshift` clusters or providers.
 
 `os2os` will be connected to two clusters (old cluster and new cluster) and with simple commands will move or copy
 an `openshift` project from the old cluster to the new cluster. `os2os` uses the command line client of `openshift`, `oc`, to 
 connect, extract the objects templates, upload the objects and optionally delete the old project. 

 A unique advantage of `os2os` is that it combines multiple steps into a convenient workflow including the migration of data volumes.

 `os2os` is a prototype from the Cloud-Native Applications research initiative of the Service Prototyping Lab at Zurich University of Applied Sciences. Use with care, things may break. We will share our findings on cloud application migration at a later point in time.

## Use Case

Use the command `help` to see the list of command and the description of them.
With `export`, you will export the templates from a cluster.
With `up`, you will uploads your templates.
With `down`, you will delete a project in a cluster.

The configuration can be added in `~/.os2os.yaml` or directly using the flags.
Use help to see all the flags.
If no configuration is given to the tool, it will take the default values.

```
    os2os help
    os2os export
    os2os up
    os2os down
    os2os ...
```

## Installation

### Install `oc`

- https://docs.openshift.org/latest/cli_reference/get_started_cli.html#installing-the-cli

### From binary

#### Install the binary: `os2os`

Download the binary from /binaries/< your operative system> and run:

```
    chmod +x os2os
    sudo mv ./os2os /usr/local/bin/os2os
```

### From source

Note: We are in the progress of converting to a default Go project structure.
Until then, please use the following commands to compile and install:
 
```
    git clone <this repository>
    mv os2os $GOPATH/src/os2os
    go get github.com/mitchellh/go-homedir
    go get github.com/spf13/cobra
    go get github.com/spf13/viper
    cd $GOPATH/src/os2os
    go build os2os
    chmod +x os2os
    sudo mv ./os2os /usr/local/bin/os2os
```

## First steps

This small example shows how to migrate an OpenShift application from a local OpenShift development cluster
to APPUiO, the Swiss Container Platform.

```
    os2os migrate \
          --clusterFrom https://127.0.0.1:8443 --clusterTo https://console.appuio.ch:443 \
          --projectFrom test --projectTo test \
          --usernameFrom user --usernameTo user \
          --passwordFrom pass --passwordTo pass
```

Considering the large number of options, it is advised to use the configuration file `~/.os2os.yaml`
to store all parameters (in YAML syntax).
