# Face

## 1

> 数据结构：

1. hashmap原理与冲突解决方法
    底层的hashMap是由数组和链表来实现的。当插入的时候，会根据key的hash值然后计算出相应的数组下标，计算方法是index = hashcode%table.length，（这个下标就是上面提到的bucket），当这个下标上面已经存在元素的时候那么就会形成链表，将后插入的元素放到尾端，若是下标上面没有存在元素的话，那么直接将元素放到这个位置上。当进行查询的时候，同样会根据key的hash值先计算相应的下标，然后到相应的位置上进行查找，若是这个下标上面有很多元素的话，那么将在这个链表上一直查找直到找到对应的元素。
    通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。这个映射函数叫做散列函数，存放记录的数组叫做散列表。
    当关键字数量比较大的时候，可能会不同的关键字映射到同一个地址上，造成hash冲突。
    1. 开放寻址法：
       1.1. 线性探测再散列；顺序查看表的下一单元，直至找到某个空单元，或查遍全表；
       1.2. 二次探测再散列；在表的左右进行跳跃式探测；
       1.3. 伪随机数序列，称伪随机探测再散列。根据产生的随机数进行探测；
    2. 再散列法：建立多个hash函数，若是当发生hash冲突的时候，使用下一个hash函数，直到找到可以存放元素的位置；
    3. 拉链法（链地址法）：在冲突的位置上建立一个链表，然后将冲突的元素插入到链表尾端；
    4. 建立公共溢出区：将哈希表分为基本表和溢出表，将与基本表发生冲突的元素放入溢出表中。
2. 单向链表反转
    func reverseList(head *ListNode)*ListNode {
        var tail *ListNode
        for head != nil {
            head.Next, tail, head = tail, head, head.Next
        }
        return tail
    }
3. 快排
    // 快速排序是一种"分治法"，将原本的问题分解成两个子问题——比基准值小的数和比基准值大的数，然后再分别解决这两个子问题。
    // 返回调整完毕之后的base的位置，便于之后对两边的数组进行快排
    // 一轮排序完成后，基准值左边的都小于基准值，基准值右边的都大于基准值
    // 排序就分成了两个子问题，分别再对基准值左边和右边的数据用快排算法进行排序
    // 直到子问题里只剩下一个数字，再也无法分解出子问题的时候，整个序列的排序才算完成
4. 红黑树的作用是什么，和平衡树有什么区别
    红黑树中，所有的节点都是标准的2-节点，为了体现出3-节点，这里将3-节点的两个元素用左斜红色的链接连接起来，即连接了两个2-节点来表示一个3-节点。这里红色节点标记就代表指向其的链接是红链接，黑色标记的节点就是普通的节点。所以才会有那样一条定义，叫“从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点”，因为红色节点是可以与其父节点合并为一个3-节点的，红黑树实现的其实是一个完美的黑色平衡，如果你将红黑树中所有的红色链接放平，那么它所有的叶子节点到根节点的距离都是相同的。所以它并不是一个严格的平衡二叉树，但是它的综合性能已经很优秀了。
    红黑树的另一种定义是满足下列条件的二叉查找树：
        ⑴红链接均为左链接。
        ⑵没有任何一个结点同时和两条红链接相连。(这样会出现4-节点)
        ⑶该树是完美黑色平衡的，即任意空链接到根结点的路径上的黑链接数量相同。

> C++：

1. 介绍c++ 智能指针: shared_ptr, weak_ptr, unique_ptr。举例一种循环引用的场景
    * 摒弃 auto_ptr 的原因：避免因潜在的内存问题导致程序崩溃。
    * unique_ptr持有对对象的独有权——两个unique_ptr 不能指向一个对象，即 unique_ptr 不共享它所管理的对象。它无法复制到其他 unique_ptr，无法通过值传递到函数，也无法用于需要副本的任何标准模板库 （STL）算法。只能移动 unique_ptr，即对资源管理权限可以实现转移。这意味着，内存资源所有权可以转移到另一个 unique_ptr，并且原始 unique_ptr 不再拥有此资源。实际使用中，建议将对象限制为由一个所有者所有，因为多个所有权会使程序逻辑变得复杂。
    * shared_ptr 是一个标准的共享所有权的智能指针，允许多个指针指向同一个对象。shared_ptr 利用引用计数的方式实现了对所管理的对象的所有权的分享，即允许多个 shared_ptr 共同管理同一个对象。像 shared_ptr 这种智能指针被称之为“引用计数型智能指针”。
    * weak_ptr 被设计为与 shared_ptr 共同工作，可以从一个 shared_ptr 或者另一个 weak_ptr 对象构造而来。weak_ptr 是为了配合 shared_ptr 而引入的一种智能指针，它更像是 shared_ptr 的一个助手而不是智能指针，因为它不具有普通指针的行为，没有重载 operator* 和 operator-> ，因此取名为 weak，表明其是功能较弱的智能指针。它的最大作用在于协助 shared_ptr 工作，可获得资源的观测权，像旁观者那样观测资源的使用情况。观察者意味着 weak_ptr 只对 shared_ptr 进行引用，而不改变其引用计数，当被观察的 shared_ptr 失效后，相应的 weak_ptr 也相应失效。

2. 右值的作用以及完美转发
    * 右值是指表达式结束时就不再存在的临时对象。所有的具名变量或者对象都是左值，而右值不具名。对表达式取地址，如果能，则为左值，否则为右值。
    int&& a = 1; // 实质上就是将不具名(匿名)变量取了个别名
    右值本来在表达式语句结束后，其生命也就该终结了（因为是临时变量），而通过右值引用，该右值又重获新生，其生命期将与右值引用类型变量的生命期一样，只要右值引用类型变量还活着，该右值临时变量将会一直存活下去。实际上就是给那个临时变量取了个名字。
    * 所谓转发，就是通过一个函数将参数继续转交给另一个函数进行处理，原参数可能是右值，可能是左值，如果还能继续保持参数的原有特征，那么它就是完美的。
    当T&&为模板参数时，在universal references和std::forward的合作下，能够完美的转发这4种类型。
