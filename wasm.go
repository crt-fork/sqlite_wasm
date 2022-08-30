package main

import (
	"os"

	"github.com/bytecodealliance/wasmtime-go"
	"go.riyazali.net/sqlite"
)

type WASMFunc struct{}

func (*WASMFunc) Deterministic() bool { return true }
func (*WASMFunc) Args() int           { return -1 }
func (*WASMFunc) Apply(c *sqlite.Context, values ...sqlite.Value) {
	store := wasmtime.NewStore(wasmtime.NewEngine())
	wasm, err := os.ReadFile(values[0].Text())
	if err != nil {
		c.ResultError(err)
	}

	if err != nil {
		c.ResultError(err)
	}

	module, err := wasmtime.NewModule(store.Engine, wasm)
	if err != nil {
		c.ResultError(err)
	}

	instance, err := wasmtime.NewInstance(store, module, nil)
	if err != nil {
		c.ResultError(err)
	}

	run := instance.GetFunc(store, values[1].Text())
	if run == nil {
		c.ResultError(err)
	}

	// Parse arguments
	argSlice := values[2:]
	args := make([]interface{}, len(argSlice))
	for i, val := range argSlice {
		var arg interface{}
		switch val.Type() {
		case sqlite.SQLITE_TEXT:
			arg = val.Text()
		case sqlite.SQLITE_INTEGER:
			arg = val.Int()
		case sqlite.SQLITE_FLOAT:
			arg = val.Float()
		case sqlite.SQLITE_BLOB:
			arg = val.Blob()
		case sqlite.SQLITE_NULL:
			arg = nil
		}

		args[i] = arg
	}

	val, err := run.Call(store, args...)
	if err != nil {
		c.ResultError(err)
	}

	c.ResultInt(int(val.(int32)))
}

func RegisterWASM(api *sqlite.ExtensionApi) error {
	if err := api.CreateFunction("wasm_exec", &WASMFunc{}); err != nil {
		return err
	}

	return nil
}
