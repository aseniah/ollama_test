using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the input argument N
if (!int.TryParse(Args[0], out int n) || n < 2)
{
    // If N is less than 2, there are no prime numbers to print.
    return;
}

// We use the Sieve of Eratosthenes algorithm to find all primes up to N.
// Create a boolean array "isPrime[0..n]" and initialize
// all entries it as true. A value in isPrime[i] will
// finally be false if i is Not a prime, else true.
bool[] isPrime = new bool[n + 1];

// Initialize array: assume all numbers are prime initially
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

// Implementation of the Sieve of Eratosthenes
for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p starting from p*p
        // Multiples less than p*p have already been marked
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Output all prime numbers
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}