class TestAddItemToCart:
	def test_add_item(self):
		# Arrange
		cart = Cart(user_type='regular')
		item_id = 1
		quantity = 2
		price = 10.0
		name = 'Test Item'
		category = 'Test Category'
		
		# Act
		cart.add_item(item_id, quantity, price, name, category)
		
		# Assert
		assert len(cart.items) == 1
		assert cart.items[0]['item_id'] == item_id
		assert cart.items[0]['quantity'] == quantity
		assert cart.items[0]['price'] == price
		assert cart.items[0]['name'] == name
		assert cart.items[0]['category'] == category
		assert cart.items[0]['user_type'] == 'regular'