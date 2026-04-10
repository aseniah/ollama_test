using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

class FibonacciSequence
{
    static void Main(string[] args)
    {
        if (args.Length < 1)
        {
            Console.WriteLine("Usage: FibonacciSequence.csx <N>");
            return;
        }

        int N = int.Parse(args[0]);
        if (N < 1)
        {
            Console.WriteLine("N must be a positive integer.");
            return;
        }

        int a = 1, b = 1;
        string output = "";

        while (a <= N)
        {
            output += $"{a}\n";
            int next = a + b;
            a = b;
            b = next;
        }

        Console.WriteLine(output);
    }
}