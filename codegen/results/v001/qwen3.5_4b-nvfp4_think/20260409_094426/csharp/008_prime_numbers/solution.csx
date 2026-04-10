int N = int.Parse(Args[0]);

var primes = new System.Collections.Generic.List<int>();

for (int i = 2; i <= N; i++)
{
    bool isPrime = true;
    for (int j = 2; j * j <= i; j++)
    {
        if (i % j == 0)
        {
            isPrime = false;
            break;
        }
    }
    if (isPrime)
    {
        primes.Add(i);
    }
}

foreach (var prime in primes)
{
    Console.WriteLine(prime);
}