using System;
using System.Collections.Generic;
using System.Linq;

// Parse the argument
if (Args.Count != 1 || !int.TryParse(Args[0], out int N))
{
    Console.Error.WriteLine("Error: provide a single integer argument N");
    return;
}

// Generate primes up to N using Sieve of Eratosthenes
List<int> primes = new List<int>();

if (N >= 2)
{
    bool[] isPrime = new bool[N + 1];
    for (int i = 2; i <= N; i++)
        isPrime[i] = true;

    for (int i = 2; i * i <= N; i++)
    {
        if (isPrime[i])
        {
            for (int j = i * i; j <= N; j += i)
                isPrime[j] = false;
        }
    }

    for (int i = 2; i <= N; i++)
    {
        if (isPrime[i])
            primes.Add(i);
    }
}

foreach (int p in primes)
    Console.WriteLine(p);