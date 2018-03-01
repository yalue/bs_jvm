BS-JVM
======

About
-----

The BS ("Blinding Speed") Java Virtual Machine is a work-in-progress JVM
implementation in the Go programming language.

Usage
-----

For now, only class file disassembly works. Usage:

```bash
cd disassemble/
go build .

# The path to RandomDots.class can be replaced with a path to any valid class
# file.
./disassemble -filename ../class_file/test_data/RandomDots.class
```
