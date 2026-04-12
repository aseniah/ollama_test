using System;
using System.Linq;

public static void Main(IList<string> Args)
{
    if (Args.Count != 1 || !int.TryParse(Args[0], out int n))
    {
        Console.WriteLine("Usage: dotnet-script script.csx <integer>");
        return;
    }

    for (int i = 2; i <= n; i++)
    {
        if (IsPrime(i))
        {
            Console.WriteLine(i);
        }
    }
}

private static bool IsPrime(int number)
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