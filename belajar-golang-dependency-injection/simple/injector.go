//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeSimpleService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository(isError bool) (*DatabaseRepository, error) {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMySQL,
		NewDatabaseRepository,
	)
	return nil, nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// func InitializeHelloService() *HelloService {
// 	wire.Build(NewSayHelloService, NewSayHelloImpl)
// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializeFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar")) //*
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "Foo", "Bar")) //*
	return nil
}

func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}


func InitializeConfiguration() *Configuration {
	wire.Build(
		NewAplication,
		wire.FieldsOf(new(*Aplication), "Configuration"),
	)
	return nil
}

func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection,
		NewFile,
	)
	return nil, nil
}
