# GO structs

### structs.go

A `struct` is a collection of fields.

### fields.go

Struct fields are accessed using a dot

### pointers.go

Struct fields can be accessed throught a struct pointer.

To access the filed `X` of a struct when we have the struct pointer `p` we could write
`(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

### literals.go

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrilevant.)

The spiacial prefix `&` returns a pointer to the struct value.
