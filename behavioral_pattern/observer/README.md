#Observer
#Khi nào sữ dụng
1. Khi chúng ta có 1 đối tượng, muốn đối tượng đó phát đi cho tất cả các thông báo cho tất cả các đôi tượng đã đăng ký nhận tin từ đối tượng đó thì ta sd design pattern này

#Thành phần
1. subject interface (phát đi thông báo và nhận đăng ký từ observer)
2. concrete subject -->item
3. observer interface (Nhận thông báo từ subject)
4. concrete observer -->customer

#Ví dụ
1.  Ta có 1 đối tượng item (cái này mình muốn bán một cái gì đó thì nó là item). Khi có 1 item mới thì customer muốn biến có một cái item mới. Thì cái item phát đi 1 cái thông báo cho các item mới cho các costomer đã đăng ký nhận tin