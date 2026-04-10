using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Try to parse the first argument as an integer
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N < 1, print nothing as per requirements
if (n < 1)
{
    return;
}

// Fibonacci sequence variables
long current = 1;
long next = 1;

// The sequence starts 1, 1, 2, 3, 5...
// We print the first 1, then loop to calculate and print subsequent numbers
while (current <= n)
{
    Console.WriteLine(current);
    
    long temp = current + next;
    current = next;
    next = temp;
}