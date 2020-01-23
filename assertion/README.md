# Assertion

## type-assertion.go

Type assertion

A _type assertion_ provides access to an interface value's underlying concrete value.

```
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T` and assings the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a pain.

To _test_ wheather an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports wheter the assertion succeded.

```
t, ok := i.(T)
```

If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.

If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

Not the similarity between this syntax and that of reading from a map.
