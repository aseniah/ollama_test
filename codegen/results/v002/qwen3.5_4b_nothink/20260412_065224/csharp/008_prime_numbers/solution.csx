using System;
using System.Collections.Generic;

var args = Args;

if (args.Count == 0 || !int.TryParse(args[0], out var n))
{
    Console.WriteLine("Error: Please provide a single integer N.");
    return;
}

n = Math.Max(n, 2);
var primes = new HashSet<int>();

for (int i = 2; i <= n; i++)
{
    if (primes.Contains(i)) continue;

    bool isPrime = true;
    for (int d = 2; d * d <= i; d++)
    {
        if (i % d == 0)
        {
            isPrime = false;
            break;
        }
    }

    if (isPrime)
    {
        primes.Add(i);
    }
}

foreach (var prime in primes.OrderBy(x => x))
{
    Console.WriteLine(prime);
}