3. stl常用容器以及底层原理
    * vector： 是动态数组。在堆上分配空间。vector 是动态空间，随着元素的加入，它的内部机制会自行扩充空间以容纳新元素。vector 动态增加大小，并不是在原空间之后持续新空间（因为无法保证原空间之后尚有可供配置的空间），而是以原大小的两倍另外配置一块较大的空间，然后将原内容拷贝过来，然后才开始在原内容之后构造新元素，并释放原空间。因此，对 vector 的任何操作，一旦引起空间重新配置，同时指向原vector 的所有迭代器就都失效了。总结：vector 常用来保存需要经常进行随机访问的内容，并且不需要经常对中间元素进行添加删除操作。
    * list： 底层是一个双向链表，而且是一个环状双向链表。这个特点使得它的随机存取非常没有效率。好处是每次插入或删除一个元素，就配置或释放一个元素空间，元素也是在堆中。因此，list 对于空间的运用有绝对的精准，一点也不浪费。而且，对于任何位置的元素插入或元素移除，永远是常数时间。总结：经常添加删除大对象的话，那么请使用 list。
    * deque： 是一种双向开口的连续线性空间，元素也是在堆中。所谓双向开口，意思是可以在队尾两端分别做元素的插入和删除操作。deque 和 vector 的最大差异，一在于 deque 允许于常数时间内对起头端进行元素的插入或移除操作，二在于deque没有所谓容量观念，因为它是动态地以分段连续空间组合而成，随时可以增加一段新的空间并链接在一起。换句话说，像 vector 那样“因旧空间不足而重新配置一块更大空间，然后复制元素，再释放旧空间”这样的事情在 deque 是不会发生的。deque 是由一段一段的定量连续空间构成。一旦有必要在 deque 的前端或尾端增加新空间，便配置一段定量连续空间，串接在整个 deque 的头端或尾端。deque 的最大任务，便是在这些分段的定量连续空间上，维护其整体连续的假象，并提供随机存取的接口。避开了“重新配置，复制，释放”的轮回，代价则是复杂的迭代器架构。因为有分段连续线性空间，就必须有中央控制器，而为了维持整体连续的假象，数据结构的设计及迭代器前进后退等操作都颇为繁琐。deque 采用一块所谓的 map 作为主控。这里的 map 是一小块连续空间，其中每个元素都是指针，指向另一段连续线性空间，称为缓冲区。缓冲区才是 deque 的存储空间主体。（ 底层数据结构为一个中央控制器和多个缓冲区）SGI STL 允许我们指定缓冲区大小，默认值 0 表示将使用 512 bytes 缓冲区。支持[]操作符，也就是支持随即存取，可以在前面快速地添加删除元素，或是在后面快速地添加删除元素，然后还可以有比较高的随机访问速度和vector 的效率相差无几。deque 支持在两端的操作：push_back,push_front,pop_back,pop_front等，并且在两端操作上与 list 的效率也差不多。在标准库中 vector 和 deque 提供几乎相同的接口，在结构上区别主要在于在组织内存上不一样，deque 是按页或块来分配存储器的，每页包含固定数目的元素；相反 vector 分配一段连续的内存，vector 只是在序列的尾段插入元素时才有效率，而 deque 的分页组织方式即使在容器的前端也可以提供常数时间的 insert 和 erase 操作，而且在体积增长方面也比 vector 更具有效率。总结：deque 在开始和最后添加元素都一样快，并提供了随机访问方法，像vector一样使用 [] 访问任意元素，但是随机访问速度比不上vector快，因为它要内部处理堆跳转。deque 也有保留空间。另外，由于 deque 不要求连续空间，所以可以保存的元素比 vector 更大，这点也要注意一下。还有就是在前面和后面添加元素时都不需要移动其它块的元素，所以性能也很高。
    * stack： 是一种先进后出（First In Last Out , FILO）的数据结构。它只有一个出口，stack 允许新增元素，移除元素，取得最顶端元素。但除了最顶端外，没有任何其它方法可以存取stack的其它元素，stack不允许遍历行为。以某种容器（ 一般用 list 或 deque 实现，封闭头部即可，不用 vector 的原因应该是容量大小有限制，扩容耗时）作为底部结构，将其接口改变，使之符合“先进后出”的特性，形成一个 stack，是很容易做到的。deque 是双向开口的数据结构，若以 deque 为底部结构并封闭其头端开口，便轻而易举地形成了一个stack。因此，SGI STL 便以 deque 作为缺省情况下的 stack 底部结构，由于 stack 系以底部容器完成其所有工作，而具有这种“修改某物接口，形成另一种风貌”之性质者，称为 adapter（配接器），因此，STL stack 往往不被归类为 container(容器)，而被归类为 container adapter。
    * queue： 是一种先进先出（First In First Out,FIFO）的数据结构。它有两个出口，queue 允许新增元素，移除元素，从最底端加入元素，取得最顶端元素。但除了最底端可以加入，最顶端可以取出外，没有任何其它方法可以存取 queue 的其它元素。以某种容器（ 一般用 list 或 deque 实现，封闭头部即可，不用 vector 的原因应该是容量大小有限制，扩容耗时）作为底部结构，将其接口改变，使之符合“先进先出”的特性，形成一个 queue，是很容易做到的。deque 是双向开口的数据结构，若以 deque 为底部结构并封闭其底部的出口和前端的入口，便轻而易举地形成了一个 queue。因此，SGI STL 便以 deque 作为缺省情况下的 queue 底部结构，由于 queue 系以底部容器完成其所有工作，而具有这种“修改某物接口，形成另一种风貌”之性质者，称为 adapter（配接器），因此，STL queue 往往不被归类为container(容器)，而被归类为 container adapter。stack 和 queue 其实是适配器，而不叫容器，因为是对容器的再封装。
    * heap： 并不归属于 STL 容器组件，它是个幕后英雄，扮演 priority queue（优先队列）的助手。priority queue 允许用户以任何次序将任何元素推入容器中，但取出时一定按从优先权最高的元素开始取。按照元素的排列方式，heap 可分为 max-heap 和 min-heap 两种，前者每个节点的键值(key)都大于或等于其子节点键值，后者的每个节点键值(key)都小于或等于其子节点键值。因此， max-heap 的最大值在根节点，并总是位于底层array或vector的起头处；min-heap 的最小值在根节点，亦总是位于底层array或vector起头处。STL 提供的是 max-heap最大堆。
    * priority_queue： 是一个拥有权值观念的 queue，它允许加入新元素，移除旧元素，审视元素值等功能。由于这是一个 queue，所以只允许在底端加入元素，并从顶端取出元素，除此之外别无其它存取元素的途径。priority_queue 带有权值观念，其内的元素并非依照被推入的次序排列，而是自动依照元素的权值排列（通常权值以实值表示）。权值最高者，排在最前面。缺省情况下 priority_queue 系利用一个 max-heap 完成，后者是一个以vector 表现的 complete binary tree.max-heap 可以满足 priority_queue 所需要的“依权值高低自动递减排序”的特性。priority_queue 完全以底部容器（一般为vector为底层容器）作为根据，再加上 heap 处理规则，所以其实现非常简单。缺省情况下是以 vector 为底部容器。queue 以底部容器完成其所有工作。具有这种“修改某物接口，形成另一种风貌“”之性质者，称为 adapter(配接器)，因此，STL priority_queue 往往不被归类为 container(容器)，而被归类为 container adapter。
    * set 和 multiset 容器： set 的特性是，所有元素都会根据元素的键值自动被排序。set 的元素不像 map 那样可以同时拥有实值(value)和键值(key)，set 元素的键值就是实值，实值就是键值，set不允许两个元素有相同的值。set 底层是通过红黑树（RB-tree）来实现的，由于红黑树是一种平衡二叉搜索树，自动排序的效果很不错，所以标准的 STL 的 set 即以 RB-Tree 为底层机制。又由于 set 所开放的各种操作接口，RB-tree 也都提供了，所以几乎所有的 set 操作行为，都只有转调用 RB-tree 的操作行为而已。multiset的特性以及用法和 set 完全相同，唯一的差别在于它允许键值重复，因此它的插入操作采用的是底层机制是 RB-tree 的 insert_equal() 而非 insert_unique()。
    * map 和 multimap 容器：map的特性是，所有元素都会根据元素的键值自动被排序。map 的所有元素都是 pair，同时拥有实值（value）和键值（key）。pair 的第一元素被视为键值，第二元素被视为实值。map不允许两个元素拥有相同的键值。由于 RB-tree 是一种平衡二叉搜索树，自动排序的效果很不错，所以标准的STL map 即以 RB-tree 为底层机制。又由于 map 所开放的各种操作接口，RB-tree 也都提供了，所以几乎所有的 map 操作行为，都只是转调 RB-tree 的操作行为。multimap 的特性以及用法与 map 完全相同，唯一的差别在于它允许键值重复，因此它的插入操作采用的是底层机制 RB-tree 的 insert_equal() 而非 insert_unique。
    * hash_set 和 hash_multiset 容器：hash_set 底层数据结构为 hash 表，无序，不重复。hash_multiset 底层数据结构为 hash 表，无序，不重复。
    * hash_map 和 hash_multimap 容器：hash_map 底层数据结构为 hash 表，无序，不重复。hash_multimap 底层数据结构为 hash 表，无序，不重复。
4. map 与 unordered_map区别，底层数据结构
    底层实现不同，map的底层是红黑树，unordered_map的底层是哈希表（散列表）；
    map是红黑树实现的，所以有排序功能，通过迭代器就可以输出排好序的序列，unordered_map没有排序的功能；
    map查找和插入的时间复杂度是O(logn)，unordered_map是接近O(1)，跟哈希函数相关，所以不需要排序的时候用Unordered_Map效率会更高；
    unordered_Map的数据，如果数据数量超过线性存储表（bucket）的数量查找效率就会降低，所以线性存储表的容量就要增加一倍(类似于vector)，是根据经验得到的。

