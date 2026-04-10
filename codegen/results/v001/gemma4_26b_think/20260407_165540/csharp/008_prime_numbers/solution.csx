using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the integer N from the first command-line argument
if (!int.TryParse(Args[0], out int n) || n < 2)
{
    // If N is less than 2, there are no prime numbers to print.
    return;
}

// Use the Sieve of Eratosthenes algorithm to find all primes up to N.
// isNotPrime[i] will be true if i is a composite number.
bool[] isNotPrime = new bool[n + 1];

for (int p = 2; p * p <= n; p++)
{
    // If isNotPrime[p] is still false, then p is a prime.
    if (!isNotPrime[p])
    {
        // Update all multiples of p starting from p*p.
        for (int i = p * p; i <= n; i += p)
        {
            isNotPrime[i] = true;
        }
    }
}

// Iterate from 2 up to N and print every number that is not marked as composite.
for (int i = 2; i <= n; i++)
{
    if (!isNotPrime[i])
    {
        Console.WriteLine(i);
    }
}