/**
 * The function to be tested.
 * Calculates a bulk discount based on quantity.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Interface representing a single test case.
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * The test suite containing the boundary conditions and specific test cases.
 */
const testCases: TestCase[] = [
    {
        description: "Boundary condition: quantity 9 (No discount)",
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "Boundary condition: quantity 10 (10% discount applied)",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "Boundary condition: quantity 49 (10% discount applied)",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "Boundary condition: quantity 50 (25% discount applied)",
        price: 100,
        quantity: 50,
        expected: 3750
    },
    {
        description: "Standard case: quantity 100 (25% discount applied)",
        price: 10,
        quantity: 100,
        expected: 750
    }
];

/**
 * Executes the test suite and prints the results to stdout.
 */
function runTests(): void {
    // Using an epsilon for floating point comparison to avoid precision errors common in JavaScript/TypeScript
    const EPSILON = 0.000001;

    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        const isPassed = Math.abs(actual - test.expected) < EPSILON;

        if (isPassed) {
            process.stdout.write(`PASS: ${test.description}\n`);
        } else {
            // Format numbers to avoid long floating point tails in error messages
            const expectedStr = Number(test.expected.toFixed(4));
            const actualStr = Number(actual.toFixed(4));
            process.stdout.write(`FAIL: ${test.description} (expected: ${expectedStr}, got: ${actualStr})\n`);
        }
    }
}

// Run the program
runTests();