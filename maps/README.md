# maps

### maps.go

A map maps keys to values.

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

The `make` function returns a map of the given type, initalized and ready for use.

### literals.go

Map literals

Map literals are like struct literals, but the keys are required.

### literals-continued.go

Map literals continued

If the top-level type is just a type name, you can omit it from the elements of the literal.

### mutating.go

Mutating maps

Inser or update an element in map `m`:

```
m[key] = elem
```

Retrive an element:

```
elem = m[key]
```

Delete an element:

```
delete(m, key)
```

Test that a key is present with a two-value assignement:

```
elem, ok = m[key]
```

If `key` is in `m`, ok is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

__Note:__ If `elem` or `ok` have not yet been declared you could use a short declaration form:

```
elem, ok := m[key]
```
