### 课后作业

---

1.用新的API修改deque程序<br>

源程序：
```java
Deque<String> deque = new LinkedList<String>();
deque.push("a");
deque.push("b");
deque.push("c");
System.out.println(deque);
String str = deque.peek();
System.out.println(str);
System.out.println(deque);
while (deque.size() > 0) {
  System.out.println(deque.pop());
}
System.out.println(deque);
```
转换后的程序：
```java
Deque<String> deque = new LinkedList<String>();
// push ---> addFirst
deque.addFirst("a");
deque.addFirst("b");
deque.addFirst("c");
System.out.println(deque);
// peek  ---> peekFirst
String str = deque.peekFirst();
System.out.println(str);
System.out.println(deque);
while (deque.size() > 0) {
  // pop ---> removeFirst
  System.out.println(deque.removeFirst());
}
System.out.println(deque);
```
说明：<br>
> 上述API的转换明显的告诉我们，deque其实在一端的操作可以实现LIFO，即stack的效果


