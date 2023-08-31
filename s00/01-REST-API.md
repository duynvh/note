### 1. RESTful API Conventions
- Use Nouns Instead of Verbs
- Plural Nouns
- Use Nesting to Show Relationships
- Versioning
- Slug-case for URL
- Snake_case for request, response body

### 2. HTTP Method
- Properties:
    + Safety: do not alter the server state
    + Idempotency: a same request is sent once or multiple times, the response is the same.
- Operations:
    + Create: POST
    + Read: GET
    + Update Totally: PUT
    + Delete/Disable: DELETE
    + Update Partially: PATCH

### 3. Problems of pagination
- 2 Problems:
  + Performance issue for a large dataset in relational DB
    + It can take quite some time to count all rows
    + Offset loops through rows to know how many should be skipped
- Solutions:
  + Cursor: 
  ```sql
  SELECT * FROM users WHERE id > last_id ORDER BY id LIMIT 10;
    ```
  + Deferred join: 
  ```sql
  SELECT * FROM  (SELECT id FROM users ORDER BY id LIMIT 100, 10) a USING id JOIN users b ON a.id = b.id;
    ```

### 4. Design API
- Design API for exporting a file with the size of 500MB.
- Process:
  + Query DB
  + Write file
  + Response file to client

- Issues:
  + Request timeout
  + Client is blocked
  + Out of Mem
  + Large result of the query

- Solutions: Use case export file
  + API request to export
    + Endpoint: GET /products/jobs/export?name=pen
  + API check status
    + Endpoint: GET /jobs/001
  + API get job result
    + Endpoint: GET /jobs/001/result

### 5. Types of Async API
- Polling:
    + Easy to implement
    + Waste resource
    + Use case: small load, import/export file
- Callback / Webhook
    + Optimize the resource
    + Complex to implement in both client side and server side
    + Use case: large load, payment

### 6. Problems
- Problem: A request might be sent twice due to network issue or replay attack. This problem is sensitive to use cases such as payment, order.
- Solution:
  + Client generates and adds an Idempotency Key to the request header.
  + Server checks Idempotency Key with unique constraint in DB.

### 7. API Document
- [REST API Document | Word](./document/REST%20API%20Document.docx)
- [REST API Document | PDF](./document/REST%20API%20Document.pdf)

### 8. References
- [Microsoft - best practice api design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)
- [Rest API best practices rest endpoint design examples](https://www.freecodecamp.org/news/rest-api-best-practices-rest-endpoint-design-examples/)
- [Google cloud - Restful web API design](https://cloud.google.com/blog/products/api-management/restful-web-api-design-best-practices)
- [Stripe API](https://stripe.com/docs/api)
- [Optimize Paging in MySQL](https://www.iheavy.com/how-to-optimize-paging-in-mysql-3-best-ways/)