> Linux：

1. 进程 线程 协程
    进程是操作系统中运行的程序，是操作系统资源管理的最小单位。一个进程可以管理多个线程，线程相对轻量，且线程共享进程地址空间。
    * 进程：为了控制和管理多个程序间的计算机资源，划分出资源管理的最小单元：进程，也就是上面说的内存中加载指令的最小单位。
    * 线程：即使划分了资源管理的最小单元，但是一个进程在运行的过程中，也不可能一直占据着CPU进行逻辑运算，运行过程中很可能在进行磁盘I/O或者网络I/O，资源还是有些浪费。为了更加充分利用CPU运算资源，提高资源利用率，设计了线程的概念。线程最大的特点是和创建它的进程共享地址空间。为何非要设计出线程呢？将进程再划分小一点(争取在进程空闲时划分为小的进程)，开多个进程也可以提升CPU的利用率，但是开多个进程的话，进程间通信又是个麻烦的事情，毕竟进程之间地址空间是独立的，没法像线程那样做到数据的共享，需要通过其他的手段来解决，比如管道等。一般操作系统都会分为内核态和用户态，用户态线程之间的地址空间是隔离的，而在内核态，所有线程都共享同一内核地址空间。不管是用户线程还是内核线程，都和进程一样，均由操作系统的调度器来统一调度（至少在Linux中是这样子）。
    * 线程池：线程本身的数据结构需要占用内存，频繁创建和销毁线程会加大系统的压力，如果开辟太多线程，系统调度的开销会很大。线程池就是在这样的场景下提出的，常见的线程池实现方案如图，线程池可以在初始化的时候批量创建线程，然后用户后续通过队列等方式提交业务逻辑，线程池中的线程进行逻辑的消费工作，可以在操作的过程中降低线程创建和销毁的开销，但是调度的开销还是存在的。
    * 协程：在多核场景下，如果是I/O密集型场景，就算开多个线程来处理，也未必能提升CPU的利用率，反而会增加线程切换的开销。另外，多线程之间假如存在临界区或者共享数据，那么同步的开销也是不可忽视的。协程恰恰就是用来解决该问题的。协程是轻量级线程，在一个用户线程上可以跑多个协程，这样可以提升单核的利用率。涉及网络IO、磁盘IO的任务都是IO密集型任务，特点：CPU消耗很少，任务的大部分时间都在等待IO操作完成(因为IO的速度远远低于CPU和内存的速度)。IO密集型任务执行期间，99%的时间花在IO上，花在CPU上的时间很少，因此用运行速度极快的C语言 ，完全无法提升运行效率。对于IO密集型任务，最合适的语言是开发效率最高(代码量最少)的语言。协程不像进程或线程那样可以让系统负责相关的调度工作，协程处于一个线程当中，系统是无感知的，如果需要在该线程中阻塞某个协程的话，需要手工进行调度。

> 数据库：

1. mysql原理
    1. MySQL逻辑架构：
    MySQL是一个关系数据库管理系统。MySQL架构可以在多种不同场景中应用并发挥良好作用。主要体现在存储引擎的架构上，插件式的存储引擎架构将查询处理和其它的系统任务以及数据的存储提取相分离。
    1).最上层：
    最上层是一些客户端和连接服务，包含本地的sock通信和大多数基于客户端/服务端工具实现的类似于tcp/ip的通信，主要完成一些类似于连接处理、授权认证及相关的安全方案，在该层上引用了线程池的概念，为通过认证安全接入的客户端提供线程。同样在该层上可以实现基于ssl的安全链接。服务器也会为安全接入的每个客户端验证它所具有的操作权限。
    2).第二层：
    第二层架构主要完成大多数的核心服务功能。如sql接口，并完成缓存的查询。sql的分析和优化以及部分内置函数的执行。所有跨存储引擎的功能也在这一层实现，如过程，函数等。在该层，服务器会解析查询并创建相应的内部解析树，并对其完成相应的优化如确定查询表的顺序，是否利用索引等。最后生成相应的执行操作。如select语句，服务器还会查询内部的缓存。如果缓存空间足够大，这样就解决大量读操作的环境中能够很好的提升系统的性能。
    3).存储引擎层：
    存储引擎真正的负责MySQL中数据的存储和提取，服务器通过API与存储引擎进行通信，不同的存储引擎具有的功能不同，这样我们可以根据自己的实际需进行选取。
    4).数据存储层：
    主要是将数据存储在运行于裸设备的文件系统之上，并完成于存储引擎的交互。
    2. 并发控制和锁的概念：
    当数据库中有多个操作需要修改同一数据时，不可避免的会产生数据的脏读。这时就需要数据库具有良好的并发控制能力，这一切在MySQL中都是由服务器和存储引擎来实现的。
    解决并发问题最有效的方案是引入了锁的机制，锁在功能上分为共享锁(shared lock)和排它锁(exclusive lock)即通常说的读锁和写锁。当一个select语句在执行时可以施加读锁，这样就可以允许其它的select操作进行，因为在这个过程中数据信息是不会被改变的这样就能够提高数据库的运行效率。当需要对数据更新时，就需要施加写锁了，不在允许其它的操作进行，以免产生数据的脏读和幻读。锁同样有粒度大小，有表级锁(table lock)和行级锁(row lock)，分别在数据操作的过程中完成行的锁定和表的锁定。这些根据不同的存储引擎所具有的特性也是不一样的。
    MySQL大多数事务型的存储引擎都不只是简单的行级锁，基于性能的考虑，他们一般在行级锁基础上实现了多版本并发控制(MVCC)。这一方案也被Oracle等主流的关系数据库采用。它是通过保存数据中某个时间点的快照来实现的，这样就保证了每个事务看到的数据都是一致的。
    3. 事务：
    事务就是一组原子性的SQL语句。可以将这组语句理解成一个工作单元，要么全部执行要么都不执行。默认MySQL中自动提交时开启的（start transaction）。
    事务具有ACID的特性：原子性：事务中的所有操作要么全部提交成功，要么全部失败回滚；一致性：数据库总是从给一个一致性的状态转换到另一个一致性的状态；隔离性：一个事务所做的修改在提交之前对其它事务是不可见的；持久性：一旦事务提交，其所做的修改便会永久保存在数据库中。
    死锁：两个或多个事务在同一资源上相互占用并请求锁定对方占用的资源，从而导致恶性循环的现象。对于死锁的处理：MySQL的部分存储引擎能够检测到死锁的循环依赖并产生相应的错误。InnoDB引擎解决的死锁的方案是将持有最少写锁的事务进行回滚。
    4. MySQL存储引擎及应用方案：
    MySQL采用插件式的存储引擎的架构，可以根据不同的需求为不同的表设置不同的存储引擎。
    常用MySQL存储引擎介绍：
    InnoDB引擎：将数据存储在表空间中，表空间由一系列的数据文件组成，由InnoDb管理；
    支持每个表的数据和索引存放在单独文件中(innodb_file_per_table)；
    支持事务，采用MVCC来控制并发，并实现标准的4个事务隔离级别，支持外键；
    索引基于聚簇索引建立，对主键查询有较高性能；
    数据文件的平台无关性，支持数据在不同的架构平台移植；
    能够通过一些工具支持真正的热备，如XtraBackup等；
    内部进行自身优化如采取可预测性预读，能够自动在内存中创建bash索引等
    MyISAM引擎：
    MySQL5.1默认，不支持事务和行级锁;
    提供大量的特性如全文索引、空间函数、压缩、延迟更新等;
    数据库故障后，安全恢复性;
    对于只读数据可以忍受故障恢复，MyISAM依然非常适用;
    日志服务器的场景也比较适用，只需插入和数据读取操作;
    不支持单表一个文件，会将所有的数据和索引内容分别存放在两个文件中;
    MyISAM对整张表加锁而不是对行，所以不适用写操作比较多的场景;
    支持索引缓存不支持数据缓存.

