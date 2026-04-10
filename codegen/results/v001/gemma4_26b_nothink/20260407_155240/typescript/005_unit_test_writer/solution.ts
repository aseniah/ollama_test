/**
 * Calculates a bulk discount based on quantity.
 * 
 * Logic:
 * - quantity < 10: no discount (100% price)
 * - 10 <= quantity < 50: 10% discount (90% price)
 * - quantity >= 50: 25% discount (75% price)
 * 
 * @param price The unit price
 * @param quantity The number of items
 * @returns The total price after applying discounts
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Interface representing a single test case
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * Main execution function
 */
function runTests(): void {
    const testCases: TestCase[] = [
        {
            description: "Boundary: quantity 9 (No discount)",
            price: 100,
            quantity: 9,
            expected: 900
        },
        {
            description: "Boundary: quantity 10 (10% discount starts)",
            price: 100,
            quantity: 10,
            expected: 900 // 10 * 100 * 0.9
        },
        {
            description: "Boundary: quantity 49 (10% discount applies)",
            price: 100,
            quantity: 49,
            expected: 4410 // 49 * 100 * 0.9
        },
        {
            description: "Boundary: quantity 50 (25% discount starts)",
            price: 100,
            quantity: 50,
            expected: 3750 // 50 * 100 * 0.75
        },
        {
            description: "Normal case: quantity 100 (25% discount applies)",
            price: 10,
            quantity: 100,
            expected: 750
        }
    ];

    let passedCount = 0;
    let failedCount = 0;

    console.log("Running Discount Calculation Tests...\n");

    for (const test of testCases) {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison
        const isCorrect = Math.abs(result - test.expected) < 0.00001;

        if (isCorrect) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
            failedCount++;
        }
    }

    console.log(`\nTest Summary: ${passedCount} passed, ${failedCount} failed.`);

    // Exit with non-zero code if there were failures
    if (failedCount > 0) {
        process.exit(1);
    } else {
        process.exit(0);
    }
}

// Execute the test suite
runTests();