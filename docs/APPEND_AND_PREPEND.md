### Quick Start

Create append rules
```
$ cat test.tmp | rulecat append
$T $h $i $s
$I $s $  $A
$T $e $s $t $1 $2 $3
```
Create prepend rules
```
$ cat test.tmp | rulecat prepend
^s ^i ^h ^T
^A ^  ^s ^I
^3 ^2 ^1 ^t ^s ^e ^T
```
>[!TIP]
>`Append` and `Prepend` modes support multibyte text.

### Creating Append Rules
Rulecat can be used to create append rules from `stdin`. This will convert
input into valid `Hashcat` rules and supports multibyte characters.
```
Example: stdin | rulecat append
Example: stdin | rulecat append remove
Example: stdin | rulecat append shift
```

The `append` mode supports three unique modes:
- Normal append
- Remove append (`]`)
- Shift then append (`}`)

When the `append` option is used with the `remove` option characters are
removed equal to the string length.
```
] ] ] ] $T $h $i $s
] ] ] ] $I $s $  $A
] ] ] ] ] ] ] $T $e $s $t $1 $2 $3
```

When the `append` option is used with the `shift` option characters are shifted
(rotated right) equal to the string length.
```
} } } } $T $h $i $s
} } } } $I $s $  $A
} } } } } } } $T $e $s $t $1 $2 $3
```

### Creating Prepend Rules
Rulecat can be used to create prepend rules from `stdin`. This will convert
input into valid `Hashcat` rules and supports multibyte characters.
```
Example: stdin | rulecat prepend
Example: stdin | rulecat prepend remove
Example: stdin | rulecat prepend shift
```

The `prepend` mode supports three unique modes:
- Normal prepend
- Remove prepend (`[`)
- Shift then prepend (`{`)

When the `prepend` option is used with the `remove` option characters are
removed equal to the string length.
```
[ [ [ [ ^s ^i ^h ^T
[ [ [ [ ^A ^  ^s ^I
[ [ [ [ [ [ [ ^3 ^2 ^1 ^t ^s ^e ^T
```

When the `prepend` option is used with the `shift` option characters are shifted
(rotated left) equal to the string length.
```
{ { { { ^s ^i ^h ^T
{ { { { ^A ^  ^s ^I
{ { { { { { { ^3 ^2 ^1 ^t ^s ^e ^T
```
