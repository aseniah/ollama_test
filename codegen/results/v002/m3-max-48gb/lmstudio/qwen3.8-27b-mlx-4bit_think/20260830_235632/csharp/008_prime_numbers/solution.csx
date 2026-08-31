using System;

int N = int.Parse(Args[0]);

for (int i = 2; i <= N; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}

static bool IsPrime(int n)
{
    if (n < 2) return false;
    if (n == 2) return true;
    if (n % 2 == 0) return false;
    for (int d = 3; (long)d * d <= n; d += 2)
    {
        if (n % d == 0) return false;
    }
    return true;
}