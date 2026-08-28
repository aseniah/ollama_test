#r "System.Text.Json"
using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

// Read the single integer argument
if (Args.Count != 1)
{
    // If no argument provided, we do nothing or exit. 
    // The problem implies valid input handling, but strictly we just proceed.
    // If the user wants error handling, they didn't specify, but usually scripts exit on bad input.
    // We will assume valid input based on "Accept a single integer argument".
    return;
}

int n = int.Parse(Args[0]);

if (n < 2)
{
    // No primes up to N if N is less than 2
    return;
}

// Sieve of Eratosthenes to find primes up to N
// We need a boolean array to mark non-prime numbers
// Index i corresponds to number i
bool[] isPrime = new bool[n + 1];

// Initialize all entries as true (potential primes)
// We can skip 0 and 1, but initialization covers them as false later
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Print all prime numbers
for (int p = 2; p <= n; p++)
{
    if (isPrime[p])
    {
        Console.WriteLine(p);
    }
}