2. redis 内部的string是如何实现的
    Redis使用自己的简单动态字符串(simple dynamic string, SDS)的抽象类型。Redis中，默认以SDS作为自己的字符串表示。只有在一些字符串不可能出现变化的地方使用C字符串。
    SDS的定义如下:

    ``` C
    struct sdshdr {
        int len;    // 用于记录buf数组中使用的字节的数目,和SDS存储的字符串的长度相等
        int free;   // 用于记录buf数组中没有使用的字节的数目
        char buf[]; // 字节数组，用于储存字符串，buf的大小等于len+free+1，其中多余的1个字节是用来存储'\0'的。
    };
    ```

    SDS除了用来保存数据库中的字符串之外，SDS还被用作缓冲区（buffer），如AOF模块中的AOF缓冲区，以及客户端状态中的输入缓冲区
    使用SDS而不使用c语言的string的好处：
    1) 常数复杂度获取字符串长度
        C语言中:字符串只是简单的字符的数组，当使用strlen获取字符串长度的时候，内部其实是直接顺序遍历数组的内容，找到对应的’\0’对应的字符，从而计算出字符串的长度。
        SDS:只需要访问SDS的len属性就能得到字符串的长度，复杂度为O(1)。
    2) 杜绝缓冲区溢出
        Redis是C语言编写的，对于字符串的拼接、复制等操作，C语言开发者必须确保目标字符串的空间足够大，不然就会出现溢出的情况。
        SDS:API内部第一步会检测字符串的大小是否满足。如果空间已经满足要求，那么就像C语言一样操作即可。如果不满足，则拓展buf的空间之后再进行操作。每次操作之后，len和free的值会做相应的修改。
    3) 减少修改字符串时带来的内存重分配次数
        字符串长度增加操作时，进行空间预分配
        字符串长度减少操作时，惰性空间释放
    4) 二进制安全
        C字符串除了末尾之外不能出现’\0’，否则会被程序认为是字符串的结尾。这就使得C字符串只能存储文本数据，而不能保存图像，音频等二进制数据。
        使用SDS就不需要依赖控制符，而是用len来指定存储数据的大小，所有的SDS API都会以处理二进制的方式来处理SDS的buf的数据。程序不会对buf的数据做任何限制、过滤或假设，数据写入的时候是什么，读取的时候依然不变。
    5) 兼容部分C字符串函数
        SDS的buf的定义（字符串末尾为’\0’)和C字符串完全相同，因此很多的C字符串的操作都是适用于SDS->buf的。比如当buf里面存的是文本字符串的时候，大多数通过调用C语言的函数就可以。

    | C字符串      | SDS |
    | ----------- | ----------- |
    | 获取字符串长度的复杂度为O(N)      | 获取字符串长度的复杂度为O(1)       |
    | API是不安全的，可能会造成缓冲区溢出   | API是安全的，不会造成缓冲区溢出        |
    | 修改字符串长度N次必然需要执行N次内存重分配   | 修改字符串长度N次最多需要执行N次内存重分配        |
    | 只能保存文本数据   | 可以保存文本或者二进制数据        |
    | 可以使用所有库中的函数   | 可以使用一部分库的函数        |

2. redis key 失效机制
    redis的Key失效机制有两种：被动方式(passive way)和主动方式(active way)
    * 被动方式
        当客户端尝试访问key时，如果发现key已经失效了，就删除该key，并且告诉客户端该key已经失效了。
    * 主动方式
        随机抽取一部分的key进行校验，如果已经失效，就删除淘汰。
        Redis每秒执行以下操作10次：
            1) 在有过期时间的key集合中随机抽取20个key。
            2) 删除所有的过期key
            3) 如果过期的key超过25%，重新执行步骤1)

3. redis zset的实现原理
    跳表（skiplist）本质上是并行的有序链表，但它克服了有序链表插入和查找性能不高的问题。它的插入和查询的时间复杂度都是O(logN)。在查询上跟平衡树的复杂度一致，因此是替代平衡树的方案。在redis的zset，leveldb都有应用。跳表为节点设置了快速访问的指针，不同在一个个遍历，可以跨越节点进行访问，这也是跳表的名字的含义。
    跳表原理：
    * 每个跳表都必须设定一个最大的连接层数MaxLevel
    * 第一层连接会连接到表中的每个元素
    * 插入一个元素会随机生成一个连接层数值[1, MaxLevel]之间，根据这个值跳表会给这元素建立N个连接
    * 插入某个元素的时候先从最高层开始，当跳到比目标值大的元素后，回退到上一个元素，用该元素的下一层连接进行遍历，周而复始直到第一层连接，最终在第一层连接中找到合适的位置

    > skiplist在redis zset的使用：
    redis中skiplist的MaxLevel设定为32层skiplist原理中提到skiplist一个元素插入后，会随机分配一个层数，而redis的实现，这个随机的规则是：
    * 一个元素拥有第1层连接的概率为100%
    * 一个元素拥有第2层连接的概率为50%
    * 一个元素拥有第3层连接的概率为25%
    * 以此类推...
    为了提高搜索效率，redis会缓存MaxLevel的值，在每次插入/删除节点后都会去更新这个值，这样每次搜索的时候不需要从32层开始搜索，而是从MaxLevel指定的层数开始搜索
    > 查找过程
    对于zrangebyscore命令：score作为查找的对象，在跳表中跳跃查询，就和上面skiplist的查询一样对于zrange还没搞清楚
    > 插入过程
    zadd [zset name] [score] [value]：
    * 在map中查找value是否已存在，如果存在现需要在skiplist中找到对应的元素删除，再在skiplist做插入
    * 插入过程也是用score来作为查询位置的依据，和skiplist插入元素方法一样。并需要更新value->score的map
    * 如果score一样怎么办？根据value再排序，按照顺序插入
    > 删除过程
    zrem [zset name] [value]：从map中找到value所对应的score，然后再在跳表中搜索这个score,value对应的节点，并删除
    > 排名是怎么算出来的
    zrank [zset name] [value]的实现依赖与一些附加在跳表上的属性：
    * 跳表的每个元素的Next指针都记录了这个指针能够跨越多少元素，redis在插入和删除元素的时候，都会更新这个值
    * 然后在搜索的过程中按经过的路径将路径中的span值相加得到rank

4. redis 主从复制的步骤
    为了分担读压力，Redis支持主从复制，Redis的主从结构可以采用一主多从或者级联结构，Redis主从复制可以根据是否是全量分为全量同步和增量同步。
    * 阶段一：建立连接阶段
    建立slave到master的连接，使master能够识别slave，并保存slave端口号
    1. 设置master的地址和端口，保存master信息
    2. 建立socket连接
    3. 发送ping命令（定时器任务）
    4. 身份验证
    5. 发送slave端口信息
    * 阶段二：数据同步阶段工作流程
    1. 在slave初次连接master后，复制master中的所有数据到slave
    2. 将slave的数据库状态更新成master当前的数据库状态
    * 阶段三：命令传播阶段
    > 当master数据库状态被修改后，导致主从服务器数据库状态不一致，此时需要让主从数据同步到一致的状态，同步的动作称为命令传播
    > master将接收到的数据变更命令发送给slave，slave接收命令后执行命令
    > 主从复制过程大体可以分为3个阶段
        建立连接阶段（即准备阶段）
        数据同步阶段
        命令传播阶段
            命令传播阶段的部分复制
                命令传播阶段出现了断网现象
                    网络闪断闪连 …忽略
                    短时间网络中断 …部分复制
                    长时间网络中断 …全量复制
                部分复制的三个核心要素
                    服务器的运行 id（run id）
                    主服务器的复制积压缓冲区
                    主从服务器的复制偏移量

