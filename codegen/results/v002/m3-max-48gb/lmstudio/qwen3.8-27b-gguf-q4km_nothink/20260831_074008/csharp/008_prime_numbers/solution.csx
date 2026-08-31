using System;
using System.Linq;

if (Args.Count == 0 || !int.TryParse(Args[0], out int N) || N < 2)
{
    return;
}

for (int i = 2; i <= N; i++)
{
    bool isPrime = true;
    for (int d = 2; d * d <= i; d++)
    {
        if (i % d == 0)
        {
            isPrime = false;
            break;
        }
    }
    if (isPrime)
    {
        Console.WriteLine(i);
    }
}