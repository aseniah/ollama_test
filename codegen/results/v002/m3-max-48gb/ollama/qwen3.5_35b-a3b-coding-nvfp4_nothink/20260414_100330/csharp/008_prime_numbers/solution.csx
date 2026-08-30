#r "System.Console.dll"

using System;
using System.Collections.Generic;

// Get the argument N
if (Args.Count < 1)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

// Handle edge cases where N < 2
if (n < 2)
{
    return;
}

// Sieve of Eratosthenes to find primes up to n
bool[] isPrime = new bool[n + 1];
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int i = 2; i * i <= n; i++)
{
    if (isPrime[i])
    {
        for (int j = i * i; j <= n; j += i)
        {
            isPrime[j] = false;
        }
    }
}

// Print primes
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}