5. 如何做持久化的
    我们需要一种持久化的机制，来保存内存中的数据，否则数据就会在redis宕机的时候直接丢失。
    Redis有两种方式来实现数据的持久化，分别是RDB（Redis Database）和AOF（Append Only File)，可以先简单的把RDB理解为某个时刻的Redis内存中的数据快照，而AOF则是所有记录了所有修改内存数据的指令的集合（也就是Redis指令的集合），而这两种方式都会生成相应的文件落地到磁盘上，实现数据的持久化，方便下次恢复使用。
    * RDB 持久化
        RDB：Redis database 的简称。Redis 的默认持久化方式。
        RDB 中持久化生成的是一个经过压缩的二进制文件。
        RDB 持久化时机：
            1. 在客户端执行 SAVE 或者 BGSAVE
            2. 根据配置规则自动快照（稍后会讲到）
            3. 执行 FLUSHALL 命令
            4. 执行复制（replication）
        RDB 持久化步骤：
            1. fork 复制出一个父进程的副本子进程
            2. 子进程将内存中的数据写入到硬盘中的临时文件
            3. 将临时文件替换旧的 rdb 文件
        自动间隔保存（配置规则）：
            save 900 1     # 每900秒检查一次，如果有1条数据修改了，那么执行 rdb
            save 300 10    # 每300秒检查一次，如果有10条数据修改了，那么执行 rdb
            save 60 10000  # 每60秒检查一次，如果有10000条数据修改了，那么执行 rdb
        RDB 文件还原
            服务器启动时，会直接载入 RDB 文件。
            但是如果 AOF 文件存在，则会载入 AOF 文件。
    * AOF 持久化
        AOF 是 Append Only File 的简称。
        AOF 通过保存客户端传过来的写命令来记录数据库的状态。
        AOF 持久化的时机
            刷新到硬盘的时机
            由于操作系统的缓存机制，每次写入到 aof 文件之后，其实是写入到了系统缓存中，默认情况下，操作系统每 30 秒同步磁盘一次
        AOF 持久化的步骤
            1. 命令追加
            2. 文件写入
            3. 文件同步
        AOF 重写的时机
            主动执行 BGREWRITEAOF 命令触发 AOF 重写。
        AOF 重写步骤
            1. 父进程写入 AOF 缓冲区和 AOF 重写缓冲区
            2. 子进程执行 AOF 重写，完成之后发送信号给父进程
            3. 父进程收到信号将 AOF 重写缓冲区的内容写入到新的 AOF 文件中，并且覆盖原有的 AOF 文件
        AOF 文件还原
            1. 创建一个伪客户端（fake client）
            2. 从 AOF 文件中分析并读取一条写命令
            3. 使用伪客户端执行写命令
            4. 一直重复步骤 2 和 3

6. Rocksdb原理
    RocksDB是facebook开源的NOSQL存储系统，其设计是基于Google开源的LevelDB，优化了LevelDB中存在的一些问题，其性能要比LevelDB强，设计与LevelDB极其类似。
    RocksDB相对传统的关系数据库的一大改进是采用LSM树存储引擎。LSM树是非常有创意的一种数据结构，它和传统的B+树不太一样。
    * B+树
    B+树根节点和枝节点分别记录每个叶子节点的最小值，并用一个指针指向叶子节点。叶子节点里每个键值都指向真正的数据块，每个叶子节点都有前指针和后指针，这是为了做范围查询时，叶子节点间可以直接跳转，从而避免再去回溯至枝和跟节点。
    B+树最大的性能问题是会产生大量的随机IO，随着新数据的插入，叶子节点会慢慢分裂，逻辑上连续的叶子节点在物理上往往不连续，甚至分离的很远，但做范围查询时，会产生大量读随机IO。
    对于大量的随机写也一样，如果新插入的数据存储在磁盘上相隔很远，会产生大量的随机写IO，低下的磁盘寻道速度严重影响性能。
    * LSM树（Log-Structured Merge Tree）
    LSM树而且通过批量存储技术规避磁盘随机写入问题。 LSM树的设计思想非常朴素, 它的原理是把一颗大树拆分成N棵小树， 它首先写入到内存中（内存没有寻道速度的问题，随机写的性能得到大幅提升），在内存中构建一颗有序小树，随着小树越来越大，内存的小树会flush到磁盘上。磁盘中的树定期可以做merge操作，合并成一棵大树，以优化读性能。
    * LevelDB特点
        1. LevelDB是一个持久化存储的KV系统，和Redis这种内存型的KV系统不同，LevelDB不会像Redis一样狂吃内存，而是将大部分数据存储到磁盘上。
        2. LevleDb在存储数据时，是根据记录的key值有序存储的，就是说相邻的key值在存储文件中是依次顺序存储的，而应用可以自定义key大小比较函数。
        3. LevelDB支持数据快照（snapshot）功能，使得读取操作不受写操作影响，可以在读操作过程中始终看到一致的数据。
        4. LevelDB还支持数据压缩等操作，这对于减小存储空间以及增快IO效率都有直接的帮助。
    * RocksDB对LevelDB的优化
        1. 增加了column family，这样有利于多个不相关的数据集存储在同一个db中，因为不同column family的数据是存储在不同的sst和memtable中，所以一定程度上起到了隔离的作用。
        2. 采用了多线程同时进行compaction的方法，优化了compact的速度。
        3. 增加了merge operator，优化了modify的效率
        4. 将flush和compaction分开不同的线程池，能有效的加快flush，防止stall。
        5. 增加了对write ahead log(WAL)的特殊管理机制，这样就能方便管理WAL文件，因为WAL是binlog文件。

7. 如何重构sql语句，为什么能优化时间复杂度
    1. 使用where条件限定要查询的数据，避免返回多余的行
    2. 尽量避免在索引列上使用mysql的内置函数
    3. 尽量避免在 where 子句中对字段进行表达式操作，这将导致系统放弃使用索引而进行全表扫描
    4. Inner join 、left join、right join，优先使用Inner join，如果是left join，左边表结果尽量小
    5. 避免在 where 子句中使用!=或<>操作符
    6. 使用联合索引时，注意索引列的顺序，一般遵循最左匹配原则
    7. 考虑在 where 及 order by 涉及的列上建立索引，尽量避免全表扫描
    8. 若插入数据过多，考虑批量插入
    9. 谨慎使用distinct关键字
    10. 若数据量较大，优化你的修改/删除语句
    11. 删除冗余和重复索引
    12. 优化like语句
    13. 优化limit分页：返回上次查询的最大记录(偏移量)；order by + 索引；在业务允许的情况下限制页数
    14. 避免在where子句中使用or来连接条件
    15. 避免使用select *，而是select具体字段
    16. 若知道查询结果只有一条或者只要最大/最小一条记录，建议用limit 1
    17. 不要有超过5个以上的表连接
    18. exist & in的合理利用
    19. where子句中考虑用默认值代替null
    20. 索引不宜太多，一般5个以内

> 分布式：

1. 分布式的三个原则
    CAP原则：在一个分布式系统中， Consistency（一致性）、 Availability（可用性）、Partition tolerance（分区容错性），三者不可得兼
    * 一致性（C）：每次读取要么获得最近写入的数据，要么获得一个错误。
    * 可用性（A）：每次请求都能获得一个（非错误）响应，但不保证返回的是最新写入的数据(对应HA，机器故障时仍然能够提供服务)。
    * 分区容错性（P）：尽管任意数量的消息被节点间的网络丢失（或延迟），系统仍继续运行（网络分区/网络故障时仍然能够提供服务）。

