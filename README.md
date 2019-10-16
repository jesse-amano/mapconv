# mapconv
Go package for flattening a structured value down to a map of paths to values.

Arrays, maps, and slices recursively produce map keys until terminal values are found.

Terminal values are:

- bool
- int (and rune, int8, int16, ...)
- uint (and byte, uint8, uint16, ...)
- float64 (and float32)
- string
