```c#
using System;
using System.Text.Json;

if (Args.Count < 1)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

// Sieve of Eratosthenes up to n
if (n < 2)
{
    return;
}

bool[] isComposite = new bool[n + 1];
isComposite[0] = true;
isComposite[1] = true;

for (int i = 2; (long)i * i <= n; i++)
{
    if (!isComposite[i])
    {
        for (int j = i * i; j <= n; j += i)
        {
            isComposite[j] = true;
        }
    }
}

for (int i = 2; i <= n; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}
```