# redirects

Small redirection service. Used mostly for QR codes in order to not statically assign a link.

Links are defined in `config.yaml`.

## Development

```bash
# https://github.com/silenceper/gowatch
# Hot reloads whenever saved
gowatch

# Or plain
go run main.go
```

## Deployment

Tag a version number in a release, the github action will create according docker images and publish them under ghcr.io/nms-ev/redirects.
