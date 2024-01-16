# Daily Coding Problem: Problem #1602 [Medium]

## Problem Statement

This problem was asked by LinkedIn.

You are given a string consisting of the letters `x` and `y`,
such as `xyxxxyxyy`.
In addition,
you have an operation called `flip`,
which changes a single `x` to `y` or vice versa.

Determine how many times you would need to apply this operation to ensure
that all `x`'s come before all `y`'s.
In the preceding example,
it suffices to flip the second and sixth characters,
so you should return 2.

## Analysis

First, it's a little tricky to examine a string and
say that all `x`'s come before all `y`'s.
If the problem guaranteed that at least 1 `x` and 1 `y`
appeared in each input string,
you could ignore the cases of all-'x' or all-'y' strings.
There's also the case of a zero-length input string,
which I'm going to say does indeed have all zero  `x`'s
before all zero `y`'s

Second, the problem statement doesn't say to return the
**minimum** number of flips,
but rather how many to **ensure** the desired condition.
The example does imply the minimum number of flips is desired,
but who really can say?

Third does an all-single-character string, `xxxx` say,
require 0 (zero) flips, or does it require 1 flip to get `xxxy`?
Another hidden imperfection in the problem statement.

Even the example input string, `xyxxxyxyy`,
can be made correct 2 different ways with 2 flips:

* `xxxxxyyyy`, flip 2nd and 7th `x` characters
* `xxxxxxxyy`, flip 2nd character `x`, and 6th character `y`

I chose to do a Go-specific algorithm:

1. Specify a data type that can contain the number of flips,
and the string that has all `x`'s before all `y`'s.
2. Create a channel that can carry pointers to the data type.
3. Write a function that checks if its input string meets the
criteria, and if so, creates an instance of the data type,
fills it out with the current string and flip count,
then puts it on the channel and returns.
Otherwise, recurse after flipping.
4. The `main()` function starts the function from (3) in a goroutine,
then read strings and flip counts off the channel until it's closed.

The function of step 3 that I wrote has a signature like this:

```
func fliprecurse(ch chan *xBeforeY, str []rune, ln int, idx int, flipCount int)
```

`ln` is the string length, pre-calculated to avoid repetitious invocations of `len()`.
`idx` is the index in the string `str` the function should work on.

Unfortunately, there's 3 places to return from recursion:

1. If `str` meets the criteria of all `x`'s before all `y`'s
2. If `idx` is past the last character in `str`
3. After recursing on `idx+1`.
I think you have to recurse twice, once without flipping character at `idx`,
once after a flip at `idx`.

## Interview Analysis

Although this problem appears to be concise and correctly phrased,
it has some hidden ambiguity.
If the interviewer is aware of this,
and wants candidates to ask clarifying questions, that's great.
The interviewer has the burden of being prepared for
questions.
If the problem statement were something like formal requirements,
it would be poorly done.

The candidate almost has to ask clarifying questions,
some of which aren't going to present themselves until it's time to
write the code for "all `x`'s before all `y`'s".

Does this rate a "[Medium]"?
I think so.
There's not much in the way  of data structures,
but there's tricky coding in assessing whether or not a string
meets the criteria.
Any way of doing the work requires some fiddly algorithm programming.
My way required 3 places to return from recursion.
