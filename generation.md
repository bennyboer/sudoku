---
title: Generation
---
## Difficulty
We calculate our difficulty value by the ammount of strategies we have to use to solve the sudoku. Using the backtracking fallback means usually a difficulty of 0.9 or higher. This strategy of calculating the difficulty may be a little flawed at the moment, as we haven't implemented all strategies yet.

## Generation

We have implemented two different generation strategies.

The `simple` generator deletes fields corresponding to the difficulty value passed to the generator. Higher difficulties will result in more deleted fields, however, at least 17 are left filled. This method does not directly correspond the passed value to the final difficulty of the generated sudoku.

The `backtracking` generator is our default generator. It tries to match the generated difficulty as closely as possible to the passed value via backtracking. This may result in longer generation times.
