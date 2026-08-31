using System;
using System.Text;

if (Args.Count < 1) return;

int n = int.Parse(Args[0]);

for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}

static bool IsPrime(int num)
{
    if (num < 2) return false;
    if (num == 2) return true;
    if (num % 2 == 0) return false;
    for (int d = 3; (long)d * d <= num; d += 2)
    {
        if (num % d == 0) return false;
    }
    return true;
}