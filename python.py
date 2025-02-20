class CreditCard:
    def __init__(self, card_number, card_holder, credit_limit=1000):
        self.card_number = card_number
        self.card_holder = card_holder
        self.credit_limit = credit_limit
        self.balance = 0
    def charge(self, amount):
        if amount > 0:
            if self.balance + amount <= self.credit_limit:
                self.balance += amount
                print(f"Charged {amount} units. Current balance is {self.balance} units.")
            else:
                print("Exceeded credit limit.")
        else:
            print("Invalid amount for charge.")
    def make_payment(self, amount):
        if amount > 0:
            self.balance -= amount
            print(f"Made payment of {amount} units. Current balance is {self.balance} units.")
        else:
            print("Invalid amount for payment.")
    def get_balance(self):
        return self.balance
    def get_credit_limit(self):
        return self.credit_limit
    def get_card_info(self):
        return f"Card Number: {self.card_number}\nCard Holder: {self.card_holder}\nCredit Limit: {self.credit_limit} units\nCurrent Balance: {self.balance} units"



