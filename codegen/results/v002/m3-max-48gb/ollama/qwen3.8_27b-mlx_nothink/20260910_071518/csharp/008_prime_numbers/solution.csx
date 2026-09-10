int N = int.Parse(Args[0]);

// Sieve of Eratosthenes or trial division
if (N >= 2)
{
    bool[] isPrime = new bool[N + 1];
    for (int i = 2; i <= N; i++)
        isPrime[i] = true;
    
    for (int i = 2; i * i <= N; i++)
    {
        if (isPrime[i])
        {
            for (int j = i * i; j <= N; j += i)
                isPrime[j] = false;
        }
    }
    
    for (int i = 2; i <= N; i++)
    {
        if (isPrime[i])
            Console.WriteLine(i);
    }
}