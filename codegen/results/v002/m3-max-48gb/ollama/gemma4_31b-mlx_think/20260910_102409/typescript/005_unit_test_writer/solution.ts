/**
 * Calculates a bulk discount based on quantity.
 * @param price The unit price of the item.
 * @param quantity The number of items being purchased.
 * @returns The total cost after applying any eligible discounts.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    price: number;
    quantity: number;
    expected: number;
    description: string;
}

function runTests() {
    const tests: TestCase[] = [
        { 
            price: 100, 
            quantity: 9, 
            expected: 900, 
            description: "Quantity just below the first threshold (9)" 
        },
        { 
            price: 100, 
            quantity: 10, 
            expected: 900, 
            description: "Quantity exactly at the first threshold (10)" 
        },
        { 
            price: 100, 
            quantity: 49, 
            expected: 4410, 
            description: "Quantity just below the second threshold (49)" 
        },
        { 
            price: 100, 
            quantity: 50, 
            expected: 3750, 
            description: "Quantity exactly at the second threshold (50)" 
        },
    ];

    tests.forEach(test => {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison to avoid precision errors
        const isCorrect = Math.abs(actual - test.expected) < 0.00001;

        if (isCorrect) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
        }
    });
}

// Execute the test suite
runTests();