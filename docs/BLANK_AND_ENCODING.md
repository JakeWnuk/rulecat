### Quick Start
Create blank lines
```
$ cat test.tmp | rulecat blank



```
Creating encoded text
```
$ cat test.lst
<hello World!>
Testing$!@%!\*()
its a 😊 day

$ cat test.lst | rulecat encode
%3Chello+World%21%3E
&lt;hello World!&gt;
Testing%24%21%40%25%21%5C%2A%28%29
its+a+%F0%9F%98%8A+day
its a \u1f60a day
```

### Creating Blank Lines
Rulecat can be used to create blank lines from `stdin`. This will create an
equal number of blank lines for every item in `stdin`. This is commonly used in
`-a9` attacks.
```
Example: stdin | rulecat blank
```

### Creating Encoded Text
Rulecat can be used to create URL, HTML, and unicode escaped text from `stdin`.
This will convert input into their escaped equivalents and will only print text
that has been transformed.
```
Example: stdin | rulecat encode
```
