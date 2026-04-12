using System;

int n = int.Parse(Args[0]);

void PrintPrimesUpTo(int limit)
{
    for (int i = 2; i <= limit; i++)
    {
        if (IsPrime(i))
        {
            Console.WriteLine(i);
        }
    }
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

PrintPrimesUpTo(n);