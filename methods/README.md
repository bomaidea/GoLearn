# GO methods

### method.go

Go does not have classes. However, you can define methods on types.

A method is a function with a special reciver argumet.

The reciver appears in its own argumennt list between the `func` keyword and the method name.

In this example, the `Abs` method has reciver of type `Vertex` named `v`.

### funcs.go

Methods are functions

Remember: a method is just a funciton with a reciver argument.

Here's `Abs` written as a regular funcion with no chage in functionally.

### continued.go

Methods continued

You can declare a method on non-struct types, too.

In this example we see a numeric type `MyFloat` with an `Abs` method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declasre a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

### pointer.go

Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to wich the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of Scale funciton on line 16 nd observe how the program's behavoir changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavoir as for any other function argument.) The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function. 

### pointers-explained

Pointers and functions

Here we see the `Abs` and `Scale` methods rewritten as funcitons.

Again, try removign the `*` from line 16. Can you see why the behavoir changes?  
What else did you need to change for the example to compile?

### indirection.go

Methods and pointer indirection

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

```
var v Vertex
ScaleFunc(v, 5) // Compile error!
ScaleFunc(&v, 5) // OK
```

while mehtods with ointer receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
v.Scale(5) // OK
p := &v
p.Scale(10) // OK
```

For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

### indirection-values.go

Methods and pointer indirection

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with values receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

### pointer-receiver.go

Chosing a value or pointer receiver

These are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

