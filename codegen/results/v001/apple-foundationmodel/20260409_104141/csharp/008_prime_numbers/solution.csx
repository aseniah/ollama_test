using System;
using System.Text.Json;
using System.Text.Json.Nodes;

void Main(string[] args)
{
    if (args.Length != 1)
    {
        Console.WriteLine("Usage: PrimeNumbers.csx <N>");
        return;
    }

    int N = int.Parse(args[0]);
    
    if (N < 2)
    {
        Console.WriteLine("No primes less than or equal to 1.");
        return;
    }

    bool[] isPrime = new bool[N + 1];
    for (int i = 2; i <= N; i++)
    {
        isPrime[i] = true;
    }

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

    for (int i = 2; i <= N; i++)
    {
        if (isPrime[i])
        {
            Console.WriteLine(i);
        }
    }
}