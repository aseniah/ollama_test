#r "System.Runtime.Numerics"

using System;
using System.Collections.Generic;
using System.Linq;

if (Args.Count == 0)
{
    return;
}

int n = int.Parse(Args[0]);

// Function to check if a number is prime
bool IsPrime(int num)
{
    if (num < 2) return false;
    if (num == 2) return true;
    if (num % 2 == 0) return false;
    for (int i = 3; i <= Math.Sqrt(num); i += 2)
    {
        if (num % i == 0) return false;
    }
    return true;
}

// Print all primes from 2 up to and including n
for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}