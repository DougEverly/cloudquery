---
title: "plugin_publish"
---
## cloudquery plugin publish

Publish to CloudQuery Hub.

### Synopsis

Publish to CloudQuery Hub.

This publishes a plugin version to CloudQuery Hub from a local dist directory.


```
cloudquery plugin publish [-D dist] [flags]
```

### Examples

```

# Publish a plugin version from a local dist directory
cloudquery plugin publish
```

### Options

```
  -D, --dist-dir string   Path to the dist directory (default "dist")
  -f, --finalize          Finalize the plugin version after publishing. If false, the plugin version will be marked as draft=true.
  -h, --help              help for publish
```

### Options inherited from parent commands

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

### SEE ALSO

* [cloudquery plugin](/docs/reference/cli/cloudquery_plugin)	 - Plugin commands

