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

// If N is less than 2, there are no prime numbers to print
if (n < 2)
{
    return;
}

// Sieve of Eratosthenes algorithm to find all primes up to N
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
        // Update all multiples of p greater than or equal to the square of it
        // numbers which are smaller than p^2 are already marked
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