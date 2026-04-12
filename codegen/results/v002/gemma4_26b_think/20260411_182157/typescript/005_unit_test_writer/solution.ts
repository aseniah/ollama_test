/**
 * The function provided in the task description.
 * Calculates a bulk discount based on the quantity of items purchased.
 * 
 * @param price - The price per unit.
 * @param quantity - The number of units purchased.
 * @returns The total price after applying potential discounts.
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
 * A collection of test cases covering various boundaries and scenarios.
 */
const testCases: TestCase[] = [
    {
        description: "Quantity below 10 (9 units)",
        price: 100,
        quantity: 9,
        expected: 900 // 100 * 9
    },
    {
        description: "Quantity at first threshold (10 units)",
        price: 100,
        quantity: 10,
        expected: 900 // 100 * 10 * 0.9
    },
    {
        description: "Quantity just below second threshold (49 units)",
        price: 100,
        quantity: 49,
        expected: 4410 // 100 * 49 * 0.9
    },
    {
        description: "Quantity at second threshold (50 units)",
        price: 100,
        quantity: 50,
        expected: 3750 // 100 * 50 * 0.75
    },
    {
        description: "Quantity well above second threshold (100 units)",
        price: 10,
        quantity: 100,
        expected: 750 // 10 * 100 * 0.75
    },
    {
        description: "Zero quantity",
        price: 100,
        quantity: 0,
        expected: 0
    }
];

/**
 * Executes the test suite and prints the results to stdout.
 */
function runTestSuite(): void {
    let passedCount = 0;
    let failedCount = 0;

    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // We use a small epsilon for floating point comparison to avoid precision issues
        // common in JavaScript/TypeScript math.
        const isPassed = Math.abs(actual - test.expected) < Number.EPSILON;

        if (isPassed) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
            failedCount++;
        }
    }

    console.log(`\nTest Summary: ${passedCount} passed, ${failedCount} failed.`);
}

// Run the program
runTestSuite();