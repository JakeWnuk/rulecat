<h1 align="center">
Rulecat
 </h1>

Rulecat (`cat` rule) performs eight (8) unique functions:
- Creates append rules from `stdin`
    - Normal append
    - Remove append (`]`)
    - Shift then append (`}`)
- Creates prepend rules from `stdin`
    - Normal prepend
    - Remove prepend (`[`)
    - Shift then prepend (`{`)
- Creates blank lines from `stdin`
    - Commonly used when doing `-a9` attacks
- Create Cartesian product of a file and `stdin`
    - For every item in `stdin` and a rule file create a combination
    - `stdin` is placed before file content
- Creates insert rules from `stdin`
    - Select starting index
- Creates overwrite rules from `stdin`
    - Select starting index
- Creates toggle rules from `stdin`
    - Select offset/starting index for toggles
- Creates custom rules per character from `stdin`
    - Inserts custom rule before each character
    - Create unique combinations: `@`, `!`, `/`, and others

- For more application examples: [blog post](https://jakewnuk.com/posts/brewing-hash-cracking-resources-w-the-twin-cats/)
- See also [maskcat](https://github.com/JakeWnuk/maskcat/tree/main).

## Getting Started

- [Install](#install)
- [Append Rules](#Append-Rules)
- [Prepend Rules](#Prepend-Rules)
- [Blank Lines](#Blank-Lines)
- [Cartesian Rules](#Cartesian-Rules)
- [Character to Rules](#Character-to-Rules)
- [Insert Rules](#Insert-Rules)
- [Overwrite Rules](#Overwrite-Rules)
- [Toggle Rules](#Toggle-Rules)

### Install

```
go install -v github.com/jakewnuk/rulecat@latest
```

```
$ cat test.tmp
This
Is A
Test123

$ cat test.tmp | rulecat
OPTIONS: append prepend insert overwrite toggle
EXAMPLE: stdin | rulecat append
EXAMPLE: stdin | rulecat prepend
EXAMPLE: stdin | rulecat append remove
EXAMPLE: stdin | rulecat prepend remove
EXAMPLE: stdin | rulecat append shift
EXAMPLE: stdin | rulecat prepend shift
EXAMPLE: stdin | rulecat blank
EXAMPLE: stdin | rulecat <RULE-FILE>
EXAMPLE: stdin | rulecat chars <RULE>
EXAMPLE: stdin | rulecat insert <START-INDEX>
EXAMPLE: stdin | rulecat overwrite <START-INDEX>
EXAMPLE: stdin | rulecat toggle <START-INDEX>
```

## Append Rules
- Creates append rules from `stdin`
- Three (3) modes: default, remove, shift
- Shift in append mode moves characters back to front
```
$ cat test.tmp | rulecat append
$T $h $i $s
$I $s $  $A
$T $e $s $t $1 $2 $3
```
```
$ cat test.tmp | rulecat append remove
] ] ] ] $T $h $i $s
] ] ] ] $I $s $  $A
] ] ] ] ] ] ] $T $e $s $t $1 $2 $3
```
```
$ cat test.tmp | rulecat append shift
} } } } $T $h $i $s
} } } } $I $s $  $A
} } } } } } } $T $e $s $t $1 $2 $3
```

## Prepend Rules
- Creates prepend rules from `stdin`
- Three (3) modes: default, remove, shift
- Shift in prepend mode moves characters front to back
```
$ cat test.tmp | rulecat prepend
^s ^i ^h ^T
^A ^  ^s ^I
^3 ^2 ^1 ^t ^s ^e ^T
```
```
$ cat test.tmp | rulecat prepend remove
[ [ [ [ ^s ^i ^h ^T
[ [ [ [ ^A ^  ^s ^I
[ [ [ [ [ [ [ ^3 ^2 ^1 ^t ^s ^e ^T
```
```
$ cat test.tmp | rulecat shift
{ { { { ^s ^i ^h ^T
{ { { { ^A ^  ^s ^I
{ { { { { { { ^3 ^2 ^1 ^t ^s ^e ^T
```

## Blank Lines
- Creates blank lines from `stdin`
- Commonly used when doing `-a9` attacks
```
$ cat test.tmp | rulecat blank



```

## Cartesian Rules
- For every item in `stdin` and a rule file create a combination
- `stdin` is placed before file content
```
$ cat test.rule
u
$1 $2 $3
```
```
$ cat test.tmp | rulecat append | rulecat test.rule
$T $h $i $s u
$T $h $i $s $1 $2 $3
$I $s $  $A u
$I $s $  $A $1 $2 $3
$T $e $s $t $1 $2 $3 u
$T $e $s $t $1 $2 $3 $1 $2 $3
```

## Character to Rules
- Inserts custom rule before each character
- Create unique combinations: `@`, `!`, `/`, and others
```
$ cat test.tmp | rulecat chars @
@T @h @i @s
@I @s @  @A
@T @e @s @t @1 @2 @3
```

## Insert Rules
- Creates insert rules from `stdin`
- Accepts index value from where to start the insert
```
$ cat test.tmp | rulecat insert
i0T i1h i2i i3s
i0I i1s i2  i3A
i0T i1e i2s i3t i41 i52 i63
```
```
$ cat test.tmp | rulecat insert 6
i6T i7h i8i i9s
i6I i7s i8  i9A
i6T i7e i8s i9t iA1 iB2 iC3
```

## Overwrite Rules
- Creates overwrite rules from `stdin`
- Accepts index value from where to start the overwriting
```
$ cat test.tmp | rulecat overwrite
o0T o1h o2i o3s
o0I o1s o2  o3A
o0T o1e o2s o3t o41 o52 o63
```
```
$ cat test.tmp | rulecat overwrite 6
o6T o7h o8i o9s
o6I o7s o8  o9A
o6T o7e o8s o9t oA1 oB2 oC3
```

## Toggle Rules
- Creates toggle rules from `stdin`
- Accepts index value from where to start the toggle
```
$ cat test.tmp | rulecat toggle
T0
T0 T3
T0
```
```
$ cat test.tmp | rulecat toggle 3
T3
T3 T6
T3
```
