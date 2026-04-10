using System;
using System.Collections.Generic;
using System.Linq;

public static class Program
{
    public static int[] PrimesUpTo(int n)
    {
        if (n < 2)
        {
            return new int[0];
        }

        // Create a boolean array "isPrime[0..n]" with all entries set to true.
        // Initialize index 0 and 1 as false (not prime).
        bool[] isPrime = new bool[n + 1];
        for (int i = 2; i <= n; i++)
            isPrime[i] = true;

        // Implement Sieve of Eratosthenes
        for (int i = 2; i * i <= n; i++)
        {
            if (isPrime[i])
            {
                for (int j = i * i; j <= n; j += i)
                {
                    isPrime[j] = false;
                }
            }
        }

        // Collect primes into a list and then to array
        List<int> primes = new List<int>();
        for (int i = 2; i <= n; i++)
        {
            if (isPrime[i])
            {
                primes.Add(i);
            }
        }

        return primes.ToArray();
    }

    public static void Main(string[] args)
    {
        int n;

        // Validate argument
        if (args.Count != 1 || !int.TryParse(args[0], out n))
        {
            Console.WriteLine("Error: Expected a single integer argument N.");
            return;
        }

        int[] primes = PrimesUpTo(n);

        // Print all primes, one per line
        foreach (var prime in primes)
        {
            Console.WriteLine(prime);
        }
    }
}