2. Paxos，raft ，gossip算法。
    * Paxos算法
    是一种基于消息传递的分布式一致性算法，解决的问题正是分布式一致性问题，即一个分布式系统中的各个进程如何就某个值（决议）达成一致。

    一个或多个提议进程 (Proposer) 可以发起提案 (Proposal)，Paxos算法使所有提案中的某一个提案，在所有进程中达成一致。系统中的多数派同时认可该提案，即达成了一致。最多只针对一个确定的提案达成一致。

    Paxos将系统中的角色分为提议者 (Proposer)，决策者 (Acceptor)，和最终决策学习者 (Learner):
        Proposer: 提出提案 (Proposal)。Proposal信息包括提案编号 (Proposal ID) 和提议的值 (Value)。
        Acceptor：参与决策，回应Proposers的提案。收到Proposal后可以接受提案，若Proposal获得多数Acceptors的接受，则称该Proposal被批准。
        Learner：不参与决策，从Proposers/Acceptors学习最新达成一致的提案（Value）。
    在多副本状态机中，每个副本同时具有Proposer、Acceptor、Learner三种角色。

    Paxos算法通过一个决议分为两个阶段（Learn阶段之前决议已经形成）：
        第一阶段：Prepare阶段。Proposer向Acceptors发出Prepare请求，Acceptors针对收到的Prepare请求进行Promise承诺。
        第二阶段：Accept阶段。Proposer收到多数Acceptors承诺的Promise后，向Acceptors发出Propose请求，Acceptors针对收到的Propose请求进行Accept处理。
        第三阶段：Learn阶段。Proposer在收到多数Acceptors的Accept之后，标志着本次Accept成功，决议形成，将形成的决议发送给所有Learners。

    Paxos算法流程中的每条消息描述如下：
        Prepare: Proposer生成全局唯一且递增的Proposal ID (可使用时间戳加Server ID)，向所有Acceptors发送Prepare请求，这里无需携带提案内容，只携带Proposal ID即可。
        Promise: Acceptors收到Prepare请求后，做出“两个承诺，一个应答”。

    * Raft算法
    Raft算法则是从多副本状态机的角度提出，用于管理多副本状态机的日志复制。Raft实现了和Paxos相同的功能，它将一致性分解为多个子问题：Leader选举（Leader election）、日志同步（Log replication）、安全性（Safety）、日志压缩（Log compaction）、成员变更（Membership change）等。同时，Raft算法使用了更强的假设来减少了需要考虑的状态，使之变的易于理解和实现。
    Raft将系统中的角色分为领导者（Leader）、跟从者（Follower）和候选人（Candidate）：
        Leader：接受客户端请求，并向Follower同步请求日志，当日志同步到大多数节点上后告诉Follower提交日志。
        Follower：接受并持久化Leader同步的日志，在Leader告之日志可以提交之后，提交日志。
        Candidate：Leader选举过程中的临时角色。
    Raft要求系统在任意时刻最多只有一个Leader，正常工作期间只有Leader和Followers。Follower只响应其他服务器的请求。如果Follower超时没有收到Leader的消息，它会成为一个Candidate并且开始一次Leader选举。收到大多数服务器投票的Candidate会成为新的Leader。Leader在宕机之前会一直保持Leader的状态。Raft算法将时间分为一个个的任期（term），每一个term的开始都是Leader选举。在成功选举Leader之后，Leader会在整个term内管理整个集群。如果Leader选举失败，该term就会因为没有Leader而结束。
    **Leader选举**：Raft 使用心跳（heartbeat）触发Leader选举。当服务器启动时，初始化为Follower。Leader向所有Followers周期性发送heartbeat。如果Follower在选举超时时间内没有收到Leader的heartbeat，就会等待一段随机的时间后发起一次Leader选举。Follower将其当前term加一然后转换为Candidate。它首先给自己投票并且给集群中的其他服务器发送 RequestVote RPC （RPC细节参见八、Raft算法总结）。结果有以下三种情况：
        赢得了多数的选票，成功选举为Leader；
        收到了Leader的消息，表示有其它服务器已经抢先当选了Leader；
        没有服务器赢得多数的选票，Leader选举失败，等待选举时间超时后发起下一次选举。
    选举出Leader后，Leader通过定期向所有Followers发送心跳信息维持其统治。若Follower一段时间未收到Leader的心跳则认为Leader可能已经挂了，再次发起Leader选举过程。Raft保证选举出的Leader上一定具有最新的已提交的日志。
    **日志同步**：Leader选出后，就开始接收客户端的请求。Leader把请求作为日志条目（Log entries）加入到它的日志中，然后并行的向其他服务器发起 AppendEntries RPC （RPC细节参见八、Raft算法总结）复制日志条目。当这条日志被复制到大多数服务器上，Leader将这条日志应用到它的状态机并向客户端返回执行结果。
    某些Followers可能没有成功的复制日志，Leader会无限的重试 AppendEntries RPC直到所有的Followers最终存储了所有的日志条目。日志由有序编号（log index）的日志条目组成。每个日志条目包含它被创建时的任期号（term），和用于状态机执行的命令。如果一个日志条目被复制到大多数服务器上，就被认为可以提交（commit）了。
    **安全性**：Raft增加了如下两条限制以保证安全性：
        拥有最新的已提交的log entry的Follower才有资格成为Leader。
        Leader只能推进commit index来提交当前term的已经复制到大多数服务器上的日志，旧term日志的提交要等到提交当前term的日志来间接提交（log index 小于 commit index的日志被间接提交）。
    **日志压缩**：在实际的系统中，不能让日志无限增长，否则系统重启时需要花很长的时间进行回放，从而影响可用性。Raft采用对整个系统进行snapshot来解决，snapshot之前的日志都可以丢弃。每个副本独立的对自己的系统状态进行snapshot，并且只能对已经提交的日志记录进行snapshot。Snapshot中包含以下内容：
        日志元数据。最后一条已提交的 log entry的 log index和term。这两个值在snapshot之后的第一条log entry的AppendEntries RPC的完整性检查的时候会被用上。
        系统当前状态。
    **成员变更**：成员变更是在集群运行过程中副本发生变化，如增加/减少副本数、节点替换等。

    * gossip算法

> 编程：

1. 要求给定一个固定数组，一个随机区间范围与一个元素值。返回区间内元素个数。

例子
list = [ 3, 3, 5, 6, 1, 2, 7, 8, 0, 1 ]
start = 4
end = 9
element = 1
表示 查询 list[4:9] 内 1的个数
返回 2
要求时间复杂度在O(logn)
func searchInArray(nums []int, start, end, elem int) int {
    i, sum := start, 0
    for start < end {
        if nums[i] == elem {
            sum++
        }
        i++
    }
    return sum
}

## 2

1. Given a singly linkedlist L : L0 → L1 → ··· → Ln−1 → Ln,
   reorder it to: L0 → Ln → L1 → Ln−1 →L2 →Ln−2 → ···

   ``` golang
    func reorderList(head *ListNode) {
        if head == nil || head.Next == nil || head.Next.Next == nil {
            return
        }

        // 找中点，链表分成两个：快指针走完链表时，慢指针刚好只走了一半
        slow, fast := head, head
        for fast != nil && fast.Next != nil {
            slow, fast = slow.Next, fast.Next.Next
        }

        // 第二个链表倒置
        var secondHalf *ListNode
        for slow != nil {
            slow.Next, secondHalf, slow = secondHalf, slow, slow.Next
        }

        // 链表节点依次连接
        current := head
        for current != secondHalf && current.Next != secondHalf {
            next1, next2 := current.Next, secondHalf.Next
            current.Next, secondHalf.Next = secondHalf, next1
            current, secondHalf = next1, next2
        }
    }

   ```

2. lr 公式，极大似然
3. tensorflow
   版本：
   训练数据量 6h 2亿
   分布式：
   数据读取：
   tf使用方面
4. 用户兴趣表示
   特征分组
   特征交差
5. 多目标
   由3->8目标
   目标互斥 拆单塔 梯队截断
   目标相近  
   目标主次
   融合 均值/方差
6. 两阶段训练

## 3

> 数据结构：

1. 单向链表反转

``` golang
func reverseList(head *ListNode) {
    if head == nil {
        return
    }

    var tail *ListNode
    for head != nil {
        head.Next, tail, head = tail, head, head.Next
    }
}
```

