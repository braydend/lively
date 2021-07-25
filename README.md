# Lively

A simple notification app to remind me to move my body when I'm working

## Running

To run from source:
1. Install Go 1.16
2. Run `go run .` in the project root.

To run the binary:
1. Download latest binary from Releases
2. Add execute permissions to binary with chmod (e.g. `chmod +x lively`)
3. Run the binary. `./lively`

## Arguments

### Config

`go run . --config {path-to-config-file}`
`./lively --config {path-to-config-file}`

The `--config` flag allows you to pass an optional YAML configuration file. The following properties are available for customisation (see defaults [here](config.yml)):
- Start and end time (in 24h format).
- Interval that the app will remind at.
- Messages that will appear in the notifications.
