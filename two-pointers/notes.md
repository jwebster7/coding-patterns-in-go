# Two-Pointers

- Two-pointers is an application of the search-and-prune strategy, where each step involves examining and eliminating possibilities.

## When does this pattern apply?

- This pattern is applicable to a problem iff two conditions are met:
  - Input data may be traversed in a linear manner (arrays, linked lists, etc.).
  - Limit the view / scope to a subset of the original data. The data considered must bee between or outside of the two pointers.
- This pattern is not applicable to a problem if either of these applies:
  - Input data cannot be traversed in a linear manner.
  - An exhaustive search is required (elements are not eliminated when one solution is).

## Real-world applications

### Memory management

Start and end pointers used to indicate a pool of memory.

### Product suggestions

Suggesting items to shopping carts to qualify for free shipping.

## Example problems

### Checking for palindromes

By having a pointer at the leftmost index and a pointer at the rightmost index, we can check if each element at each respective matches. Increment the leftmost, decrement the rightmost until they meet.

### Finding three values that add to a number 'x' in a sorted array

By iterating though an array and considering each value, we can calculate the difference between 'x' and the current element, then check if the array[left] + array[right] add up to the remainder. We can then increment or decrement the pointers accordingly.