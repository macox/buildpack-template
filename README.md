## Overview

### Prerequisites
[pack cli](https://github.com/buildpacks/pack)

### Building

```shell script
./scripts/build.sh
```
This creates bin/detect and bin/build binaries

### Packaging

#### local
```shell script
pack package-buildpack buildpack-helloworld --config package/package.toml
```

#### Publishing
```shell script
pack package-buildpack --publish <DOCKER_REGISTRY>/buildpack-helloworld  --config package/package.toml
```