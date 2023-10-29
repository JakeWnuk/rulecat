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
Creating combo rules
```
$ echo 'this-Test123' | rulecat combo toggle insert
T5 i4-

$ echo 'this-Test123' | rulecat combo prepend append
^s ^i ^h ^t $- $1 $2 $3
```

### Creating Cartesian Products
Rulecat can be used to create the cartesian product of `stdin` and a provided
`FILE`. The content from `stdin` is placed before the `FILE` content.
```
Example: stdin | rulecat [FILE]
```

### Creating Combo Rules
Rulecat can be used to create combinations of different modes for each item
from `stdin`. 

The valid mode options for `combo` are:
    - toggle
    - prepend
    - append
    - insert

```
Example: stdin | rulecat combo [MODE-A] [MODE-B]
```
