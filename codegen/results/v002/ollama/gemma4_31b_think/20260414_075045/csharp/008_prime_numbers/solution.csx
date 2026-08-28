using System;

if (Args.Length == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

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
    if (isPrime[p])
    {
        // Update all multiples of p starting from p*p
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