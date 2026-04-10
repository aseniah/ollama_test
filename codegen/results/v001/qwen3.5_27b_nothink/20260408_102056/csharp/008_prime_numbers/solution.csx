#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Error: Exactly one integer argument N is required.");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine($"Error: '{Args[0]}' is not a valid integer.");
    return;
}

if (n < 2)
{
    // No primes less than 2, print nothing
    return;
}

// Sieve of Eratosthenes
bool[] isPrime = new bool[n + 1];

// Initialize all numbers >= 2 as potential primes
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++)
{
    if (isPrime[p])
    {
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

for (int p = 2; p <= n; p++)
{
    if (isPrime[p])
    {
        Console.WriteLine(p);
    }
}