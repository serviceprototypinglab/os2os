
# os2os

`os2os` is a command line interface tool which help `openshift` users to move their apps
 between different `openshift` clusters.
 
 `os2os` will be connected to two clusters (old cluster and new cluster) and with simple commands will move
 an openshift project from the old cluster to the new cluster. `os2os` use the command line of openshift, `oc`, to 
 connect, extract the objects templates, upload the objects and delete the old project. 


## Use Case

Use the command `help` to see the list of command and the description of them.
With `export`, you will export the templates from a cluster.
With `up`, you will uploads your templates.
With `down`, you will delete a project in a cluster.

The configuration can be added in ~/.os2os.yaml or directly with the flags.
Use help to see all the flags.
If no configuration is given to the tool, it will take the default values.

```
    os2os help
    os2os export
    os2os up
    os2os down
```

## Installation

### Install `oc`

- https://docs.openshift.org/latest/cli_reference/get_started_cli.html#installing-the-cli

### From binary

### Install the binary: `os2os`

Download the binary from /binaries/<your operative system> and run:

```
    chmod +x os2os
    sudo mv ./os2os /usr/local/bin/os2os
```

### From source
 
```
    git clone
    go build os2os
```