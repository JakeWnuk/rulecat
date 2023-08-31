<h1 align="center">
Rulecat
 </h1>

Rulecat (`cat` rule) performs eleven (11) unique functions:
- Creates append rules from `stdin`
    - Normal append
    - Remove append (`]`)
    - Shift then append (`}`)
    - Multibyte character support
- Creates prepend rules from `stdin`
    - Normal prepend
    - Remove prepend (`[`)
    - Shift then prepend (`{`)
    - Multibyte character support
- Creates blank lines from `stdin`
    - Commonly used when doing `-a9` attacks
- Create Cartesian product of a file and `stdin`
    - For every item in `stdin` and a rule file create a combination
    - `stdin` is placed before file content
- Creates custom rules per character from `stdin`
    - Inserts custom rule before each character
    - Create unique combinations: `@`, `!`, `/`, and others
    - Multibyte character support
- Creates insert rules from `stdin`
    - Select starting index
- Creates overwrite rules from `stdin`
    - Select starting index
- Creates toggle rules from `stdin`
    - Select offset/starting index for toggles
- Creates custom text rehashing based on an expression
    - Create a custom expression like `100xmd5(sha1(p))`
- Creates dehexed text from `stdin`
    - Only prints dehexed text
- URL and HTML encodes text from `stdin`
    - Only prints encoded text

- For more application examples: 
    - [Rulecat Examples](https://jakewnuk.com/posts/brewing-hash-cracking-resources-w-the-twin-cats/)
    - [Rulecat Usages](https://jakewnuk.com/posts/how-to-use-rulecat-to-crack-the-perfect-eggs-every-time/)
- See also [maskcat](https://github.com/JakeWnuk/maskcat/tree/main).

## Getting Started

- [Install](#install)
- [Append Rules](#append-rules)
- [Prepend Rules](#prepend-rules)
- [Blank Lines](#blank-lines)
- [Cartesian Rules](#cartesian-rules)
- [Character to Rules](#character-to-rules)
- [Insert Rules](#insert-rules)
- [Overwrite Rules](#overwrite-rules)
- [Toggle Rules](#toggle-rules)
- [Custom Expressions](#custom-expressions)
- [Dehex Text](#dehex-text)
- [URL & HTML Encode Text](#url-&-html-encode-text)
- [Emoji Text](#emoji-text)

### Install

#### Go
```
go install -v github.com/jakewnuk/rulecat/cmd/rulecat@latest
```
#### From Source
```
git clone https://github.com/JakeWnuk/rulecat && cd rulecat && go build ./cmd/rulecat && mv ./rulecat ~/go/bin/
```
```
$ cat test.tmp
This
Is A
Test123

$ cat test.tmp | rulecat
Modes for rulecat (version 1.2.0):

  append        Creates append rules from text
                Example: stdin | rulecat append
                Example: stdin | rulecat append remove
                Example: stdin | rulecat append shift

  prepend       Creates prepend rules from text
                Example: stdin | rulecat prepend
                Example: stdin | rulecat prepend remove
                Example: stdin | rulecat prepend shift

  blank         Creates blank lines from text
                Example: stdin | rulecat blank

  [RULE-FILE]   Create Cartesian product of a file and text
                Example: stdin | rulecat [FILE]

  chars         Creates custom rules per character from text
                Example: stdin | rulecat chars [RULE]

  insert        Creates insert rules from from text
                Example: stdin | rulecat insert [START-INDEX]

  overwrite     Creates overwrite rules from from text
                Example: stdin | rulecat overwrite [START-INDEX]

  toggle        Creates toggle rules from from text
                Example: stdin | rulecat toggle [START-INDEX]

  custom        Creates custom text rehashing from an expression
                Example: stdin | rulecat custom [EXPRESSION]
                Example: stdin | rulecat custom 2xmd5(sha1(p))

  dehex         Dehexes $HEX[...] input to standard out
                Example: stdin | rulecat dehex

  encode        URL and HTML encodes input and prints new output
                Example: stdin | rulecat encode
```

## Append Rules
- Creates append rules from `stdin`
- Three (3) modes: default, remove, shift
- Shift in append mode moves characters back to front
- Multibyte character support
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
- Multibyte character support
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
- Multibyte character support
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

## Custom Expressions
- Creates custom text rehashing based on a given expression
- Examples: `sha256(md5(p))` or `100xmd5(p)`
- Currently supports: `md5`, `sha1`, `sha256`, and `sha512`
```
$ cat test.tmp | rulecat custom "100xmd5(p)"
f6a11b053985c4b9ee9eb8d867fd566f
39840dd1dea35531cd02746bf84c8f6e
1123d54890652bd74f2adcf104dbd4a3
```

## Dehex Text
- Dehexes text from `stdin` to `stdout`
- Only prints dehexed text
```
$ cat test.lst
$2a$10youkinz
$HEX[2121212126233033363a68616e64]
$HEX[21212133333a32343638]

$ cat test.lst | rulecat dehex
!!!!&#036:hand
!!!33:2468
```

## URL & HTML Encode Text
- URL & HTML encodes text from `stdin` to `stdout`
- Only prints encoded text
```
$ cat test.lst
<hello World!>
Testing$!@%!\*()

$ cat test.lst | rulecat encode
%3Chello+World%21%3E
&lt;hello World!&gt;
Testing%24%21%40%25%21%2A%28%29
```

