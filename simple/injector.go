//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)

	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository)

	return nil
}

var fooset = wire.NewSet(NewFooRepository, NewFooService)
var barset = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooset, barset, NewFooBarService)

	return nil
}

// Binding Interface
// buat binding set apabila kita ingin inject interface dengan struct
var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

// var FooBarSet = wire.NewSet(
// 	NewFoo,
// 	NewBar,
// )

// func InitializedFooBar() *FooBar {
// 	wire.Build(FooBarSet, wire.Struct(new(FooBar), "Foo", "Bar"))
// 	return nil
// }

// kalau tidak menggunakan struct provider
var FooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializedFooBar() *FooBar {
	wire.Build(FooBarSet, NewFooBar)
	return nil
}

// Binding Value
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), NewFooBar)
	return nil
}

// binding interface value
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// struct field provider
func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
