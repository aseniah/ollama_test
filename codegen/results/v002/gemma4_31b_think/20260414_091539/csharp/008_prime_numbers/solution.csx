using System;

// Check if an argument was provided
if (Args.Length == 0)
{
    return;
}

// Parse the argument N
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// Prime numbers start at 2
if (n < 2)
{
    return;
}

// Sieve of Eratosthenes to find all primes up to N
bool[] isPrime = new bool[n + 1];
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p starting from p*p
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Print all prime numbers found
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}