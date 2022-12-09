# Reading notes:
- Concurrency không phải Parallel, nhưng để parallel đc thì chương trình phải là concurrency.
- Về Thread thì cần phân biệt OS Thread và Application threads. Application threads có thể hiểu là nhiều thread hơn OS Thread rất nhiều, và mapping M to N. nếu kernel level thì 1:1, user level thì 1:N
- 1 process, có nhiều thread, OS sẽ execute process và trong lúc chạy cái process đó sẽ có các thread được chạy, nhưng tại 1 thời điểm chỉ có 1 thread/core
- Process có thể coi là một máy tính ảo, có không gian địa chỉ riêng. Thread mới là đc OS lập lịch trực tiếp.
- Process quản lý việc cũng cấp tài nguyên, nên mình có thể coi là VM. Thread mới là cái thực thi chương trình. Mỗi process có ít nhất 1 thread là main thread