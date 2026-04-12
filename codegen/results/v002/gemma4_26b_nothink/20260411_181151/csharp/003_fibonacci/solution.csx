using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Try to parse the argument as an integer
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N < 1, print nothing as per requirements
if (n < 1)
{
    return;
}

// Fibonacci sequence starting with 1, 1, 2, 3, 5...
// We use long to prevent overflow during calculation, 
// although N is limited to int range.
long a = 1;
long b = 1;

// Handle the first number in the sequence
Console.WriteLine(a);

// If N is 1, we only print the first 1. 
// However, the sequence 1, 1, 2... implies the second 1 is also part of it.
// Since 1 <= N, we check if the next number (also 1) is <= N.
if (b <= n)
{
    Console.WriteLine(b);
}
else
{
    return;
}

// Calculate subsequent numbers
while (true)
{
    long next = a + b;
    
    // If the next Fibonacci number exceeds N, we stop
    if (next > n)
    {
        break;
    }
    
    Console.WriteLine(next);
    
    // Update pointers for the next iteration
    a = b;
    b = next;
}