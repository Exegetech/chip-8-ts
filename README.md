# Chip-8

Chip-8 emulator written in Go and TypeScript. CPU is implemented in Go, then compiled to WebAssembly. GUI is written in TypeScript with React.

Requirements:
- Go 1.19.4
- TinyGo 0.26.0
- NodeJS 18.11.0
- NPM 8.19.2

Getting started:
- Install dependencies
  - `cd core && go mod tidy`
  - `cd web && npm install`

- Build
  - `cd core && make build`
  - `cd web && npm run dev`

See `core/Makefile` and `web/package.json` for commands.
