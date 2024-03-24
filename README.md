# cantor

[![Go Reference](https://pkg.go.dev/badge/github.com/frederik-jatzkowski/cantor.svg)](https://pkg.go.dev/github.com/frederik-jatzkowski/cantor)
[![Go Report Card](https://goreportcard.com/badge/github.com/frederik-jatzkowski/cantor)](https://goreportcard.com/report/github.com/frederik-jatzkowski/cantor)

The `cantor` package is a comprehensive toolkit for set operations for the Go Programming Language.

Features:

- **Basic and Advanced Set Operations**:

  - [Union](<https://en.wikipedia.org/wiki/Union_(set_theory)>)
  - [Intersection](<https://en.wikipedia.org/wiki/Intersection_(set_theory)>)
  - [Complement](<https://en.wikipedia.org/wiki/Complement_(set_theory)>)
  <!-- - [Difference](<https://en.wikipedia.org/wiki/Complement_(set_theory)#Relative_complement>) -->
  <!-- - [Symmetric Difference](https://en.wikipedia.org/wiki/Symmetric_difference)
  - [Equality](https://proofwiki.org/wiki/Definition:Set_Equality) and [Comparisons](https://en.wikipedia.org/wiki/Subset) -->

- **Infinite, Implicit Sets**:
  In addition to [HashSets](https://go.dev/blog/maps), you can implicitly define sets using a [Predicate](https://proofwiki.org/wiki/Definition:Set/Definition_by_Predicate).
  Such sets can represent infinitely many elements and be used within certain limitations together with other types of sets.

- **Dynamic Data Views**:
  If you use set operations in `cantor`, you will define a dynamic data view on the underlying sets without doing any evaluation upfront.
  Such a derived set can be used for lookups, iteration or be used to construct new HashSets.

- **Performance by Design**:
  Due to the powerful yet simple [lazy evaluation](https://en.wikipedia.org/wiki/Lazy_evaluation) using boolean expressions, no intermediate results have to be allocated.
  This makes the usage of complex set operations highly performant and keeps pressure off the garbage collector.

- **Type Safety and Generics**:
  Utilizes Go's generics and a system of interfaces to ensure type safety across set operations. The type system also ensures correct and sensible usage of the components of this package.

- **Stability and Confidence**:

  - An ever-growing test suite including more than `175` tests.
  - `100%` code coverage enforced by `CI`.
  - High code quality, enforced by `golangci-lint`.
  - Guaranteed compatibility with Go `v1.18` onwards.

<!-- - **Extensibility**:
The exposed interfaces can be implemented  -->

## Architecture

The following diagram represents the public facing API of `cantor`. The transparent parts of the diagram are not yet implemented but planned.

![Architecture of the public API](/docs/media/public_architecture.svg)

## Release Notes

Here you will find a list of the release notes for all versions.

- [v0.1.0](docs/releases/v0.1.0.md)
- [v0.2.0](docs/releases/v0.2.0.md)

## Development Guide

Here, you can find the [development guide](docs/development/guide.md).
