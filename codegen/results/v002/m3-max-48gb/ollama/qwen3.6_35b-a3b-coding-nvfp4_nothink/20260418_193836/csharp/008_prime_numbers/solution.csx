using System;
using System.Collections.Generic;

// Get N from command line arguments
if (Args.Count < 1)
{
    return;
}

int N;
if (!int.TryParse(Args[0], out N))
{
    return;
}

// Generate primes up to N using Sieve of Eratosthenes
if (N < 2)
{
    // No primes up to N
    return;
}

bool[] isComposite = new bool[N + 1];

// 0 and 1 are not prime, but we won't check them since we start from 2

for (int i = 2; i * i <= N; i++)
{
    if (!isComposite[i])
    {
        for (int j = i * i; j <= N; j += i)
        {
            isComposite[j] = true;
        }
    }
}

for (int i = 2; i <= N; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}