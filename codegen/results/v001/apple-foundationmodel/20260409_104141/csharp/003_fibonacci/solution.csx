using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Function to generate and return the Fibonacci sequence up to the specified limit
public static List<int> GenerateFibonacciSequence(int limit)
{
    List<int> fibonacciSequence = new List<int>();
    if (limit < 1) return fibonacciSequence; // Return empty list if N < 1

    int a = 1;
    int b = 1;

    while (b <= limit)
    {
        fibonacciSequence.Add(b);
        int next = a + b;
        a = b;
        b = next;
    }

    return fibonacciSequence;
}

// Main script execution
if (Args.Length == 1)
{
    try
    {
        int N = int.Parse(Args[0]);
        if (N < 1)
        {
            Console.WriteLine(""); // Print nothing if N is less than 1
            return;
        }

        List<int> fibonacci = GenerateFibonacciSequence(N);
        foreach (int number in fibonacci)
        {
            Console.WriteLine(number);
        }
    }
    catch (Exception)
    {
        Console.WriteLine("Invalid input. Please provide a valid integer.");
    }
}