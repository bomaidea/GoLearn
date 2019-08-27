# Arrays

### arrays.go

The type `[n]T` is an array of `n` values of type `T`.

The expression

```
var a [10]int
```

declares a variable `a` as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

### slices.go

An array has a fixed size. A slice on the other hand, is a dinamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and hig bound, separated by a colon:

```
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

```
a[1:4]
```

### slices-pointers.go

Slices are like references to array

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

### slices-literals.go

A slice literal is like an array literal without the length.

This is an array literal:

```
[3]bool{true, true, false}
```

And this creates the same array as abobe, then builds a slice that references it: 

```
[]bool{true, true, false}
```

### slices-bounds.go

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the lenght of the slice for the high bound.

For the array

```
var a [10]int
```

these slice expressions are equivalent:

```
a[0:10]
a[:10]
a[0:]
a[:]
```

### slices-len-cap.go

Slice length and capacity

A slice has both a _length_ and a _capacity_.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expression `len(s)` and `cap(s)`.

You can extend a slice's length by re`slicing it, provided it has sufficent capacity. Try changing one of the slice operation in the example program to extend it beyond its capacity and see what happens.

### nil slices

The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

### making-slices.go

Creating a slice with make

Slices can be created with the built-in `make` functions; this is how you create dynamilcally-sized arrays.

The `make` function allocates a zeroed array and returns a slice refers to that array:

```
a := make([]int, 5) // len(a)=5
```

To specify a capacity, pass a third argument to make:

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]	   // len(b)=4, cpa(b)=4
```

### slices-of-slices.go

Slices can contain any type, including other slices.

### append.go

Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in `append` function. The documentation of the built-in package describes `append`.

```
func append(s []T, vs ...T) []T
```

The first parameter `s` of `append` is a slice of type `T`, and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too smal to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

### range.go

The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index. and the second is a copy of the element at that index.

### range-continued.go

You can skip the index or value by assigning to `_`.

```
for i, _ := range pow
for _, value := pow
```

If you only want the index, you can omit the second variable.

```
for i := range pow
```

