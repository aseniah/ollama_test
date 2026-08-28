using System;

// Check if an argument was provided
if (Args.Length == 0)
{
    return;
}

// Parse the input argument N as a long to handle large numbers
if (!long.TryParse(Args[0], out long n))
{
    return;
}

// If N < 1, print nothing as per requirements
if (n < 1)
{
    return;
}

// Initialize the first two Fibonacci numbers
long a = 1;
long b = 1;

// Generate and print the Fibonacci sequence until the current number exceeds N
while (a <= n)
{
    Console.WriteLine(a);
    
    // Calculate the next number in the sequence
    long next = a + b;
    a = b;
    b = next;
}