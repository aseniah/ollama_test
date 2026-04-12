using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the input argument N. 
// We use long to handle potential large values for N and Fibonacci calculations.
if (!long.TryParse(Args[0], out long n) || n < 1)
{
    // If N < 1, the requirements state to print nothing.
    return;
}

// The Fibonacci sequence defined by the user starts: 1, 1, 2, 3, 5, ...
// We initialize the first two numbers of the sequence.
long f1 = 1;
long f2 = 1;

// The first number in the sequence is 1. 
// Since N >= 1, this will always be printed.
Console.WriteLine(f1);

// Check if the second number in the sequence (also 1) is within the limit N.
if (f2 <= n)
{
    Console.WriteLine(f2);
    
    // Calculate the next number in the sequence.
    long next = f1 + f2;
    
    // Continue generating and printing Fibonacci numbers as long as they do not exceed N.
    while (next <= n)
    {
        Console.WriteLine(next);
        
        // Update the previous two numbers to shift the window forward.
        f1 = f2;
        f2 = next;
        
        // Calculate the next number in the sequence.
        next = f1 + f2;
    }
}