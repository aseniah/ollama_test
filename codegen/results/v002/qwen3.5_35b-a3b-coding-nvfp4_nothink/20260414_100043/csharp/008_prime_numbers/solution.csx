// Check if an argument was provided
if (Args.Count < 1)
{
    // No argument provided, exit silently as per "print nothing" if no primes (and no input means no primes)
    return;
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int n))
{
    // Invalid input, exit silently
    return;
}

// Handle edge cases where N is less than 2 (no primes exist)
if (n < 2)
{
    return;
}

// Sieve of Eratosthenes to find primes up to n
bool[] isPrime = new bool[n + 1];
for (int i = 2; i <= n; i++)
{
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++)
{
    if (isPrime[p])
    {
        for (int i = p * p; i <= n; i += p)
        {
            isPrime[i] = false;
        }
    }
}

// Output the prime numbers
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}