# algorithms-notes-for-professionals
Algorithm: Notes for professionals

# Algorithm

An algorithmic problem is specified by describing the complete set of instances it must work on and of its output
after running on one of these instances. This distinction, between a problem and an instance of a problem, is
fundamental. The algorithmic problem known as sorting is defined as follows: [Skiena:2008:ADM:1410219]

**Problem:** Sorting

**Input:** A sequence of n keys, a_1, a_2, ..., a_n.

**Output:** The reordering of the input sequence such that a'_1 <= a'_2 <= ... <= a'_{n-1} <= a'_n

An instance of sorting might be an array of strings, such as { Haskell, Emacs } or a sequence of numbers such as
{ 154, 245, 1337 }.

# Fizz buzz

You may have seen Fizz Buzz written as Fizz Buzz, FizzBuzz, or Fizz-Buzz; they're all referring to the same thing. That
"thing" is the main topic of discussion today. First, what is FizzBuzz?

Fizz and Buzz refer to any number that's a multiple of 3 and 5 respectively. In other words, if a number is divisible
by 3, it is substituted with fizz; if a number is divisible by 5, it is substituted with buzz. If a number is simultaneously
a multiple of 3 AND 5, the number is replaced with "fizz buzz." In essence, it emulates the famous children game
"fizz buzz".

To work on this problem, open up Xcode to create a new playground and initialize an array like below:
// for example
`let number = [1,2,3,4,5]`
// here 3 is fizz and 5 is buzz
To find all the fizz and buzz, we must iterate through the array and check which numbers are fizz and which are
buzz. To do this, create a for loop to iterate through the array we have initialised:
`
for num in number {
    // Body and calculation goes here
}
`

Here is the solution:
`
for num in number {
    // num % 3 == 0 && num % 5 == 0
    // This way is most efficient.
    if num % 15 == 0 {
        print("\(num) fizz buzz")
    } else if num % 3 == 0 {
        print("\(num) fizz")
    } else if num % 5 == 0 {
        print("\(num) buzz")
    } else {
        print(num)
    }
}
`
