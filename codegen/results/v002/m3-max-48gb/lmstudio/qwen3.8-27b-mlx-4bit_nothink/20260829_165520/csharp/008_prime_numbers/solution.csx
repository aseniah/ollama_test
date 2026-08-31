using System;
using System.Collections.Generic;

int N = int.Parse(Args[0]);

// Sieve of Eratosthenes to find all primes up to N
bool isComposite = false;
List<int> primes = new List<int>();

if (N >= 2)
{
    bool[] isPrime = new bool[N + 1];
    for (int i = 0; i <= N; i++)
        isPrime[i] = true;
    isPrime[0] = false;
    isPrime[1] = false;

    for (int i = 2; i * i <= N; i++)
    {
        if (isPrime[i])
        {
            for (int j = i * i; j <= N; j += i)
            {
                isPrime[j] = false;
            }
        }
    }

    for (int i = 2; i <= N; i++)
    {
        if (isPrime[i])
        {
            primes.Add(i);
        }
    }
}

foreach (int p in primes)
{
    Console.WriteLine(p);
}