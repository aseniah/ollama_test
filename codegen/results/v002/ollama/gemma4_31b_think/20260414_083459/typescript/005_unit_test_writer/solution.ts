/**
 * Calculates a bulk discount based on quantity.
 * 
 * @param price - The unit price of the item.
 * @param quantity - The number of items being purchased.
 * @returns The total price after applying discounts.
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

function runTests() {
    const testCases: TestCase[] = [
        {
            description: "Quantity just below the first discount threshold (9)",
            price: 100,
            quantity: 9,
            expected: 900, // 100 * 9
        },
        {
            description: "Quantity at the first discount threshold (10)",
            price: 100,
            quantity: 10,
            expected: 900, // 100 * 10 * 0.9
        },
        {
            description: "Quantity just below the second discount threshold (49)",
            price: 100,
            quantity: 49,
            expected: 4410, // 100 * 49 * 0.9
        },
        {
            description: "Quantity at the second discount threshold (50)",
            price: 100,
            quantity: 50,
            expected: 3750, // 100 * 50 * 0.75
        },
        {
            description: "Large quantity well above thresholds (100)",
            price: 10,
            quantity: 100,
            expected: 750, // 10 * 100 * 0.75
        }
    ];

    testCases.forEach((test) => {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon to handle potential floating point precision issues
        const isPass = Math.abs(result - test.expected) < 0.0001;

        if (isPass) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        }
    });
}

// Execute the test runner
runTests();