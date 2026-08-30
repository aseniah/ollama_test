using System;

int N = int.Parse(Args[0]);

if (N < 2)
{
    // No primes up to 1
    return;
}

// Sieve of Eratosthenes
bool[] isComposite = new bool[N + 1];

for (int i = 2; i * i <= N; i++)
{
    if (!isComposite[i])
    {
        for (int j = i * i; j <= N; j += i)
        {
            isComposite[j] = true;
        }
    }
}

for (int i = 2; i <= N; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}