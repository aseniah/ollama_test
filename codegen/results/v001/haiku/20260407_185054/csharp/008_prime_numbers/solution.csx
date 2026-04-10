if (Args.Count == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n) || n < 2)
{
    return;
}

bool IsPrime(int num)
{
    if (num < 2) return false;
    if (num == 2) return true;
    if (num % 2 == 0) return false;

    for (int i = 3; i * i <= num; i += 2)
    {
        if (num % i == 0) return false;
    }

    return true;
}

for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}
