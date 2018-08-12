# builder

[![Build Status](https://travis-ci.com/aufaitio/builder.svg?branch=master)](https://travis-ci.com/aufaitio/builder)
[![Coverage Status](https://coveralls.io/repos/github/aufaitio/builder/badge.svg?branch=master)](https://coveralls.io/github/aufaitio/builder?branch=master)

Micro service responsible for performing updates and providing notifications around dependency updates and build status.

## Config

The builder config should be in Yaml format and named app.yaml.

```yaml
# Default config values set by application. Outlined to illustrate config structure.
db:
    host: localhost
    port: 27017
    name: aufait
    username: null
    password: null
```

## Development

### CLI

#### Usage

`./server [--configPath=<path>]`

#### Options

```
-h --help			 Show this message
--version			 Show version info
--configPath=<path>  Path to app.yaml config file [default: config]
```

#### Examples

```bash
# Pass --config if you need to override config options.
mongod --dbpath <dbPath>
./server
```
