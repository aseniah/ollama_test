/**
 * Calculates the bulk discount based on quantity.
 * 
 * @param price - The price of a single unit.
 * @param quantity - The number of units being purchased.
 * @returns The total price after applying applicable discounts.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Represents a single test case for the discount calculation.
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * Runs the test suite and prints results to stdout.
 */
function runTests(): void {
    const testCases: TestCase[] = [
        {
            description: "Quantity below 10 (No discount)",
            price: 100,
            quantity: 9,
            expected: 900
        },
        {
            description: "Quantity exactly 10 (10% discount threshold)",
            price: 100,
            quantity: 10,
            expected: 900 // 100 * 10 * 0.9
        },
        {
            description: "Quantity just below 50 (10% discount)",
            price: 10,
            quantity: 49,
            expected: 441 // 10 * 49 * 0.9
        },
        {
            description: "Quantity exactly 50 (25% discount threshold)",
            price: 10,
            quantity: 50,
            expected: 375 // 10 * 50 * 0.75
        },
        {
            description: "Large quantity (25% discount)",
            price: 2,
            quantity: 100,
            expected: 150 // 2 * 100 * 0.75
        }
    ];

    testCases.forEach((test) => {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // We use a small epsilon to handle floating point precision errors in comparison
        const epsilon = 0.000001;
        const passed = Math.abs(actual - test.expected) < epsilon;

        if (passed) {
            process.stdout.write(`PASS: ${test.description}\n`);
        } else {
            // Formatted to 2 decimal places to ensure clean output for the user
            const expectedFormatted = Number(test.expected.toFixed(2));
            const actualFormatted = Number(actual.toFixed(2));
            process.stdout.write(`FAIL: ${test.description} (expected: ${expectedFormatted}, got: ${actualFormatted})\n`);
        }
    });
}

// Execute the test runner
runTests();