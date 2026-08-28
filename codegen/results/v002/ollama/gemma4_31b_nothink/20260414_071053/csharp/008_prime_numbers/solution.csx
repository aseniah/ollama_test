using System;

if (Args.Length == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

// Prime numbers start from 2
for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}

bool IsPrime(int number)
{
    if (number < 2) return false;
    if (number == 2) return true;
    if (number % 2 == 0) return false;

    int boundary = (int)Math.Floor(Math.Sqrt(number));

    for (int i = 3; i <= boundary; i += 2)
    {
        if (number % i == 0)
        {
            return false;
        }
    }

    return true;
}