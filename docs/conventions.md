# Conventions

## Language differences

Java is an OOP language, which has concepts such as overloading and inheritance. This is not
the case for Golang, which relies on composition and duck typing. While porting, we aim to
maintain as close resemblance to Java code as possible; however, it will not be always possible.
The rules set in these document shall help a developer who is coming from JTS documentation or 
developed in Java.

## Constructors

Since Golang does not provide either constructor or overloading, we use `New[...]` functions,
which return the resulting struct. Rules for naming these functions:

1. Function with the most primitive type shall be the one with the shortest name, e.g. `NewCoordinate(x,y float64)`
1. Function, which creates an empty or default constructor shall contain the keyword in the name, e.g.
   `NewDefaultPrecisionModel`
1. Function, which uses complex data type shall be named after that type with the `from` keyword, e.g.
   `NewEnvelopeFromCoordinates`