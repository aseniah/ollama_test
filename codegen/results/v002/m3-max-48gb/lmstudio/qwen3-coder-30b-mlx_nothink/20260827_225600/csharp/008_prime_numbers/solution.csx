using System;
using System.Collections.Generic;

int N = int.Parse(Args[0]);

if (N < 2)
{
    return;
}

List<int> primes = new();

for (int i = 2; i <= N; i++)
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
    if (number < 2)
        return false;
    
    if (number == 2)
        return true;
    
    if (number % 2 == 0)
        return false;
    
    for (int i = 3; i * i <= number; i += 2)
    {
        if (number % i == 0)
            return false;
    }
    
    return true;
}