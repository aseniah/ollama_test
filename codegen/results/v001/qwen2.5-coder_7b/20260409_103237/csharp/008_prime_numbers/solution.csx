using System;

public static bool IsPrime(int number)
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

if (Args.Count != 1 || !int.TryParse(Args[0], out int N))
{
    Console.WriteLine("Please provide a valid integer argument.");
}
else
{
    for (int i = 2; i <= N; i++)
    {
        if (IsPrime(i))
        {
            Console.WriteLine(i);
        }
    }
}