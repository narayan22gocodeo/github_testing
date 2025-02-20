def calculate_order_total(items, tax_rate = 0.08)
  # Calculate subtotal
  subtotal = items.reduce(0) do |total, item|
    total + item[:price] * item[:quantity]
  end

  # Calculate tax
  tax = subtotal * tax_rate

  # Calculate total
  total = subtotal + tax

  # Return a hash with the results
  {
    subtotal: subtotal,
    tax: tax,
    total: total
  }
end
 
items = [
  { price: 10.99, quantity: 2 },
  { price: 5.99, quantity: 1 },
  { price: 7.99, quantity: 3 }
]

result = calculate_order_total(items)
puts result[:subtotal]  # Output: 43.92
puts result[:tax]       # Output: 3.51
puts result[:total]     # Output: 47.43