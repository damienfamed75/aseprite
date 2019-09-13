# aseprite

Aseprite JSON loader

<p align="center">
  <a href="https://godoc.org/github.com/damienfamed75/aseprite"><img src="https://godoc.org/github.com/damienfamed75/aseprite?status.svg" alt="GoDoc"></a>
  <img src="https://goreportcard.com/badge/github.com/damienfamed75/aseprite" alt="Go Report Card" /></a>
  <a href="https://github.com/damienfamed75/aseprite/actions"><img src="https://github.com/damienfamed75/aseprite/workflows/Pipeline/badge.svg" alt="Pipeline" /></a>
</p>

Why use this package?
---
Well the mission of this repository was to improve performance and extensability of the [goaseprite](https://github.com/SolarLune/goaseprite) package.

So far this package allows for two different types of aseprite JSON files and more formats to come.

When it came to improving performance it was the smaller changes that differ this one from the others. For instance I'm utilizing Rob Pike's [stringer](https://github.com/golang/tools/tree/master/cmd/stringer) tool and unmarshalling JSONs with builtin Golang "overriding"

Examples
---
For examples on how to use this package please check out the `examples/` directory.

Future documentation to come for how to export JSONs from aseprite.

Special Thanks
---
All credit to [SolarLune](https://github.com/SolarLune) for making [goaseprite](https://github.com/SolarLune/goaseprite).
