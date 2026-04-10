using System;

// Check if an argument was provided and if it is a valid integer
if (Args.Count != 1 || !int.TryParse(Args[0], out int n) || n < 2)
{
    // If n < 2, there are no prime numbers, so we print nothing as per instructions.
    // If the input is invalid, we exit silently.
    return;
}

// Implementation of the Sieve of Eratosthenes algorithm
// We create a boolean array "isPrime[0..n]" and initialize
// all entries it as true. A value in isPrime[i] will
// finally be false if i is Not a prime, else true.
bool[] isPrime = new bool[n + 1];

// Initialize array: assume all numbers >= 2 are prime initially
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

// Start sieving from the first prime, 2
for (int p = 2; p * p <= n; p++)
{
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p])
    {
        // Update all multiples of p greater than or equal to the square of it
        // Numbers which are multiples of p and are already marked as false
        // will be skipped naturally by the inner loop logic.
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