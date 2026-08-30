using System;
using System.Collections.Generic;
using System.Linq;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the input N
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N is less than 2, there are no prime numbers to print
if (n < 2)
{
    return;
}

// Use the Sieve of Eratosthenes algorithm to find all primes up to N
// This is efficient for large values of N.
bool[] isPrime = new bool[n + 1];

// Initialize all entries as true. 
// A value in isPrime[i] will finally be false if i is Not a prime, else true.
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

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

// Print all prime numbers
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}