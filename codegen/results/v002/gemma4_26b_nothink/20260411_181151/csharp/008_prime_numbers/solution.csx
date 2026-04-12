using System;
using System.Collections.Generic;

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

// Prime numbers are greater than 1. 
// If n < 2, there are no primes to print.
if (n < 2)
{
    return;
}

// Use the Sieve of Eratosthenes algorithm for efficiency
// Create a boolean array "isPrime[0..n]" and initialize
// all entries it as true. A value in isPrime[i] will
// finally be false if i is Not a prime, else true.
bool[] isPrime = new bool[n + 1];
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p greater than or equal to the square of it
        // numbers which are multiple of p and are smaller than p^2 are already marked.
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Print all prime numbers
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}