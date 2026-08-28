// Check if we have an argument
if (Args.Count == 0)
{
    return;
}

// Parse the argument
int n = int.Parse(Args[0]);

// Function to check if a number is prime
bool IsPrime(int num)
{
    if (num < 2) return false;
    if (num == 2) return true;
    if (num % 2 == 0) return false;
    
    int sqrt = (int)Math.Sqrt(num);
    for (int i = 3; i <= sqrt; i += 2)
    {
        if (num % i == 0) return false;
    }
    return true;
}

// Print all primes up to n
for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}