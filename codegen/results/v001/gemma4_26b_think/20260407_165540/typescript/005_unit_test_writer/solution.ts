/**
 * Calculates a bulk discount based on quantity.
 * 
 * @param price - The unit price of the item.
 * @param quantity - The number of items being purchased.
 * @returns The total price after applying applicable discounts.
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
 * The test suite containing boundary conditions and standard cases.
 */
const testSuite: TestCase[] = [
    {
        description: "Quantity 9 (Boundary: below first discount tier)",
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "Quantity 10 (Boundary: start of 10% discount tier)",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "Quantity 49 (Boundary: end of 10% discount tier)",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "Quantity 50 (Boundary: start of 25% discount tier)",
        price: 100,
        quantity: 50,
        expected: 3750
    },
    {
        description: "Quantity 100 (Standard case: deep discount tier)",
        price: 10,
        quantity: 100,
        expected: 750
    }
];

/**
 * Executes the test suite and prints results to stdout.
 */
function runTests(): void {
    for (const test of testSuite) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison to avoid precision issues
        const epsilon = 0.000001;
        const passed = Math.abs(actual - test.expected) < epsilon;

        if (passed) {
            process.stdout.write(`PASS: ${test.description}\n`);
        } else {
            // Format numbers to avoid scientific notation for readability in error logs
            const expectedStr = Number.isInteger(test.expected) ? test.expected.toString() : test.expected.toFixed(2);
            const actualStr = Number.isInteger(actual) ? actual.toString() : actual.toFixed(2);
            
            process.stdout.write(`FAIL: ${test.description} (expected: ${expectedStr}, got: ${actualStr})\n`);
        }
    }
}

// Start the test execution
runTests();