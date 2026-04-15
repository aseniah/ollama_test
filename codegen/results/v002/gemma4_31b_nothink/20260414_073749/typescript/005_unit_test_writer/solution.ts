/**
 * Calculates the bulk discount based on quantity.
 * - Under 10: No discount
 * - 10 to 49: 10% discount
 * - 50 and above: 25% discount
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
            description: "No discount for quantity < 10 (Boundary 9)",
            price: 100,
            quantity: 9,
            expected: 900,
        },
        {
            description: "10% discount for quantity >= 10 (Boundary 10)",
            price: 100,
            quantity: 10,
            expected: 900, // 100 * 10 * 0.9
        },
        {
            description: "10% discount for quantity < 50 (Boundary 49)",
            price: 100,
            quantity: 49,
            expected: 4410, // 100 * 49 * 0.9
        },
        {
            description: "25% discount for quantity >= 50 (Boundary 50)",
            price: 100,
            quantity: 50,
            expected: 3750, // 100 * 50 * 0.75
        },
        {
            description: "High volume discount for quantity > 50",
            price: 10,
            quantity: 100,
            expected: 750, // 10 * 100 * 0.75
        }
    ];

    let passedCount = 0;

    testCases.forEach((test) => {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison if necessary, 
        // though these specific numbers are precise.
        if (Math.abs(result - test.expected) < 0.0001) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        }
    });

    console.log(`\nSummary: ${passedCount}/${testCases.length} tests passed.`);
}

// Execute the test suite
runTests();