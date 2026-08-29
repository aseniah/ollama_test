```csharp
using System;
using System.Globalization;

double CalculateDiscount(double price, int quantity)
{
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

int failures = 0;

void Check(double price, int quantity, double expected, string description)
{
    double got = CalculateDiscount(price, quantity);

    if (Math.Abs(got - expected) < 1e-9)
    {
        Console.WriteLine($"PASS: {description}");
    }
    else
    {
        Console.WriteLine(
            $"FAIL: {description} (expected: {expected.ToString(CultureInfo.InvariantCulture)}, got: {got.ToString(CultureInfo.InvariantCulture)})"
        );
        failures++;
    }
}

Check(10, 9, 90, "quantity 9 is below the discount threshold");
Check(10, 10, 90, "quantity 10 gets a 10% discount");
Check(10