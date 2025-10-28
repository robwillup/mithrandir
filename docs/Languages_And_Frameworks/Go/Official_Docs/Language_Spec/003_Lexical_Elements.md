# Lexical Elements

## Comments

> Comments serve as program documentation. There are two forms:
>
> 1. _Line comments_ start with the character sequence `//` and stop at the end of the line.
> 2. _General comments_ start with the character sequence `/*` and stop with the first subsequent character sequence
> `*/`

A comment cannot start inside a `rune` or `literal string`, or inside a comment.

## Tokens

> Tokens form the vocabulary of the Go language. There are four classes:
>
> - identifiers, keywords, operators and punctuation, and literals.

## Semicolons

> The formal syntax uses semicolons `";"` as terminators in a number of productions. Go programs may omit most of these
> semicolons.

## Identifiers

> Identifiers name program entities such as variables and types. The first character in an identifier must be a
> character.

## Keyworkds

> The following keywords are reserved and may not be used as identifiers
>
> - break
> - case
> - chan
> - const
> - continue
> - default
> - defer
> - else
> - fallthrough
> - for
> - func
> - go
> - goto
> - if 
> - import
> - interface
> - map
> - package
> - range
> - return
> - select
> - struct
> - switch
> - type
> - var

## Operators and Punctuation

> +    &     +=    &=     &&    ==    !=    (    )
> -    |     -=    |=     ||    <     <=    [    ]
> *    ^     *=    ^=     <-    >     >=    {    }
> /    <<    /=    <<=    ++    =     :=    ,    ;
> %    >>    %=    >>=    --    !     ...   .    :
> &^          &^=          ~

## Integer literals

> An integer literal is a sequence of digits representing an integer constant.
