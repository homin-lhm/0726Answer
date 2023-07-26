import test_pb2

contact = test_pb2.Contact()
contact.name = "homin"
contact.id = 123
contact.email = "lhmingwork@163.com"

data = contact.SerializeToString()

with open('./test.bin', 'wb') as f:
    f.write(data)

new_contact = test_pb2.Contact()
with open('./test.bin', 'rb') as f:
    new_contact.ParseFromString(f.read())

print(new_contact.name)
print(new_contact.id)
print(new_contact.email)
