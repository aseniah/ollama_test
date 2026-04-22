using System;
using System.Collections.Generic;

if (Args.Count < 1)
{
    Console.WriteLine("Usage: prime-sieve <N>");
    Environment.Exit(1);
}

int n = int.Parse(Args[0]);

if (n < 2)
{
    // No primes up to 1 or less
    return;
}

// Sieve of Eratosthenes
bool[] isComposite = new bool[n + 1];
isComposite[0] = true;
isComposite[1] = true;

for (int i = 2; i * i <= n; i++)
{
    if (!isComposite[i])
    {
        for (int j = i * i; j <= n; j += i)
        {
            isComposite[j] = true;
        }
    }
}

for (int i = 2; i <= n; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}