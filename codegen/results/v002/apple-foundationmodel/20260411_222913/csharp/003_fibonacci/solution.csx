using System;
using System.Text.Json;

class FibonacciPrinter
{
    static void Main(string[] args)
    {
        if (args.Length < 1)
        {
            Console.WriteLine("Usage: Program.csx <N>");
            return;
        }

        int N = int.Parse(args[0]);

        if (N < 1)
        {
            Console.WriteLine("N must be a positive integer.");
            return;
        }

        int a = 1, b = 1;
        Console.WriteLine(a);

        while (b <= N)
        {
            Console.WriteLine(b);
            int next = a + b;
            a = b;
            b = next;
        }
    }
}