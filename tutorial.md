#os2os tutorial

## Prerequisites

One cluster with a project.
Another cluster where to migrate the project.

## Instructions

1. Change the config file to your cluster and project.
2. Run the next sequence to migrate the project.

```
os2os export
os2os up

os2os exportData
os2os upData
os2os downData

os2os down
```