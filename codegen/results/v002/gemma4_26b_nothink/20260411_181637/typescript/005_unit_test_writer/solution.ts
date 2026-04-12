/**
 * Calculates the total price after applying bulk discounts.
 * - Quantity < 10: No discount.
 * - 10 <= Quantity < 50: 10% discount.
 * - Quantity >= 50: 25% discount.
 * 
 * @param price The unit price of the item.
 * @param quantity The number of items being purchased.
 * @returns The total price after discount.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Represents a single test case structure.
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
function runTests() {
    const testCases: TestCase[] = [
        {
            description: "Boundary: quantity 9 (No discount)",
            price: 100,
            quantity: 9,
            expected: 900
        },
        {
            description: "Boundary: quantity 10 (10% discount applied)",
            price: 100,
            quantity: 10,
            expected: 900 // 10 * 100 * 0.9
        },
        {
            description: "Boundary: quantity 49 (10% discount applied)",
            price: 100,
            quantity: 49,
            expected: 4410 // 49 * 100 * 0.9
        },
        {
            description: "Boundary: quantity 50 (25% discount applied)",
            price: 100,
            quantity: 50,
            expected: 3750 // 50 * 100 * 0.75
        },
        {
            description: "Standard Case: quantity 100 (25% discount applied)",
            price: 10,
            quantity: 100,
            expected: 750
        }
    ];

    let passedCount = 0;
    let failedCount = 0;

    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Use a small epsilon for floating point comparison
        const epsilon = 0.00001;
        const isPass = Math.abs(actual - test.expected) < epsilon;

        if (isPass) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
            failedCount++;
        }
    }

    console.log(`\nTest Summary: ${passedCount} passed, ${failedCount} failed.`);

    // Exit with error code if tests failed
    if (failedCount > 0) {
        process.exit(1);
    }
}

// Execute the tests
runTests();