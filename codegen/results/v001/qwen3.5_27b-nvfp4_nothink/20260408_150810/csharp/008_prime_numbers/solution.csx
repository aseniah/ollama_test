#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Validate argument count
if (Args.Count != 1)
{
    Console.WriteLine("Error: Expected exactly one integer argument N.");
    Environment.Exit(1);
}

// Parse the integer argument
int N = int.Parse(Args[0]);

// Handle edge cases where no primes exist (N < 2)
if (N < 2)
{
    // Print nothing
    return;
}

// Sieve of Eratosthenes to find primes up to N
bool[] isComposite = new bool[N + 1];
var primes = new List<int>();

for (int i = 2; i <= N; i++)
{
    if (!isComposite[i])
    {
        primes.Add(i);
        
        // Optimization: start at i*i, but check for overflow
        if ((long)i * i <= N)
        {
            for (int j = i * i; j <= N; j += i)
            {
                isComposite[j] = true;
            }
        }
    }
}

// Print primes, one per line
foreach (var prime in primes)
{
    Console.WriteLine(prime);
}