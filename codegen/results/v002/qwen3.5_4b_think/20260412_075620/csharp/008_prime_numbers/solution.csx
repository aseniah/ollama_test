int n = int.Parse(Args[0]);

for (int i = 2; i <= n; i++)
{
    bool isPrime = true;
    int j = 2;
    while (j * j <= i)
    {
        if (i % j == 0)
        {
            isPrime = false;
            break;
        }
        j++;
    }
    if (isPrime)
    {
        Console.WriteLine(i);
    }
}