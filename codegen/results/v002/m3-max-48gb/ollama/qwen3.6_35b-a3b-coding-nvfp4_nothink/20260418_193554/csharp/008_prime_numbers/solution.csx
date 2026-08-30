using System;
using System.Collections.Generic;

// Parse the command line argument N
if (Args.Count < 1)
{
    return;
}

int N;
try
{
    N = int.Parse(Args[0]);
}
catch
{
    return;
}

// Print all prime numbers up to and including N
if (N >= 2)
{
    for (int i = 2; i <= N; i++)
    {
        if (IsPrime(i))
        {
            Console.WriteLine(i);
        }
    }
}

bool IsPrime(int n)
{
    if (n < 2)
        return false;
    if (n == 2)
        return true;
    if (n % 2 == 0)
        return false;
    
    int limit = (int)Math.Sqrt(n);
    for (int i = 3; i <= limit; i += 2)
    {
        if (n % i == 0)
            return false;
    }
    return true;
}