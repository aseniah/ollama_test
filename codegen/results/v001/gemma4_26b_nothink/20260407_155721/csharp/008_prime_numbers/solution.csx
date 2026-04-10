using System;
using System.Collections.Generic;

// Check if argument is provided
if (Args.Count == 0)
{
    return;
}

// Attempt to parse the integer N
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N is less than 2, there are no primes to print
if (n < 2)
{
    return;
}

// Sieve of Eratosthenes algorithm to find primes up to N
// We use a bool array where index represents the number
bool[] isPrime = new bool[n + 1];

// Initialize all entries as true from 2 to n
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

// Implementation of the Sieve
for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p starting from p*p
        // Multiples smaller than p*p have already been marked
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