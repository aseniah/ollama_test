```c#
using System;
using System.Collections.Generic;

int N = int.Parse(Args[0]);

if (N >= 2)
{
    var sieve = new bool[N + 1];
    for (int i = 2; i * i <= N; i++)
    {
        if (!sieve[i])
        {
            for (int j = i * i; j <= N; j += i)
            {
                sieve[j] = true;
            }
        }
    }
    for (int i = 2; i <= N; i++)
    {
        if (!sieve[i])
        {
            Console.WriteLine(i);
        }
    }
}
```