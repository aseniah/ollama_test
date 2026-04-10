using System;

if (Args.Count != 1)
{
    Console.WriteLine("Usage: script.csx <integer>");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Invalid input. Please provide a valid integer.");
    return;
}

for (int num = 2; num <= n; num++)
{
    if (IsPrime(num))
    {
        Console.WriteLine(num);
    }
}

bool IsPrime(int number)
{
    if (number <= 1) return false;
    if (number == 2) return true;
    if (number % 2 == 0) return false;

    for (int i = 3; i * i <= number; i += 2)
    {
        if (number % i == 0) return false;
    }

    return true;
}