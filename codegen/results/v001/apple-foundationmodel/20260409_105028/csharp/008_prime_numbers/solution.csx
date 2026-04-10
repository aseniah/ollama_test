using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

namespace PrimePrinter
{
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
            var primes = GetPrimes(N);

            foreach (var prime in primes)
            {
                Console.WriteLine(prime);
            }
        }

        static List<int> GetPrimes(int n)
        {
            var primes = new List<int>();

            if (n < 2)
            {
                return primes;
            }

            // Use Sieve of Eratosthenes
            for (int i = 2; i <= n; i++)
            {
                if (IsPrime(i))
                {
                    primes.Add(i);
                }
            }

            return primes;
        }

        static bool IsPrime(int num)
        {
            if (num <= 1)
            {
                return false;
            }

            if (num <= 3)
            {
                return true;
            }

            if (num % 2 == 0 || num % 3 == 0)
            {
                return false;
            }

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
}