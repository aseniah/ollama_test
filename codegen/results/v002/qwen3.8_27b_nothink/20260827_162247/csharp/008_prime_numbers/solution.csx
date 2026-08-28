using System;
using System.Collections.Generic;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: provide a single integer argument N");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Invalid input: not an integer");
    return;
}

if (n < 2)
{
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