using System;
using System.Collections.Generic;
using System.Text;

int N = int.Parse(Args[0]);

if (N < 2)
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