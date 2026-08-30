/**
 * The function provided in the requirements.
 * Calculates a bulk discount based on the quantity of items.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Interface representing a single test case structure.
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * Main execution function to run the test suite.
 */
function runTestSuite(): void {
    const testCases: TestCase[] = [
        {
            description: "Boundary condition: quantity 9 (no discount expected)",
            price: 100,
            quantity: 9,
            expected: 900
        },
        {
            description: "Boundary condition: quantity 10 (10% discount start)",
            price: 100,
            quantity: 10,
            expected: 900 // 100 * 10 * 0.9
        },
        {
            description: "Boundary condition: quantity 49 (10% discount end)",
            price: 100,
            quantity: 49,
            expected: 4410 // 100 * 49 * 0.9
        },
        {
            description: "Boundary condition: quantity 50 (25% discount start)",
            price: 100,
            quantity: 50,
            expected: 3750 // 100 * 50 * 0.75
        },
        {
            description: "Standard case: quantity 20 (10% discount)",
            price: 50,
            quantity: 20,
            expected: 900 // 50 * 20 * 0.9
        }
    ];

    // Epsilon value to handle potential floating point precision issues during comparison
    const EPSILON = 0.000001;

    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Check if the difference between actual and expected is within the acceptable epsilon
        const passed = Math.abs(actual - test.expected) < EPSILON;

        if (passed) {
            process.stdout.write(`PASS: ${test.description}\n`);
        } else {
            // Formatting numbers to remove trailing decimals for cleaner error messages if necessary,
            // but keeping them as provided by the calculation.
            process.stdout.write(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})\n`);
        }
    }
}

// Entry point
runTestSuite();