2. 快排

``` golang
// 快速排序是一种"分治法"，将原本的问题分解成两个子问题——比基准值小的数和比基准值大的数，然后再分别解决这两个子问题。
// 解决子问题的时候会再次使用快速排序，只有在子问题里只剩下一个数字的时候，排序才算完成。

func QuickSort(nums []int) []int {
 quickSort(nums, 0, len(nums)-1)
 return nums
}

func quickSort(nums []int, left, right int) {
 if left < right {
  mid := partition(nums, left, right)
  quickSort(nums, left, mid-1)
  quickSort(nums, mid+1, right)
 }
}

// 返回调整完毕之后的base的位置，便于之后对两边的数组进行快排
// 一轮排序完成后，基准值左边的都小于基准值，基准值右边的都大于基准值
// 排序就分成了两个子问题，分别再对基准值左边和右边的数据用快排算法进行排序
// 直到子问题里只剩下一个数字，再也无法分解出子问题的时候，整个序列的排序才算完成
func partition(nums []int, left, right int) int {
 base := nums[left] // 选择一个基准值
 for left < right {
  for left < right && nums[right] >= base { // 大于基准值的往左移
   right--
  }
  // 填补left位置空值
  // right指针值 < base， right值移到left位置
  // right位置值空
  nums[left] = nums[right]

  for left < right && nums[left] <= base { // 小于基准值的往右移
   left++
  }
  // 填补right位置空值
  // left指针值 > base，left值移到right位置
  // left位置值空
  nums[right] = nums[left]
 }
 // base填补left位置的空值
 nums[left] = base

 return left
}
```

> 数据库：

1. InnoDB数据落盘要经历哪些缓存。
    数据InnoDB到磁盘需要经过:
    1. InnoDB buffer pool， Redo log buffer。InnoDB应用系统本身的缓冲。该层的缓冲都放在主机内存中，它的目的主要是在应用层管理自己的数据，避免慢速的读写操作影响了InnoDB的响应时间。
    2. page cache /Buffer cache。VFS层的缓冲。该层的缓冲都放在主机内存中，它的目的主要是在操作系统层缓冲数据，避免慢速块设备读写操作影响了IO的响应时间。
    3. Inode cache/directory buffer。这个也是vfs层的缓冲。需要通过O_SYNC或者fsync()来刷新。
    4. Write-Back buffer。存储控制器层。该层的缓冲都放在存储控制器的对应板载cache中，它的目的主要是在存储控制器层缓冲数据，避免慢速块设备读写操作影响了IO的响应时间。当数据被fsync()等刷到存储层时，首先会发送到存储控制器层。常见的存储控制器就是Raid卡，而目前大部分的Raid卡都有1G或者更大的存储容量。这个缓冲一般为易失性的存储，通过板载电池/电容来保证该“易失性的存储”的数据在机器断电以后仍然会同步到底层的磁盘存储介质上。
    5. Disk on-borad buffer。磁盘控制器层。该层的缓冲都放在磁盘控制器的对应板载cache中。存储设备固件(firmware)会按规则排序将写操作真正同步到介质中去。
2. InnoDB为什么使用b+树，b+树相比较其他树的优点。
    各种树解决的问题以及面临的新问题：
       1) 二叉查找树(BST)：解决了排序的基本问题，但是由于无法保证平衡，可能退化为链表；
       2) 平衡二叉树(AVL)：通过旋转解决了平衡的问题，但是旋转操作效率太低；
       3) 红黑树：通过舍弃严格的平衡和引入红黑节点，解决了AVL旋转效率过低的问题，但是在磁盘等场景下，树仍然太高，IO次数太多；
       4) B树：通过将二叉树改为多路平衡查找树，解决了树过高的问题；
       5) B+树：在B树的基础上，将非叶节点改造为不存储数据的纯索引节点，进一步降低了树的高度；此外将叶节点使用指针连接成链表，范围查询更加高效。

3. mysql的事务隔离级别, MVCC, WAL，几种日志文件的作用
    * MySQL的事务隔离级别一共有四个，分别是读未提交、读已提交、可重复读以及可串行化。
    MySQL的隔离级别的作用就是让事务之间互相隔离，互不影响，这样可以保证事务的一致性。
    隔离级别比较：可串行化 > 可重复读 > 读已提交 > 读未提交
    隔离级别对性能的影响比较：可串行化 > 可重复读 > 读已提交 > 读未提交
    隔离级别越高，所需要消耗的MySQL性能越大（如事务并发严重性），为了平衡二者，一般建议设置的隔离级别为可重复读，MySQL默认的隔离级别也是可重复读。
    * 在MySQL中，主要有5种日志文件：
        1. 错误日志(error log)：记录mysql服务的启停时正确和错误的信息，还记录启动、停止、运行过程中的错误信息。
        2. 查询日志(general log)：记录建立的客户端连接和执行的语句。
        3. 二进制日志(bin log)：记录所有更改数据的语句，可用于数据复制。
        4. 慢查询日志(slow log)：记录所有执行时间超过long_query_time的所有查询或不使用索引的查询。
        5. 中继日志(relay log)：主从复制时使用的日志。

> 编程：

1. 列出hostname字符串可以对应的全部合法ip地址

## 4

1. hashmap怎么实现，怎么解决冲突
    见上。
2. paddlepaddle/tf
  tensorflow使用版本 1.14
  特征处理使用tf.feature_column
  Embedding维度过大问题怎么解决
  embedding最大维度，5千万
  低频id过滤问题
3. 双塔召回
  流程
  spark处理数据转tfrecord
  天级训练模型
  Ad侧天级更新模型，向量
  User实时更新
  向量检索 手写
  整个流程用shell pipeline实现
  效果
  与场景有关，有些场景无效
  样本选取，正负样本选取
4. 排序
   xgboost升级到dnn，为什么
   din序列特征怎么处理的，tf.feature_column.sequence_xxx
5. ctr/cvr
  mmoe两个任务如何关联，有没有侧重
  mmoe比essm效果好，为什么
6. 链表旋转最后k个节点
1->2->3->4->5, 3, 1->2->5->4->3

    ``` golang

// 链表旋转最后k个节点
// head: 1->2->3->4->5, k: 3, want: 1->2->5->4->3
func rotateKthFromEnd(head *ListNode, k int)*ListNode {
 if head == nil || head.Next == nil {
  return head
 }

 // 找k点，链表分成两个
 h1 := head
 firstHalf := &ListNode{Next: h1}
 for h1 != nil && k > 0 {
  h1 = h1.Next
  firstHalf.Next = h1.Next
  k--
 }

 // 找到后部分链表
 h2 := head
 for h1 != nil {
  h1, h2 = h1.Next, h2.Next
 }

 // 反转后部分链表
 var tail *ListNode
 for h2 != nil {
  h2.Next, tail, h2 = tail, h2, h2.Next
 }

 // 两部分连接到一起
 firstHalf.Next = tail

 return firstHalf
}
    ```

## 5

1. 在寒武纪做芯片适配tf，包括tf device/kernel等
2. 描述下Session.run流程
3. 线程池何时创建，有哪些线程池
    TODO
4. 怎么创建op，写过哪些op，op在何时实例化
5. 训练io优化   prefetch
6. horovod架构
7. 何时进行梯度规约，怎么保证不同worker因规约死锁
8. estimator做了哪些开发  
9. 大规模embedding

## 6

1. 编程题：路径简化
    TODO
2. hashmap实现
3. lr loss function
   为什么要用simoid
   lr怎么实现多分类  1.softmax 2.多个lr
   两种多分类问题区别
   初始化方法
   优化器  怎么算梯度
   过拟合 怎么解决
4. xgboost
   基分类器
   损失函数
   树分裂方式
   常用的参数
   错误点：将样本采样说为xgboost样本加权

## 7

1. 编程题：分糖果问题
2. hashmap，冲突怎么解决，二叉树的遍历
3. lr loss function
   怎么优化
   哪些场景 广告召回 一致性/
