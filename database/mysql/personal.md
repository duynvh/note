# Reading notes:
## 1. High Performance MYSQL
### Chapter 6 - Schema Design and Management
- Try to avoid extremes in your design, such as a schema that will force enormously complex queries, or tables with oodles and oodlres of columns
- Use small, simple, appropriate data types, and avoid NULL unless it’s actually the right way to model your data’s reality.
- Try to use the same data types to store similar or related values, especially if they’ll be used in a join condition.
- Watch out for variable-length strings, which might cause pessimistic full-length memory allocation for temporary tables and sorting.
- Try to use integers for identifiers if you can.
- Avoid the legacy MySQL-isms such as specifying precisions for floating-point numbers or display widths for integers.
- Be careful with ENUM and SET. They’re handy, but they can be abused, and they’re tricky sometimes. BIT is best avoided
- Normalization is good, but denormalization (duplication of data, in most cases) is sometimes actually necessary and beneficial. And precomputing, caching, or generating summary tables can also be a big win. Justin Swanhart’s Flexviews tool can help maintain summary tables.
- Finally, ALTER TABLE can be painful because in most cases, it locks and rebuilds the whole table. We showed a number of workarounds for specific cases; for the general case, you’ll have to use other techniques, such as performing the ALTER on a replica and then promoting it to master.

### Chapter 7 - Indexing for high performance
- Unused Indexes: `mysql> SELECT * FROM sys.schema_unused_indexes;`
- Most of the time you'll use B-tree indexes with MySQL. The other types of indexes are rather more suitable for speical purposes, and it will generally be obvious when you ought to use them and how they can improve query response times. We'll say no more about them in this chapter, but it's worth wrapping up with a review of the properties and uses of B-tree indexes
- Here are three principles to keep in mind as you choose indexes and write queries to take advantage of them:
  + Single-row access is slow, especially on spindle-based storage. (SSDs are faster at random I/O but this point remains true). If the server reads a block of data from storage and then accesses only one row in it, it waste a lot of work. It's much better to read in a block that contains lots of rows you need
  + Accessing ranges of rows in order is fast, for two reasons. First, sequential I/O doesn’t require disk seeks, so it is faster than random I/O, especially on spindle-based storage. Second, if the server can read the data in the order you need it, it doesn’t need to perform any follow-up work to sort it, and GROUP BY queries don’t need to sort and group rows together to compute aggregates over them.
  +  Index-only access is fast. If an index contains all the columns that the query needs, the storage engine doesn’t need to find the other columns by looking up rows in the table. This avoids lots of single-row access, which as we know from the first point is slow
- In sum, try to choose indexes and write queries so that you can avoid single-row lookups, use the inherent ordering of the data to avoid sorting operations, and exploit index-only access. This corresponds to the three-star ranking system set out in Lahdenmaki and Leach’s book, mentioned at the beginning of this chapter
- It would be great to be able to create perfect indexes for every query against your tables. Unfortunately, sometimes this would require an impractically large number of indexes, and at other times there simply is no way to create a three-star index for a
given query (for example, if the query is ordered by two columns, one ascending and the other descending). In these cases, you have to settle for the best you can do or pursue alternative strategies, such as denormalization or summary tables.
- It’s very important to be able to reason through how indexes work and to choose them based on that understanding, not on rules of thumb or heuristics such as “place the most selective columns first in multicolumn indexes” or “you should index all of the columns that appear in the WHERE clause
- How do you know whether your schema is indexed well enough? As always, we suggest that you frame the question in terms of response time. Find queries that are either taking too long or contributing too much load to the server. Examine the schema, SQL, and index structures for the queries that need attention. Determine whether the query has to examine too many rows, perform postretrieval sorting or use temporary tables, access data with random I/O, or look up full rows from the table to retrieve columns not included in the index
- If you find a query that doesn’t benefit from all of the possible advantages of indexes, see if a better index can be created to improve performance. If not, perhaps a rewrite
can transform the query so that it can use an index that either already exists or could be created. That’s what the next chapter is about.