using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Try to parse the first argument as an integer N
if (int.TryParse(Args[0], out int n))
{
    // If N is less than 2, there are no prime numbers to print
    if (n < 2)
    {
        return;
    }

    // We use the Sieve of Eratosthenes algorithm to find all primes up to N.
    // isPrime[i] will be true if i is a prime number.
    bool[] isPrime = new bool[n + 1];
    
    // Initialize all entries from 2 to n as true
    for (int i = 2; i <= n; i++)
    {
        isPrime[i] = true;
    }

    // Sieve process: Iterate from the first prime (2) up to the square root of n
    for (int p = 2; p * p <= n; p++)
    {
        // If isPrime[p] is not changed, then it is a prime
        if (isPrime[p])
        {
            // Update all multiples of p starting from p*p
            // Multiples of p less than p*p have already been marked
            for (int i = p * p; i <= n; i += p)
            {
                isPrime[i] = false;
            }
        }
    }

    // Print all indices that are still marked as true
    for (int i = 2; i <= n; i++)
    {
        if (isPrime[i])
        {
            Console.WriteLine(i);
        }
    }
}