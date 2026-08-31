using System;
using System.Collections.Generic;

if (Args.Count == 0) return;
int N = int.Parse(Args[0]);

// Sieve of Eratosthenes
if (N >= 2)
{
    bool[] isComposite = new bool[N + 1];
    isComposite[0] = true;
    isComposite[1] = true;

    for (int i = 2; (long)i * i <= N; i++)
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
}