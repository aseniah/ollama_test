using System;
using System.Text.Json;
using System.Text.Json.Nodes;

class Program
{
    static void Main(string[] args)
    {
        if (args.Length < 1)
        {
            Console.WriteLine("Usage: PrimePrinter.csx <integer>");
            return;
        }

        int N = int.Parse(args[0]);

        if (N < 2)
        {
            Console.WriteLine("No prime numbers less than 2.");
            return;
        }

        PrimePrinter.PrintPrimes(N);
    }
}

class PrimePrinter
{
    static void PrintPrimes(int N)
    {
        for (int i = 2; i <= N; i++)
        {
            if (IsPrime(i))
            {
                Console.WriteLine(i);
            }
        }
    }

    static bool IsPrime(int num)
    {
        if (num <= 1) return false;
        if (num <= 3) return true;
        if (num % 2 == 0 || num % 3 == 0) return false;

        for (int i = 5; i * i <= num; i += 6)
        {
            if (num % i == 0 || num % (i + 2) == 0)
            {
                return false;
            }
        }
        return true;
    }
}