using System;
using System.Linq;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the input argument N
if (!int.TryParse(Args[0], out int n) || n < 1)
{
    // If N < 1, print nothing as per requirements
    return;
}

// Fibonacci sequence variables
long current = 1;
long next = 1;

// We start with the sequence 1, 1, 2, 3, 5...
// The logic handles the first two 1s explicitly or via a loop structure.
// Since the sequence starts 1, 1..., we can iterate while the current number <= N.

while (current <= n)
{
    // Print the current Fibonacci number
    Console.WriteLine(current);

    // Calculate the next number in the sequence
    long temp = current + next;
    current = next;
    next = temp;
}