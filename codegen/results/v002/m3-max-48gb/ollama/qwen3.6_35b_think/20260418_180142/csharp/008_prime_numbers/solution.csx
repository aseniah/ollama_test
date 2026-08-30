using System;

if (Args.Count < 1 || !int.TryParse(Args[0], out int n))
    return;

for (int i = 2; i <= n; i++)
{
    bool isPrime = true;
    for (int j = 2; j * j <= i; j++)
    {
        if (i % j == 0)
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