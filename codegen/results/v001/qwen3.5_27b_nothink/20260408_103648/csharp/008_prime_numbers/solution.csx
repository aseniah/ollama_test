#r "nuget: System.Text.Json"

using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Validate input
if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: dotnet-script prime.csx <integer N>");
    Environment.Exit(1);
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Error: Argument must be a valid integer.");
    Environment.Exit(1);
}

// Handle edge cases
if (n < 2)
{
    // No primes below 2, print nothing
    return;
}

// Sieve of Eratosthenes
// Create a boolean array to mark primes. 
// index i is true if i is prime, false if composite.
// We need size n + 1 to include n.
var isPrime = new bool[n + 1];

// Initialize all entries to true
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Output primes
for (int p = 2; p <= n; p++)
{
    if (isPrime[p])
    {
        Console.WriteLine(p);
    }
}