4. gbdt 原理，基分类器，剪枝
   怎么调参，
   xgboost
5. attention结构
6. 图模型
7. tf版本
8. esmm

## 8

1. 链表partition  数组partition
    TODO
2. hashmap，二叉树遍历
3. 选一个自己最熟悉的算法
   Svm 最大化间隔
   为什么>=1
   为什么用拉格朗日
   对偶问题
   Xgboost 特点 加入正则；二阶泰勒展开
   基分类树 CART

## 9

1. 找到全部的从根节点到叶子的和为指定值的序列
    TODO
2. hashmap原理，怎么解决冲突，排序算法有哪些，时间复杂度
3. 京东机器学习平台
   分布式任务调度  k8s
   数据读取 worker 动态获取文件
   分布式预估
   任务状态检测 检测任务是否成功
4. 图计算框架
5. tf，自定义op
6. lr，fm，gbdt等机器学习算法

## 10

> 基础题

1. java 1.8 新增加的特性，对函数式编程的理解
2. 对final的理解
    TODO
3. 线程池参数设置
    TODO
4. java中Lock和synchronized
    TODO
5. 什么是AOP?使用场景?实现原理是什么
6. 一致性hash算法
    TODO
7. kafka，使用场景，如何保证顺序消费
    TODO
8. InnoDB为什么使用b+树，b+树相比较其他树的优点

> 算法题

1. 根据给定值旋转数组
    TODO
2. 返回两数之和为某值的数组下标
    TODO

## 11

1. 二分查找
2. 上述的二分查找题变形了一下，在排序数组中找小于目标数的最大的数的位置
    TODO
3. 最大子数组和
    TODO
4. 环数组最大子数组和
    TODO
5. 判断链表中是否有环
    TODO

## 12

> java基础

1. jdk1.8特性与函数式编程的理解
2. 线程池的使用，类型和参数
3. hashmap线程安全问题与内部数据结构
4. java常用的原子类与介绍
5. jvm内存区域介绍
6. java类加载的过程
7. spring框架ioc和aop原理和用法
8. mysql索引优化与分库分表

> 算法

1. 冒泡排序
    TODO
2. 单链表反转
    TODO

## 13

>主要考察了java基础、数据结构、redis、mysql、算法

>基础

1. java中hashmap线程安全问题和底层实现
2. java中final关键字的理解
3. jvm 运行时内存区域介绍
4. 双亲委派机制
5. 匿名内部类
6. spring中aop在工作中的用法
7. redis数据结构和用法
8. mysql优化和底层数据结构

> 算法题

1. 单向链表反转
    TODO

## 14

1. java中list，set和map的区别及底层实现，及动态数组扩容，是否线程安全
    哈希碰撞
2. 进程、线程、进程、线程间通信
    TODO

> 算法题：

1. 二叉树，平衡二叉树，先、中、后遍历、按层序遍历（BFS）、左视图
    TODO
2. 快排
    TODO
3. 用两个栈实现一个队列
    TODO

## 15

1. list，set和map的区别及底层实现
2. udp，tcp，三次握手四次挥手
    TODO
3. TIME_WAIT状态，CLOSE_WAIT状态
    TODO
4. io复用
    TODO
5. 锁，死锁，redis，设计模式
    TODO

> 算法题：

1. 快排，二叉树先、中、后序、层序遍历，左视图
    TODO
2. 一个有序数组，后M个数通过转换排到前面，例如：
            1 2 3 4 5 6 7 8 9
            7 8 9 1 2 3 4 5 6
            时间复杂度<=O(N)
            空间复杂度=O(1)
    TODO

## 16

1. 交叉墒，lr极大似然
2. mmoe 3个塔

## 17

1. 详细介绍摇钱树Flink实时项目数据源接入到数据处理再到数据的落地，日处理量2千万
2. Canal同步mysql binlog 如何保证数据一致性？
3. Kafka leader挂掉后，如何选举新的leader?
    TODO
4. 消费者消费Kafka partition如何分配？
5. 生产者写入Kafka ack策略？
6. Flink Sql使用过哪些算子？基本的sum count  join group 等。是否自定义过UDF UDTF？
7. Flink sql 如何关联基础信息表？
    1.经常更新的可以用CDC 2.不频繁更新的直接映射到mysql 定时更新
8. 是否自定义过input sink?
9. Flink Streaming 自定义聚合函数和reduce的区别？
10. Flink有哪些state类型？
11. Flink 任务反压问题如何排查？
12. 数据倾斜如何优化？
    两阶段处理
    TODO
13. Flink 监控信息如何配置？如何查看？
14. Flink 配置文件有改过什么？
15. 机器负载高如何排查？
    TODO
16. Jdk 1.8垃圾回收
17. Redis常用数据结构

## 18

> c++：

1. 虚析构函数：虚析构函数的使用，虚函数实现方式
    TODO
2. 容器：常用vector、list、map、queue等，遍历迭代器删除元素
3. c++ 智能指针的种类和各自特点
4. map/unordered_map的特点和使用场景

> 网络框架：

1. epoll+多线程模型实现
    TODO

> 代码

1. 实现二叉树层次遍历
    TODO

## 19

1. ceph读写的io路径，osd端的处理流程
2. ceph存储引擎：filestore和bluestore的原理
3. rocksdb的读写流程
    TODO
4. c++ 智能指针
5. 右值引用
6. 代码实现BlockingQueue
    TODO

## 20

1. redis rdb的原理
    TODO
2. 存储概念的理解：块、对象、文件
    TODO
3. C++ 智能指针: unique_ptr shared_ptr的使用
4. C++ 右值引用，移动构造函数
5. 虚析构函数的作用
    TODO
6. 大端小端
    TODO
7. Leveldb的原理，布隆过滤器、compact的原理
    TODO
8. 写代码：实现一个BlockingQueue
    TODO

## 21

1. redis 常用数据结构以及底层实现
2. redis 分布式锁的原理
    TODO
3. zk分布式锁的方案
    TODO
4. 如何解决缓存雪崩、击穿、穿透问题
    TODO
5. mysql分库分表
    TODO
6. 数据库ID策略
    snowflake
    TODO
7. mysql B+树索引的优势
8. 限流策略
    TODO

> Java基础

1. hashmap实现原理
2. 线程池原理
    TODO
3. volatile关键字原理

> JVM相关

1. 垃圾收集算法以及优缺点适用场景

## 22

1. 分布式全局递增Id实现方式，
    雪花算法模式，mysql ,redis
    TODO
2. 数据落mysql定时同步到hive，其他方案
    canal
3. Java线上排查工具
    JStack ，火焰图
4. zookeeper 选举算法
5. 线程池的使用和优劣对比
6. nio和bio有什么区别
7. 实现top100订单销售额
    两种方案：1.使用redis zset做，维护一个长度 2. 小顶堆
    TODO
8. A和B值互换
9. 数组数字 组成最大的数值
    TODO

## 23

1. 贵司服务发现怎么解决
2. k8s核心组件原理, apiserver数据的处理流程
3. kube-scheduler 怎么选主，多个schduler工作流程
    分布式锁来实现选主
4. 集群部署和升级怎么做的
5. scheduler 如果想让某些pod优先调度怎么做
6. k8s rbac机制怎么实现用户权限控制

## 24

1. Redis常用数据结构使用和数据清除策略
    TODO
2. 实现top100订单销售额，在数据不断增加的情况下优化
3. Mysql索引通过B+和红黑树实现的差异和优劣
4. GC算法和生命周期
    TODO
5. 多线程原理和并发锁的介绍
    TODO
6. 算法题：a和b值互换
    将两者之和存入a，值互换 b 即和减去自身，a 即和减去 b
    a = a + b; b = a - b; a = a - b;
    Go编译器帮我们在栈上创建了一个临时变量 temp, 然后按顺序交换其他各个变量的值。
