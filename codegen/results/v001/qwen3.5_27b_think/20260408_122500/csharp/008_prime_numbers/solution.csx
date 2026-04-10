using System;
using System.Collections.Generic;
using System.Linq;

// Parse the integer argument N
if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: script.csx <integer N>");
    Environment.Exit(1);
}

int N = int.Parse(Args[0]);

// Function to check if a number is prime
bool IsPrime(int num)
{
    if (num < 2) return false;
    if (num == 2) return true;
    if (num % 2 == 0) return false;
    
    int limit = (int)Math.Sqrt(num);
    for (int i = 3; i <= limit; i += 2)
    {
        if (num % i == 0) return false;
    }
    return true;
}

// Print all primes up to and including N
for (int i = 2; i <= N; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}