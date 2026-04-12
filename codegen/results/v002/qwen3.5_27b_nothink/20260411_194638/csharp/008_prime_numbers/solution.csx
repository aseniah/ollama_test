#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.Linq;

// Check if we have exactly one argument
if (Args.Count != 1)
{
    return;
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int N))
{
    return;
}

// If N is less than 2, there are no primes
if (N < 2)
{
    return;
}

// Sieve of Eratosthenes to find all primes up to N
bool[] isPrime = new bool[N + 1];
for (int i = 2; i <= N; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= N; p++)
{
    if (isPrime[p])
    {
        for (int multiple = p * p; multiple <= N; multiple += p)
        {
            isPrime[multiple] = false;
        }
    }
}

// Output all prime numbers
for (int i = 2; i <= N; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}