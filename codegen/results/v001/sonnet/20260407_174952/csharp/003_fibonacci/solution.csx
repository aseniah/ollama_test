int n = int.Parse(Args[0]);

if (n >= 1)
{
    long a = 1, b = 1;
    while (a <= n)
    {
        Console.WriteLine(a);
        (a, b) = (b, a + b);
    }
}
