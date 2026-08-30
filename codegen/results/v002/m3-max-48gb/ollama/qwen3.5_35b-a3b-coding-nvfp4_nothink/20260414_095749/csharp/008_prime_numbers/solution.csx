#r "System.Core"
#r "System.Linq"

using System;
using System.Collections.Generic;
using System.Linq;

// Check if command line arguments are provided
if (Args.Count == 0)
{
    // No argument provided, exit silently
    return;
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int n))
{
    // Invalid integer provided, exit silently
    return;
}

// Handle edge cases where no primes exist
if (n < 2)
{
    return;
}

// Generate primes using Sieve of Eratosthenes
var limit = n + 1;
var isPrime = new bool[limit];
Array.Fill(isPrime, true);
isPrime[0] = false;
isPrime[1] = false;

for (int i = 2; i * i < limit; i++)
{
    if (isPrime[i])
    {
        for (int j = i * i; j < limit; j += i)
        {
            isPrime[j] = false;
        }
    }
}

// Print primes up to and including N
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}