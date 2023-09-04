### Quick Start
Create insert rules
```
$ cat test.tmp | rulecat insert
i0T i1h i2i i3s
i0I i1s i2  i3A
i0T i1e i2s i3t i41 i52 i63
```
Create overwrite rules
```
$ cat test.tmp | rulecat overwrite
o0T o1h o2i o3s
o0I o1s o2  o3A
o0T o1e o2s o3t o41 o52 o63
```

### Creating Insert Rules
Rulecat can be used to create insert rules from `stdin`. This will convert
input into valid `Hashcat` rules and does not support multibyte text.
```
Example: stdin | rulecat insert [START-INDEX]
```

When the `insert` option is used with a valid `START-INDEX` value the starting
index of the insert rule can be changed.
```
$ cat test.tmp | rulecat insert 6
i6T i7h i8i i9s
i6I i7s i8  i9A
i6T i7e i8s i9t iA1 iB2 iC3
```

### Creating Overwrite Rules
Rulecat can be used to create overwrite rules from `stdin`. This will convert
input into valid `Hashcat` rules and does not support multibyte text.
```
Example: stdin | rulecat overwrite [START-INDEX]
```

When the `overwrite` option is used with a valid `START-INDEX` value the starting
index of the overwrite rule can be changed.
```
$ cat test.tmp | rulecat overwrite 6
o6T o7h o8i o9s
o6I o7s o8  o9A
o6T o7e o8s o9t oA1 oB2 oC3
```
