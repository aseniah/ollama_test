using System;
using System.Collections.Generic;

IList<string> Args;

if (!Args.Any() || !int.TryParse(Args[0], out int N))
{
    return;
}

if (N < 2)
{
    return;
}

bool[] isPrime = new bool[N + 1];
for (int i = 0; i <= N; i++)
{
    isPrime[i] = true;
}

isPrime[0] = false;
isPrime[1] = false;

for (int i = 2; i <= N; i++)
{
    if (isPrime[i])
    {
        for (int j = i * i; j <= N; j += i)
        {
            isPrime[j] = false;
        }
    }
}

int count = 0;
for (int i = 2; i <= N; i++)
{
    if (isPrime[i])
    {
        count++;
    }
}

int start = 2;
int end = 2;
int step = count / N;

for (int i = 0; i <= N; i++)
{
    if (isPrime[i])
    {
        Console.WriteLine(i);
    }
}