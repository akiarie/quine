with open('rep', 'r') as f:
    data = f.read()

print([f'{c}' for c in data])
