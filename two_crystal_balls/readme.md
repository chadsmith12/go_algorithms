## Two Crystal Balls Problem

Given two crystal balls that will break if dropped from a high enough distance, determine the exact spot in which it will break in the most optimized way.

### Options:
* Linear
* Binary Search

We know that we really just have an array full of falses, until we get our first true. We also have two crystal balls. When our first crystal ball breaks, then we go jump back and walk until we find our break.

If we jump using a different unit, `Sqrt(n)`, where `n` is the length of the array. A jump in that cause would be `Sqrt(n)`, then we walk a `sqrt(n)` space to get a running time of `sqrt(n)` times.


