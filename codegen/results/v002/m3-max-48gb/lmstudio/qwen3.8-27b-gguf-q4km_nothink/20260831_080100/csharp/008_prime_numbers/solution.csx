using System;

int N = int.Parse(Args[0]);

// Sieve of Eratosthenes
bool[] isPrime = new bool[N + 1];
if (N >= 2)
{
    Array.Fill(isPrime, true);
    isPrime[0] = false;
    isPrime[1] = false;
    for (int i = 2; i * i <= N; i++)
    {
        if (isPrime[i])
        {
            for (int j = i * i; j <= N; j += i)
            {
                isPrime[j] = false;
            }
        }
    }
}

for (int i = 2; i <= N; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}