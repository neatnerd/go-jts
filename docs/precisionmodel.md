# Precision Model
## Introduction

There are two main purposes for the precision model in the JTS framework:
 - Rounding up value, which effectively restricts range of values for (x,y)
 - Providing indication for serialization. It is necessary to know how many digits in
a mantissa will be stored

## Translating types

JTS uses Java's `double` value for storing coordinate values. The closest type for s Go is
`float64`. According to official documentation these types must be identical:

**Java**:

[data type](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/datatypes.html)
```
The double data type is a double-precision 64-bit IEEE 754 floating point
```

**Go**:
[data type](https://golang.org/pkg/builtin/#float64)
```
float64 is the set of all IEEE-754 64-bit floating-point numbers.
```

## Precision Truncating

For fixed precision operations