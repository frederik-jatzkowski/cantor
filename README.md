# cantor

[![Go Reference](https://pkg.go.dev/badge/github.com/frederik-jatzkowski/cantor.svg)](https://pkg.go.dev/github.com/frederik-jatzkowski/cantor)
[![Go Report Card](https://goreportcard.com/badge/github.com/frederik-jatzkowski/cantor)](https://goreportcard.com/report/github.com/frederik-jatzkowski/cantor)

The `cantor` package is a versatile library, focusing on providing comprehensive support for set operations based on set theory principles. Its core functionality revolves around the manipulation and management of sets, offering a rich set of operations such as [union](https://en.wikipedia.org/wiki/Union_(set_theory)), [intersection](https://en.wikipedia.org/wiki/Intersection_(set_theory)), [complement](https://en.wikipedia.org/wiki/Complement_(set_theory)), and more. The package is designed with performance and flexibility in mind, leveraging Go's type system and [generic programming features](https://go.dev/doc/tutorial/generics) introduced in Go 1.18.

Key Features of the `cantor` Package:
- **Type Safety and Generics**: Utilizes Go's generics to ensure type safety across set operations, allowing for the creation of sets containing elements of any comparable type. The type system also prevents improper use such as the attempt to iterate infinite sets.
- **Performance Optimization**: Designed with performance considerations, such as minimal memory overhead and efficient computation, making it suitable for high-performance applications.
- **Implicit and Derived Sets**: Allows for the definition of sets using [predicates](https://proofwiki.org/wiki/Definition:Set/Definition_by_Predicate) (implicit sets) and the derivation of new sets from existing ones, enabling complex set expressions and dynamic data views.
- **Real-world Applicability**: Through its comprehensive set operations and dynamic set views, the package is applicable to a wide range of real-world scenarios, such as data filtering, in-memory data management, and more.

The `cantor` package stands out for its implementation of set theory concepts in a way that is both accessible and powerful, providing developers with the tools needed to handle complex data structures and operations efficiently. Whether for managing collections of data in real-time, performing complex set operations, or implementing business logic that requires dynamic data views, `cantor` offers a robust solution.

## Architecture

The following diagram represents the public facing API of `cantor`.

![Architecture of the public API](/docs/media/public_architecture.svg)

## Release Notes

Here you will find a list of the release notes for all versions.

- [v0.1.0](docs/releases/v0.1.0.md)

A list of ideas for future releases can be found [here](docs/roadmap/ideas.md).

## Development Guide

Here, you can find the [development guide](docs/development/guide.md).