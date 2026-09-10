using System;

int N = int.Parse(Args[0]);

for (int i = 2; i <= N; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}

bool IsPrime(int n)
{
    if (n < 2) return false;
    if (n == 2 || n == 3) return true;
    if (n % 2 == 0 || n % 3 == 0) return false;
    for (int j = 5; j * j <= n; j += 6)
    {
        if (n % j == 0 || n % (j + 2) == 0) return false;
    }
    return true;
}