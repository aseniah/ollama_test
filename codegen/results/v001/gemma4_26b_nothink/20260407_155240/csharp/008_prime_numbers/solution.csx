using System;
using System.Collections.Generic;
using System.Linq;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the input argument N
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N is less than 2, there are no prime numbers
if (n < 2)
{
    return;
}

// Implement Sieve of Eratosthenes to find all primes up to N
// We use a bool array where true means the index is prime
bool[] isPrime = new bool[n + 1];

// Initialize all entries as true from index 2 to n
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

// Sieve algorithm logic
for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p starting from p*p
        // (smaller multiples will have already been marked)
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Output all indices that remained true
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}