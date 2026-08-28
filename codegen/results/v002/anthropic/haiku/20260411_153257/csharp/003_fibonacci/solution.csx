int n = int.Parse(Args[0]);
if (n < 1) return;

long a = 1, b = 1;
while (a <= n)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}
