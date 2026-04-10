using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the argument to an integer
if (!int.TryParse(Args[0], out int n) || n < 1)
{
    // If N < 1 or not a valid integer, print nothing as per requirements
    return;
}

// Initialize the first two numbers of the Fibonacci sequence
long a = 1;
long b = 1;

// Handle the case for N >= 1 specifically for the first number
// Since the sequence starts 1, 1, 2... we print the first '1'
Console.WriteLine(a);

// If N is at least 1, we also check if we can print the second '1'
if (n >= 1)
{
    // We only print the second '1' if it doesn't exceed N
    // Since we already printed the first '1', we check if the next number (also 1) is <= N
    // This is always true if N >= 1.
    Console.WriteLine(b);
}

// Calculate subsequent Fibonacci numbers
while (true)
{
    long next = a + b;
    
    // If the next number exceeds N, we stop
    if (next > n)
    {
        break;
    }

    Console.WriteLine(next);

    // Update values for the next iteration
    a = b;
    b = next;
}