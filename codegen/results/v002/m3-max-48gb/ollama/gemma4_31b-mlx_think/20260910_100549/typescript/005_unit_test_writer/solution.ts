/**
 * Calculates the total price based on bulk discount rules:
 * - Less than 10 items: No discount
 * - 10 to 49 items: 10% discount
 * - 50 or more items: 25% discount
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
    const tests: TestCase[] = [
        {
            description: "No discount for quantity < 10 (Boundary 9)",
            price: 100,
            quantity: 9,
            expected: 900,
        },
        {
            description: "10% discount for quantity >= 10 (Boundary 10)",
            price: 100,
            quantity: 10,
            expected: 900, // 100 * 10 * 0.9 = 900
        },
        {
            description: "10% discount for quantity < 50 (Boundary 49)",
            price: 100,
            quantity: 49,
            expected: 4410, // 100 * 49 * 0.9 = 4410
        },
        {
            description: "25% discount for quantity >= 50 (Boundary 50)",
            price: 100,
            quantity: 50,
            expected: 3750, // 100 * 50 * 0.75 = 3750
        },
        {
            description: "No discount for small quantity (Case 1)",
            price: 10,
            quantity: 1,
            expected: 10,
        },
        {
            description: "25% discount for very large quantity",
            price: 10,
            quantity: 100,
            expected: 750, // 10 * 100 * 0.75 = 750
        },
    ];

    tests.forEach((test) => {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison, 
        // although these specific calculations result in clean numbers.
        if (Math.abs(actual - test.expected) < 0.0001) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
        }
    });
}

runTests();