/**
 * Calculates the bulk discount based on quantity.
 * 
 * Logic:
 * - quantity < 10: no discount (100% of price)
 * - 10 <= quantity < 50: 10% discount (90% of price)
 * - quantity >= 50: 25% discount (75% of price)
 * 
 * @param price - The unit price of the item
 * @param quantity - The number of items being purchased
 * @returns The total price after applying discounts
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Represents a single test case
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * Executes the test suite and prints results to stdout
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
            description: "Standard case: quantity 1 (No discount)",
            price: 50,
            quantity: 1,
            expected: 50
        },
        {
            description: "Large quantity: quantity 100 (25% discount applied)",
            price: 10,
            quantity: 100,
            expected: 750 // 100 * 10 * 0.75
        }
    ];

    let passedCount = 0;
    let failedCount = 0;

    for (const test of testCases) {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison
        const isCorrect = Math.abs(actual - test.expected) < 0.000001;

        if (isCorrect) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
            failedCount++;
        }
    }

    console.log(`\nTest Summary: ${passedCount} passed, ${failedCount} failed.`);
}

// Execute the program
runTests();