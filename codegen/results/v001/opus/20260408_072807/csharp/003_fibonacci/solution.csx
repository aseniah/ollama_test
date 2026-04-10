if (Args.Count == 0) return;

if (!int.TryParse(Args[0], out int n) || n < 1) return;

long a = 1, b = 1;
while (a <= n)
{
    Console.WriteLine(a);
    (a, b) = (b, a + b);
}
