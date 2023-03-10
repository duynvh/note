# Question
#####1. What's different between pessimistic lock and optimistic lock
- Pessimistic lock: Who came before the first person who locked the door? The people at the back have to wait for them to release the lock before going in to check if the item is still in stock. If the door is locked, please do not allow anyone else to come out and check
- Optimistic lock: Put everything in your cat at the same time, and don't wait for each orther (everyone can select one item from the stock). However, if you're in a hurry to buy, make sure to check if the status of the goods has changed. If you successfully purcharse an item, its status will be updated

#####2. What's shared lock?
- Used when a transaction wants to read data
- Because only read is allowed, any other transaction that wants to write during the read period is not allowed
- In addition, share lock allows multiple transactions can read data at the same time

#####3. What's exclusive lock?
- Used when a transaction want read or write data
- Only one transaction has the right to obtain an exclusive lock on that data
- Other, no other transaction can apply share lock
- Therefore, it prevents other transactions from taking over the execution time.