#Adapter (bộ chuyển đổi)
#Khi nào dùng
1. Dùng với mục đích dùng dể chuyển đổi
2. Khi nào chúng ta chuyển 1 hành vi của một đối tượng mà khi chúng ta gọi thì no qua lớp trung gian lớp trung gian đó gọi là adapter.
Lớp trung gian đos sẽ giúp chúng ta cái input đầu vào đến 1 cái hành vi của một đối tượng khác.

#Thành phần
 1.Target interface (sẽ chứa các hành vi chung của 1 hoặc nhiều đối tượng)
 2. concrete prototype 1 thằng này là cài đặt bình thường (Đối tượng cài đặt tất cả các hành vi của của interface trên) --> Cài đặt bình thường 
 3. concrete prototype 2 -->(adapter) khi chúng ta gọi các hành vi của pt2, nó sẽ ko phải đi thực thi nó chuyển các lời gọi vào adaptee
 4. adaptee --> nó ko có phương thức trong interface, nó là phương thức khác hoàng toàn, người gọi vào sẽ ko biết được chuyển đi đâu. trong này chúng ta cài đặt chuyển cái hành vi đó về 1 cái hành vi của adaptee, cái adaptee no sẽ là người thực thi cuối cùng
 5. client (thành phần kết nối sữ dụng các concrete prototype, nó sẽ ko biết bên trong prototype 2  có cài đặt adapter
 
#Ví dụ:
1. computer
2. window
3. (adapter) ->macAdapter -->giúp sữ dụng cổng usb
4. Macbook
5. client insert usb

Giải thích: cumputer là đại diện cho các hành vi của cái máy tính.
-Các kiểu máy tính có cổng usb có cổng vuông là window
-MAcbbok thì nó có lỗ cắm hình tròn
-Chúng ta phải mua adapter để cắm vào sd mới sd được cổng usb của macbook
-adaptee ở đây là macbook
- client  thì người cắp: cắm usb vào mac hay win


#Let's go