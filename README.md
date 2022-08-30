# sqlite_wasm

⚠️ **This library is experimental, and should not be used under any
circumstances at this time** ⚠️

A SQLite extension for running arbitrary WASM bundles via Wasmtime.

## Usage

### Prerequisites

- Go
- SQLite3 library/CLI with runtime extensions turned on. The default SQLite on
  MacOS does not ship with runtime extension support, the homebrew version does.
- Language that compiles to WASM. AssemblyScript examples are available in the
  `examples/` folder.

### Building

```
$ asc examples/add.ts -o add.wasm
$ go build -buildmode=c-shared -o sqlite_wasm.so # .dylib/.dll/.so depending on platform
$ sqlite3
SQLite version 3.39.2 2022-07-21 15:24:47
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .load sqlite_wasm.so
sqlite> select wasm_exec("add.wasm", "add", 5, 7);
12
sqlite>
```

## Current Status

- Working with an arbitrary number of numeric arguments
- Higher level types aren't supported yet, and requires further consideration.
  Passing strings requires binding memory between the host. Issues/PRs around
  this are welcome!
- Needs tests
