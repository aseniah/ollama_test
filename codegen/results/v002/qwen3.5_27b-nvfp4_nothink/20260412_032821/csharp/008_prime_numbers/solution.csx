#r "System.Console"

using System;
using System.Collections.Generic;
using System.Linq;

// Parse the argument N
if (Args.Count != 1 || !int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Usage: dotnet-script primes.csx <integer N>");
    Environment.Exit(1);
}

// If n is less than 2, no primes exist
if (n < 2)
{
    // Print nothing as requested
    return;
}

// Sieve of Eratosthenes to find primes up to N
bool[] isPrime = new bool[n + 1];
// Initialize all numbers as prime
Array.Fill(isPrime, true);
isPrime[0] = false;
isPrime[1] = false;

for (int p = 2; p * p <= n; p++)
{
    if (isPrime[p])
    {
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Collect and print primes
var primes = new List<int>();
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        primes.Add(i);
    }
}

foreach (var prime in primes)
{
    Console.WriteLine(prime);
}