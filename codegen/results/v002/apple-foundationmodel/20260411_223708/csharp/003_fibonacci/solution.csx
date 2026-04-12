using System;
using System.IO;
using System.Text.Json;

class FibonacciProcessor
{
    static void Main()
    {
        if (Args.Length < 1 || !int.TryParse(Args[0], out int N))
        {
            Console.WriteLine("Invalid input. Please provide a single integer argument.");
            return;
        }

        if (N < 1)
        {
            Console.WriteLine("N must be at least 1.");
            return;
        }

        PrintFibonacciNumbers(N);
    }

    static void PrintFibonacciNumbers(int N)
    {
        if (N < 1)
        {
            return;
        }

        int a = 1, b = 1, current = 1;
        while (current <= N)
        {
            Console.WriteLine(current);
            int next = a + b;
            a = b;
            b = next;
            current = next;
        }
    }
}

class Program
{
    static void Main(string[] args)
    {
        FibonacciProcessor.Main(args);
    }
}