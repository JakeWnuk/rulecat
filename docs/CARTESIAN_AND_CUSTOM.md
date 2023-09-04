### Quick Start
Creating cartesian rules
```
$ cat test.rule
u
$1 $2 $3

$ cat test.tmp | rulecat append | rulecat test.rule
$T $h $i $s u
$T $h $i $s $1 $2 $3
$I $s $  $A u
$I $s $  $A $1 $2 $3
$T $e $s $t $1 $2 $3 u
$T $e $s $t $1 $2 $3 $1 $2 $3
```
Creating custom hashing expressions
```
$ cat test.tmp | rulecat custom "100xmd5(p)"
f6a11b053985c4b9ee9eb8d867fd566f
39840dd1dea35531cd02746bf84c8f6e
1123d54890652bd74f2adcf104dbd4a3
```

### Creating Cartesian Products
Rulecat can be used to create the cartesian product of `stdin` and a provided
`FILE`. The content from `stdin` is placed before the `FILE` content.
```
Example: stdin | rulecat [FILE]
```

### Creating Custom Hashing Expressions
Rulecat can be used to create custom hashing expressions of `stdin` based on
a provided expression. The content from `stdin` is rehashed based on the
provided expression.
```
Example: stdin | rulecat custom [EXPRESSION]
Example: stdin | rulecat custom 2xmd5(sha1(p))
```

The `EXPRESSION` will use `(...)` to identify nested hashing algorithms and
uses the `NxALGO(...)` syntax to to rehash input by the provided `ALGO` `N`
number of times.

The `custom` option support the following algorithms:
- `md5`
- `sha1`
- `sha256`
- `sha512`
