# Fast & Slow Pointers

- The fast and slow pointers pattern utilizes two pointers to traverse an iterable data structure.
- The primary difference between this pattern and two pointers is that two pointers is concerned with data values, while fast & slow is concerned with determining traits or finding patterns in a data structure.
- Another difference is that these two pointers will almost always have two different, adjustable speeds.
- Typically useful for bidirectional data structures (linked lists, arrays).
- Commonly used to detect cycles in a data structure. If the slow and fast pointer ever meet and are not at the head (or 0th index of array), then a cycle in that data structure has been detected.

## When does this pattern apply?

- A problem may be solved using this pattern iff either of the following conditions are met:
  - The problem requires identifying:
    - The first kth% elements within a linked list.
    - The element residing at the k-way point of a linked list (middle, start of second quartile, etc.).
  - The solution requires detecting a cycle in a linked list.
  - The solution requires detecting a cycle in a sequence.
- A problem may not (generally) be solved using this pattern if the following conditions are met:
  - The input data mayy not be traversed in a linear fashion. If it not a data structure that supports linear traversal, this pattern may not work.
  - The problem may be solved with two pointers traversing at the same pace (which would be two pointers, not fast & slow!).

## Example problems

### Finding the middle node of a linked list

To find the middle node of a linked list, a slow and fast pointer can be set to the head. The fast pointer should be 2x as fast as the slow pointer. Once the fast pointer hits the tail of the list, the slow pointer should be on the middle of the linked list.

### Determining if an integer is a happy number

A single step is represented by the sum of squares of the individual digits that comprise a number. For 27, it would be: (2^2)+(7^2), which results in 18. The next step would be: (1^2)+(8^2) which results in 65.

Example:
    27 -> 18 -> 65 -> ...

The slow pointer moves 1 step a time, while the fast pointer moves 2 steps a time. This process repeats until either:

- The fast pointer points to a value of 1, which means the origin is a happy number.
- The slow and fast pointer meet, which means that some cycles exists which means the origin is not a happy number.

## Real-world applications

### Symlink verification

Fast and slow pointers can be used in a symlink verification utility in an operating system. A symbolic link, or symlink, is simply a shortcut to another file. This is a file that points to another file.  

Symlinks can easily create loops or cycles where shortcuts point to each other. To avoid symlink cycles, a symlink verification utility may be used. Fast & slow pointers can detect a loop in symlinks by moving along the connected files at differing speeds.

### Compiling object oriented programs

Typically speaking, programs are not contained within a single file. Large applications can have modules and components divided into different files. Dependencies must be defined ahead of time so that the compiler can compile certain modules first. There may, however, be instances where cyclic dependencies (import cycles) are introduced. The fast & slow pattern may be used to detect and remove these cycles to allow successful compilation.  
