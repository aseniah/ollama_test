/**
 * Calculates a bulk discount based on the quantity of items.
 * @param price The unit price of the item.
 * @param quantity The number of items purchased.
 * @returns The total price after applying the discount.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * Test runner for the calculateDiscount function.
 */
function runTests() {
    const tests: TestCase[] = [
        {
            description: "No discount for quantity < 10 (boundary 9)",
            price: 100,
            quantity: 9,
            expected: 900, // 100 * 9
        },
        {
            description: "10% discount for quantity >= 10 (boundary 10)",
            price: 100,
            quantity: 10,
            expected: 900, // 100 * 10 * 0.9
        },
        {
            description: "10% discount for quantity < 50 (boundary 49)",
            price: 100,
            quantity: 49,
            expected: 4410, // 100 * 49 * 0.9
        },
        {
            description: "25% discount for quantity >= 50 (boundary 50)",
            price: 100,
            quantity: 50,
            expected: 3750, // 100 * 50 * 0.75
        },
    ];

    tests.forEach((test) => {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison if necessary, 
        // although these specific numbers should be precise.
        if (Math.abs(result - test.expected) < 0.0001) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        }
    });
}

// Execute the tests
runTests();