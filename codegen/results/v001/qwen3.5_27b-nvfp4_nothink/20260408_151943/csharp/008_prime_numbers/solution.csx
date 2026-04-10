#r "System.Numerics"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.WriteLine("Error: Exactly one integer argument N is required.");
    Environment.Exit(1);
}

if (!int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Error: The provided argument is not a valid integer.");
    Environment.Exit(1);
}

if (n < 2)
{
    // No primes less than 2, print nothing
    return;
}

bool IsPrime(int number)
{
    if (number < 2) return false;
    if (number == 2) return true;
    if (number % 2 == 0) return false;

    int limit = (int)Math.Sqrt(number);
    for (int i = 3; i <= limit; i += 2)
    {
        if (number % i == 0) return false;
    }
    return true;
}

for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}