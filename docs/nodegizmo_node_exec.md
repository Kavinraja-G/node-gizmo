## nodegizmo node exec

Spawns a 'nsenter' pod to exec into the provided node

```
nodegizmo node exec nodeName [flags]
```

### Options

```
  -h, --help               help for exec
  -i, --image string       Image used by nsenter pod (default "docker.io/alpine:3.18")
  -n, --namespace string   Namespace where nsenter pod to be created (default "kube-system")
  -t, --ttl string         Time to live (seconds) for the exec container. Defaults to 3600s (default "3600")
```

### Options inherited from parent commands

```
      --sort-by string   Sorts output using a valid Column name. Defaults to 'name' if the column name is not valid (default "name")
```

### SEE ALSO

* [nodegizmo node](nodegizmo_node.md)	 - Displays generic node related information in the cluster

###### Auto generated by spf13/cobra on 24-Dec-2023
