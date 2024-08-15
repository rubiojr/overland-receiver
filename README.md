# Overland Receiver

This is a simple receiver for the [Overland](https://github.com/aaronpk/Overland-iOS) project.

## Installation

### For Linux systems with systemd

```bash
make install
```

It will install a systemd user service unit that will start the receiver on boot.

## Usage

The service listens on `:3111` by default. You can change this using the `--address` flag.
