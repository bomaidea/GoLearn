# Interfaces

## interface.go

An _interface type_ is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

**Note:** There is an error in the example code on line 22. `Vertex` (the value type) doesn't implement `Abser` because the `Abs` mehtod is defined only on `*Vertex` (the pointer type).

## are-satisfied-implicitly.go

Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

## values.go

Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```
  (value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

## values-with-nil.go

Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in go it is common to write methods that gracefully handle being called with a nil receiver (as with the method `M` in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

## nil-interface-values.go

Nil interface values

A nil interface value holds niether value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which _concrete_ method to call.

## empty-interface.go

The empty interface

The interface type that specifies zero methods is known as the _empty interface_:

```
  interface{}
```

An empty interface may hohld values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unkonwn type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.


