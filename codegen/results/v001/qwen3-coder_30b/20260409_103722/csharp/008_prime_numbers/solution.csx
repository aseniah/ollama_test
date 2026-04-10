using System;
using System.Collections.Generic;

if (Args.Count == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n) || n < 2)
{
    return;
}

var primes = new List<int>();

for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        primes.Add(i);
    }
}

foreach (int prime in primes)
{
    Console.WriteLine(prime);
}

bool IsPrime(int number)
{
    if (number < 2) return false;
    if (number == 2) return true;
    if (number % 2 == 0) return false;
    
    for (int i = 3; i * i <= number; i += 2)
    {
        if (number % i == 0)
            return false;
    }
    
    return true;
}