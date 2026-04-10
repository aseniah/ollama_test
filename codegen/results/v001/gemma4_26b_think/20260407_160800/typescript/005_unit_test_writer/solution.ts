/**
 * Calculates the bulk discount based on the quantity ordered.
 * 
 * @param price - The unit price of the item.
 * @param quantity - The number of items ordered.
 * @returns The total price after applying the appropriate discount.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Represents a single test case for the calculateDiscount function.
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
const testCases: TestCase[] = [
    {
        description: "Boundary condition: quantity 9 (no discount)",
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "Boundary condition: quantity 10 (10% discount starts)",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "Boundary condition: quantity 49 (10% discount applies)",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "Boundary condition: quantity 50 (25% discount starts)",
        price: 100,
        quantity: 50,
        expected: 3750
    }
];

/**
 * Executes the test suite and prints results to stdout.
 */
function runTests(): void {
    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Using an epsilon comparison to handle potential floating point precision issues
        // though for these specific inputs exact equality usually holds in IEEE 754.
        const epsilon = 0.000001;
        const passed = Math.abs(actual - test.expected) < epsilon;

        if (passed) {
            process.stdout.write(`PASS: ${test.description}\n`);
        } else {
            process.stdout.write(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})\n`);
        }
    }
}

// Execute the test runner
runTests();