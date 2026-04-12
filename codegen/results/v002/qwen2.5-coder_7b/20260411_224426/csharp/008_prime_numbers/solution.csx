using System;
using System.Linq;

public static void Main(string[] args)
{
    if (args.Length == 0 || !int.TryParse(args[0], out int n))
    {
        Console.WriteLine("Please provide a valid integer argument.");
        return;
    }

    var primes = GeneratePrimes(n);

    foreach (var prime in primes)
    {
        Console.WriteLine(prime);
    }
}

public static IEnumerable<int> GeneratePrimes(int maxNumber)
{
    if (maxNumber < 2) yield break;

    for (int number = 2; number <= maxNumber; number++)
    {
        if (IsPrime(number))
        {
            yield return number;
        }
    }
}

private static bool IsPrime(int number)
{
    if (number == 2) return true;
    if (number % 2 == 0) return false;

    var boundary = (int)Math.Floor(Math.Sqrt(number));

    for (int i = 3; i <= boundary; i += 2)
    {
        if (number % i == 0) return false;
    }

    return true;
}