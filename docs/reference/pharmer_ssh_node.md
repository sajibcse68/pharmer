## pharmer ssh node

SSH into a Kubernetes cluster instance

### Synopsis


SSH into a cluster instance.

```
pharmer ssh node [flags]
```

### Examples

```
pharmer ssh node -k cluster-name node-name
```

### Options

```
  -k, --cluster string   Name of cluster
  -h, --help             help for node
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Guard (default true)
      --config-file string               Path to Pharmer config file
      --env string                       Environment used to enable debugging (default "dev")
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default true)
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [pharmer ssh](pharmer_ssh.md)	 - 

