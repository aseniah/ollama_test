/**
 * Calculates the total price after applying bulk discounts.
 * 
 * Business Rules:
 * - Quantity < 10: No discount (100% of price)
 * - 10 <= Quantity < 50: 10% discount (90% of price)
 * - Quantity >= 50: 25% discount (75% of price)
 * 
 * @param price - The unit price of the item
 * @param quantity - The number of items purchased
 * @returns The total price after applying applicable discounts
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Interface representing a test case
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * The main execution function
 */
function runTests() {
    const testCases: TestCase[] = [
        {
            description: "Boundary: Quantity 9 (No discount)",
            price: 100,
            quantity: 9,
            expected: 900
        },
        {
            description: "Boundary: Quantity 10 (Start of 10% discount)",
            price: 100,
            quantity: 10,
            expected: 900 // 100 * 10 * 0.9
        },
        {
            description: "Boundary: Quantity 49 (End of 10% discount)",
            price: 100,
            quantity: 49,
            expected: 4410 // 100 * 49 * 0.9
        },
        {
            description: "Boundary: Quantity 50 (Start of 25% discount)",
            price: 100,
            quantity: 50,
            expected: 3750 // 100 * 50 * 0.75
        },
        {
            description: "Standard case: High quantity (25% discount)",
            price: 10,
            quantity: 100,
            expected: 750 // 10 * 100 * 0.75
        }
    ];

    let passedCount = 0;
    let failedCount = 0;

    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Use a small epsilon for floating point comparison
        const isPassed = Math.abs(actual - test.expected) < 0.000001;

        if (isPassed) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
            failedCount++;
        }
    }

    console.log("\n--- Test Summary ---");
    console.log(`Total: ${testCases.length} | Passed: ${passedCount} | Failed: ${failedCount}`);

    // Exit with non-zero status code if tests failed
    if (failedCount > 0) {
        process.exit(1);
    }
}

// Execute the test suite
runTests();