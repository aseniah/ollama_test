using System;
using System.Collections.Generic;

// Get N from command line arguments
if (Args.Count < 1)
{
    return;
}

int N = int.Parse(Args[0]);

// Find all prime numbers up to and including N
List<int> primes = new List<int>();

if (N >= 2)
{
    bool[] isComposite = new bool[N + 1];
    
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
            primes.Add(i);
        }
    }
}

// Print each prime number on its own line
foreach (int prime in primes)
{
    Console.WriteLine(prime);
}