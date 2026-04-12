using System;
using System.Collections.Generic;

// Read command line argument
if (Args.Count < 1)
{
    return;
}

int N;
if (!int.TryParse(Args[0], out N))
{
    return;
}

// Generate and print primes up to N
for (int i = 2; i <= N; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}

bool IsPrime(int n)
{
    if (n < 2) return false;
    if (n == 2) return true;
    if (n % 2 == 0) return false;
    
    for (int i = 3; i * i <= n; i += 2)
    {
        if (n % i == 0) return false;
    }
    
    return true;
}