#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.WriteLine("Usage: dotnet-script script.csx N");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Invalid integer input.");
    return;
}

if (n < 2)
{
    // No primes less than 2, print nothing.
    return;
}

// Sieve of Eratosthenes
var sieve = new bool[n + 1];

for (int i = 2; i * i <= n; i++)
{
    if (!sieve[i])
    {
        for (int j = i * i; j <= n; j += i)
        {
            sieve[j] = true;
        }
    }
}

for (int i = 2; i <= n; i++)
{
    if (!sieve[i])
    {
        Console.WriteLine(i);
    }
}