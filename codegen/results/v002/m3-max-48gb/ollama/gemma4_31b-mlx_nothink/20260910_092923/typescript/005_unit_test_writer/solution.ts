/**
 * Calculates a bulk discount based on the quantity ordered.
 * - Less than 10: No discount
 * - 10 to 49: 10% discount
 * - 50 or more: 25% discount
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
            description: "No discount for quantity < 10 (boundary 9)", 
            price: 100, 
            quantity: 9, 
            expected: 900 
        },
        { 
            description: "10% discount for quantity 10 (lower boundary)", 
            price: 100, 
            quantity: 10, 
            expected: 900 // 100 * 10 * 0.9
        },
        { 
            description: "10% discount for quantity 49 (upper boundary)", 
            price: 100, 
            quantity: 49, 
            expected: 4410 // 100 * 49 * 0.9
        },
        { 
            description: "25% discount for quantity 50 (lower boundary)", 
            price: 100, 
            quantity: 50, 
            expected: 3750 // 100 * 50 * 0.75
        },
        { 
            description: "25% discount for quantity > 50", 
            price: 100, 
            quantity: 100, 
            expected: 7500 // 100 * 100 * 0.75
        },
    ];

    testCases.forEach(tc => {
        const result = calculateDiscount(tc.price, tc.quantity);
        
        // Using a small epsilon for floating point comparison
        if (Math.abs(result - tc.expected) < 0.000001) {
            console.log(`PASS: ${tc.description}`);
        } else {
            console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
        }
    });
}

// Execute the tests
runTests();