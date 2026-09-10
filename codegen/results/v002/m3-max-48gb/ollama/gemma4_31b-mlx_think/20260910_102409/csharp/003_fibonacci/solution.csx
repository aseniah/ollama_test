using System;

// Check if an argument was provided
if (Args.Length == 0)
{
    return;
}

// Parse the input N. Using long to handle larger potential input values.
if (!long.TryParse(Args[0], out long n))
{
    return;
}

// If N < 1, print nothing.
if (n < 1)
{
    return;
}

// Initialize the first two numbers of the Fibonacci sequence
long current = 1;
long next = 1;

// Continue printing Fibonacci numbers as long as the current number does not exceed N
while (current <= n)
{
    Console.WriteLine(current);
    
    // Calculate the next Fibonacci number
    long temp = current + next;
    current = next;
    next = temp;
}