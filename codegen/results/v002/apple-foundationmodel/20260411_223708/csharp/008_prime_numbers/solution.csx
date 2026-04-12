using System;
using System.Collections.Generic;

// Function to check if a number is prime
bool IsPrime(int n)
{
    if (n <= 1) return false;
    if (n <= 3) return true;
    if (n % 2 == 0 || n % 3 == 0) return false;

    for (int i = 5; i * i <= n; i += 6)
    {
        if (n % i == 0 || n % (i + 2) == 0)
        {
            return false;
        }
    }
    return true;
}

// Main entry point
int Main(IList<string> args)
{
    if (args.Count != 1)
    {
        Console.WriteLine("Usage: Program.csx <N>");
        return 1;
    }

    int N = int.Parse(args[0]);

    for (int i = 2; i <= N; i++)
    {
        if (IsPrime(i))
        {
            Console.WriteLine(i);
        }
    }

    return 0;
}