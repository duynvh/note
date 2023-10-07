### 1. Table Design

#### Data Types
- Smaller is usually better.
- Simple is good.
- Avoid NULL if possible. It’s harder for MySQL to optimize queries that refer to nullable columns, because they make indexes, index statistics, and value comparisons more complicated.

- <strong>Data Type / String</strong>
    + VARCHAR
        - Uses only as much space as it needs. Uses 1 or 2 extra bytes to record the value’s length
        - Best: Max Length < 255, frequently retrieved, infrequently update
        - Use cases: user name, subject, …
    + TEXT, BLOB
        - InnoDB may use a separate “external” storage area for TEXT, BLOB
        - MySQL can not index the full length of these data types and can’t use the indexes for sorting
        - Use case: logs, message, comments, …

- <strong>Data Type / Number</strong>
    + Integer
        - Use small if possible
        - Misunderstand: For storage and computational purposes, INT(1) is identical to INT(20)
    + Real number: 
        - FLOAT, DOUBLE consume less bytes than DECIMAL
        - To storage money, do not use FLOAT, DOUBLE, use DECIMAL and fractional number → accuracy

- <strong>Data Type / Date</strong>
    + Timestamp:
        - Pros: consume less space
        - Cons: limitation of value range
        - Use cases: to record a (more or less) fixed point in time. Example: created_at
    + Datetime:
        - Pros: readable, no limitation of value range
        - Cons: consume more space, depend time zone
        - Use cases: time can be set and changed arbitrarily. Example: appointment time
    + Practices:
        - Use a fractional seconds
        - Save time zone in addition
        - JVM time zone = OS time zone = DB time zone = time zone 0

#### Table Format
- Problem 1:
    + Creating another columns → Locking table → Downtime

- Format: Common
    + Id
    + Other columns
        - Grouping columns
    + status
    + data: json
    + created_at
    + created_by
    + updated_at
    + updated_by

- Problem 2
    + Need to index a field in JSON.
    + In older versions, DB do not support index on function
    +  → We need columns to index

- Format: Redundant Column
    + Create redundant columns when creating tables
    + Adding new data into the unused column
    + Document the meaning of columns

- Problem 3
    + Creating index → locking table
    + Solution: 
        - ALTER TABLE my_table ADD INDEX my_table__idx (my_column), 
ALGORITHM=INPLACE, LOCK=NONE;
        - [Create an index on a huge MySQL production table without table locking - Stack Overflow ](https://stackoverflow.com/questions/4244685/create-an-index-on-a-huge-mysql-production-table-without-table-locking)
        - Stil lock table
        - Cause: Foreign Key
        - Remove constraint foreign key, implement this constraint in app layer
        - [Adding indexes to very large tables in MySQL - Database Administrators Stack Exchange](https://dba.stackexchange.com/questions/261752/adding-indexes-to-very-large-tables-in-mysql)

#### Schema
- Entity-Relationship Diagram (ERD)
    + Entity: object has identity and life cycle
    + Entity-Relationship Diagram (ERD) is a visual diagram to illustrate the relationships between different entities or tables
- Normalization
    + Normalization is the process of organizing data to reduce redundancy and improve data integrity.
    + Normalization breaking down large tables into smaller, related tables and establishing relationships between them
    + Pros:
        - Reduce redundancy
        - Data integrity
    + Cons:
        - Increase complexity to schema
        - Join to get needed data
        - Some cases can not take effect of indexes
- Denormalization
    + Denormalization is the inverse process of normalization.
    + Denormalization combine tables or adding redundant data.
    + Pros:
        - Improved query performance
        - Index
    + Cons:
        - Increased Storage
        - Data Consistency

### 2. View
- Introduction
    + A view is a virtual table that doesn’t store any data itself.
    + Instead, the data “in” the base table.
    + Syntax: 
        CREATE VIEW simple_bookings AS 
        SELECT book_ref, date(book_date), total_amount
        FROM bookings;
    + Normally, view is not updatable.
    + There is a type of updatable views, but this is not recommended.
- How View Works?
    + Processing Algorithm:
        - MERGE (by default): merge user’s SQL + view’s SQL → final SQL → execute on base table (index).
        - TEMPTABLE: execute view’s SQL on base table 
            → temporary table 
            → execute user’s SQL on the temporary table 
            (no index in some case).

- Limitation
    + Views might trick developers into thinking they’re simple, when in fact they’re very complicated under the hood.
    + In some case, no index in used
    + MySQL does not support the materialized views. A materialized view generally stores its results in an invisible table behind the scenes, with periodic updates to refresh the invisible table from the source data

- Practices
    + Use Cases:
        - Simplify complex query (join, function,…)
        - Add extra security layers
            + HR can only view user profile
            + Accountant can view user balance
        - Enable backward compatibility
    + Best Practices:
        - Simple view → merge view
        - Check execution plan of SQL on views

### 3. Case Studies
- Ordering
    + Context: Reorder tasks in to-do list
    + Requirements:
        - Easy to order
        - Less operations to update the order

- Report
    + Context: Count total clicks in last day
    + Requirements 
        - Response time
        - Accuracy

### 4. Recap
- Leverage JSON if possible
- View is not magic. Use it in the right way.
- Technical issues might related to locking, ordering

### 5. References
- [MySQL view performance TEMPTABLE or MERGE?](https://stackoverflow.com/questions/42513839/mysql-view-performance-temptable-or-merge)
- [MySQL: Large VARCHAR vs. TEXT?](https://stackoverflow.com/questions/2023481/mysql-large-varchar-vs-text)
- [Create an index on a huge MySQL production table without table locking](https://stackoverflow.com/questions/4244685/create-an-index-on-a-huge-mysql-production-table-without-table-locking)
- [Adding indexes to very large tables in MySQL](https://dba.stackexchange.com/questions/261752/adding-indexes-to-very-large-tables-in-mysql)