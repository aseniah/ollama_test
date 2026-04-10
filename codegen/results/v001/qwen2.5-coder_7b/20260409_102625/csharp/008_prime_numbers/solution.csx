using System;

if (Args.Count != 1)
{
    Console.WriteLine("Usage: dotnet-script script.csx <integer>");
    return;
}

int n = int.Parse(Args[0]);

for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}

bool IsPrime(int number)
{
    if (number <= 1) return false;
    if (number == 2) return true;
    if (number % 2 == 0) return false;

    var boundary = (int)Math.Floor(Math.Sqrt(number));

    for (int i = 3; i <= boundary; i += 2)
    {
        if (number % i == 0) return false;
    }

    return true;
}