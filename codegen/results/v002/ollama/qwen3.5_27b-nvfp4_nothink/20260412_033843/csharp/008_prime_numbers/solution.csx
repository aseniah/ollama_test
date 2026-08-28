#!/usr/bin/env dotnet-script

using System;
using System.Collections.Generic;
using System.Linq;

// Validate input count
if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: script.csx <integer N>");
    Environment.Exit(1);
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine($"Error: '{Args[0]}' is not a valid integer.");
    Environment.Exit(1);
}

// Handle edge cases where N < 2 (no primes exist)
if (n < 2)
{
    return;
}

// Sieve of Eratosthenes to find primes up to n
bool[] isComposite = new bool[n + 1];
List<int> primes = new List<int>();

// We only need to check factors up to sqrt(n)
for (int i = 2; i * i <= n; i++)
{
    if (!isComposite[i])
    {
        // Mark multiples of i starting from i*i
        for (int j = i * i; j <= n; j += i)
        {
            isComposite[j] = true;
        }
    }
}

// Collect and print primes
for (int i = 2; i <= n; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}