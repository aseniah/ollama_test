using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Attempt to parse the first argument as an integer
if (int.TryParse(Args[0], out int n) && n >= 2)
{
    // We use the Sieve of Eratosthenes algorithm to find primes up to N.
    // isNotPrime[i] will be true if i is a composite number.
    // bool arrays in C# are initialized to false by default.
    bool[] isNotPrime = new bool[n + 1];

    // 0 and 1 are not prime numbers
    isNotPrime[0] = true;
    isNotPrime[1] = true;

    // Iterate from 2 up to the square root of n
    for (int p = 2; p * p <= n; p++)
    {
        // If isNotPrime[p] is false, then p is a prime
        if (!isNotPrime[p])
        {
            // Mark all multiples of p starting from p*p as not prime
            for (int i = p * p; i <= n; i += p)
            {
                isNotPrime[i] = true;
            }
        }
    }

    // Print all indices that are still marked as false (meaning they are prime)
    for (int i = 2; i <= n; i++)
    {
        if (!isNotPrime[i])
        {
            Console.WriteLine(i);
        }
    }
}
// If n < 2, there are no prime numbers, so we print nothing as per requirements.