# Reading notes:
### 1. Multitasking:
- Multitasking allows several activities to occur concurrently on the computer
    + Process-based multitasking
    + Thread-based multitasking
### 2. Thread and process
- Two threads share the same address space
- Context switching between threads is usually less expensive than between processes
- The cost of communication between threads in relatively low
### 3. Thread in JAVA:
- A thread is an independent sequential path of execution within a program
- Many threads can run concurrently within a program
- At runtime, threads in a program exist in a common memory space and can, therefore share both data and code (i.e., they are lightweight compared to processes)
- They also share the process running the program
- 3 Important conceps related to Multithreading in Java
    + Creating threads and providing the code that gets executed by a thread
    + Accessing common data and code through synchronization
    + Transitioning between thread states.
### 4. The main thread
- When a standablone application is run, a user thread is created to execute the main() method of the application is called the main thread
- If no other user threads are spawned, the program terminate the main() method finishes executing
- All other threads, called child threads, are spawned from the main thread
- The main() method can then finish, but the program will keep running until all user threads have completed.
- The runtime enviroment distinguishes between user threads and daemon threads.
- Calling the setDaemon(boolean) method in the Thread class marks the status of the threads as either daemon or user, but this must be done before the thread is started
- As long as a user thread is alive, the JVM does not terminate
- A daemon thread is at the mercy of the runtime system: it is stopped if there are no more user threads runnings, thus termination the program
### 5. Synchronization:
- Threads share the same memory space, i.e they can share resources (objects)
- However, there are critical situations where it is desirable that only one thread at a time has access to a shared resource.
### 6. Synchronized Methods:
- While a thread is inside a synchronized method of an object, all other threads that wish to execute this synchronized method or any other synchronized method of the object will have to wait
- This restriction does not apply to the thread that already has the lock and is executing a synchronized method of the object
- Such a method can invoke other synchronized methods of the object without being blocked
- The non-synchronized methods of the object can always be called at any time by any thread
### 7. Rules of synchronization:
- A thread must acquire the object lock associated with a shared resource before it can enter the shared resource.
- The runtime system ensures that no other thread can enter a shared resource if another thread already holds the object lock associated with it
- If a thread cannot immediately acquire the object lock, it is blocked, i.e, it must wait for the lock to become available
- When a thread exits a shared resource, the runtime system ensures that the object lock is also relinquished. If another thread is waiting for this object lock, it can try to acquire the lock in order to gain access to the shared resource.