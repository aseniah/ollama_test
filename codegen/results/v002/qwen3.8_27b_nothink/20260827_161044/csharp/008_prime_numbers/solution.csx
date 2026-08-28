using System;
using System.Collections.Generic;
using System.Numerics;

int N = int.Parse(Args[0]);

// Sieve of Eratosthenes for all primes up to N
if (N >= 2)
{
    bool[] isPrime = new bool[N + 1];
    // Initialize all as true (prime)
    for (int i = 0; i <= N; i++)
    {
        isPrime[i] = true;
    }

    // 0 and 1 are not prime
    isPrime[0] = false;
    isPrime[1] = false;

    for (int i = 2; i <= (int)Math.Sqrt(N); i++)
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
            Console.WriteLine(i);
        }
    }
}