if (Args.Count != 1)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

if (n >= 2)
{
    bool[] isPrime = new bool[n + 1];
    // Initialize all as true except 0 and 1
    for (int i = 2; i <= n; i++)
    {
        isPrime[i] = true;
    }
    
    // Sieve of Eratosthenes
    for (int i = 2; i * i <= n; i++)
    {
        if (isPrime[i])
        {
            for (int j = i * i; j <= n; j += i)
            {
                isPrime[j] = false;
            }
        }
    }
    
    // Print all primes
    for (int i = 2; i <= n; i++)
    {
        if (isPrime[i])
        {
            Console.WriteLine(i);
        }
    }
}