#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.WriteLine("Usage: dotnet-script script.csx <integer N>");
    return;
}

if (!int.TryParse(Args[0], out int N))
{
    Console.WriteLine("Error: Argument must be a valid integer.");
    return;
}

if (N < 2)
{
    // No primes less than 2
    return;
}

// Sieve of Eratosthenes implementation
bool[] isPrime = new bool[N + 1];

// Initialize all entries as true
for (int i = 2; i <= N; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= N; p++)
{
    if (isPrime[p])
    {
        for (int multiple = p * p; multiple <= N; multiple += p)
        {
            isPrime[multiple] = false;
        }
    }
}

// Print all prime numbers
for (int i = 2; i <= N; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}