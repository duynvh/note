# Question
#####1. What's a process in OS?
- In OS, a process is a program in execution. For example, when we write our program in C, we save, compile and run it, the moment we run it, it is transferred to the main memory and becomes a process. When a process executes, it passes through many different stages like Start, Ready, Running, Waiting and Terminated state
#####2. What's a thread?
- A thread is the smallest unit of execution within a process. A process can have multiple threads.
- It breaks up a task into smaller units that can be executed concurrently and thus provides concurrency within a process.
- It exists within a process and uses process resources such as code, data and files.
- Since, a thread is lightweight as compared to a process and helps to improve the execution performance by running in parallel with other threads.
- And just like a process, each thread have its own attributes such as thread ID, program counter, register set and a stack. 
- Threads enhance the throughput of the system because when a process is split into threads, the function of each thread is considered as one job and then the number of jobs done in the unit time increases
- In a multiprocessor architecture, threads can be distributed over a series of processors and can run in parallel on different processors
#####3. What is a deadlock?
- Deadlock is a situation where two processes are waiting for each other to complete in order to start their operation
- This happens when first process is holding a resource and waiting for another resource that is currently being held by the second processes and the second processes is waiting for the resource being held by the first process. 
- This will cause them to wait forever and be in a state of deadlock