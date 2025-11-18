# Getting Started with Go and Debugging

## Debugging Methods

### Print Statement

- Pros
  - Part of fmt package
  - Very simple to use
- Cons
  - Concurrency is not supported
  - Can't handle complex scenarios

### Log package

- Pros
  - Part of log package
  - Relatively simple to use
  - Supports concurrency

- Cons
  - Can't handle complex scenarios

### Delve Utility

- Pros
  - Command Line
  - Integrates with IDE
  - Full-featured debugging tooll
  - Supports remote-debug
  - Preferred over GDB

- Cons
  - Lots of features can be overwhelming

### More on Delve

Delve can be found at:

> https://github.com/go-delve/delve

### Running Delve on a Container / remotely


