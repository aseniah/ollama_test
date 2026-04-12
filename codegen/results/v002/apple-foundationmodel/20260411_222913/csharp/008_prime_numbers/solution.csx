using System;

// Function to check if a number is prime
bool IsPrime(int num)
{
    if (num <= 1) return false;
    for (int i = 2; i <= Math.Sqrt(num); i++)
    {
        if (num % i == 0)
        {
            return false;
        }
    }
    return true;
}

// Function to print all prime numbers up to and including N
void PrintPrimes(int N)
{
    for (int i = 2; i <= N; i++)
    {
        if (IsPrime(i))
        {
            Console.WriteLine(i);
        }
    }
}

// Main function
void Main()
{
    if (Args.Length < 1)
    {
        Console.WriteLine("Please provide an integer argument N.");
        return;
    }

    try
    {
        int N = int.Parse(Args[0]);
        PrintPrimes(N);
    }
    catch (FormatException)
    {
        Console.WriteLine("Invalid input. Please provide a valid integer.");
    }
}